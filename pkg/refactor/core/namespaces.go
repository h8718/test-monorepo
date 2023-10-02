package core

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Namespace struct {
	ID uuid.UUID

	Name   string
	Config string

	RootsInfo string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RootInfo struct {
	Name   string
	RootID uuid.UUID
}

//nolint:musttag
type RootsInfo struct {
	Default RootInfo
}

func (ri *RootsInfo) Marshal() string {
	data, err := json.Marshal(ri)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func (ns *Namespace) Roots() (*RootsInfo, error) {
	ri := new(RootsInfo)
	err := json.Unmarshal([]byte(ns.RootsInfo), ri)
	if err != nil {
		return nil, err
	}

	return ri, nil
}

func (ns *Namespace) GetAttributes() map[string]string {
	return map[string]string{
		"namespace":    ns.Name,
		"namespace-id": ns.ID.String(),
	}
}

var (
	ErrInvalidNamespaceName    = errors.New("ErrInvalidNamespaceName")
	ErrDuplicatedNamespaceName = errors.New("ErrDuplicatedNamespaceName")
)

// NamespacesStore responsible for fetching and setting namespaces from datastore.
type NamespacesStore interface {
	// GetByID gets a single namespace object from store. if no record found,
	// it returns datastore.ErrNotFound error.
	GetByID(ctx context.Context, id uuid.UUID) (*Namespace, error)

	// GetByName gets a single namespace object from store. if no record found,
	// it returns datastore.ErrNotFound error.
	GetByName(ctx context.Context, name string) (*Namespace, error)

	// GetAll gets all namespaces from store.
	GetAll(ctx context.Context) ([]*Namespace, error)

	// Update changes a namespace data.
	Update(ctx context.Context, namespace *Namespace) (*Namespace, error)

	// Delete deletes a single namespace. if no record found,
	//	// it returns datastore.ErrNotFound error.
	Delete(ctx context.Context, id uuid.UUID) error

	// Create creates a new namespace. Returned errors could be ErrDuplicatedNamespaceName when namespace name is
	// already exists or ErrInvalidNamespaceName or when namespace name is invalid, too short or too long.
	Create(ctx context.Context, namespace *Namespace) (*Namespace, error)
}

const DefaultNamespaceConfig = `
{
	"broadcast": {
	  "workflow.create": false,
	  "workflow.update": false,
	  "workflow.delete": false,
	  "directory.create": false,
	  "directory.delete": false,
	  "workflow.variable.create": false,
	  "workflow.variable.update": false,
	  "workflow.variable.delete": false,
	  "namespace.variable.create": false,
	  "namespace.variable.update": false,
	  "namespace.variable.delete": false,
	  "instance.variable.create": false,
	  "instance.variable.update": false,
	  "instance.variable.delete": false,
	  "instance.started": false,
	  "instance.success": false,
	  "instance.failed": false
	}
  }
`
