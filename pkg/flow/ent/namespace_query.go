// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudeventfilters"
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudevents"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// NamespaceQuery is the builder for querying Namespace entities.
type NamespaceQuery struct {
	config
	ctx                    *QueryContext
	order                  []OrderFunc
	inters                 []Interceptor
	predicates             []predicate.Namespace
	withLogs               *LogMsgQuery
	withCloudevents        *CloudEventsQuery
	withNamespacelisteners *EventsQuery
	withCloudeventfilters  *CloudEventFiltersQuery
	modifiers              []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NamespaceQuery builder.
func (nq *NamespaceQuery) Where(ps ...predicate.Namespace) *NamespaceQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NamespaceQuery) Limit(limit int) *NamespaceQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NamespaceQuery) Offset(offset int) *NamespaceQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NamespaceQuery) Unique(unique bool) *NamespaceQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NamespaceQuery) Order(o ...OrderFunc) *NamespaceQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryLogs chains the current query on the "logs" edge.
func (nq *NamespaceQuery) QueryLogs() *LogMsgQuery {
	query := (&LogMsgClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(namespace.Table, namespace.FieldID, selector),
			sqlgraph.To(logmsg.Table, logmsg.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, namespace.LogsTable, namespace.LogsColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCloudevents chains the current query on the "cloudevents" edge.
func (nq *NamespaceQuery) QueryCloudevents() *CloudEventsQuery {
	query := (&CloudEventsClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(namespace.Table, namespace.FieldID, selector),
			sqlgraph.To(cloudevents.Table, cloudevents.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, namespace.CloudeventsTable, namespace.CloudeventsColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNamespacelisteners chains the current query on the "namespacelisteners" edge.
func (nq *NamespaceQuery) QueryNamespacelisteners() *EventsQuery {
	query := (&EventsClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(namespace.Table, namespace.FieldID, selector),
			sqlgraph.To(events.Table, events.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, namespace.NamespacelistenersTable, namespace.NamespacelistenersColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCloudeventfilters chains the current query on the "cloudeventfilters" edge.
func (nq *NamespaceQuery) QueryCloudeventfilters() *CloudEventFiltersQuery {
	query := (&CloudEventFiltersClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(namespace.Table, namespace.FieldID, selector),
			sqlgraph.To(cloudeventfilters.Table, cloudeventfilters.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, namespace.CloudeventfiltersTable, namespace.CloudeventfiltersColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Namespace entity from the query.
// Returns a *NotFoundError when no Namespace was found.
func (nq *NamespaceQuery) First(ctx context.Context) (*Namespace, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{namespace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NamespaceQuery) FirstX(ctx context.Context) *Namespace {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Namespace ID from the query.
// Returns a *NotFoundError when no Namespace ID was found.
func (nq *NamespaceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{namespace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NamespaceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Namespace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Namespace entity is found.
// Returns a *NotFoundError when no Namespace entities are found.
func (nq *NamespaceQuery) Only(ctx context.Context) (*Namespace, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{namespace.Label}
	default:
		return nil, &NotSingularError{namespace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NamespaceQuery) OnlyX(ctx context.Context) *Namespace {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Namespace ID in the query.
// Returns a *NotSingularError when more than one Namespace ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NamespaceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{namespace.Label}
	default:
		err = &NotSingularError{namespace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NamespaceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Namespaces.
func (nq *NamespaceQuery) All(ctx context.Context) ([]*Namespace, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Namespace, *NamespaceQuery]()
	return withInterceptors[[]*Namespace](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NamespaceQuery) AllX(ctx context.Context) []*Namespace {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Namespace IDs.
func (nq *NamespaceQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(namespace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NamespaceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NamespaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NamespaceQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NamespaceQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NamespaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NamespaceQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NamespaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NamespaceQuery) Clone() *NamespaceQuery {
	if nq == nil {
		return nil
	}
	return &NamespaceQuery{
		config:                 nq.config,
		ctx:                    nq.ctx.Clone(),
		order:                  append([]OrderFunc{}, nq.order...),
		inters:                 append([]Interceptor{}, nq.inters...),
		predicates:             append([]predicate.Namespace{}, nq.predicates...),
		withLogs:               nq.withLogs.Clone(),
		withCloudevents:        nq.withCloudevents.Clone(),
		withNamespacelisteners: nq.withNamespacelisteners.Clone(),
		withCloudeventfilters:  nq.withCloudeventfilters.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithLogs tells the query-builder to eager-load the nodes that are connected to
// the "logs" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NamespaceQuery) WithLogs(opts ...func(*LogMsgQuery)) *NamespaceQuery {
	query := (&LogMsgClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withLogs = query
	return nq
}

// WithCloudevents tells the query-builder to eager-load the nodes that are connected to
// the "cloudevents" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NamespaceQuery) WithCloudevents(opts ...func(*CloudEventsQuery)) *NamespaceQuery {
	query := (&CloudEventsClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withCloudevents = query
	return nq
}

// WithNamespacelisteners tells the query-builder to eager-load the nodes that are connected to
// the "namespacelisteners" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NamespaceQuery) WithNamespacelisteners(opts ...func(*EventsQuery)) *NamespaceQuery {
	query := (&EventsClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNamespacelisteners = query
	return nq
}

// WithCloudeventfilters tells the query-builder to eager-load the nodes that are connected to
// the "cloudeventfilters" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NamespaceQuery) WithCloudeventfilters(opts ...func(*CloudEventFiltersQuery)) *NamespaceQuery {
	query := (&CloudEventFiltersClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withCloudeventfilters = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Namespace.Query().
//		GroupBy(namespace.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NamespaceQuery) GroupBy(field string, fields ...string) *NamespaceGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NamespaceGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = namespace.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Namespace.Query().
//		Select(namespace.FieldCreatedAt).
//		Scan(ctx, &v)
func (nq *NamespaceQuery) Select(fields ...string) *NamespaceSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NamespaceSelect{NamespaceQuery: nq}
	sbuild.label = namespace.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NamespaceSelect configured with the given aggregations.
func (nq *NamespaceQuery) Aggregate(fns ...AggregateFunc) *NamespaceSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NamespaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !namespace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NamespaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Namespace, error) {
	var (
		nodes       = []*Namespace{}
		_spec       = nq.querySpec()
		loadedTypes = [4]bool{
			nq.withLogs != nil,
			nq.withCloudevents != nil,
			nq.withNamespacelisteners != nil,
			nq.withCloudeventfilters != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Namespace).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Namespace{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withLogs; query != nil {
		if err := nq.loadLogs(ctx, query, nodes,
			func(n *Namespace) { n.Edges.Logs = []*LogMsg{} },
			func(n *Namespace, e *LogMsg) { n.Edges.Logs = append(n.Edges.Logs, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withCloudevents; query != nil {
		if err := nq.loadCloudevents(ctx, query, nodes,
			func(n *Namespace) { n.Edges.Cloudevents = []*CloudEvents{} },
			func(n *Namespace, e *CloudEvents) { n.Edges.Cloudevents = append(n.Edges.Cloudevents, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withNamespacelisteners; query != nil {
		if err := nq.loadNamespacelisteners(ctx, query, nodes,
			func(n *Namespace) { n.Edges.Namespacelisteners = []*Events{} },
			func(n *Namespace, e *Events) { n.Edges.Namespacelisteners = append(n.Edges.Namespacelisteners, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withCloudeventfilters; query != nil {
		if err := nq.loadCloudeventfilters(ctx, query, nodes,
			func(n *Namespace) { n.Edges.Cloudeventfilters = []*CloudEventFilters{} },
			func(n *Namespace, e *CloudEventFilters) {
				n.Edges.Cloudeventfilters = append(n.Edges.Cloudeventfilters, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NamespaceQuery) loadLogs(ctx context.Context, query *LogMsgQuery, nodes []*Namespace, init func(*Namespace), assign func(*Namespace, *LogMsg)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Namespace)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.LogMsg(func(s *sql.Selector) {
		s.Where(sql.InValues(namespace.LogsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.namespace_logs
		if fk == nil {
			return fmt.Errorf(`foreign-key "namespace_logs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "namespace_logs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NamespaceQuery) loadCloudevents(ctx context.Context, query *CloudEventsQuery, nodes []*Namespace, init func(*Namespace), assign func(*Namespace, *CloudEvents)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Namespace)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.InValues(namespace.CloudeventsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.namespace_cloudevents
		if fk == nil {
			return fmt.Errorf(`foreign-key "namespace_cloudevents" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "namespace_cloudevents" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NamespaceQuery) loadNamespacelisteners(ctx context.Context, query *EventsQuery, nodes []*Namespace, init func(*Namespace), assign func(*Namespace, *Events)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Namespace)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Events(func(s *sql.Selector) {
		s.Where(sql.InValues(namespace.NamespacelistenersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.namespace_namespacelisteners
		if fk == nil {
			return fmt.Errorf(`foreign-key "namespace_namespacelisteners" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "namespace_namespacelisteners" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NamespaceQuery) loadCloudeventfilters(ctx context.Context, query *CloudEventFiltersQuery, nodes []*Namespace, init func(*Namespace), assign func(*Namespace, *CloudEventFilters)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Namespace)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CloudEventFilters(func(s *sql.Selector) {
		s.Where(sql.InValues(namespace.CloudeventfiltersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.namespace_cloudeventfilters
		if fk == nil {
			return fmt.Errorf(`foreign-key "namespace_cloudeventfilters" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "namespace_cloudeventfilters" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nq *NamespaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NamespaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(namespace.Table, namespace.Columns, sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeUUID))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, namespace.FieldID)
		for i := range fields {
			if fields[i] != namespace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NamespaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(namespace.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = namespace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range nq.modifiers {
		m(selector)
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (nq *NamespaceQuery) ForUpdate(opts ...sql.LockOption) *NamespaceQuery {
	if nq.driver.Dialect() == dialect.Postgres {
		nq.Unique(false)
	}
	nq.modifiers = append(nq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return nq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (nq *NamespaceQuery) ForShare(opts ...sql.LockOption) *NamespaceQuery {
	if nq.driver.Dialect() == dialect.Postgres {
		nq.Unique(false)
	}
	nq.modifiers = append(nq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return nq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (nq *NamespaceQuery) Modify(modifiers ...func(s *sql.Selector)) *NamespaceSelect {
	nq.modifiers = append(nq.modifiers, modifiers...)
	return nq.Select()
}

// NamespaceGroupBy is the group-by builder for Namespace entities.
type NamespaceGroupBy struct {
	selector
	build *NamespaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NamespaceGroupBy) Aggregate(fns ...AggregateFunc) *NamespaceGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NamespaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NamespaceQuery, *NamespaceGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NamespaceGroupBy) sqlScan(ctx context.Context, root *NamespaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NamespaceSelect is the builder for selecting fields of Namespace entities.
type NamespaceSelect struct {
	*NamespaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NamespaceSelect) Aggregate(fns ...AggregateFunc) *NamespaceSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NamespaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NamespaceQuery, *NamespaceSelect](ctx, ns.NamespaceQuery, ns, ns.inters, v)
}

func (ns *NamespaceSelect) sqlScan(ctx context.Context, root *NamespaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ns *NamespaceSelect) Modify(modifiers ...func(s *sql.Selector)) *NamespaceSelect {
	ns.modifiers = append(ns.modifiers, modifiers...)
	return ns
}
