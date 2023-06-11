// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// Namespace is the model entity for the Namespace schema.
type Namespace struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Config holds the value of the "config" field.
	Config string `json:"config,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NamespaceQuery when eager-loading is set.
	Edges NamespaceEdges `json:"edges"`
}

// NamespaceEdges holds the relations/edges for other nodes in the graph.
type NamespaceEdges struct {
	// Logs holds the value of the logs edge.
	Logs []*LogMsg `json:"logs,omitempty"`
	// Cloudevents holds the value of the cloudevents edge.
	Cloudevents []*CloudEvents `json:"cloudevents,omitempty"`
	// Namespacelisteners holds the value of the namespacelisteners edge.
	Namespacelisteners []*Events `json:"namespacelisteners,omitempty"`
	// Annotations holds the value of the annotations edge.
	Annotations []*Annotation `json:"annotations,omitempty"`
	// Cloudeventfilters holds the value of the cloudeventfilters edge.
	Cloudeventfilters []*CloudEventFilters `json:"cloudeventfilters,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// LogsOrErr returns the Logs value or an error if the edge
// was not loaded in eager-loading.
func (e NamespaceEdges) LogsOrErr() ([]*LogMsg, error) {
	if e.loadedTypes[0] {
		return e.Logs, nil
	}
	return nil, &NotLoadedError{edge: "logs"}
}

// CloudeventsOrErr returns the Cloudevents value or an error if the edge
// was not loaded in eager-loading.
func (e NamespaceEdges) CloudeventsOrErr() ([]*CloudEvents, error) {
	if e.loadedTypes[1] {
		return e.Cloudevents, nil
	}
	return nil, &NotLoadedError{edge: "cloudevents"}
}

// NamespacelistenersOrErr returns the Namespacelisteners value or an error if the edge
// was not loaded in eager-loading.
func (e NamespaceEdges) NamespacelistenersOrErr() ([]*Events, error) {
	if e.loadedTypes[2] {
		return e.Namespacelisteners, nil
	}
	return nil, &NotLoadedError{edge: "namespacelisteners"}
}

// AnnotationsOrErr returns the Annotations value or an error if the edge
// was not loaded in eager-loading.
func (e NamespaceEdges) AnnotationsOrErr() ([]*Annotation, error) {
	if e.loadedTypes[3] {
		return e.Annotations, nil
	}
	return nil, &NotLoadedError{edge: "annotations"}
}

// CloudeventfiltersOrErr returns the Cloudeventfilters value or an error if the edge
// was not loaded in eager-loading.
func (e NamespaceEdges) CloudeventfiltersOrErr() ([]*CloudEventFilters, error) {
	if e.loadedTypes[4] {
		return e.Cloudeventfilters, nil
	}
	return nil, &NotLoadedError{edge: "cloudeventfilters"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Namespace) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case namespace.FieldConfig, namespace.FieldName:
			values[i] = new(sql.NullString)
		case namespace.FieldCreatedAt, namespace.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case namespace.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Namespace", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Namespace fields.
func (n *Namespace) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case namespace.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				n.ID = *value
			}
		case namespace.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case namespace.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case namespace.FieldConfig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value.Valid {
				n.Config = value.String
			}
		case namespace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		}
	}
	return nil
}

// QueryLogs queries the "logs" edge of the Namespace entity.
func (n *Namespace) QueryLogs() *LogMsgQuery {
	return NewNamespaceClient(n.config).QueryLogs(n)
}

// QueryCloudevents queries the "cloudevents" edge of the Namespace entity.
func (n *Namespace) QueryCloudevents() *CloudEventsQuery {
	return NewNamespaceClient(n.config).QueryCloudevents(n)
}

// QueryNamespacelisteners queries the "namespacelisteners" edge of the Namespace entity.
func (n *Namespace) QueryNamespacelisteners() *EventsQuery {
	return NewNamespaceClient(n.config).QueryNamespacelisteners(n)
}

// QueryAnnotations queries the "annotations" edge of the Namespace entity.
func (n *Namespace) QueryAnnotations() *AnnotationQuery {
	return NewNamespaceClient(n.config).QueryAnnotations(n)
}

// QueryCloudeventfilters queries the "cloudeventfilters" edge of the Namespace entity.
func (n *Namespace) QueryCloudeventfilters() *CloudEventFiltersQuery {
	return NewNamespaceClient(n.config).QueryCloudeventfilters(n)
}

// Update returns a builder for updating this Namespace.
// Note that you need to call Namespace.Unwrap() before calling this method if this Namespace
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Namespace) Update() *NamespaceUpdateOne {
	return NewNamespaceClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Namespace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Namespace) Unwrap() *Namespace {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Namespace is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Namespace) String() string {
	var builder strings.Builder
	builder.WriteString("Namespace(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(n.Config)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(n.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Namespaces is a parsable slice of Namespace.
type Namespaces []*Namespace
