// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
)

// EventsWaitDelete is the builder for deleting a EventsWait entity.
type EventsWaitDelete struct {
	config
	hooks    []Hook
	mutation *EventsWaitMutation
}

// Where appends a list predicates to the EventsWaitDelete builder.
func (ewd *EventsWaitDelete) Where(ps ...predicate.EventsWait) *EventsWaitDelete {
	ewd.mutation.Where(ps...)
	return ewd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ewd *EventsWaitDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, EventsWaitMutation](ctx, ewd.sqlExec, ewd.mutation, ewd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ewd *EventsWaitDelete) ExecX(ctx context.Context) int {
	n, err := ewd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ewd *EventsWaitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(eventswait.Table, sqlgraph.NewFieldSpec(eventswait.FieldID, field.TypeUUID))
	if ps := ewd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ewd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ewd.mutation.done = true
	return affected, err
}

// EventsWaitDeleteOne is the builder for deleting a single EventsWait entity.
type EventsWaitDeleteOne struct {
	ewd *EventsWaitDelete
}

// Where appends a list predicates to the EventsWaitDelete builder.
func (ewdo *EventsWaitDeleteOne) Where(ps ...predicate.EventsWait) *EventsWaitDeleteOne {
	ewdo.ewd.mutation.Where(ps...)
	return ewdo
}

// Exec executes the deletion query.
func (ewdo *EventsWaitDeleteOne) Exec(ctx context.Context) error {
	n, err := ewdo.ewd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{eventswait.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ewdo *EventsWaitDeleteOne) ExecX(ctx context.Context) {
	if err := ewdo.Exec(ctx); err != nil {
		panic(err)
	}
}
