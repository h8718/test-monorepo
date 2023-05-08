// Code generated by ent, DO NOT EDIT.

package vardata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.VarData {
	return predicate.VarData(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldUpdatedAt, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldSize, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldHash, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v []byte) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldData, v))
}

// MimeType applies equality check predicate on the "mime_type" field. It's identical to MimeTypeEQ.
func MimeType(v string) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldMimeType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VarData {
	return predicate.VarData(sql.FieldLTE(FieldUpdatedAt, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.VarData {
	return predicate.VarData(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.VarData {
	return predicate.VarData(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.VarData {
	return predicate.VarData(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.VarData {
	return predicate.VarData(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.VarData {
	return predicate.VarData(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.VarData {
	return predicate.VarData(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.VarData {
	return predicate.VarData(sql.FieldLTE(FieldSize, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.VarData {
	return predicate.VarData(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.VarData {
	return predicate.VarData(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.VarData {
	return predicate.VarData(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.VarData {
	return predicate.VarData(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.VarData {
	return predicate.VarData(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.VarData {
	return predicate.VarData(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.VarData {
	return predicate.VarData(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.VarData {
	return predicate.VarData(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.VarData {
	return predicate.VarData(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.VarData {
	return predicate.VarData(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.VarData {
	return predicate.VarData(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.VarData {
	return predicate.VarData(sql.FieldContainsFold(FieldHash, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v []byte) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v []byte) predicate.VarData {
	return predicate.VarData(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...[]byte) predicate.VarData {
	return predicate.VarData(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...[]byte) predicate.VarData {
	return predicate.VarData(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v []byte) predicate.VarData {
	return predicate.VarData(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v []byte) predicate.VarData {
	return predicate.VarData(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v []byte) predicate.VarData {
	return predicate.VarData(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v []byte) predicate.VarData {
	return predicate.VarData(sql.FieldLTE(FieldData, v))
}

// MimeTypeEQ applies the EQ predicate on the "mime_type" field.
func MimeTypeEQ(v string) predicate.VarData {
	return predicate.VarData(sql.FieldEQ(FieldMimeType, v))
}

// MimeTypeNEQ applies the NEQ predicate on the "mime_type" field.
func MimeTypeNEQ(v string) predicate.VarData {
	return predicate.VarData(sql.FieldNEQ(FieldMimeType, v))
}

// MimeTypeIn applies the In predicate on the "mime_type" field.
func MimeTypeIn(vs ...string) predicate.VarData {
	return predicate.VarData(sql.FieldIn(FieldMimeType, vs...))
}

// MimeTypeNotIn applies the NotIn predicate on the "mime_type" field.
func MimeTypeNotIn(vs ...string) predicate.VarData {
	return predicate.VarData(sql.FieldNotIn(FieldMimeType, vs...))
}

// MimeTypeGT applies the GT predicate on the "mime_type" field.
func MimeTypeGT(v string) predicate.VarData {
	return predicate.VarData(sql.FieldGT(FieldMimeType, v))
}

// MimeTypeGTE applies the GTE predicate on the "mime_type" field.
func MimeTypeGTE(v string) predicate.VarData {
	return predicate.VarData(sql.FieldGTE(FieldMimeType, v))
}

// MimeTypeLT applies the LT predicate on the "mime_type" field.
func MimeTypeLT(v string) predicate.VarData {
	return predicate.VarData(sql.FieldLT(FieldMimeType, v))
}

// MimeTypeLTE applies the LTE predicate on the "mime_type" field.
func MimeTypeLTE(v string) predicate.VarData {
	return predicate.VarData(sql.FieldLTE(FieldMimeType, v))
}

// MimeTypeContains applies the Contains predicate on the "mime_type" field.
func MimeTypeContains(v string) predicate.VarData {
	return predicate.VarData(sql.FieldContains(FieldMimeType, v))
}

// MimeTypeHasPrefix applies the HasPrefix predicate on the "mime_type" field.
func MimeTypeHasPrefix(v string) predicate.VarData {
	return predicate.VarData(sql.FieldHasPrefix(FieldMimeType, v))
}

// MimeTypeHasSuffix applies the HasSuffix predicate on the "mime_type" field.
func MimeTypeHasSuffix(v string) predicate.VarData {
	return predicate.VarData(sql.FieldHasSuffix(FieldMimeType, v))
}

// MimeTypeEqualFold applies the EqualFold predicate on the "mime_type" field.
func MimeTypeEqualFold(v string) predicate.VarData {
	return predicate.VarData(sql.FieldEqualFold(FieldMimeType, v))
}

// MimeTypeContainsFold applies the ContainsFold predicate on the "mime_type" field.
func MimeTypeContainsFold(v string) predicate.VarData {
	return predicate.VarData(sql.FieldContainsFold(FieldMimeType, v))
}

// HasVarrefs applies the HasEdge predicate on the "varrefs" edge.
func HasVarrefs() predicate.VarData {
	return predicate.VarData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VarrefsTable, VarrefsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVarrefsWith applies the HasEdge predicate on the "varrefs" edge with a given conditions (other predicates).
func HasVarrefsWith(preds ...predicate.VarRef) predicate.VarData {
	return predicate.VarData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VarrefsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VarrefsTable, VarrefsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VarData) predicate.VarData {
	return predicate.VarData(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VarData) predicate.VarData {
	return predicate.VarData(func(s *sql.Selector) {
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
func Not(p predicate.VarData) predicate.VarData {
	return predicate.VarData(func(s *sql.Selector) {
		p(s.Not())
	})
}
