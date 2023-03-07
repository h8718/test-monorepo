// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirroractivity"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// LogMsg is the model entity for the LogMsg schema.
type LogMsg struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"-"`
	// T holds the value of the "t" field.
	T time.Time `json:"t,omitempty"`
	// Msg holds the value of the "msg" field.
	Msg string `json:"msg,omitempty"`
	// Root holds the value of the "root" field.
	Root string `json:"root,omitempty"`
	// Callpath holds the value of the "callpath" field.
	Callpath string `json:"callpath,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LogMsgQuery when eager-loading is set.
	Edges                LogMsgEdges `json:"edges"`
	instance_logs        *uuid.UUID
	mirror_activity_logs *uuid.UUID
	namespace_logs       *uuid.UUID
	workflow_logs        *uuid.UUID
}

// LogMsgEdges holds the relations/edges for other nodes in the graph.
type LogMsgEdges struct {
	// Namespace holds the value of the namespace edge.
	Namespace *Namespace `json:"namespace,omitempty"`
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// Instance holds the value of the instance edge.
	Instance *Instance `json:"instance,omitempty"`
	// Activity holds the value of the activity edge.
	Activity *MirrorActivity `json:"activity,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LogMsgEdges) NamespaceOrErr() (*Namespace, error) {
	if e.loadedTypes[0] {
		if e.Namespace == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: namespace.Label}
		}
		return e.Namespace, nil
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LogMsgEdges) WorkflowOrErr() (*Workflow, error) {
	if e.loadedTypes[1] {
		if e.Workflow == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: workflow.Label}
		}
		return e.Workflow, nil
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// InstanceOrErr returns the Instance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LogMsgEdges) InstanceOrErr() (*Instance, error) {
	if e.loadedTypes[2] {
		if e.Instance == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: instance.Label}
		}
		return e.Instance, nil
	}
	return nil, &NotLoadedError{edge: "instance"}
}

// ActivityOrErr returns the Activity value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LogMsgEdges) ActivityOrErr() (*MirrorActivity, error) {
	if e.loadedTypes[3] {
		if e.Activity == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mirroractivity.Label}
		}
		return e.Activity, nil
	}
	return nil, &NotLoadedError{edge: "activity"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LogMsg) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case logmsg.FieldMsg, logmsg.FieldRoot, logmsg.FieldCallpath:
			values[i] = new(sql.NullString)
		case logmsg.FieldT:
			values[i] = new(sql.NullTime)
		case logmsg.FieldID:
			values[i] = new(uuid.UUID)
		case logmsg.ForeignKeys[0]: // instance_logs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case logmsg.ForeignKeys[1]: // mirror_activity_logs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case logmsg.ForeignKeys[2]: // namespace_logs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case logmsg.ForeignKeys[3]: // workflow_logs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type LogMsg", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LogMsg fields.
func (lm *LogMsg) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case logmsg.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				lm.ID = *value
			}
		case logmsg.FieldT:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field t", values[i])
			} else if value.Valid {
				lm.T = value.Time
			}
		case logmsg.FieldMsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field msg", values[i])
			} else if value.Valid {
				lm.Msg = value.String
			}
		case logmsg.FieldRoot:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field root", values[i])
			} else if value.Valid {
				lm.Root = value.String
			}
		case logmsg.FieldCallpath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field callpath", values[i])
			} else if value.Valid {
				lm.Callpath = value.String
			}
		case logmsg.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field instance_logs", values[i])
			} else if value.Valid {
				lm.instance_logs = new(uuid.UUID)
				*lm.instance_logs = *value.S.(*uuid.UUID)
			}
		case logmsg.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field mirror_activity_logs", values[i])
			} else if value.Valid {
				lm.mirror_activity_logs = new(uuid.UUID)
				*lm.mirror_activity_logs = *value.S.(*uuid.UUID)
			}
		case logmsg.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field namespace_logs", values[i])
			} else if value.Valid {
				lm.namespace_logs = new(uuid.UUID)
				*lm.namespace_logs = *value.S.(*uuid.UUID)
			}
		case logmsg.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_logs", values[i])
			} else if value.Valid {
				lm.workflow_logs = new(uuid.UUID)
				*lm.workflow_logs = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryNamespace queries the "namespace" edge of the LogMsg entity.
func (lm *LogMsg) QueryNamespace() *NamespaceQuery {
	return (&LogMsgClient{config: lm.config}).QueryNamespace(lm)
}

// QueryWorkflow queries the "workflow" edge of the LogMsg entity.
func (lm *LogMsg) QueryWorkflow() *WorkflowQuery {
	return (&LogMsgClient{config: lm.config}).QueryWorkflow(lm)
}

// QueryInstance queries the "instance" edge of the LogMsg entity.
func (lm *LogMsg) QueryInstance() *InstanceQuery {
	return (&LogMsgClient{config: lm.config}).QueryInstance(lm)
}

// QueryActivity queries the "activity" edge of the LogMsg entity.
func (lm *LogMsg) QueryActivity() *MirrorActivityQuery {
	return (&LogMsgClient{config: lm.config}).QueryActivity(lm)
}

// Update returns a builder for updating this LogMsg.
// Note that you need to call LogMsg.Unwrap() before calling this method if this LogMsg
// was returned from a transaction, and the transaction was committed or rolled back.
func (lm *LogMsg) Update() *LogMsgUpdateOne {
	return (&LogMsgClient{config: lm.config}).UpdateOne(lm)
}

// Unwrap unwraps the LogMsg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lm *LogMsg) Unwrap() *LogMsg {
	_tx, ok := lm.config.driver.(*txDriver)
	if !ok {
		panic("ent: LogMsg is not a transactional entity")
	}
	lm.config.driver = _tx.drv
	return lm
}

// String implements the fmt.Stringer.
func (lm *LogMsg) String() string {
	var builder strings.Builder
	builder.WriteString("LogMsg(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lm.ID))
	builder.WriteString("t=")
	builder.WriteString(lm.T.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("msg=")
	builder.WriteString(lm.Msg)
	builder.WriteString(", ")
	builder.WriteString("root=")
	builder.WriteString(lm.Root)
	builder.WriteString(", ")
	builder.WriteString("callpath=")
	builder.WriteString(lm.Callpath)
	builder.WriteByte(')')
	return builder.String()
}

// LogMsgs is a parsable slice of LogMsg.
type LogMsgs []*LogMsg

func (lm LogMsgs) config(cfg config) {
	for _i := range lm {
		lm[_i].config = cfg
	}
}
