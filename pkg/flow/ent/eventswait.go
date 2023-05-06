// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/google/uuid"
)

// EventsWait is the model entity for the EventsWait schema.
type EventsWait struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id"`
	// Events holds the value of the "events" field.
	Events map[string]interface{} `json:"events,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventsWaitQuery when eager-loading is set.
	Edges               EventsWaitEdges `json:"edges"`
	events_wfeventswait *uuid.UUID
}

// EventsWaitEdges holds the relations/edges for other nodes in the graph.
type EventsWaitEdges struct {
	// Workflowevent holds the value of the workflowevent edge.
	Workflowevent *Events `json:"workflowevent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorkfloweventOrErr returns the Workflowevent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventsWaitEdges) WorkfloweventOrErr() (*Events, error) {
	if e.loadedTypes[0] {
		if e.Workflowevent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: events.Label}
		}
		return e.Workflowevent, nil
	}
	return nil, &NotLoadedError{edge: "workflowevent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EventsWait) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case eventswait.FieldEvents:
			values[i] = new([]byte)
		case eventswait.FieldID:
			values[i] = new(uuid.UUID)
		case eventswait.ForeignKeys[0]: // events_wfeventswait
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type EventsWait", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EventsWait fields.
func (ew *EventsWait) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case eventswait.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ew.ID = *value
			}
		case eventswait.FieldEvents:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field events", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ew.Events); err != nil {
					return fmt.Errorf("unmarshal field events: %w", err)
				}
			}
		case eventswait.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field events_wfeventswait", values[i])
			} else if value.Valid {
				ew.events_wfeventswait = new(uuid.UUID)
				*ew.events_wfeventswait = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryWorkflowevent queries the "workflowevent" edge of the EventsWait entity.
func (ew *EventsWait) QueryWorkflowevent() *EventsQuery {
	return NewEventsWaitClient(ew.config).QueryWorkflowevent(ew)
}

// Update returns a builder for updating this EventsWait.
// Note that you need to call EventsWait.Unwrap() before calling this method if this EventsWait
// was returned from a transaction, and the transaction was committed or rolled back.
func (ew *EventsWait) Update() *EventsWaitUpdateOne {
	return NewEventsWaitClient(ew.config).UpdateOne(ew)
}

// Unwrap unwraps the EventsWait entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ew *EventsWait) Unwrap() *EventsWait {
	_tx, ok := ew.config.driver.(*txDriver)
	if !ok {
		panic("ent: EventsWait is not a transactional entity")
	}
	ew.config.driver = _tx.drv
	return ew
}

// String implements the fmt.Stringer.
func (ew *EventsWait) String() string {
	var builder strings.Builder
	builder.WriteString("EventsWait(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ew.ID))
	builder.WriteString("events=")
	builder.WriteString(fmt.Sprintf("%v", ew.Events))
	builder.WriteByte(')')
	return builder.String()
}

// EventsWaits is a parsable slice of EventsWait.
type EventsWaits []*EventsWait
