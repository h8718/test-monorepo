// Code generated by ent, DO NOT EDIT.

package namespace

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldUpdatedAt, v))
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldConfig, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldUpdatedAt, v))
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldConfig, v))
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldConfig, v))
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...string) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldConfig, vs...))
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...string) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldConfig, vs...))
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldConfig, v))
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldConfig, v))
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldConfig, v))
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldConfig, v))
}

// ConfigContains applies the Contains predicate on the "config" field.
func ConfigContains(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldContains(FieldConfig, v))
}

// ConfigHasPrefix applies the HasPrefix predicate on the "config" field.
func ConfigHasPrefix(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldHasPrefix(FieldConfig, v))
}

// ConfigHasSuffix applies the HasSuffix predicate on the "config" field.
func ConfigHasSuffix(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldHasSuffix(FieldConfig, v))
}

// ConfigEqualFold applies the EqualFold predicate on the "config" field.
func ConfigEqualFold(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEqualFold(FieldConfig, v))
}

// ConfigContainsFold applies the ContainsFold predicate on the "config" field.
func ConfigContainsFold(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldContainsFold(FieldConfig, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldContainsFold(FieldName, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
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
func Not(p predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		p(s.Not())
	})
}
