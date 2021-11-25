// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/instanceruntime"
	"github.com/google/uuid"
)

// InstanceRuntime is the model entity for the InstanceRuntime schema.
type InstanceRuntime struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"-"`
	// Input holds the value of the "input" field.
	Input []byte `json:"input,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
	// Controller holds the value of the "controller" field.
	Controller string `json:"controller,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
	// Flow holds the value of the "flow" field.
	Flow []string `json:"flow,omitempty"`
	// Output holds the value of the "output" field.
	Output string `json:"output,omitempty"`
	// StateBeginTime holds the value of the "stateBeginTime" field.
	StateBeginTime time.Time `json:"stateBeginTime,omitempty"`
	// Deadline holds the value of the "deadline" field.
	Deadline time.Time `json:"deadline,omitempty"`
	// Attempts holds the value of the "attempts" field.
	Attempts int `json:"attempts,omitempty"`
	// CallerData holds the value of the "caller_data" field.
	CallerData string `json:"caller_data,omitempty"`
	// InstanceContext holds the value of the "instanceContext" field.
	InstanceContext string `json:"instanceContext,omitempty"`
	// StateContext holds the value of the "stateContext" field.
	StateContext string `json:"stateContext,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstanceRuntimeQuery when eager-loading is set.
	Edges             InstanceRuntimeEdges `json:"edges"`
	instance_runtime  *uuid.UUID
	instance_children *uuid.UUID
}

// InstanceRuntimeEdges holds the relations/edges for other nodes in the graph.
type InstanceRuntimeEdges struct {
	// Instance holds the value of the instance edge.
	Instance *Instance `json:"instance,omitempty"`
	// Caller holds the value of the caller edge.
	Caller *Instance `json:"caller,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// InstanceOrErr returns the Instance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceRuntimeEdges) InstanceOrErr() (*Instance, error) {
	if e.loadedTypes[0] {
		if e.Instance == nil {
			// The edge instance was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instance.Label}
		}
		return e.Instance, nil
	}
	return nil, &NotLoadedError{edge: "instance"}
}

// CallerOrErr returns the Caller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceRuntimeEdges) CallerOrErr() (*Instance, error) {
	if e.loadedTypes[1] {
		if e.Caller == nil {
			// The edge caller was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instance.Label}
		}
		return e.Caller, nil
	}
	return nil, &NotLoadedError{edge: "caller"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InstanceRuntime) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case instanceruntime.FieldInput, instanceruntime.FieldFlow:
			values[i] = new([]byte)
		case instanceruntime.FieldAttempts:
			values[i] = new(sql.NullInt64)
		case instanceruntime.FieldData, instanceruntime.FieldController, instanceruntime.FieldMemory, instanceruntime.FieldOutput, instanceruntime.FieldCallerData, instanceruntime.FieldInstanceContext, instanceruntime.FieldStateContext:
			values[i] = new(sql.NullString)
		case instanceruntime.FieldStateBeginTime, instanceruntime.FieldDeadline:
			values[i] = new(sql.NullTime)
		case instanceruntime.FieldID:
			values[i] = new(uuid.UUID)
		case instanceruntime.ForeignKeys[0]: // instance_runtime
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case instanceruntime.ForeignKeys[1]: // instance_children
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type InstanceRuntime", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InstanceRuntime fields.
func (ir *InstanceRuntime) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case instanceruntime.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ir.ID = *value
			}
		case instanceruntime.FieldInput:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field input", values[i])
			} else if value != nil {
				ir.Input = *value
			}
		case instanceruntime.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				ir.Data = value.String
			}
		case instanceruntime.FieldController:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field controller", values[i])
			} else if value.Valid {
				ir.Controller = value.String
			}
		case instanceruntime.FieldMemory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				ir.Memory = value.String
			}
		case instanceruntime.FieldFlow:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field flow", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ir.Flow); err != nil {
					return fmt.Errorf("unmarshal field flow: %w", err)
				}
			}
		case instanceruntime.FieldOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[i])
			} else if value.Valid {
				ir.Output = value.String
			}
		case instanceruntime.FieldStateBeginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field stateBeginTime", values[i])
			} else if value.Valid {
				ir.StateBeginTime = value.Time
			}
		case instanceruntime.FieldDeadline:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deadline", values[i])
			} else if value.Valid {
				ir.Deadline = value.Time
			}
		case instanceruntime.FieldAttempts:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attempts", values[i])
			} else if value.Valid {
				ir.Attempts = int(value.Int64)
			}
		case instanceruntime.FieldCallerData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field caller_data", values[i])
			} else if value.Valid {
				ir.CallerData = value.String
			}
		case instanceruntime.FieldInstanceContext:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instanceContext", values[i])
			} else if value.Valid {
				ir.InstanceContext = value.String
			}
		case instanceruntime.FieldStateContext:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stateContext", values[i])
			} else if value.Valid {
				ir.StateContext = value.String
			}
		case instanceruntime.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field instance_runtime", values[i])
			} else if value.Valid {
				ir.instance_runtime = new(uuid.UUID)
				*ir.instance_runtime = *value.S.(*uuid.UUID)
			}
		case instanceruntime.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field instance_children", values[i])
			} else if value.Valid {
				ir.instance_children = new(uuid.UUID)
				*ir.instance_children = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryInstance queries the "instance" edge of the InstanceRuntime entity.
func (ir *InstanceRuntime) QueryInstance() *InstanceQuery {
	return (&InstanceRuntimeClient{config: ir.config}).QueryInstance(ir)
}

// QueryCaller queries the "caller" edge of the InstanceRuntime entity.
func (ir *InstanceRuntime) QueryCaller() *InstanceQuery {
	return (&InstanceRuntimeClient{config: ir.config}).QueryCaller(ir)
}

// Update returns a builder for updating this InstanceRuntime.
// Note that you need to call InstanceRuntime.Unwrap() before calling this method if this InstanceRuntime
// was returned from a transaction, and the transaction was committed or rolled back.
func (ir *InstanceRuntime) Update() *InstanceRuntimeUpdateOne {
	return (&InstanceRuntimeClient{config: ir.config}).UpdateOne(ir)
}

// Unwrap unwraps the InstanceRuntime entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ir *InstanceRuntime) Unwrap() *InstanceRuntime {
	tx, ok := ir.config.driver.(*txDriver)
	if !ok {
		panic("ent: InstanceRuntime is not a transactional entity")
	}
	ir.config.driver = tx.drv
	return ir
}

// String implements the fmt.Stringer.
func (ir *InstanceRuntime) String() string {
	var builder strings.Builder
	builder.WriteString("InstanceRuntime(")
	builder.WriteString(fmt.Sprintf("id=%v", ir.ID))
	builder.WriteString(", input=")
	builder.WriteString(fmt.Sprintf("%v", ir.Input))
	builder.WriteString(", data=")
	builder.WriteString(ir.Data)
	builder.WriteString(", controller=")
	builder.WriteString(ir.Controller)
	builder.WriteString(", memory=")
	builder.WriteString(ir.Memory)
	builder.WriteString(", flow=")
	builder.WriteString(fmt.Sprintf("%v", ir.Flow))
	builder.WriteString(", output=")
	builder.WriteString(ir.Output)
	builder.WriteString(", stateBeginTime=")
	builder.WriteString(ir.StateBeginTime.Format(time.ANSIC))
	builder.WriteString(", deadline=")
	builder.WriteString(ir.Deadline.Format(time.ANSIC))
	builder.WriteString(", attempts=")
	builder.WriteString(fmt.Sprintf("%v", ir.Attempts))
	builder.WriteString(", caller_data=")
	builder.WriteString(ir.CallerData)
	builder.WriteString(", instanceContext=")
	builder.WriteString(ir.InstanceContext)
	builder.WriteString(", stateContext=")
	builder.WriteString(ir.StateContext)
	builder.WriteByte(')')
	return builder.String()
}

// InstanceRuntimes is a parsable slice of InstanceRuntime.
type InstanceRuntimes []*InstanceRuntime

func (ir InstanceRuntimes) config(cfg config) {
	for _i := range ir {
		ir[_i].config = cfg
	}
}
