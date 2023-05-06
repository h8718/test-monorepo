// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
)

// LogMsgDelete is the builder for deleting a LogMsg entity.
type LogMsgDelete struct {
	config
	hooks    []Hook
	mutation *LogMsgMutation
}

// Where appends a list predicates to the LogMsgDelete builder.
func (lmd *LogMsgDelete) Where(ps ...predicate.LogMsg) *LogMsgDelete {
	lmd.mutation.Where(ps...)
	return lmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lmd *LogMsgDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LogMsgMutation](ctx, lmd.sqlExec, lmd.mutation, lmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lmd *LogMsgDelete) ExecX(ctx context.Context) int {
	n, err := lmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lmd *LogMsgDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(logmsg.Table, sqlgraph.NewFieldSpec(logmsg.FieldID, field.TypeUUID))
	if ps := lmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lmd.mutation.done = true
	return affected, err
}

// LogMsgDeleteOne is the builder for deleting a single LogMsg entity.
type LogMsgDeleteOne struct {
	lmd *LogMsgDelete
}

// Where appends a list predicates to the LogMsgDelete builder.
func (lmdo *LogMsgDeleteOne) Where(ps ...predicate.LogMsg) *LogMsgDeleteOne {
	lmdo.lmd.mutation.Where(ps...)
	return lmdo
}

// Exec executes the deletion query.
func (lmdo *LogMsgDeleteOne) Exec(ctx context.Context) error {
	n, err := lmdo.lmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{logmsg.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lmdo *LogMsgDeleteOne) ExecX(ctx context.Context) {
	if err := lmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
