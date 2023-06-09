// Code generated by ent, DO NOT EDIT.

package annotation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldUpdatedAt, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldSize, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldHash, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v []byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldData, v))
}

// MimeType applies equality check predicate on the "mime_type" field. It's identical to MimeTypeEQ.
func MimeType(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldMimeType, v))
}

// InstanceID applies equality check predicate on the "instance_id" field. It's identical to InstanceIDEQ.
func InstanceID(v uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldInstanceID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldContainsFold(FieldName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldUpdatedAt, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldSize, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldContainsFold(FieldHash, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v []byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v []byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...[]byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...[]byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v []byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v []byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v []byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v []byte) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldData, v))
}

// MimeTypeEQ applies the EQ predicate on the "mime_type" field.
func MimeTypeEQ(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldMimeType, v))
}

// MimeTypeNEQ applies the NEQ predicate on the "mime_type" field.
func MimeTypeNEQ(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldMimeType, v))
}

// MimeTypeIn applies the In predicate on the "mime_type" field.
func MimeTypeIn(vs ...string) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldMimeType, vs...))
}

// MimeTypeNotIn applies the NotIn predicate on the "mime_type" field.
func MimeTypeNotIn(vs ...string) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldMimeType, vs...))
}

// MimeTypeGT applies the GT predicate on the "mime_type" field.
func MimeTypeGT(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldMimeType, v))
}

// MimeTypeGTE applies the GTE predicate on the "mime_type" field.
func MimeTypeGTE(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldMimeType, v))
}

// MimeTypeLT applies the LT predicate on the "mime_type" field.
func MimeTypeLT(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldMimeType, v))
}

// MimeTypeLTE applies the LTE predicate on the "mime_type" field.
func MimeTypeLTE(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldMimeType, v))
}

// MimeTypeContains applies the Contains predicate on the "mime_type" field.
func MimeTypeContains(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldContains(FieldMimeType, v))
}

// MimeTypeHasPrefix applies the HasPrefix predicate on the "mime_type" field.
func MimeTypeHasPrefix(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldHasPrefix(FieldMimeType, v))
}

// MimeTypeHasSuffix applies the HasSuffix predicate on the "mime_type" field.
func MimeTypeHasSuffix(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldHasSuffix(FieldMimeType, v))
}

// MimeTypeEqualFold applies the EqualFold predicate on the "mime_type" field.
func MimeTypeEqualFold(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldEqualFold(FieldMimeType, v))
}

// MimeTypeContainsFold applies the ContainsFold predicate on the "mime_type" field.
func MimeTypeContainsFold(v string) predicate.Annotation {
	return predicate.Annotation(sql.FieldContainsFold(FieldMimeType, v))
}

// InstanceIDEQ applies the EQ predicate on the "instance_id" field.
func InstanceIDEQ(v uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldEQ(FieldInstanceID, v))
}

// InstanceIDNEQ applies the NEQ predicate on the "instance_id" field.
func InstanceIDNEQ(v uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNEQ(FieldInstanceID, v))
}

// InstanceIDIn applies the In predicate on the "instance_id" field.
func InstanceIDIn(vs ...uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldIn(FieldInstanceID, vs...))
}

// InstanceIDNotIn applies the NotIn predicate on the "instance_id" field.
func InstanceIDNotIn(vs ...uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldNotIn(FieldInstanceID, vs...))
}

// InstanceIDGT applies the GT predicate on the "instance_id" field.
func InstanceIDGT(v uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGT(FieldInstanceID, v))
}

// InstanceIDGTE applies the GTE predicate on the "instance_id" field.
func InstanceIDGTE(v uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldGTE(FieldInstanceID, v))
}

// InstanceIDLT applies the LT predicate on the "instance_id" field.
func InstanceIDLT(v uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLT(FieldInstanceID, v))
}

// InstanceIDLTE applies the LTE predicate on the "instance_id" field.
func InstanceIDLTE(v uuid.UUID) predicate.Annotation {
	return predicate.Annotation(sql.FieldLTE(FieldInstanceID, v))
}

// InstanceIDIsNil applies the IsNil predicate on the "instance_id" field.
func InstanceIDIsNil() predicate.Annotation {
	return predicate.Annotation(sql.FieldIsNull(FieldInstanceID))
}

// InstanceIDNotNil applies the NotNil predicate on the "instance_id" field.
func InstanceIDNotNil() predicate.Annotation {
	return predicate.Annotation(sql.FieldNotNull(FieldInstanceID))
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
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
func And(predicates ...predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
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
func Not(p predicate.Annotation) predicate.Annotation {
	return predicate.Annotation(func(s *sql.Selector) {
		p(s.Not())
	})
}
