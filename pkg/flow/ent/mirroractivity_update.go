// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirror"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirroractivity"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// MirrorActivityUpdate is the builder for updating MirrorActivity entities.
type MirrorActivityUpdate struct {
	config
	hooks     []Hook
	mutation  *MirrorActivityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MirrorActivityUpdate builder.
func (mau *MirrorActivityUpdate) Where(ps ...predicate.MirrorActivity) *MirrorActivityUpdate {
	mau.mutation.Where(ps...)
	return mau
}

// SetType sets the "type" field.
func (mau *MirrorActivityUpdate) SetType(s string) *MirrorActivityUpdate {
	mau.mutation.SetType(s)
	return mau
}

// SetStatus sets the "status" field.
func (mau *MirrorActivityUpdate) SetStatus(s string) *MirrorActivityUpdate {
	mau.mutation.SetStatus(s)
	return mau
}

// SetUpdatedAt sets the "updated_at" field.
func (mau *MirrorActivityUpdate) SetUpdatedAt(t time.Time) *MirrorActivityUpdate {
	mau.mutation.SetUpdatedAt(t)
	return mau
}

// SetEndAt sets the "end_at" field.
func (mau *MirrorActivityUpdate) SetEndAt(t time.Time) *MirrorActivityUpdate {
	mau.mutation.SetEndAt(t)
	return mau
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (mau *MirrorActivityUpdate) SetNillableEndAt(t *time.Time) *MirrorActivityUpdate {
	if t != nil {
		mau.SetEndAt(*t)
	}
	return mau
}

// ClearEndAt clears the value of the "end_at" field.
func (mau *MirrorActivityUpdate) ClearEndAt() *MirrorActivityUpdate {
	mau.mutation.ClearEndAt()
	return mau
}

// SetController sets the "controller" field.
func (mau *MirrorActivityUpdate) SetController(s string) *MirrorActivityUpdate {
	mau.mutation.SetController(s)
	return mau
}

// SetNillableController sets the "controller" field if the given value is not nil.
func (mau *MirrorActivityUpdate) SetNillableController(s *string) *MirrorActivityUpdate {
	if s != nil {
		mau.SetController(*s)
	}
	return mau
}

// ClearController clears the value of the "controller" field.
func (mau *MirrorActivityUpdate) ClearController() *MirrorActivityUpdate {
	mau.mutation.ClearController()
	return mau
}

// SetDeadline sets the "deadline" field.
func (mau *MirrorActivityUpdate) SetDeadline(t time.Time) *MirrorActivityUpdate {
	mau.mutation.SetDeadline(t)
	return mau
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (mau *MirrorActivityUpdate) SetNillableDeadline(t *time.Time) *MirrorActivityUpdate {
	if t != nil {
		mau.SetDeadline(*t)
	}
	return mau
}

// ClearDeadline clears the value of the "deadline" field.
func (mau *MirrorActivityUpdate) ClearDeadline() *MirrorActivityUpdate {
	mau.mutation.ClearDeadline()
	return mau
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (mau *MirrorActivityUpdate) SetNamespaceID(id uuid.UUID) *MirrorActivityUpdate {
	mau.mutation.SetNamespaceID(id)
	return mau
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (mau *MirrorActivityUpdate) SetNamespace(n *Namespace) *MirrorActivityUpdate {
	return mau.SetNamespaceID(n.ID)
}

// SetMirrorID sets the "mirror" edge to the Mirror entity by ID.
func (mau *MirrorActivityUpdate) SetMirrorID(id uuid.UUID) *MirrorActivityUpdate {
	mau.mutation.SetMirrorID(id)
	return mau
}

// SetNillableMirrorID sets the "mirror" edge to the Mirror entity by ID if the given value is not nil.
func (mau *MirrorActivityUpdate) SetNillableMirrorID(id *uuid.UUID) *MirrorActivityUpdate {
	if id != nil {
		mau = mau.SetMirrorID(*id)
	}
	return mau
}

// SetMirror sets the "mirror" edge to the Mirror entity.
func (mau *MirrorActivityUpdate) SetMirror(m *Mirror) *MirrorActivityUpdate {
	return mau.SetMirrorID(m.ID)
}

// AddLogIDs adds the "logs" edge to the LogMsg entity by IDs.
func (mau *MirrorActivityUpdate) AddLogIDs(ids ...uuid.UUID) *MirrorActivityUpdate {
	mau.mutation.AddLogIDs(ids...)
	return mau
}

// AddLogs adds the "logs" edges to the LogMsg entity.
func (mau *MirrorActivityUpdate) AddLogs(l ...*LogMsg) *MirrorActivityUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return mau.AddLogIDs(ids...)
}

// Mutation returns the MirrorActivityMutation object of the builder.
func (mau *MirrorActivityUpdate) Mutation() *MirrorActivityMutation {
	return mau.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (mau *MirrorActivityUpdate) ClearNamespace() *MirrorActivityUpdate {
	mau.mutation.ClearNamespace()
	return mau
}

// ClearMirror clears the "mirror" edge to the Mirror entity.
func (mau *MirrorActivityUpdate) ClearMirror() *MirrorActivityUpdate {
	mau.mutation.ClearMirror()
	return mau
}

// ClearLogs clears all "logs" edges to the LogMsg entity.
func (mau *MirrorActivityUpdate) ClearLogs() *MirrorActivityUpdate {
	mau.mutation.ClearLogs()
	return mau
}

// RemoveLogIDs removes the "logs" edge to LogMsg entities by IDs.
func (mau *MirrorActivityUpdate) RemoveLogIDs(ids ...uuid.UUID) *MirrorActivityUpdate {
	mau.mutation.RemoveLogIDs(ids...)
	return mau
}

// RemoveLogs removes "logs" edges to LogMsg entities.
func (mau *MirrorActivityUpdate) RemoveLogs(l ...*LogMsg) *MirrorActivityUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return mau.RemoveLogIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mau *MirrorActivityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mau.defaults()
	if len(mau.hooks) == 0 {
		if err = mau.check(); err != nil {
			return 0, err
		}
		affected, err = mau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MirrorActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mau.check(); err != nil {
				return 0, err
			}
			mau.mutation = mutation
			affected, err = mau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mau.hooks) - 1; i >= 0; i-- {
			if mau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mau *MirrorActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := mau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mau *MirrorActivityUpdate) Exec(ctx context.Context) error {
	_, err := mau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mau *MirrorActivityUpdate) ExecX(ctx context.Context) {
	if err := mau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mau *MirrorActivityUpdate) defaults() {
	if _, ok := mau.mutation.UpdatedAt(); !ok {
		v := mirroractivity.UpdateDefaultUpdatedAt()
		mau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mau *MirrorActivityUpdate) check() error {
	if _, ok := mau.mutation.NamespaceID(); mau.mutation.NamespaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MirrorActivity.namespace"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mau *MirrorActivityUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MirrorActivityUpdate {
	mau.modifiers = append(mau.modifiers, modifiers...)
	return mau
}

func (mau *MirrorActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mirroractivity.Table,
			Columns: mirroractivity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mirroractivity.FieldID,
			},
		},
	}
	if ps := mau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mau.mutation.GetType(); ok {
		_spec.SetField(mirroractivity.FieldType, field.TypeString, value)
	}
	if value, ok := mau.mutation.Status(); ok {
		_spec.SetField(mirroractivity.FieldStatus, field.TypeString, value)
	}
	if value, ok := mau.mutation.UpdatedAt(); ok {
		_spec.SetField(mirroractivity.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mau.mutation.EndAt(); ok {
		_spec.SetField(mirroractivity.FieldEndAt, field.TypeTime, value)
	}
	if mau.mutation.EndAtCleared() {
		_spec.ClearField(mirroractivity.FieldEndAt, field.TypeTime)
	}
	if value, ok := mau.mutation.Controller(); ok {
		_spec.SetField(mirroractivity.FieldController, field.TypeString, value)
	}
	if mau.mutation.ControllerCleared() {
		_spec.ClearField(mirroractivity.FieldController, field.TypeString)
	}
	if value, ok := mau.mutation.Deadline(); ok {
		_spec.SetField(mirroractivity.FieldDeadline, field.TypeTime, value)
	}
	if mau.mutation.DeadlineCleared() {
		_spec.ClearField(mirroractivity.FieldDeadline, field.TypeTime)
	}
	if mau.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirroractivity.NamespaceTable,
			Columns: []string{mirroractivity.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mau.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirroractivity.NamespaceTable,
			Columns: []string{mirroractivity.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mau.mutation.MirrorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirroractivity.MirrorTable,
			Columns: []string{mirroractivity.MirrorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirror.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mau.mutation.MirrorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirroractivity.MirrorTable,
			Columns: []string{mirroractivity.MirrorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirror.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mau.mutation.LogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirroractivity.LogsTable,
			Columns: []string{mirroractivity.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logmsg.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mau.mutation.RemovedLogsIDs(); len(nodes) > 0 && !mau.mutation.LogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirroractivity.LogsTable,
			Columns: []string{mirroractivity.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logmsg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mau.mutation.LogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirroractivity.LogsTable,
			Columns: []string{mirroractivity.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logmsg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mirroractivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MirrorActivityUpdateOne is the builder for updating a single MirrorActivity entity.
type MirrorActivityUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MirrorActivityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetType sets the "type" field.
func (mauo *MirrorActivityUpdateOne) SetType(s string) *MirrorActivityUpdateOne {
	mauo.mutation.SetType(s)
	return mauo
}

// SetStatus sets the "status" field.
func (mauo *MirrorActivityUpdateOne) SetStatus(s string) *MirrorActivityUpdateOne {
	mauo.mutation.SetStatus(s)
	return mauo
}

// SetUpdatedAt sets the "updated_at" field.
func (mauo *MirrorActivityUpdateOne) SetUpdatedAt(t time.Time) *MirrorActivityUpdateOne {
	mauo.mutation.SetUpdatedAt(t)
	return mauo
}

// SetEndAt sets the "end_at" field.
func (mauo *MirrorActivityUpdateOne) SetEndAt(t time.Time) *MirrorActivityUpdateOne {
	mauo.mutation.SetEndAt(t)
	return mauo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (mauo *MirrorActivityUpdateOne) SetNillableEndAt(t *time.Time) *MirrorActivityUpdateOne {
	if t != nil {
		mauo.SetEndAt(*t)
	}
	return mauo
}

// ClearEndAt clears the value of the "end_at" field.
func (mauo *MirrorActivityUpdateOne) ClearEndAt() *MirrorActivityUpdateOne {
	mauo.mutation.ClearEndAt()
	return mauo
}

// SetController sets the "controller" field.
func (mauo *MirrorActivityUpdateOne) SetController(s string) *MirrorActivityUpdateOne {
	mauo.mutation.SetController(s)
	return mauo
}

// SetNillableController sets the "controller" field if the given value is not nil.
func (mauo *MirrorActivityUpdateOne) SetNillableController(s *string) *MirrorActivityUpdateOne {
	if s != nil {
		mauo.SetController(*s)
	}
	return mauo
}

// ClearController clears the value of the "controller" field.
func (mauo *MirrorActivityUpdateOne) ClearController() *MirrorActivityUpdateOne {
	mauo.mutation.ClearController()
	return mauo
}

// SetDeadline sets the "deadline" field.
func (mauo *MirrorActivityUpdateOne) SetDeadline(t time.Time) *MirrorActivityUpdateOne {
	mauo.mutation.SetDeadline(t)
	return mauo
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (mauo *MirrorActivityUpdateOne) SetNillableDeadline(t *time.Time) *MirrorActivityUpdateOne {
	if t != nil {
		mauo.SetDeadline(*t)
	}
	return mauo
}

// ClearDeadline clears the value of the "deadline" field.
func (mauo *MirrorActivityUpdateOne) ClearDeadline() *MirrorActivityUpdateOne {
	mauo.mutation.ClearDeadline()
	return mauo
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (mauo *MirrorActivityUpdateOne) SetNamespaceID(id uuid.UUID) *MirrorActivityUpdateOne {
	mauo.mutation.SetNamespaceID(id)
	return mauo
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (mauo *MirrorActivityUpdateOne) SetNamespace(n *Namespace) *MirrorActivityUpdateOne {
	return mauo.SetNamespaceID(n.ID)
}

// SetMirrorID sets the "mirror" edge to the Mirror entity by ID.
func (mauo *MirrorActivityUpdateOne) SetMirrorID(id uuid.UUID) *MirrorActivityUpdateOne {
	mauo.mutation.SetMirrorID(id)
	return mauo
}

// SetNillableMirrorID sets the "mirror" edge to the Mirror entity by ID if the given value is not nil.
func (mauo *MirrorActivityUpdateOne) SetNillableMirrorID(id *uuid.UUID) *MirrorActivityUpdateOne {
	if id != nil {
		mauo = mauo.SetMirrorID(*id)
	}
	return mauo
}

// SetMirror sets the "mirror" edge to the Mirror entity.
func (mauo *MirrorActivityUpdateOne) SetMirror(m *Mirror) *MirrorActivityUpdateOne {
	return mauo.SetMirrorID(m.ID)
}

// AddLogIDs adds the "logs" edge to the LogMsg entity by IDs.
func (mauo *MirrorActivityUpdateOne) AddLogIDs(ids ...uuid.UUID) *MirrorActivityUpdateOne {
	mauo.mutation.AddLogIDs(ids...)
	return mauo
}

// AddLogs adds the "logs" edges to the LogMsg entity.
func (mauo *MirrorActivityUpdateOne) AddLogs(l ...*LogMsg) *MirrorActivityUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return mauo.AddLogIDs(ids...)
}

// Mutation returns the MirrorActivityMutation object of the builder.
func (mauo *MirrorActivityUpdateOne) Mutation() *MirrorActivityMutation {
	return mauo.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (mauo *MirrorActivityUpdateOne) ClearNamespace() *MirrorActivityUpdateOne {
	mauo.mutation.ClearNamespace()
	return mauo
}

// ClearMirror clears the "mirror" edge to the Mirror entity.
func (mauo *MirrorActivityUpdateOne) ClearMirror() *MirrorActivityUpdateOne {
	mauo.mutation.ClearMirror()
	return mauo
}

// ClearLogs clears all "logs" edges to the LogMsg entity.
func (mauo *MirrorActivityUpdateOne) ClearLogs() *MirrorActivityUpdateOne {
	mauo.mutation.ClearLogs()
	return mauo
}

// RemoveLogIDs removes the "logs" edge to LogMsg entities by IDs.
func (mauo *MirrorActivityUpdateOne) RemoveLogIDs(ids ...uuid.UUID) *MirrorActivityUpdateOne {
	mauo.mutation.RemoveLogIDs(ids...)
	return mauo
}

// RemoveLogs removes "logs" edges to LogMsg entities.
func (mauo *MirrorActivityUpdateOne) RemoveLogs(l ...*LogMsg) *MirrorActivityUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return mauo.RemoveLogIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mauo *MirrorActivityUpdateOne) Select(field string, fields ...string) *MirrorActivityUpdateOne {
	mauo.fields = append([]string{field}, fields...)
	return mauo
}

// Save executes the query and returns the updated MirrorActivity entity.
func (mauo *MirrorActivityUpdateOne) Save(ctx context.Context) (*MirrorActivity, error) {
	var (
		err  error
		node *MirrorActivity
	)
	mauo.defaults()
	if len(mauo.hooks) == 0 {
		if err = mauo.check(); err != nil {
			return nil, err
		}
		node, err = mauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MirrorActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mauo.check(); err != nil {
				return nil, err
			}
			mauo.mutation = mutation
			node, err = mauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mauo.hooks) - 1; i >= 0; i-- {
			if mauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MirrorActivity)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MirrorActivityMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mauo *MirrorActivityUpdateOne) SaveX(ctx context.Context) *MirrorActivity {
	node, err := mauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mauo *MirrorActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := mauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mauo *MirrorActivityUpdateOne) ExecX(ctx context.Context) {
	if err := mauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mauo *MirrorActivityUpdateOne) defaults() {
	if _, ok := mauo.mutation.UpdatedAt(); !ok {
		v := mirroractivity.UpdateDefaultUpdatedAt()
		mauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mauo *MirrorActivityUpdateOne) check() error {
	if _, ok := mauo.mutation.NamespaceID(); mauo.mutation.NamespaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MirrorActivity.namespace"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mauo *MirrorActivityUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MirrorActivityUpdateOne {
	mauo.modifiers = append(mauo.modifiers, modifiers...)
	return mauo
}

func (mauo *MirrorActivityUpdateOne) sqlSave(ctx context.Context) (_node *MirrorActivity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mirroractivity.Table,
			Columns: mirroractivity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mirroractivity.FieldID,
			},
		},
	}
	id, ok := mauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MirrorActivity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mirroractivity.FieldID)
		for _, f := range fields {
			if !mirroractivity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mirroractivity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mauo.mutation.GetType(); ok {
		_spec.SetField(mirroractivity.FieldType, field.TypeString, value)
	}
	if value, ok := mauo.mutation.Status(); ok {
		_spec.SetField(mirroractivity.FieldStatus, field.TypeString, value)
	}
	if value, ok := mauo.mutation.UpdatedAt(); ok {
		_spec.SetField(mirroractivity.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mauo.mutation.EndAt(); ok {
		_spec.SetField(mirroractivity.FieldEndAt, field.TypeTime, value)
	}
	if mauo.mutation.EndAtCleared() {
		_spec.ClearField(mirroractivity.FieldEndAt, field.TypeTime)
	}
	if value, ok := mauo.mutation.Controller(); ok {
		_spec.SetField(mirroractivity.FieldController, field.TypeString, value)
	}
	if mauo.mutation.ControllerCleared() {
		_spec.ClearField(mirroractivity.FieldController, field.TypeString)
	}
	if value, ok := mauo.mutation.Deadline(); ok {
		_spec.SetField(mirroractivity.FieldDeadline, field.TypeTime, value)
	}
	if mauo.mutation.DeadlineCleared() {
		_spec.ClearField(mirroractivity.FieldDeadline, field.TypeTime)
	}
	if mauo.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirroractivity.NamespaceTable,
			Columns: []string{mirroractivity.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mauo.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirroractivity.NamespaceTable,
			Columns: []string{mirroractivity.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mauo.mutation.MirrorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirroractivity.MirrorTable,
			Columns: []string{mirroractivity.MirrorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirror.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mauo.mutation.MirrorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirroractivity.MirrorTable,
			Columns: []string{mirroractivity.MirrorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirror.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mauo.mutation.LogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirroractivity.LogsTable,
			Columns: []string{mirroractivity.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logmsg.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mauo.mutation.RemovedLogsIDs(); len(nodes) > 0 && !mauo.mutation.LogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirroractivity.LogsTable,
			Columns: []string{mirroractivity.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logmsg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mauo.mutation.LogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirroractivity.LogsTable,
			Columns: []string{mirroractivity.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logmsg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mauo.modifiers...)
	_node = &MirrorActivity{config: mauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mirroractivity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
