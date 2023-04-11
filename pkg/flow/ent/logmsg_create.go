// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// LogMsgCreate is the builder for creating a LogMsg entity.
type LogMsgCreate struct {
	config
	mutation *LogMsgMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetT sets the "t" field.
func (lmc *LogMsgCreate) SetT(t time.Time) *LogMsgCreate {
	lmc.mutation.SetT(t)
	return lmc
}

// SetMsg sets the "msg" field.
func (lmc *LogMsgCreate) SetMsg(s string) *LogMsgCreate {
	lmc.mutation.SetMsg(s)
	return lmc
}

// SetLevel sets the "level" field.
func (lmc *LogMsgCreate) SetLevel(s string) *LogMsgCreate {
	lmc.mutation.SetLevel(s)
	return lmc
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableLevel(s *string) *LogMsgCreate {
	if s != nil {
		lmc.SetLevel(*s)
	}
	return lmc
}

// SetRootInstanceId sets the "rootInstanceId" field.
func (lmc *LogMsgCreate) SetRootInstanceId(s string) *LogMsgCreate {
	lmc.mutation.SetRootInstanceId(s)
	return lmc
}

// SetNillableRootInstanceId sets the "rootInstanceId" field if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableRootInstanceId(s *string) *LogMsgCreate {
	if s != nil {
		lmc.SetRootInstanceId(*s)
	}
	return lmc
}

// SetLogInstanceCallPath sets the "logInstanceCallPath" field.
func (lmc *LogMsgCreate) SetLogInstanceCallPath(s string) *LogMsgCreate {
	lmc.mutation.SetLogInstanceCallPath(s)
	return lmc
}

// SetNillableLogInstanceCallPath sets the "logInstanceCallPath" field if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableLogInstanceCallPath(s *string) *LogMsgCreate {
	if s != nil {
		lmc.SetLogInstanceCallPath(*s)
	}
	return lmc
}

// SetTags sets the "tags" field.
func (lmc *LogMsgCreate) SetTags(m map[string]string) *LogMsgCreate {
	lmc.mutation.SetTags(m)
	return lmc
}

// SetWorkflowID sets the "workflow_id" field.
func (lmc *LogMsgCreate) SetWorkflowID(u uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetWorkflowID(u)
	return lmc
}

// SetNillableWorkflowID sets the "workflow_id" field if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableWorkflowID(u *uuid.UUID) *LogMsgCreate {
	if u != nil {
		lmc.SetWorkflowID(*u)
	}
	return lmc
}

// SetMirrorActivityID sets the "mirror_activity_id" field.
func (lmc *LogMsgCreate) SetMirrorActivityID(u uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetMirrorActivityID(u)
	return lmc
}

// SetNillableMirrorActivityID sets the "mirror_activity_id" field if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableMirrorActivityID(u *uuid.UUID) *LogMsgCreate {
	if u != nil {
		lmc.SetMirrorActivityID(*u)
	}
	return lmc
}

// SetID sets the "id" field.
func (lmc *LogMsgCreate) SetID(u uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetID(u)
	return lmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableID(u *uuid.UUID) *LogMsgCreate {
	if u != nil {
		lmc.SetID(*u)
	}
	return lmc
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (lmc *LogMsgCreate) SetNamespaceID(id uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetNamespaceID(id)
	return lmc
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableNamespaceID(id *uuid.UUID) *LogMsgCreate {
	if id != nil {
		lmc = lmc.SetNamespaceID(*id)
	}
	return lmc
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (lmc *LogMsgCreate) SetNamespace(n *Namespace) *LogMsgCreate {
	return lmc.SetNamespaceID(n.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (lmc *LogMsgCreate) SetInstanceID(id uuid.UUID) *LogMsgCreate {
	lmc.mutation.SetInstanceID(id)
	return lmc
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (lmc *LogMsgCreate) SetNillableInstanceID(id *uuid.UUID) *LogMsgCreate {
	if id != nil {
		lmc = lmc.SetInstanceID(*id)
	}
	return lmc
}

// SetInstance sets the "instance" edge to the Instance entity.
func (lmc *LogMsgCreate) SetInstance(i *Instance) *LogMsgCreate {
	return lmc.SetInstanceID(i.ID)
}

// Mutation returns the LogMsgMutation object of the builder.
func (lmc *LogMsgCreate) Mutation() *LogMsgMutation {
	return lmc.mutation
}

// Save creates the LogMsg in the database.
func (lmc *LogMsgCreate) Save(ctx context.Context) (*LogMsg, error) {
	var (
		err  error
		node *LogMsg
	)
	lmc.defaults()
	if len(lmc.hooks) == 0 {
		if err = lmc.check(); err != nil {
			return nil, err
		}
		node, err = lmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LogMsgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmc.check(); err != nil {
				return nil, err
			}
			lmc.mutation = mutation
			if node, err = lmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lmc.hooks) - 1; i >= 0; i-- {
			if lmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lmc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LogMsg)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LogMsgMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lmc *LogMsgCreate) SaveX(ctx context.Context) *LogMsg {
	v, err := lmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lmc *LogMsgCreate) Exec(ctx context.Context) error {
	_, err := lmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmc *LogMsgCreate) ExecX(ctx context.Context) {
	if err := lmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lmc *LogMsgCreate) defaults() {
	if _, ok := lmc.mutation.Level(); !ok {
		v := logmsg.DefaultLevel
		lmc.mutation.SetLevel(v)
	}
	if _, ok := lmc.mutation.RootInstanceId(); !ok {
		v := logmsg.DefaultRootInstanceId
		lmc.mutation.SetRootInstanceId(v)
	}
	if _, ok := lmc.mutation.LogInstanceCallPath(); !ok {
		v := logmsg.DefaultLogInstanceCallPath
		lmc.mutation.SetLogInstanceCallPath(v)
	}
	if _, ok := lmc.mutation.ID(); !ok {
		v := logmsg.DefaultID()
		lmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmc *LogMsgCreate) check() error {
	if _, ok := lmc.mutation.T(); !ok {
		return &ValidationError{Name: "t", err: errors.New(`ent: missing required field "LogMsg.t"`)}
	}
	if _, ok := lmc.mutation.Msg(); !ok {
		return &ValidationError{Name: "msg", err: errors.New(`ent: missing required field "LogMsg.msg"`)}
	}
	if _, ok := lmc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "LogMsg.level"`)}
	}
	if _, ok := lmc.mutation.RootInstanceId(); !ok {
		return &ValidationError{Name: "rootInstanceId", err: errors.New(`ent: missing required field "LogMsg.rootInstanceId"`)}
	}
	if _, ok := lmc.mutation.LogInstanceCallPath(); !ok {
		return &ValidationError{Name: "logInstanceCallPath", err: errors.New(`ent: missing required field "LogMsg.logInstanceCallPath"`)}
	}
	return nil
}

func (lmc *LogMsgCreate) sqlSave(ctx context.Context) (*LogMsg, error) {
	_node, _spec := lmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (lmc *LogMsgCreate) createSpec() (*LogMsg, *sqlgraph.CreateSpec) {
	var (
		_node = &LogMsg{config: lmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: logmsg.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: logmsg.FieldID,
			},
		}
	)
	_spec.OnConflict = lmc.conflict
	if id, ok := lmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lmc.mutation.T(); ok {
		_spec.SetField(logmsg.FieldT, field.TypeTime, value)
		_node.T = value
	}
	if value, ok := lmc.mutation.Msg(); ok {
		_spec.SetField(logmsg.FieldMsg, field.TypeString, value)
		_node.Msg = value
	}
	if value, ok := lmc.mutation.Level(); ok {
		_spec.SetField(logmsg.FieldLevel, field.TypeString, value)
		_node.Level = value
	}
	if value, ok := lmc.mutation.RootInstanceId(); ok {
		_spec.SetField(logmsg.FieldRootInstanceId, field.TypeString, value)
		_node.RootInstanceId = value
	}
	if value, ok := lmc.mutation.LogInstanceCallPath(); ok {
		_spec.SetField(logmsg.FieldLogInstanceCallPath, field.TypeString, value)
		_node.LogInstanceCallPath = value
	}
	if value, ok := lmc.mutation.Tags(); ok {
		_spec.SetField(logmsg.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := lmc.mutation.WorkflowID(); ok {
		_spec.SetField(logmsg.FieldWorkflowID, field.TypeUUID, value)
		_node.WorkflowID = value
	}
	if value, ok := lmc.mutation.MirrorActivityID(); ok {
		_spec.SetField(logmsg.FieldMirrorActivityID, field.TypeUUID, value)
		_node.MirrorActivityID = value
	}
	if nodes := lmc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.NamespaceTable,
			Columns: []string{logmsg.NamespaceColumn},
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
		_node.namespace_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lmc.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logmsg.InstanceTable,
			Columns: []string{logmsg.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.instance_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LogMsg.Create().
//		SetT(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LogMsgUpsert) {
//			SetT(v+v).
//		}).
//		Exec(ctx)
func (lmc *LogMsgCreate) OnConflict(opts ...sql.ConflictOption) *LogMsgUpsertOne {
	lmc.conflict = opts
	return &LogMsgUpsertOne{
		create: lmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lmc *LogMsgCreate) OnConflictColumns(columns ...string) *LogMsgUpsertOne {
	lmc.conflict = append(lmc.conflict, sql.ConflictColumns(columns...))
	return &LogMsgUpsertOne{
		create: lmc,
	}
}

type (
	// LogMsgUpsertOne is the builder for "upsert"-ing
	//  one LogMsg node.
	LogMsgUpsertOne struct {
		create *LogMsgCreate
	}

	// LogMsgUpsert is the "OnConflict" setter.
	LogMsgUpsert struct {
		*sql.UpdateSet
	}
)

// SetT sets the "t" field.
func (u *LogMsgUpsert) SetT(v time.Time) *LogMsgUpsert {
	u.Set(logmsg.FieldT, v)
	return u
}

// UpdateT sets the "t" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateT() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldT)
	return u
}

// SetMsg sets the "msg" field.
func (u *LogMsgUpsert) SetMsg(v string) *LogMsgUpsert {
	u.Set(logmsg.FieldMsg, v)
	return u
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateMsg() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldMsg)
	return u
}

// SetLevel sets the "level" field.
func (u *LogMsgUpsert) SetLevel(v string) *LogMsgUpsert {
	u.Set(logmsg.FieldLevel, v)
	return u
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateLevel() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldLevel)
	return u
}

// SetRootInstanceId sets the "rootInstanceId" field.
func (u *LogMsgUpsert) SetRootInstanceId(v string) *LogMsgUpsert {
	u.Set(logmsg.FieldRootInstanceId, v)
	return u
}

// UpdateRootInstanceId sets the "rootInstanceId" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateRootInstanceId() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldRootInstanceId)
	return u
}

// SetLogInstanceCallPath sets the "logInstanceCallPath" field.
func (u *LogMsgUpsert) SetLogInstanceCallPath(v string) *LogMsgUpsert {
	u.Set(logmsg.FieldLogInstanceCallPath, v)
	return u
}

// UpdateLogInstanceCallPath sets the "logInstanceCallPath" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateLogInstanceCallPath() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldLogInstanceCallPath)
	return u
}

// SetTags sets the "tags" field.
func (u *LogMsgUpsert) SetTags(v map[string]string) *LogMsgUpsert {
	u.Set(logmsg.FieldTags, v)
	return u
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateTags() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldTags)
	return u
}

// ClearTags clears the value of the "tags" field.
func (u *LogMsgUpsert) ClearTags() *LogMsgUpsert {
	u.SetNull(logmsg.FieldTags)
	return u
}

// SetWorkflowID sets the "workflow_id" field.
func (u *LogMsgUpsert) SetWorkflowID(v uuid.UUID) *LogMsgUpsert {
	u.Set(logmsg.FieldWorkflowID, v)
	return u
}

// UpdateWorkflowID sets the "workflow_id" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateWorkflowID() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldWorkflowID)
	return u
}

// ClearWorkflowID clears the value of the "workflow_id" field.
func (u *LogMsgUpsert) ClearWorkflowID() *LogMsgUpsert {
	u.SetNull(logmsg.FieldWorkflowID)
	return u
}

// SetMirrorActivityID sets the "mirror_activity_id" field.
func (u *LogMsgUpsert) SetMirrorActivityID(v uuid.UUID) *LogMsgUpsert {
	u.Set(logmsg.FieldMirrorActivityID, v)
	return u
}

// UpdateMirrorActivityID sets the "mirror_activity_id" field to the value that was provided on create.
func (u *LogMsgUpsert) UpdateMirrorActivityID() *LogMsgUpsert {
	u.SetExcluded(logmsg.FieldMirrorActivityID)
	return u
}

// ClearMirrorActivityID clears the value of the "mirror_activity_id" field.
func (u *LogMsgUpsert) ClearMirrorActivityID() *LogMsgUpsert {
	u.SetNull(logmsg.FieldMirrorActivityID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(logmsg.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LogMsgUpsertOne) UpdateNewValues() *LogMsgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(logmsg.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LogMsgUpsertOne) Ignore() *LogMsgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LogMsgUpsertOne) DoNothing() *LogMsgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LogMsgCreate.OnConflict
// documentation for more info.
func (u *LogMsgUpsertOne) Update(set func(*LogMsgUpsert)) *LogMsgUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LogMsgUpsert{UpdateSet: update})
	}))
	return u
}

// SetT sets the "t" field.
func (u *LogMsgUpsertOne) SetT(v time.Time) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetT(v)
	})
}

// UpdateT sets the "t" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateT() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateT()
	})
}

// SetMsg sets the "msg" field.
func (u *LogMsgUpsertOne) SetMsg(v string) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetMsg(v)
	})
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateMsg() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateMsg()
	})
}

// SetLevel sets the "level" field.
func (u *LogMsgUpsertOne) SetLevel(v string) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetLevel(v)
	})
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateLevel() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateLevel()
	})
}

// SetRootInstanceId sets the "rootInstanceId" field.
func (u *LogMsgUpsertOne) SetRootInstanceId(v string) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetRootInstanceId(v)
	})
}

// UpdateRootInstanceId sets the "rootInstanceId" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateRootInstanceId() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateRootInstanceId()
	})
}

// SetLogInstanceCallPath sets the "logInstanceCallPath" field.
func (u *LogMsgUpsertOne) SetLogInstanceCallPath(v string) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetLogInstanceCallPath(v)
	})
}

// UpdateLogInstanceCallPath sets the "logInstanceCallPath" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateLogInstanceCallPath() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateLogInstanceCallPath()
	})
}

// SetTags sets the "tags" field.
func (u *LogMsgUpsertOne) SetTags(v map[string]string) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetTags(v)
	})
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateTags() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateTags()
	})
}

// ClearTags clears the value of the "tags" field.
func (u *LogMsgUpsertOne) ClearTags() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.ClearTags()
	})
}

// SetWorkflowID sets the "workflow_id" field.
func (u *LogMsgUpsertOne) SetWorkflowID(v uuid.UUID) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetWorkflowID(v)
	})
}

// UpdateWorkflowID sets the "workflow_id" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateWorkflowID() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateWorkflowID()
	})
}

// ClearWorkflowID clears the value of the "workflow_id" field.
func (u *LogMsgUpsertOne) ClearWorkflowID() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.ClearWorkflowID()
	})
}

// SetMirrorActivityID sets the "mirror_activity_id" field.
func (u *LogMsgUpsertOne) SetMirrorActivityID(v uuid.UUID) *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetMirrorActivityID(v)
	})
}

// UpdateMirrorActivityID sets the "mirror_activity_id" field to the value that was provided on create.
func (u *LogMsgUpsertOne) UpdateMirrorActivityID() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateMirrorActivityID()
	})
}

// ClearMirrorActivityID clears the value of the "mirror_activity_id" field.
func (u *LogMsgUpsertOne) ClearMirrorActivityID() *LogMsgUpsertOne {
	return u.Update(func(s *LogMsgUpsert) {
		s.ClearMirrorActivityID()
	})
}

// Exec executes the query.
func (u *LogMsgUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LogMsgCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LogMsgUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LogMsgUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: LogMsgUpsertOne.ID is not supported by MySQL driver. Use LogMsgUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LogMsgUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LogMsgCreateBulk is the builder for creating many LogMsg entities in bulk.
type LogMsgCreateBulk struct {
	config
	builders []*LogMsgCreate
	conflict []sql.ConflictOption
}

// Save creates the LogMsg entities in the database.
func (lmcb *LogMsgCreateBulk) Save(ctx context.Context) ([]*LogMsg, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lmcb.builders))
	nodes := make([]*LogMsg, len(lmcb.builders))
	mutators := make([]Mutator, len(lmcb.builders))
	for i := range lmcb.builders {
		func(i int, root context.Context) {
			builder := lmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LogMsgMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lmcb *LogMsgCreateBulk) SaveX(ctx context.Context) []*LogMsg {
	v, err := lmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lmcb *LogMsgCreateBulk) Exec(ctx context.Context) error {
	_, err := lmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmcb *LogMsgCreateBulk) ExecX(ctx context.Context) {
	if err := lmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LogMsg.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LogMsgUpsert) {
//			SetT(v+v).
//		}).
//		Exec(ctx)
func (lmcb *LogMsgCreateBulk) OnConflict(opts ...sql.ConflictOption) *LogMsgUpsertBulk {
	lmcb.conflict = opts
	return &LogMsgUpsertBulk{
		create: lmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lmcb *LogMsgCreateBulk) OnConflictColumns(columns ...string) *LogMsgUpsertBulk {
	lmcb.conflict = append(lmcb.conflict, sql.ConflictColumns(columns...))
	return &LogMsgUpsertBulk{
		create: lmcb,
	}
}

// LogMsgUpsertBulk is the builder for "upsert"-ing
// a bulk of LogMsg nodes.
type LogMsgUpsertBulk struct {
	create *LogMsgCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(logmsg.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LogMsgUpsertBulk) UpdateNewValues() *LogMsgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(logmsg.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LogMsg.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LogMsgUpsertBulk) Ignore() *LogMsgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LogMsgUpsertBulk) DoNothing() *LogMsgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LogMsgCreateBulk.OnConflict
// documentation for more info.
func (u *LogMsgUpsertBulk) Update(set func(*LogMsgUpsert)) *LogMsgUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LogMsgUpsert{UpdateSet: update})
	}))
	return u
}

// SetT sets the "t" field.
func (u *LogMsgUpsertBulk) SetT(v time.Time) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetT(v)
	})
}

// UpdateT sets the "t" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateT() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateT()
	})
}

// SetMsg sets the "msg" field.
func (u *LogMsgUpsertBulk) SetMsg(v string) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetMsg(v)
	})
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateMsg() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateMsg()
	})
}

// SetLevel sets the "level" field.
func (u *LogMsgUpsertBulk) SetLevel(v string) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetLevel(v)
	})
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateLevel() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateLevel()
	})
}

// SetRootInstanceId sets the "rootInstanceId" field.
func (u *LogMsgUpsertBulk) SetRootInstanceId(v string) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetRootInstanceId(v)
	})
}

// UpdateRootInstanceId sets the "rootInstanceId" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateRootInstanceId() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateRootInstanceId()
	})
}

// SetLogInstanceCallPath sets the "logInstanceCallPath" field.
func (u *LogMsgUpsertBulk) SetLogInstanceCallPath(v string) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetLogInstanceCallPath(v)
	})
}

// UpdateLogInstanceCallPath sets the "logInstanceCallPath" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateLogInstanceCallPath() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateLogInstanceCallPath()
	})
}

// SetTags sets the "tags" field.
func (u *LogMsgUpsertBulk) SetTags(v map[string]string) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetTags(v)
	})
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateTags() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateTags()
	})
}

// ClearTags clears the value of the "tags" field.
func (u *LogMsgUpsertBulk) ClearTags() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.ClearTags()
	})
}

// SetWorkflowID sets the "workflow_id" field.
func (u *LogMsgUpsertBulk) SetWorkflowID(v uuid.UUID) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetWorkflowID(v)
	})
}

// UpdateWorkflowID sets the "workflow_id" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateWorkflowID() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateWorkflowID()
	})
}

// ClearWorkflowID clears the value of the "workflow_id" field.
func (u *LogMsgUpsertBulk) ClearWorkflowID() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.ClearWorkflowID()
	})
}

// SetMirrorActivityID sets the "mirror_activity_id" field.
func (u *LogMsgUpsertBulk) SetMirrorActivityID(v uuid.UUID) *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.SetMirrorActivityID(v)
	})
}

// UpdateMirrorActivityID sets the "mirror_activity_id" field to the value that was provided on create.
func (u *LogMsgUpsertBulk) UpdateMirrorActivityID() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.UpdateMirrorActivityID()
	})
}

// ClearMirrorActivityID clears the value of the "mirror_activity_id" field.
func (u *LogMsgUpsertBulk) ClearMirrorActivityID() *LogMsgUpsertBulk {
	return u.Update(func(s *LogMsgUpsert) {
		s.ClearMirrorActivityID()
	})
}

// Exec executes the query.
func (u *LogMsgUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LogMsgCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LogMsgCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LogMsgUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
