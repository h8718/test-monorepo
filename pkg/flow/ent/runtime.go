// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/direktiv/direktiv/pkg/flow/ent/annotation"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudeventfilters"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudevents"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	annotationFields := schema.Annotation{}.Fields()
	_ = annotationFields
	// annotationDescName is the schema descriptor for name field.
	annotationDescName := annotationFields[1].Descriptor()
	// annotation.NameValidator is a validator for the "name" field. It is called by the builders before save.
	annotation.NameValidator = annotationDescName.Validators[0].(func(string) error)
	// annotationDescCreatedAt is the schema descriptor for created_at field.
	annotationDescCreatedAt := annotationFields[2].Descriptor()
	// annotation.DefaultCreatedAt holds the default value on creation for the created_at field.
	annotation.DefaultCreatedAt = annotationDescCreatedAt.Default.(func() time.Time)
	// annotationDescUpdatedAt is the schema descriptor for updated_at field.
	annotationDescUpdatedAt := annotationFields[3].Descriptor()
	// annotation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	annotation.DefaultUpdatedAt = annotationDescUpdatedAt.Default.(func() time.Time)
	// annotation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	annotation.UpdateDefaultUpdatedAt = annotationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// annotationDescID is the schema descriptor for id field.
	annotationDescID := annotationFields[0].Descriptor()
	// annotation.DefaultID holds the default value on creation for the id field.
	annotation.DefaultID = annotationDescID.Default.(func() uuid.UUID)
	cloudeventfiltersFields := schema.CloudEventFilters{}.Fields()
	_ = cloudeventfiltersFields
	// cloudeventfiltersDescName is the schema descriptor for name field.
	cloudeventfiltersDescName := cloudeventfiltersFields[0].Descriptor()
	// cloudeventfilters.NameValidator is a validator for the "name" field. It is called by the builders before save.
	cloudeventfilters.NameValidator = cloudeventfiltersDescName.Validators[0].(func(string) error)
	// cloudeventfiltersDescJscode is the schema descriptor for jscode field.
	cloudeventfiltersDescJscode := cloudeventfiltersFields[1].Descriptor()
	// cloudeventfilters.JscodeValidator is a validator for the "jscode" field. It is called by the builders before save.
	cloudeventfilters.JscodeValidator = cloudeventfiltersDescJscode.Validators[0].(func(string) error)
	cloudeventsFields := schema.CloudEvents{}.Fields()
	_ = cloudeventsFields
	// cloudeventsDescEventId is the schema descriptor for eventId field.
	cloudeventsDescEventId := cloudeventsFields[1].Descriptor()
	// cloudevents.EventIdValidator is a validator for the "eventId" field. It is called by the builders before save.
	cloudevents.EventIdValidator = cloudeventsDescEventId.Validators[0].(func(string) error)
	// cloudeventsDescFire is the schema descriptor for fire field.
	cloudeventsDescFire := cloudeventsFields[3].Descriptor()
	// cloudevents.DefaultFire holds the default value on creation for the fire field.
	cloudevents.DefaultFire = cloudeventsDescFire.Default.(func() time.Time)
	// cloudeventsDescCreated is the schema descriptor for created field.
	cloudeventsDescCreated := cloudeventsFields[4].Descriptor()
	// cloudevents.DefaultCreated holds the default value on creation for the created field.
	cloudevents.DefaultCreated = cloudeventsDescCreated.Default.(func() time.Time)
	// cloudeventsDescID is the schema descriptor for id field.
	cloudeventsDescID := cloudeventsFields[0].Descriptor()
	// cloudevents.DefaultID holds the default value on creation for the id field.
	cloudevents.DefaultID = cloudeventsDescID.Default.(func() uuid.UUID)
	eventsFields := schema.Events{}.Fields()
	_ = eventsFields
	// eventsDescCreatedAt is the schema descriptor for created_at field.
	eventsDescCreatedAt := eventsFields[5].Descriptor()
	// events.DefaultCreatedAt holds the default value on creation for the created_at field.
	events.DefaultCreatedAt = eventsDescCreatedAt.Default.(func() time.Time)
	// eventsDescUpdatedAt is the schema descriptor for updated_at field.
	eventsDescUpdatedAt := eventsFields[6].Descriptor()
	// events.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	events.DefaultUpdatedAt = eventsDescUpdatedAt.Default.(func() time.Time)
	// events.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	events.UpdateDefaultUpdatedAt = eventsDescUpdatedAt.UpdateDefault.(func() time.Time)
	// eventsDescID is the schema descriptor for id field.
	eventsDescID := eventsFields[0].Descriptor()
	// events.DefaultID holds the default value on creation for the id field.
	events.DefaultID = eventsDescID.Default.(func() uuid.UUID)
	eventswaitFields := schema.EventsWait{}.Fields()
	_ = eventswaitFields
	// eventswaitDescID is the schema descriptor for id field.
	eventswaitDescID := eventswaitFields[0].Descriptor()
	// eventswait.DefaultID holds the default value on creation for the id field.
	eventswait.DefaultID = eventswaitDescID.Default.(func() uuid.UUID)
	logmsgFields := schema.LogMsg{}.Fields()
	_ = logmsgFields
	// logmsgDescLevel is the schema descriptor for level field.
	logmsgDescLevel := logmsgFields[3].Descriptor()
	// logmsg.DefaultLevel holds the default value on creation for the level field.
	logmsg.DefaultLevel = logmsgDescLevel.Default.(string)
	// logmsgDescRootInstanceId is the schema descriptor for rootInstanceId field.
	logmsgDescRootInstanceId := logmsgFields[4].Descriptor()
	// logmsg.DefaultRootInstanceId holds the default value on creation for the rootInstanceId field.
	logmsg.DefaultRootInstanceId = logmsgDescRootInstanceId.Default.(string)
	// logmsgDescLogInstanceCallPath is the schema descriptor for logInstanceCallPath field.
	logmsgDescLogInstanceCallPath := logmsgFields[5].Descriptor()
	// logmsg.DefaultLogInstanceCallPath holds the default value on creation for the logInstanceCallPath field.
	logmsg.DefaultLogInstanceCallPath = logmsgDescLogInstanceCallPath.Default.(string)
	// logmsgDescID is the schema descriptor for id field.
	logmsgDescID := logmsgFields[0].Descriptor()
	// logmsg.DefaultID holds the default value on creation for the id field.
	logmsg.DefaultID = logmsgDescID.Default.(func() uuid.UUID)
	namespaceFields := schema.Namespace{}.Fields()
	_ = namespaceFields
	// namespaceDescCreatedAt is the schema descriptor for created_at field.
	namespaceDescCreatedAt := namespaceFields[1].Descriptor()
	// namespace.DefaultCreatedAt holds the default value on creation for the created_at field.
	namespace.DefaultCreatedAt = namespaceDescCreatedAt.Default.(func() time.Time)
	// namespaceDescUpdatedAt is the schema descriptor for updated_at field.
	namespaceDescUpdatedAt := namespaceFields[2].Descriptor()
	// namespace.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	namespace.DefaultUpdatedAt = namespaceDescUpdatedAt.Default.(func() time.Time)
	// namespace.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	namespace.UpdateDefaultUpdatedAt = namespaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// namespaceDescConfig is the schema descriptor for config field.
	namespaceDescConfig := namespaceFields[3].Descriptor()
	// namespace.DefaultConfig holds the default value on creation for the config field.
	namespace.DefaultConfig = namespaceDescConfig.Default.(string)
	// namespaceDescName is the schema descriptor for name field.
	namespaceDescName := namespaceFields[4].Descriptor()
	// namespace.NameValidator is a validator for the "name" field. It is called by the builders before save.
	namespace.NameValidator = func() func(string) error {
		validators := namespaceDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
			validators[3].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// namespaceDescID is the schema descriptor for id field.
	namespaceDescID := namespaceFields[0].Descriptor()
	// namespace.DefaultID holds the default value on creation for the id field.
	namespace.DefaultID = namespaceDescID.Default.(func() uuid.UUID)
}
