// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NamespacesColumns holds the columns for the "namespaces" table.
	NamespacesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "config", Type: field.TypeString, Default: "\n{\n\t\"broadcast\": {\n\t  \"workflow.create\": false,\n\t  \"workflow.update\": false,\n\t  \"workflow.delete\": false,\n\t  \"directory.create\": false,\n\t  \"directory.delete\": false,\n\t  \"workflow.variable.create\": false,\n\t  \"workflow.variable.update\": false,\n\t  \"workflow.variable.delete\": false,\n\t  \"namespace.variable.create\": false,\n\t  \"namespace.variable.update\": false,\n\t  \"namespace.variable.delete\": false,\n\t  \"instance.variable.create\": false,\n\t  \"instance.variable.update\": false,\n\t  \"instance.variable.delete\": false,\n\t  \"instance.started\": false,\n\t  \"instance.success\": false,\n\t  \"instance.failed\": false\n\t}\n  }\n"},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 64},
	}
	// NamespacesTable holds the schema information for the "namespaces" table.
	NamespacesTable = &schema.Table{
		Name:       "namespaces",
		Columns:    NamespacesColumns,
		PrimaryKey: []*schema.Column{NamespacesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NamespacesTable,
	}
)

func init() {
}
