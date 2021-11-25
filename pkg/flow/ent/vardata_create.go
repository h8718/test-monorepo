// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/vardata"
	"github.com/direktiv/direktiv/pkg/flow/ent/varref"
	"github.com/google/uuid"
)

// VarDataCreate is the builder for creating a VarData entity.
type VarDataCreate struct {
	config
	mutation *VarDataMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (vdc *VarDataCreate) SetCreatedAt(t time.Time) *VarDataCreate {
	vdc.mutation.SetCreatedAt(t)
	return vdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vdc *VarDataCreate) SetNillableCreatedAt(t *time.Time) *VarDataCreate {
	if t != nil {
		vdc.SetCreatedAt(*t)
	}
	return vdc
}

// SetUpdatedAt sets the "updated_at" field.
func (vdc *VarDataCreate) SetUpdatedAt(t time.Time) *VarDataCreate {
	vdc.mutation.SetUpdatedAt(t)
	return vdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vdc *VarDataCreate) SetNillableUpdatedAt(t *time.Time) *VarDataCreate {
	if t != nil {
		vdc.SetUpdatedAt(*t)
	}
	return vdc
}

// SetSize sets the "size" field.
func (vdc *VarDataCreate) SetSize(i int) *VarDataCreate {
	vdc.mutation.SetSize(i)
	return vdc
}

// SetHash sets the "hash" field.
func (vdc *VarDataCreate) SetHash(s string) *VarDataCreate {
	vdc.mutation.SetHash(s)
	return vdc
}

// SetData sets the "data" field.
func (vdc *VarDataCreate) SetData(b []byte) *VarDataCreate {
	vdc.mutation.SetData(b)
	return vdc
}

// SetMimeType sets the "mime_type" field.
func (vdc *VarDataCreate) SetMimeType(s string) *VarDataCreate {
	vdc.mutation.SetMimeType(s)
	return vdc
}

// SetNillableMimeType sets the "mime_type" field if the given value is not nil.
func (vdc *VarDataCreate) SetNillableMimeType(s *string) *VarDataCreate {
	if s != nil {
		vdc.SetMimeType(*s)
	}
	return vdc
}

// SetID sets the "id" field.
func (vdc *VarDataCreate) SetID(u uuid.UUID) *VarDataCreate {
	vdc.mutation.SetID(u)
	return vdc
}

// AddVarrefIDs adds the "varrefs" edge to the VarRef entity by IDs.
func (vdc *VarDataCreate) AddVarrefIDs(ids ...uuid.UUID) *VarDataCreate {
	vdc.mutation.AddVarrefIDs(ids...)
	return vdc
}

// AddVarrefs adds the "varrefs" edges to the VarRef entity.
func (vdc *VarDataCreate) AddVarrefs(v ...*VarRef) *VarDataCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vdc.AddVarrefIDs(ids...)
}

// Mutation returns the VarDataMutation object of the builder.
func (vdc *VarDataCreate) Mutation() *VarDataMutation {
	return vdc.mutation
}

// Save creates the VarData in the database.
func (vdc *VarDataCreate) Save(ctx context.Context) (*VarData, error) {
	var (
		err  error
		node *VarData
	)
	vdc.defaults()
	if len(vdc.hooks) == 0 {
		if err = vdc.check(); err != nil {
			return nil, err
		}
		node, err = vdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VarDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vdc.check(); err != nil {
				return nil, err
			}
			vdc.mutation = mutation
			if node, err = vdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vdc.hooks) - 1; i >= 0; i-- {
			if vdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vdc *VarDataCreate) SaveX(ctx context.Context) *VarData {
	v, err := vdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vdc *VarDataCreate) Exec(ctx context.Context) error {
	_, err := vdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vdc *VarDataCreate) ExecX(ctx context.Context) {
	if err := vdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vdc *VarDataCreate) defaults() {
	if _, ok := vdc.mutation.CreatedAt(); !ok {
		v := vardata.DefaultCreatedAt()
		vdc.mutation.SetCreatedAt(v)
	}
	if _, ok := vdc.mutation.UpdatedAt(); !ok {
		v := vardata.DefaultUpdatedAt()
		vdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vdc.mutation.MimeType(); !ok {
		v := vardata.DefaultMimeType
		vdc.mutation.SetMimeType(v)
	}
	if _, ok := vdc.mutation.ID(); !ok {
		v := vardata.DefaultID()
		vdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vdc *VarDataCreate) check() error {
	if _, ok := vdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := vdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := vdc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "size"`)}
	}
	if _, ok := vdc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "hash"`)}
	}
	if _, ok := vdc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "data"`)}
	}
	if _, ok := vdc.mutation.MimeType(); !ok {
		return &ValidationError{Name: "mime_type", err: errors.New(`ent: missing required field "mime_type"`)}
	}
	return nil
}

func (vdc *VarDataCreate) sqlSave(ctx context.Context) (*VarData, error) {
	_node, _spec := vdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (vdc *VarDataCreate) createSpec() (*VarData, *sqlgraph.CreateSpec) {
	var (
		_node = &VarData{config: vdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vardata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vardata.FieldID,
			},
		}
	)
	if id, ok := vdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vdc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vardata.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vdc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vardata.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := vdc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vardata.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := vdc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vardata.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := vdc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: vardata.FieldData,
		})
		_node.Data = value
	}
	if value, ok := vdc.mutation.MimeType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vardata.FieldMimeType,
		})
		_node.MimeType = value
	}
	if nodes := vdc.mutation.VarrefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vardata.VarrefsTable,
			Columns: []string{vardata.VarrefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: varref.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VarDataCreateBulk is the builder for creating many VarData entities in bulk.
type VarDataCreateBulk struct {
	config
	builders []*VarDataCreate
}

// Save creates the VarData entities in the database.
func (vdcb *VarDataCreateBulk) Save(ctx context.Context) ([]*VarData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vdcb.builders))
	nodes := make([]*VarData, len(vdcb.builders))
	mutators := make([]Mutator, len(vdcb.builders))
	for i := range vdcb.builders {
		func(i int, root context.Context) {
			builder := vdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VarDataMutation)
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
					_, err = mutators[i+1].Mutate(root, vdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, vdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vdcb *VarDataCreateBulk) SaveX(ctx context.Context) []*VarData {
	v, err := vdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vdcb *VarDataCreateBulk) Exec(ctx context.Context) error {
	_, err := vdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vdcb *VarDataCreateBulk) ExecX(ctx context.Context) {
	if err := vdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
