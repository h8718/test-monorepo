// Code generated by ent, DO NOT EDIT.

package events

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Signature applies equality check predicate on the "signature" field. It's identical to SignatureEQ.
func Signature(v []byte) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSignature), v))
	})
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// WorkflowID applies equality check predicate on the "workflow_id" field. It's identical to WorkflowIDEQ.
func WorkflowID(v uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkflowID), v))
	})
}

// SignatureEQ applies the EQ predicate on the "signature" field.
func SignatureEQ(v []byte) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSignature), v))
	})
}

// SignatureNEQ applies the NEQ predicate on the "signature" field.
func SignatureNEQ(v []byte) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSignature), v))
	})
}

// SignatureIn applies the In predicate on the "signature" field.
func SignatureIn(vs ...[]byte) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSignature), v...))
	})
}

// SignatureNotIn applies the NotIn predicate on the "signature" field.
func SignatureNotIn(vs ...[]byte) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSignature), v...))
	})
}

// SignatureGT applies the GT predicate on the "signature" field.
func SignatureGT(v []byte) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSignature), v))
	})
}

// SignatureGTE applies the GTE predicate on the "signature" field.
func SignatureGTE(v []byte) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSignature), v))
	})
}

// SignatureLT applies the LT predicate on the "signature" field.
func SignatureLT(v []byte) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSignature), v))
	})
}

// SignatureLTE applies the LTE predicate on the "signature" field.
func SignatureLTE(v []byte) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSignature), v))
	})
}

// SignatureIsNil applies the IsNil predicate on the "signature" field.
func SignatureIsNil() predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSignature)))
	})
}

// SignatureNotNil applies the NotNil predicate on the "signature" field.
func SignatureNotNil() predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSignature)))
	})
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCount), v))
	})
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCount), v...))
	})
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCount), v...))
	})
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCount), v))
	})
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCount), v))
	})
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCount), v))
	})
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// WorkflowIDEQ applies the EQ predicate on the "workflow_id" field.
func WorkflowIDEQ(v uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDNEQ applies the NEQ predicate on the "workflow_id" field.
func WorkflowIDNEQ(v uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDIn applies the In predicate on the "workflow_id" field.
func WorkflowIDIn(vs ...uuid.UUID) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWorkflowID), v...))
	})
}

// WorkflowIDNotIn applies the NotIn predicate on the "workflow_id" field.
func WorkflowIDNotIn(vs ...uuid.UUID) predicate.Events {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWorkflowID), v...))
	})
}

// WorkflowIDGT applies the GT predicate on the "workflow_id" field.
func WorkflowIDGT(v uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDGTE applies the GTE predicate on the "workflow_id" field.
func WorkflowIDGTE(v uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDLT applies the LT predicate on the "workflow_id" field.
func WorkflowIDLT(v uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDLTE applies the LTE predicate on the "workflow_id" field.
func WorkflowIDLTE(v uuid.UUID) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkflowID), v))
	})
}

// WorkflowIDIsNil applies the IsNil predicate on the "workflow_id" field.
func WorkflowIDIsNil() predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWorkflowID)))
	})
}

// WorkflowIDNotNil applies the NotNil predicate on the "workflow_id" field.
func WorkflowIDNotNil() predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWorkflowID)))
	})
}

// HasWfeventswait applies the HasEdge predicate on the "wfeventswait" edge.
func HasWfeventswait() predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WfeventswaitTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WfeventswaitTable, WfeventswaitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWfeventswaitWith applies the HasEdge predicate on the "wfeventswait" edge with a given conditions (other predicates).
func HasWfeventswaitWith(preds ...predicate.EventsWait) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WfeventswaitInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WfeventswaitTable, WfeventswaitColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstance applies the HasEdge predicate on the "instance" edge.
func HasInstance() predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstanceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstanceWith applies the HasEdge predicate on the "instance" edge with a given conditions (other predicates).
func HasInstanceWith(preds ...predicate.Instance) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstanceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Events) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Events) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Events) predicate.Events {
	return predicate.Events(func(s *sql.Selector) {
		p(s.Not())
	})
}
