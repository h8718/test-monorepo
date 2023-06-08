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
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// EventsQuery is the builder for querying Events entities.
type EventsQuery struct {
	config
	ctx              *QueryContext
	order            []OrderFunc
	inters           []Interceptor
	predicates       []predicate.Events
	withWfeventswait *EventsWaitQuery
	withNamespace    *NamespaceQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EventsQuery builder.
func (eq *EventsQuery) Where(ps ...predicate.Events) *EventsQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EventsQuery) Limit(limit int) *EventsQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EventsQuery) Offset(offset int) *EventsQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EventsQuery) Unique(unique bool) *EventsQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EventsQuery) Order(o ...OrderFunc) *EventsQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryWfeventswait chains the current query on the "wfeventswait" edge.
func (eq *EventsQuery) QueryWfeventswait() *EventsWaitQuery {
	query := (&EventsWaitClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(events.Table, events.FieldID, selector),
			sqlgraph.To(eventswait.Table, eventswait.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, events.WfeventswaitTable, events.WfeventswaitColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNamespace chains the current query on the "namespace" edge.
func (eq *EventsQuery) QueryNamespace() *NamespaceQuery {
	query := (&NamespaceClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(events.Table, events.FieldID, selector),
			sqlgraph.To(namespace.Table, namespace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, events.NamespaceTable, events.NamespaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Events entity from the query.
// Returns a *NotFoundError when no Events was found.
func (eq *EventsQuery) First(ctx context.Context) (*Events, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{events.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EventsQuery) FirstX(ctx context.Context) *Events {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Events ID from the query.
// Returns a *NotFoundError when no Events ID was found.
func (eq *EventsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{events.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EventsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Events entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Events entity is found.
// Returns a *NotFoundError when no Events entities are found.
func (eq *EventsQuery) Only(ctx context.Context) (*Events, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{events.Label}
	default:
		return nil, &NotSingularError{events.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EventsQuery) OnlyX(ctx context.Context) *Events {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Events ID in the query.
// Returns a *NotSingularError when more than one Events ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EventsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{events.Label}
	default:
		err = &NotSingularError{events.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EventsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EventsSlice.
func (eq *EventsQuery) All(ctx context.Context) ([]*Events, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Events, *EventsQuery]()
	return withInterceptors[[]*Events](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EventsQuery) AllX(ctx context.Context) []*Events {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Events IDs.
func (eq *EventsQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(events.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EventsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EventsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EventsQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EventsQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EventsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EventsQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EventsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EventsQuery) Clone() *EventsQuery {
	if eq == nil {
		return nil
	}
	return &EventsQuery{
		config:           eq.config,
		ctx:              eq.ctx.Clone(),
		order:            append([]OrderFunc{}, eq.order...),
		inters:           append([]Interceptor{}, eq.inters...),
		predicates:       append([]predicate.Events{}, eq.predicates...),
		withWfeventswait: eq.withWfeventswait.Clone(),
		withNamespace:    eq.withNamespace.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithWfeventswait tells the query-builder to eager-load the nodes that are connected to
// the "wfeventswait" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EventsQuery) WithWfeventswait(opts ...func(*EventsWaitQuery)) *EventsQuery {
	query := (&EventsWaitClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withWfeventswait = query
	return eq
}

// WithNamespace tells the query-builder to eager-load the nodes that are connected to
// the "namespace" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EventsQuery) WithNamespace(opts ...func(*NamespaceQuery)) *EventsQuery {
	query := (&NamespaceClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withNamespace = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Events []map[string]interface {} `json:"events,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Events.Query().
//		GroupBy(events.FieldEvents).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EventsQuery) GroupBy(field string, fields ...string) *EventsGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EventsGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = events.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Events []map[string]interface {} `json:"events,omitempty"`
//	}
//
//	client.Events.Query().
//		Select(events.FieldEvents).
//		Scan(ctx, &v)
func (eq *EventsQuery) Select(fields ...string) *EventsSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EventsSelect{EventsQuery: eq}
	sbuild.label = events.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EventsSelect configured with the given aggregations.
func (eq *EventsQuery) Aggregate(fns ...AggregateFunc) *EventsSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EventsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !events.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EventsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Events, error) {
	var (
		nodes       = []*Events{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [2]bool{
			eq.withWfeventswait != nil,
			eq.withNamespace != nil,
		}
	)
	if eq.withNamespace != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, events.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Events).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Events{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withWfeventswait; query != nil {
		if err := eq.loadWfeventswait(ctx, query, nodes,
			func(n *Events) { n.Edges.Wfeventswait = []*EventsWait{} },
			func(n *Events, e *EventsWait) { n.Edges.Wfeventswait = append(n.Edges.Wfeventswait, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withNamespace; query != nil {
		if err := eq.loadNamespace(ctx, query, nodes, nil,
			func(n *Events, e *Namespace) { n.Edges.Namespace = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EventsQuery) loadWfeventswait(ctx context.Context, query *EventsWaitQuery, nodes []*Events, init func(*Events), assign func(*Events, *EventsWait)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Events)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.EventsWait(func(s *sql.Selector) {
		s.Where(sql.InValues(events.WfeventswaitColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.events_wfeventswait
		if fk == nil {
			return fmt.Errorf(`foreign-key "events_wfeventswait" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "events_wfeventswait" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EventsQuery) loadNamespace(ctx context.Context, query *NamespaceQuery, nodes []*Events, init func(*Events), assign func(*Events, *Namespace)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Events)
	for i := range nodes {
		if nodes[i].namespace_namespacelisteners == nil {
			continue
		}
		fk := *nodes[i].namespace_namespacelisteners
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(namespace.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "namespace_namespacelisteners" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eq *EventsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EventsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(events.Table, events.Columns, sqlgraph.NewFieldSpec(events.FieldID, field.TypeUUID))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, events.FieldID)
		for i := range fields {
			if fields[i] != events.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EventsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(events.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = events.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range eq.modifiers {
		m(selector)
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (eq *EventsQuery) ForUpdate(opts ...sql.LockOption) *EventsQuery {
	if eq.driver.Dialect() == dialect.Postgres {
		eq.Unique(false)
	}
	eq.modifiers = append(eq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return eq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (eq *EventsQuery) ForShare(opts ...sql.LockOption) *EventsQuery {
	if eq.driver.Dialect() == dialect.Postgres {
		eq.Unique(false)
	}
	eq.modifiers = append(eq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return eq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (eq *EventsQuery) Modify(modifiers ...func(s *sql.Selector)) *EventsSelect {
	eq.modifiers = append(eq.modifiers, modifiers...)
	return eq.Select()
}

// EventsGroupBy is the group-by builder for Events entities.
type EventsGroupBy struct {
	selector
	build *EventsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EventsGroupBy) Aggregate(fns ...AggregateFunc) *EventsGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EventsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EventsQuery, *EventsGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EventsGroupBy) sqlScan(ctx context.Context, root *EventsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EventsSelect is the builder for selecting fields of Events entities.
type EventsSelect struct {
	*EventsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EventsSelect) Aggregate(fns ...AggregateFunc) *EventsSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EventsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EventsQuery, *EventsSelect](ctx, es.EventsQuery, es, es.inters, v)
}

func (es *EventsSelect) sqlScan(ctx context.Context, root *EventsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (es *EventsSelect) Modify(modifiers ...func(s *sql.Selector)) *EventsSelect {
	es.modifiers = append(es.modifiers, modifiers...)
	return es
}
