package mirror

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
)

// TODO: validate credentials helper
// TODO: failed mirror garbage collector?

type Manager struct {
	callbacks Callbacks
}

func NewManager(callbacks Callbacks) *Manager {
	mgr := &Manager{
		callbacks: callbacks,
	}

	go mgr.gc()

	return mgr
}

// garbage collector
func (d *Manager) gc() {
	ctx := context.Background()

	jitter := 1000 // milliseconds
	interval := time.Second * 10
	maxRunTime := 5 * time.Minute
	maxRecordTime := time.Hour * 48

	// TODO: gracefully close the loop
	for {
		a, _ := rand.Int(rand.Reader, big.NewInt(int64(jitter)*2))
		delta := int(a.Int64()) - jitter // this gets a value between +/- jitter
		time.Sleep(interval + time.Duration(delta*int(time.Millisecond)))

		// this first loop looks for operations that seem to have timed out
		procs, err := d.callbacks.Store().GetUnfinishedProcesses(ctx)
		if err != nil {
			d.callbacks.SysLogCrit(fmt.Sprintf("Failed to query unfinished mirror processes: %v", err))
			continue
		}

		for _, proc := range procs {
			if time.Since(proc.CreatedAt) > maxRunTime {
				d.callbacks.SysLogCrit(fmt.Sprintf("Detected an old unfinished mirror process '%s' for namespace '%s'. Terminating...", proc.ID.String(), proc.NamespaceID.String()))
				c, cancel := context.WithTimeout(ctx, 5*time.Second)
				err = d.Cancel(c, proc.ID)
				cancel()
				if err != nil {
					d.callbacks.SysLogCrit(fmt.Sprintf("Error cancelling old unfinished mirror process '%s' for namespace '%s': %v", proc.ID.String(), proc.NamespaceID.String(), err))
				}

				p, err := d.callbacks.Store().GetProcess(ctx, proc.ID)
				if err != nil {
					d.callbacks.SysLogCrit(fmt.Sprintf("Error requerying old unfinished mirror process '%s' for namespace '%s': %v", proc.ID.String(), proc.NamespaceID.String(), err))
					continue
				}

				if p.Status != ProcessStatusFailed && p.Status != ProcessStatusComplete {
					d.failProcess(p, errors.New("timed out"))
				}
			}
		}

		// this second loop deletes really old processes from the database so that it doesn't grow endlessly
		err = d.callbacks.Store().DeleteOldProcesses(ctx, time.Now().Add(-1*maxRecordTime))
		if err != nil {
			d.callbacks.SysLogCrit(fmt.Sprintf("Failed to query old mirror processes: %v", err))
			continue
		}
	}
}

// Cancel stops a currently running mirroring process.
func (d *Manager) Cancel(_ context.Context, _ uuid.UUID) error {
	// TODO

	return nil
}

func (d *Manager) silentFailProcess(p *Process) {
	p.Status = ProcessStatusFailed
	p.EndedAt = time.Now().UTC()
	_, e := d.callbacks.Store().UpdateProcess(context.Background(), p)
	if e != nil {
		d.callbacks.SysLogCrit(fmt.Sprintf("Error updating failed mirror process record in database: %v", e))

		return
	}
}

func (d *Manager) failProcess(p *Process, err error) {
	d.silentFailProcess(p)
	d.callbacks.ProcessLogger().Error(p.ID, fmt.Sprintf("Mirroring process failed %v", err), "process_id", p.ID)
}

func (d *Manager) setProcessStatus(ctx context.Context, process *Process, status string) error {
	process.Status = status
	if status == ProcessStatusComplete || status == ProcessStatusFailed {
		process.EndedAt = time.Now().UTC()
	}

	_, err := d.callbacks.Store().UpdateProcess(ctx, process)
	if err != nil {
		return err
	}

	return nil
}

// Execute ..
func (d *Manager) Execute(ctx context.Context, p *Process, get func(ctx context.Context) (Source, error), applyer Applyer) {
	err := d.setProcessStatus(ctx, p, ProcessStatusExecuting)
	if err != nil {
		//nolint:contextcheck
		d.failProcess(p, fmt.Errorf("updating process status: %w", err))

		return
	}

	src, err := get(ctx)
	if err != nil {
		//nolint:contextcheck
		d.failProcess(p, fmt.Errorf("initializing source: %w", err))

		return
	}
	defer func() {
		err := src.Free()
		if err != nil {
			d.callbacks.SysLogCrit(fmt.Sprintf("Error freeing mirror source: %v", err))
		}
	}()

	parser, err := NewParser(newPIDFormatLogger(d.callbacks.ProcessLogger(), p.ID), src)
	if err != nil {
		//nolint:contextcheck
		d.silentFailProcess(p)

		return
	}
	defer func() {
		err := parser.Close()
		if err != nil {
			d.callbacks.SysLogCrit(fmt.Sprintf("Error freeing mirror temporary files: %v", err))
		}
	}()

	err = applyer.apply(ctx, d.callbacks, p, parser, src.Notes())
	if err != nil {
		//nolint:contextcheck
		d.failProcess(p, fmt.Errorf("applying changes: %w", err))

		return
	}

	err = d.setProcessStatus(ctx, p, ProcessStatusComplete)
	if err != nil {
		//nolint:contextcheck
		d.failProcess(p, fmt.Errorf("updating process status: %w", err))

		return
	}
}
