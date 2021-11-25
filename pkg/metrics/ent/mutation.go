// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/direktiv/direktiv/pkg/metrics/ent/metrics"
	"github.com/direktiv/direktiv/pkg/metrics/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMetrics = "Metrics"
)

// MetricsMutation represents an operation that mutates the Metrics nodes in the graph.
type MetricsMutation struct {
	config
	op             Op
	typ            string
	id             *int
	namespace      *string
	workflow       *string
	instance       *string
	state          *string
	timestamp      *time.Time
	workflow_ms    *int64
	addworkflow_ms *int64
	isolate_ms     *int64
	addisolate_ms  *int64
	error_code     *string
	invoker        *string
	next           *int8
	addnext        *int8
	transition     *string
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Metrics, error)
	predicates     []predicate.Metrics
}

var _ ent.Mutation = (*MetricsMutation)(nil)

// metricsOption allows management of the mutation configuration using functional options.
type metricsOption func(*MetricsMutation)

// newMetricsMutation creates new mutation for the Metrics entity.
func newMetricsMutation(c config, op Op, opts ...metricsOption) *MetricsMutation {
	m := &MetricsMutation{
		config:        c,
		op:            op,
		typ:           TypeMetrics,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMetricsID sets the ID field of the mutation.
func withMetricsID(id int) metricsOption {
	return func(m *MetricsMutation) {
		var (
			err   error
			once  sync.Once
			value *Metrics
		)
		m.oldValue = func(ctx context.Context) (*Metrics, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Metrics.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMetrics sets the old Metrics of the mutation.
func withMetrics(node *Metrics) metricsOption {
	return func(m *MetricsMutation) {
		m.oldValue = func(context.Context) (*Metrics, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MetricsMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MetricsMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *MetricsMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetNamespace sets the "namespace" field.
func (m *MetricsMutation) SetNamespace(s string) {
	m.namespace = &s
}

// Namespace returns the value of the "namespace" field in the mutation.
func (m *MetricsMutation) Namespace() (r string, exists bool) {
	v := m.namespace
	if v == nil {
		return
	}
	return *v, true
}

// OldNamespace returns the old "namespace" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldNamespace(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNamespace is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNamespace requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNamespace: %w", err)
	}
	return oldValue.Namespace, nil
}

// ResetNamespace resets all changes to the "namespace" field.
func (m *MetricsMutation) ResetNamespace() {
	m.namespace = nil
}

// SetWorkflow sets the "workflow" field.
func (m *MetricsMutation) SetWorkflow(s string) {
	m.workflow = &s
}

// Workflow returns the value of the "workflow" field in the mutation.
func (m *MetricsMutation) Workflow() (r string, exists bool) {
	v := m.workflow
	if v == nil {
		return
	}
	return *v, true
}

// OldWorkflow returns the old "workflow" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldWorkflow(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldWorkflow is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldWorkflow requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWorkflow: %w", err)
	}
	return oldValue.Workflow, nil
}

// ResetWorkflow resets all changes to the "workflow" field.
func (m *MetricsMutation) ResetWorkflow() {
	m.workflow = nil
}

// SetInstance sets the "instance" field.
func (m *MetricsMutation) SetInstance(s string) {
	m.instance = &s
}

// Instance returns the value of the "instance" field in the mutation.
func (m *MetricsMutation) Instance() (r string, exists bool) {
	v := m.instance
	if v == nil {
		return
	}
	return *v, true
}

// OldInstance returns the old "instance" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldInstance(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldInstance is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldInstance requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInstance: %w", err)
	}
	return oldValue.Instance, nil
}

// ResetInstance resets all changes to the "instance" field.
func (m *MetricsMutation) ResetInstance() {
	m.instance = nil
}

// SetState sets the "state" field.
func (m *MetricsMutation) SetState(s string) {
	m.state = &s
}

// State returns the value of the "state" field in the mutation.
func (m *MetricsMutation) State() (r string, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old "state" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldState(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldState is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// ResetState resets all changes to the "state" field.
func (m *MetricsMutation) ResetState() {
	m.state = nil
}

// SetTimestamp sets the "timestamp" field.
func (m *MetricsMutation) SetTimestamp(t time.Time) {
	m.timestamp = &t
}

// Timestamp returns the value of the "timestamp" field in the mutation.
func (m *MetricsMutation) Timestamp() (r time.Time, exists bool) {
	v := m.timestamp
	if v == nil {
		return
	}
	return *v, true
}

// OldTimestamp returns the old "timestamp" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldTimestamp(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTimestamp is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTimestamp requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTimestamp: %w", err)
	}
	return oldValue.Timestamp, nil
}

// ResetTimestamp resets all changes to the "timestamp" field.
func (m *MetricsMutation) ResetTimestamp() {
	m.timestamp = nil
}

// SetWorkflowMs sets the "workflow_ms" field.
func (m *MetricsMutation) SetWorkflowMs(i int64) {
	m.workflow_ms = &i
	m.addworkflow_ms = nil
}

// WorkflowMs returns the value of the "workflow_ms" field in the mutation.
func (m *MetricsMutation) WorkflowMs() (r int64, exists bool) {
	v := m.workflow_ms
	if v == nil {
		return
	}
	return *v, true
}

// OldWorkflowMs returns the old "workflow_ms" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldWorkflowMs(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldWorkflowMs is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldWorkflowMs requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWorkflowMs: %w", err)
	}
	return oldValue.WorkflowMs, nil
}

// AddWorkflowMs adds i to the "workflow_ms" field.
func (m *MetricsMutation) AddWorkflowMs(i int64) {
	if m.addworkflow_ms != nil {
		*m.addworkflow_ms += i
	} else {
		m.addworkflow_ms = &i
	}
}

// AddedWorkflowMs returns the value that was added to the "workflow_ms" field in this mutation.
func (m *MetricsMutation) AddedWorkflowMs() (r int64, exists bool) {
	v := m.addworkflow_ms
	if v == nil {
		return
	}
	return *v, true
}

// ResetWorkflowMs resets all changes to the "workflow_ms" field.
func (m *MetricsMutation) ResetWorkflowMs() {
	m.workflow_ms = nil
	m.addworkflow_ms = nil
}

// SetIsolateMs sets the "isolate_ms" field.
func (m *MetricsMutation) SetIsolateMs(i int64) {
	m.isolate_ms = &i
	m.addisolate_ms = nil
}

// IsolateMs returns the value of the "isolate_ms" field in the mutation.
func (m *MetricsMutation) IsolateMs() (r int64, exists bool) {
	v := m.isolate_ms
	if v == nil {
		return
	}
	return *v, true
}

// OldIsolateMs returns the old "isolate_ms" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldIsolateMs(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldIsolateMs is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldIsolateMs requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsolateMs: %w", err)
	}
	return oldValue.IsolateMs, nil
}

// AddIsolateMs adds i to the "isolate_ms" field.
func (m *MetricsMutation) AddIsolateMs(i int64) {
	if m.addisolate_ms != nil {
		*m.addisolate_ms += i
	} else {
		m.addisolate_ms = &i
	}
}

// AddedIsolateMs returns the value that was added to the "isolate_ms" field in this mutation.
func (m *MetricsMutation) AddedIsolateMs() (r int64, exists bool) {
	v := m.addisolate_ms
	if v == nil {
		return
	}
	return *v, true
}

// ResetIsolateMs resets all changes to the "isolate_ms" field.
func (m *MetricsMutation) ResetIsolateMs() {
	m.isolate_ms = nil
	m.addisolate_ms = nil
}

// SetErrorCode sets the "error_code" field.
func (m *MetricsMutation) SetErrorCode(s string) {
	m.error_code = &s
}

// ErrorCode returns the value of the "error_code" field in the mutation.
func (m *MetricsMutation) ErrorCode() (r string, exists bool) {
	v := m.error_code
	if v == nil {
		return
	}
	return *v, true
}

// OldErrorCode returns the old "error_code" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldErrorCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldErrorCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldErrorCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldErrorCode: %w", err)
	}
	return oldValue.ErrorCode, nil
}

// ClearErrorCode clears the value of the "error_code" field.
func (m *MetricsMutation) ClearErrorCode() {
	m.error_code = nil
	m.clearedFields[metrics.FieldErrorCode] = struct{}{}
}

// ErrorCodeCleared returns if the "error_code" field was cleared in this mutation.
func (m *MetricsMutation) ErrorCodeCleared() bool {
	_, ok := m.clearedFields[metrics.FieldErrorCode]
	return ok
}

// ResetErrorCode resets all changes to the "error_code" field.
func (m *MetricsMutation) ResetErrorCode() {
	m.error_code = nil
	delete(m.clearedFields, metrics.FieldErrorCode)
}

// SetInvoker sets the "invoker" field.
func (m *MetricsMutation) SetInvoker(s string) {
	m.invoker = &s
}

// Invoker returns the value of the "invoker" field in the mutation.
func (m *MetricsMutation) Invoker() (r string, exists bool) {
	v := m.invoker
	if v == nil {
		return
	}
	return *v, true
}

// OldInvoker returns the old "invoker" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldInvoker(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldInvoker is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldInvoker requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInvoker: %w", err)
	}
	return oldValue.Invoker, nil
}

// ResetInvoker resets all changes to the "invoker" field.
func (m *MetricsMutation) ResetInvoker() {
	m.invoker = nil
}

// SetNext sets the "next" field.
func (m *MetricsMutation) SetNext(i int8) {
	m.next = &i
	m.addnext = nil
}

// Next returns the value of the "next" field in the mutation.
func (m *MetricsMutation) Next() (r int8, exists bool) {
	v := m.next
	if v == nil {
		return
	}
	return *v, true
}

// OldNext returns the old "next" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldNext(ctx context.Context) (v int8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNext is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNext requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNext: %w", err)
	}
	return oldValue.Next, nil
}

// AddNext adds i to the "next" field.
func (m *MetricsMutation) AddNext(i int8) {
	if m.addnext != nil {
		*m.addnext += i
	} else {
		m.addnext = &i
	}
}

// AddedNext returns the value that was added to the "next" field in this mutation.
func (m *MetricsMutation) AddedNext() (r int8, exists bool) {
	v := m.addnext
	if v == nil {
		return
	}
	return *v, true
}

// ResetNext resets all changes to the "next" field.
func (m *MetricsMutation) ResetNext() {
	m.next = nil
	m.addnext = nil
}

// SetTransition sets the "transition" field.
func (m *MetricsMutation) SetTransition(s string) {
	m.transition = &s
}

// Transition returns the value of the "transition" field in the mutation.
func (m *MetricsMutation) Transition() (r string, exists bool) {
	v := m.transition
	if v == nil {
		return
	}
	return *v, true
}

// OldTransition returns the old "transition" field's value of the Metrics entity.
// If the Metrics object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MetricsMutation) OldTransition(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTransition is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTransition requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTransition: %w", err)
	}
	return oldValue.Transition, nil
}

// ClearTransition clears the value of the "transition" field.
func (m *MetricsMutation) ClearTransition() {
	m.transition = nil
	m.clearedFields[metrics.FieldTransition] = struct{}{}
}

// TransitionCleared returns if the "transition" field was cleared in this mutation.
func (m *MetricsMutation) TransitionCleared() bool {
	_, ok := m.clearedFields[metrics.FieldTransition]
	return ok
}

// ResetTransition resets all changes to the "transition" field.
func (m *MetricsMutation) ResetTransition() {
	m.transition = nil
	delete(m.clearedFields, metrics.FieldTransition)
}

// Op returns the operation name.
func (m *MetricsMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Metrics).
func (m *MetricsMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MetricsMutation) Fields() []string {
	fields := make([]string, 0, 11)
	if m.namespace != nil {
		fields = append(fields, metrics.FieldNamespace)
	}
	if m.workflow != nil {
		fields = append(fields, metrics.FieldWorkflow)
	}
	if m.instance != nil {
		fields = append(fields, metrics.FieldInstance)
	}
	if m.state != nil {
		fields = append(fields, metrics.FieldState)
	}
	if m.timestamp != nil {
		fields = append(fields, metrics.FieldTimestamp)
	}
	if m.workflow_ms != nil {
		fields = append(fields, metrics.FieldWorkflowMs)
	}
	if m.isolate_ms != nil {
		fields = append(fields, metrics.FieldIsolateMs)
	}
	if m.error_code != nil {
		fields = append(fields, metrics.FieldErrorCode)
	}
	if m.invoker != nil {
		fields = append(fields, metrics.FieldInvoker)
	}
	if m.next != nil {
		fields = append(fields, metrics.FieldNext)
	}
	if m.transition != nil {
		fields = append(fields, metrics.FieldTransition)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MetricsMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case metrics.FieldNamespace:
		return m.Namespace()
	case metrics.FieldWorkflow:
		return m.Workflow()
	case metrics.FieldInstance:
		return m.Instance()
	case metrics.FieldState:
		return m.State()
	case metrics.FieldTimestamp:
		return m.Timestamp()
	case metrics.FieldWorkflowMs:
		return m.WorkflowMs()
	case metrics.FieldIsolateMs:
		return m.IsolateMs()
	case metrics.FieldErrorCode:
		return m.ErrorCode()
	case metrics.FieldInvoker:
		return m.Invoker()
	case metrics.FieldNext:
		return m.Next()
	case metrics.FieldTransition:
		return m.Transition()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MetricsMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case metrics.FieldNamespace:
		return m.OldNamespace(ctx)
	case metrics.FieldWorkflow:
		return m.OldWorkflow(ctx)
	case metrics.FieldInstance:
		return m.OldInstance(ctx)
	case metrics.FieldState:
		return m.OldState(ctx)
	case metrics.FieldTimestamp:
		return m.OldTimestamp(ctx)
	case metrics.FieldWorkflowMs:
		return m.OldWorkflowMs(ctx)
	case metrics.FieldIsolateMs:
		return m.OldIsolateMs(ctx)
	case metrics.FieldErrorCode:
		return m.OldErrorCode(ctx)
	case metrics.FieldInvoker:
		return m.OldInvoker(ctx)
	case metrics.FieldNext:
		return m.OldNext(ctx)
	case metrics.FieldTransition:
		return m.OldTransition(ctx)
	}
	return nil, fmt.Errorf("unknown Metrics field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MetricsMutation) SetField(name string, value ent.Value) error {
	switch name {
	case metrics.FieldNamespace:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNamespace(v)
		return nil
	case metrics.FieldWorkflow:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWorkflow(v)
		return nil
	case metrics.FieldInstance:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInstance(v)
		return nil
	case metrics.FieldState:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	case metrics.FieldTimestamp:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTimestamp(v)
		return nil
	case metrics.FieldWorkflowMs:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWorkflowMs(v)
		return nil
	case metrics.FieldIsolateMs:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsolateMs(v)
		return nil
	case metrics.FieldErrorCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetErrorCode(v)
		return nil
	case metrics.FieldInvoker:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInvoker(v)
		return nil
	case metrics.FieldNext:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNext(v)
		return nil
	case metrics.FieldTransition:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTransition(v)
		return nil
	}
	return fmt.Errorf("unknown Metrics field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MetricsMutation) AddedFields() []string {
	var fields []string
	if m.addworkflow_ms != nil {
		fields = append(fields, metrics.FieldWorkflowMs)
	}
	if m.addisolate_ms != nil {
		fields = append(fields, metrics.FieldIsolateMs)
	}
	if m.addnext != nil {
		fields = append(fields, metrics.FieldNext)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MetricsMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case metrics.FieldWorkflowMs:
		return m.AddedWorkflowMs()
	case metrics.FieldIsolateMs:
		return m.AddedIsolateMs()
	case metrics.FieldNext:
		return m.AddedNext()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MetricsMutation) AddField(name string, value ent.Value) error {
	switch name {
	case metrics.FieldWorkflowMs:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddWorkflowMs(v)
		return nil
	case metrics.FieldIsolateMs:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIsolateMs(v)
		return nil
	case metrics.FieldNext:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddNext(v)
		return nil
	}
	return fmt.Errorf("unknown Metrics numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MetricsMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(metrics.FieldErrorCode) {
		fields = append(fields, metrics.FieldErrorCode)
	}
	if m.FieldCleared(metrics.FieldTransition) {
		fields = append(fields, metrics.FieldTransition)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MetricsMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MetricsMutation) ClearField(name string) error {
	switch name {
	case metrics.FieldErrorCode:
		m.ClearErrorCode()
		return nil
	case metrics.FieldTransition:
		m.ClearTransition()
		return nil
	}
	return fmt.Errorf("unknown Metrics nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MetricsMutation) ResetField(name string) error {
	switch name {
	case metrics.FieldNamespace:
		m.ResetNamespace()
		return nil
	case metrics.FieldWorkflow:
		m.ResetWorkflow()
		return nil
	case metrics.FieldInstance:
		m.ResetInstance()
		return nil
	case metrics.FieldState:
		m.ResetState()
		return nil
	case metrics.FieldTimestamp:
		m.ResetTimestamp()
		return nil
	case metrics.FieldWorkflowMs:
		m.ResetWorkflowMs()
		return nil
	case metrics.FieldIsolateMs:
		m.ResetIsolateMs()
		return nil
	case metrics.FieldErrorCode:
		m.ResetErrorCode()
		return nil
	case metrics.FieldInvoker:
		m.ResetInvoker()
		return nil
	case metrics.FieldNext:
		m.ResetNext()
		return nil
	case metrics.FieldTransition:
		m.ResetTransition()
		return nil
	}
	return fmt.Errorf("unknown Metrics field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MetricsMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MetricsMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MetricsMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MetricsMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MetricsMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MetricsMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MetricsMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Metrics unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MetricsMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Metrics edge %s", name)
}
