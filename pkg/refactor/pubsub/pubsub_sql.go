package pubsub

type Bus interface {
	Publish(channel string, data string) error
	Subscribe(handler func(data string), channels ...string)
}

var (
	WorkflowCreate = "workflow_create"
	WorkflowDelete = "workflow_delete"
	WorkflowUpdate = "workflow_update"

	ServiceCreate = "service_create"
	ServiceDelete = "service_delete"
	ServiceUpdate = "service_update"

	MirrorSync = "mirror_sync"
)
