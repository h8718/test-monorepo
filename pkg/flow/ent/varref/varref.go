// Code generated by entc, DO NOT EDIT.

package varref

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the varref type in the database.
	Label = "var_ref"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBehaviour holds the string denoting the behaviour field in the database.
	FieldBehaviour = "behaviour"
	// EdgeVardata holds the string denoting the vardata edge name in mutations.
	EdgeVardata = "vardata"
	// EdgeNamespace holds the string denoting the namespace edge name in mutations.
	EdgeNamespace = "namespace"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// EdgeInstance holds the string denoting the instance edge name in mutations.
	EdgeInstance = "instance"
	// Table holds the table name of the varref in the database.
	Table = "var_refs"
	// VardataTable is the table that holds the vardata relation/edge.
	VardataTable = "var_refs"
	// VardataInverseTable is the table name for the VarData entity.
	// It exists in this package in order to avoid circular dependency with the "vardata" package.
	VardataInverseTable = "var_data"
	// VardataColumn is the table column denoting the vardata relation/edge.
	VardataColumn = "var_data_varrefs"
	// NamespaceTable is the table that holds the namespace relation/edge.
	NamespaceTable = "var_refs"
	// NamespaceInverseTable is the table name for the Namespace entity.
	// It exists in this package in order to avoid circular dependency with the "namespace" package.
	NamespaceInverseTable = "namespaces"
	// NamespaceColumn is the table column denoting the namespace relation/edge.
	NamespaceColumn = "namespace_vars"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "var_refs"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "workflow_vars"
	// InstanceTable is the table that holds the instance relation/edge.
	InstanceTable = "var_refs"
	// InstanceInverseTable is the table name for the Instance entity.
	// It exists in this package in order to avoid circular dependency with the "instance" package.
	InstanceInverseTable = "instances"
	// InstanceColumn is the table column denoting the instance relation/edge.
	InstanceColumn = "instance_vars"
)

// Columns holds all SQL columns for varref fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldBehaviour,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "var_refs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"instance_vars",
	"namespace_vars",
	"var_data_varrefs",
	"workflow_vars",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
