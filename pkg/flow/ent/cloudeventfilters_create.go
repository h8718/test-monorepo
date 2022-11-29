// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudeventfilters"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// CloudEventFiltersCreate is the builder for creating a CloudEventFilters entity.
type CloudEventFiltersCreate struct {
	config
	mutation *CloudEventFiltersMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (cefc *CloudEventFiltersCreate) SetName(s string) *CloudEventFiltersCreate {
	cefc.mutation.SetName(s)
	return cefc
}

// SetJscode sets the "jscode" field.
func (cefc *CloudEventFiltersCreate) SetJscode(s string) *CloudEventFiltersCreate {
	cefc.mutation.SetJscode(s)
	return cefc
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (cefc *CloudEventFiltersCreate) SetNamespaceID(id uuid.UUID) *CloudEventFiltersCreate {
	cefc.mutation.SetNamespaceID(id)
	return cefc
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (cefc *CloudEventFiltersCreate) SetNamespace(n *Namespace) *CloudEventFiltersCreate {
	return cefc.SetNamespaceID(n.ID)
}

// Mutation returns the CloudEventFiltersMutation object of the builder.
func (cefc *CloudEventFiltersCreate) Mutation() *CloudEventFiltersMutation {
	return cefc.mutation
}

// Save creates the CloudEventFilters in the database.
func (cefc *CloudEventFiltersCreate) Save(ctx context.Context) (*CloudEventFilters, error) {
	var (
		err  error
		node *CloudEventFilters
	)
	if len(cefc.hooks) == 0 {
		if err = cefc.check(); err != nil {
			return nil, err
		}
		node, err = cefc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CloudEventFiltersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cefc.check(); err != nil {
				return nil, err
			}
			cefc.mutation = mutation
			if node, err = cefc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cefc.hooks) - 1; i >= 0; i-- {
			if cefc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cefc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cefc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CloudEventFilters)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CloudEventFiltersMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cefc *CloudEventFiltersCreate) SaveX(ctx context.Context) *CloudEventFilters {
	v, err := cefc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cefc *CloudEventFiltersCreate) Exec(ctx context.Context) error {
	_, err := cefc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cefc *CloudEventFiltersCreate) ExecX(ctx context.Context) {
	if err := cefc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cefc *CloudEventFiltersCreate) check() error {
	if _, ok := cefc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CloudEventFilters.name"`)}
	}
	if v, ok := cefc.mutation.Name(); ok {
		if err := cloudeventfilters.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CloudEventFilters.name": %w`, err)}
		}
	}
	if _, ok := cefc.mutation.Jscode(); !ok {
		return &ValidationError{Name: "jscode", err: errors.New(`ent: missing required field "CloudEventFilters.jscode"`)}
	}
	if v, ok := cefc.mutation.Jscode(); ok {
		if err := cloudeventfilters.JscodeValidator(v); err != nil {
			return &ValidationError{Name: "jscode", err: fmt.Errorf(`ent: validator failed for field "CloudEventFilters.jscode": %w`, err)}
		}
	}
	if _, ok := cefc.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required edge "CloudEventFilters.namespace"`)}
	}
	return nil
}

func (cefc *CloudEventFiltersCreate) sqlSave(ctx context.Context) (*CloudEventFilters, error) {
	_node, _spec := cefc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cefc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cefc *CloudEventFiltersCreate) createSpec() (*CloudEventFilters, *sqlgraph.CreateSpec) {
	var (
		_node = &CloudEventFilters{config: cefc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cloudeventfilters.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cloudeventfilters.FieldID,
			},
		}
	)
	_spec.OnConflict = cefc.conflict
	if value, ok := cefc.mutation.Name(); ok {
		_spec.SetField(cloudeventfilters.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cefc.mutation.Jscode(); ok {
		_spec.SetField(cloudeventfilters.FieldJscode, field.TypeString, value)
		_node.Jscode = value
	}
	if nodes := cefc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cloudeventfilters.NamespaceTable,
			Columns: []string{cloudeventfilters.NamespaceColumn},
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
		_node.namespace_cloudeventfilters = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CloudEventFilters.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CloudEventFiltersUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (cefc *CloudEventFiltersCreate) OnConflict(opts ...sql.ConflictOption) *CloudEventFiltersUpsertOne {
	cefc.conflict = opts
	return &CloudEventFiltersUpsertOne{
		create: cefc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CloudEventFilters.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cefc *CloudEventFiltersCreate) OnConflictColumns(columns ...string) *CloudEventFiltersUpsertOne {
	cefc.conflict = append(cefc.conflict, sql.ConflictColumns(columns...))
	return &CloudEventFiltersUpsertOne{
		create: cefc,
	}
}

type (
	// CloudEventFiltersUpsertOne is the builder for "upsert"-ing
	//  one CloudEventFilters node.
	CloudEventFiltersUpsertOne struct {
		create *CloudEventFiltersCreate
	}

	// CloudEventFiltersUpsert is the "OnConflict" setter.
	CloudEventFiltersUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *CloudEventFiltersUpsert) SetName(v string) *CloudEventFiltersUpsert {
	u.Set(cloudeventfilters.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CloudEventFiltersUpsert) UpdateName() *CloudEventFiltersUpsert {
	u.SetExcluded(cloudeventfilters.FieldName)
	return u
}

// SetJscode sets the "jscode" field.
func (u *CloudEventFiltersUpsert) SetJscode(v string) *CloudEventFiltersUpsert {
	u.Set(cloudeventfilters.FieldJscode, v)
	return u
}

// UpdateJscode sets the "jscode" field to the value that was provided on create.
func (u *CloudEventFiltersUpsert) UpdateJscode() *CloudEventFiltersUpsert {
	u.SetExcluded(cloudeventfilters.FieldJscode)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.CloudEventFilters.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CloudEventFiltersUpsertOne) UpdateNewValues() *CloudEventFiltersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CloudEventFilters.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CloudEventFiltersUpsertOne) Ignore() *CloudEventFiltersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CloudEventFiltersUpsertOne) DoNothing() *CloudEventFiltersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CloudEventFiltersCreate.OnConflict
// documentation for more info.
func (u *CloudEventFiltersUpsertOne) Update(set func(*CloudEventFiltersUpsert)) *CloudEventFiltersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CloudEventFiltersUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CloudEventFiltersUpsertOne) SetName(v string) *CloudEventFiltersUpsertOne {
	return u.Update(func(s *CloudEventFiltersUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CloudEventFiltersUpsertOne) UpdateName() *CloudEventFiltersUpsertOne {
	return u.Update(func(s *CloudEventFiltersUpsert) {
		s.UpdateName()
	})
}

// SetJscode sets the "jscode" field.
func (u *CloudEventFiltersUpsertOne) SetJscode(v string) *CloudEventFiltersUpsertOne {
	return u.Update(func(s *CloudEventFiltersUpsert) {
		s.SetJscode(v)
	})
}

// UpdateJscode sets the "jscode" field to the value that was provided on create.
func (u *CloudEventFiltersUpsertOne) UpdateJscode() *CloudEventFiltersUpsertOne {
	return u.Update(func(s *CloudEventFiltersUpsert) {
		s.UpdateJscode()
	})
}

// Exec executes the query.
func (u *CloudEventFiltersUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CloudEventFiltersCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CloudEventFiltersUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CloudEventFiltersUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CloudEventFiltersUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CloudEventFiltersCreateBulk is the builder for creating many CloudEventFilters entities in bulk.
type CloudEventFiltersCreateBulk struct {
	config
	builders []*CloudEventFiltersCreate
	conflict []sql.ConflictOption
}

// Save creates the CloudEventFilters entities in the database.
func (cefcb *CloudEventFiltersCreateBulk) Save(ctx context.Context) ([]*CloudEventFilters, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cefcb.builders))
	nodes := make([]*CloudEventFilters, len(cefcb.builders))
	mutators := make([]Mutator, len(cefcb.builders))
	for i := range cefcb.builders {
		func(i int, root context.Context) {
			builder := cefcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CloudEventFiltersMutation)
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
					_, err = mutators[i+1].Mutate(root, cefcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cefcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cefcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, cefcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cefcb *CloudEventFiltersCreateBulk) SaveX(ctx context.Context) []*CloudEventFilters {
	v, err := cefcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cefcb *CloudEventFiltersCreateBulk) Exec(ctx context.Context) error {
	_, err := cefcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cefcb *CloudEventFiltersCreateBulk) ExecX(ctx context.Context) {
	if err := cefcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CloudEventFilters.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CloudEventFiltersUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (cefcb *CloudEventFiltersCreateBulk) OnConflict(opts ...sql.ConflictOption) *CloudEventFiltersUpsertBulk {
	cefcb.conflict = opts
	return &CloudEventFiltersUpsertBulk{
		create: cefcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CloudEventFilters.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cefcb *CloudEventFiltersCreateBulk) OnConflictColumns(columns ...string) *CloudEventFiltersUpsertBulk {
	cefcb.conflict = append(cefcb.conflict, sql.ConflictColumns(columns...))
	return &CloudEventFiltersUpsertBulk{
		create: cefcb,
	}
}

// CloudEventFiltersUpsertBulk is the builder for "upsert"-ing
// a bulk of CloudEventFilters nodes.
type CloudEventFiltersUpsertBulk struct {
	create *CloudEventFiltersCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CloudEventFilters.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CloudEventFiltersUpsertBulk) UpdateNewValues() *CloudEventFiltersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CloudEventFilters.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CloudEventFiltersUpsertBulk) Ignore() *CloudEventFiltersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CloudEventFiltersUpsertBulk) DoNothing() *CloudEventFiltersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CloudEventFiltersCreateBulk.OnConflict
// documentation for more info.
func (u *CloudEventFiltersUpsertBulk) Update(set func(*CloudEventFiltersUpsert)) *CloudEventFiltersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CloudEventFiltersUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CloudEventFiltersUpsertBulk) SetName(v string) *CloudEventFiltersUpsertBulk {
	return u.Update(func(s *CloudEventFiltersUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CloudEventFiltersUpsertBulk) UpdateName() *CloudEventFiltersUpsertBulk {
	return u.Update(func(s *CloudEventFiltersUpsert) {
		s.UpdateName()
	})
}

// SetJscode sets the "jscode" field.
func (u *CloudEventFiltersUpsertBulk) SetJscode(v string) *CloudEventFiltersUpsertBulk {
	return u.Update(func(s *CloudEventFiltersUpsert) {
		s.SetJscode(v)
	})
}

// UpdateJscode sets the "jscode" field to the value that was provided on create.
func (u *CloudEventFiltersUpsertBulk) UpdateJscode() *CloudEventFiltersUpsertBulk {
	return u.Update(func(s *CloudEventFiltersUpsert) {
		s.UpdateJscode()
	})
}

// Exec executes the query.
func (u *CloudEventFiltersUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CloudEventFiltersCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CloudEventFiltersCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CloudEventFiltersUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}