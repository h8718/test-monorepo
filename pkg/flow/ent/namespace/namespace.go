// Code generated by ent, DO NOT EDIT.

package namespace

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the namespace type in the database.
	Label = "namespace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeLogs holds the string denoting the logs edge name in mutations.
	EdgeLogs = "logs"
	// EdgeCloudevents holds the string denoting the cloudevents edge name in mutations.
	EdgeCloudevents = "cloudevents"
	// EdgeNamespacelisteners holds the string denoting the namespacelisteners edge name in mutations.
	EdgeNamespacelisteners = "namespacelisteners"
	// EdgeAnnotations holds the string denoting the annotations edge name in mutations.
	EdgeAnnotations = "annotations"
	// EdgeCloudeventfilters holds the string denoting the cloudeventfilters edge name in mutations.
	EdgeCloudeventfilters = "cloudeventfilters"
	// EdgeServices holds the string denoting the services edge name in mutations.
	EdgeServices = "services"
	// CloudEventFiltersFieldID holds the string denoting the ID field of the CloudEventFilters.
	CloudEventFiltersFieldID = "id"
	// Table holds the table name of the namespace in the database.
	Table = "namespaces"
	// LogsTable is the table that holds the logs relation/edge.
	LogsTable = "log_msgs"
	// LogsInverseTable is the table name for the LogMsg entity.
	// It exists in this package in order to avoid circular dependency with the "logmsg" package.
	LogsInverseTable = "log_msgs"
	// LogsColumn is the table column denoting the logs relation/edge.
	LogsColumn = "namespace_logs"
	// CloudeventsTable is the table that holds the cloudevents relation/edge.
	CloudeventsTable = "cloud_events"
	// CloudeventsInverseTable is the table name for the CloudEvents entity.
	// It exists in this package in order to avoid circular dependency with the "cloudevents" package.
	CloudeventsInverseTable = "cloud_events"
	// CloudeventsColumn is the table column denoting the cloudevents relation/edge.
	CloudeventsColumn = "namespace_cloudevents"
	// NamespacelistenersTable is the table that holds the namespacelisteners relation/edge.
	NamespacelistenersTable = "events"
	// NamespacelistenersInverseTable is the table name for the Events entity.
	// It exists in this package in order to avoid circular dependency with the "events" package.
	NamespacelistenersInverseTable = "events"
	// NamespacelistenersColumn is the table column denoting the namespacelisteners relation/edge.
	NamespacelistenersColumn = "namespace_namespacelisteners"
	// AnnotationsTable is the table that holds the annotations relation/edge.
	AnnotationsTable = "annotations"
	// AnnotationsInverseTable is the table name for the Annotation entity.
	// It exists in this package in order to avoid circular dependency with the "annotation" package.
	AnnotationsInverseTable = "annotations"
	// AnnotationsColumn is the table column denoting the annotations relation/edge.
	AnnotationsColumn = "namespace_annotations"
	// CloudeventfiltersTable is the table that holds the cloudeventfilters relation/edge.
	CloudeventfiltersTable = "cloud_event_filters"
	// CloudeventfiltersInverseTable is the table name for the CloudEventFilters entity.
	// It exists in this package in order to avoid circular dependency with the "cloudeventfilters" package.
	CloudeventfiltersInverseTable = "cloud_event_filters"
	// CloudeventfiltersColumn is the table column denoting the cloudeventfilters relation/edge.
	CloudeventfiltersColumn = "namespace_cloudeventfilters"
	// ServicesTable is the table that holds the services relation/edge.
	ServicesTable = "services"
	// ServicesInverseTable is the table name for the Services entity.
	// It exists in this package in order to avoid circular dependency with the "services" package.
	ServicesInverseTable = "services"
	// ServicesColumn is the table column denoting the services relation/edge.
	ServicesColumn = "namespace_services"
)

// Columns holds all SQL columns for namespace fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldConfig,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultConfig holds the default value on creation for the "config" field.
	DefaultConfig string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
