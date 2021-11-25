// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instanceruntime"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
)

// InstanceRuntimeDelete is the builder for deleting a InstanceRuntime entity.
type InstanceRuntimeDelete struct {
	config
	hooks    []Hook
	mutation *InstanceRuntimeMutation
}

// Where appends a list predicates to the InstanceRuntimeDelete builder.
func (ird *InstanceRuntimeDelete) Where(ps ...predicate.InstanceRuntime) *InstanceRuntimeDelete {
	ird.mutation.Where(ps...)
	return ird
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ird *InstanceRuntimeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ird.hooks) == 0 {
		affected, err = ird.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceRuntimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ird.mutation = mutation
			affected, err = ird.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ird.hooks) - 1; i >= 0; i-- {
			if ird.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ird.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ird.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ird *InstanceRuntimeDelete) ExecX(ctx context.Context) int {
	n, err := ird.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ird *InstanceRuntimeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: instanceruntime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instanceruntime.FieldID,
			},
		},
	}
	if ps := ird.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ird.driver, _spec)
}

// InstanceRuntimeDeleteOne is the builder for deleting a single InstanceRuntime entity.
type InstanceRuntimeDeleteOne struct {
	ird *InstanceRuntimeDelete
}

// Exec executes the deletion query.
func (irdo *InstanceRuntimeDeleteOne) Exec(ctx context.Context) error {
	n, err := irdo.ird.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{instanceruntime.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (irdo *InstanceRuntimeDeleteOne) ExecX(ctx context.Context) {
	irdo.ird.ExecX(ctx)
}
