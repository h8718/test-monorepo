// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/instanceruntime"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// InstanceRuntimeQuery is the builder for querying InstanceRuntime entities.
type InstanceRuntimeQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.InstanceRuntime
	withInstance *InstanceQuery
	withCaller   *InstanceQuery
	withFKs      bool
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InstanceRuntimeQuery builder.
func (irq *InstanceRuntimeQuery) Where(ps ...predicate.InstanceRuntime) *InstanceRuntimeQuery {
	irq.predicates = append(irq.predicates, ps...)
	return irq
}

// Limit the number of records to be returned by this query.
func (irq *InstanceRuntimeQuery) Limit(limit int) *InstanceRuntimeQuery {
	irq.ctx.Limit = &limit
	return irq
}

// Offset to start from.
func (irq *InstanceRuntimeQuery) Offset(offset int) *InstanceRuntimeQuery {
	irq.ctx.Offset = &offset
	return irq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (irq *InstanceRuntimeQuery) Unique(unique bool) *InstanceRuntimeQuery {
	irq.ctx.Unique = &unique
	return irq
}

// Order specifies how the records should be ordered.
func (irq *InstanceRuntimeQuery) Order(o ...OrderFunc) *InstanceRuntimeQuery {
	irq.order = append(irq.order, o...)
	return irq
}

// QueryInstance chains the current query on the "instance" edge.
func (irq *InstanceRuntimeQuery) QueryInstance() *InstanceQuery {
	query := (&InstanceClient{config: irq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := irq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := irq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instanceruntime.Table, instanceruntime.FieldID, selector),
			sqlgraph.To(instance.Table, instance.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, instanceruntime.InstanceTable, instanceruntime.InstanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(irq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCaller chains the current query on the "caller" edge.
func (irq *InstanceRuntimeQuery) QueryCaller() *InstanceQuery {
	query := (&InstanceClient{config: irq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := irq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := irq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instanceruntime.Table, instanceruntime.FieldID, selector),
			sqlgraph.To(instance.Table, instance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instanceruntime.CallerTable, instanceruntime.CallerColumn),
		)
		fromU = sqlgraph.SetNeighbors(irq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InstanceRuntime entity from the query.
// Returns a *NotFoundError when no InstanceRuntime was found.
func (irq *InstanceRuntimeQuery) First(ctx context.Context) (*InstanceRuntime, error) {
	nodes, err := irq.Limit(1).All(setContextOp(ctx, irq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{instanceruntime.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (irq *InstanceRuntimeQuery) FirstX(ctx context.Context) *InstanceRuntime {
	node, err := irq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InstanceRuntime ID from the query.
// Returns a *NotFoundError when no InstanceRuntime ID was found.
func (irq *InstanceRuntimeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = irq.Limit(1).IDs(setContextOp(ctx, irq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{instanceruntime.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (irq *InstanceRuntimeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := irq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InstanceRuntime entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InstanceRuntime entity is found.
// Returns a *NotFoundError when no InstanceRuntime entities are found.
func (irq *InstanceRuntimeQuery) Only(ctx context.Context) (*InstanceRuntime, error) {
	nodes, err := irq.Limit(2).All(setContextOp(ctx, irq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{instanceruntime.Label}
	default:
		return nil, &NotSingularError{instanceruntime.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (irq *InstanceRuntimeQuery) OnlyX(ctx context.Context) *InstanceRuntime {
	node, err := irq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InstanceRuntime ID in the query.
// Returns a *NotSingularError when more than one InstanceRuntime ID is found.
// Returns a *NotFoundError when no entities are found.
func (irq *InstanceRuntimeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = irq.Limit(2).IDs(setContextOp(ctx, irq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{instanceruntime.Label}
	default:
		err = &NotSingularError{instanceruntime.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (irq *InstanceRuntimeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := irq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InstanceRuntimes.
func (irq *InstanceRuntimeQuery) All(ctx context.Context) ([]*InstanceRuntime, error) {
	ctx = setContextOp(ctx, irq.ctx, "All")
	if err := irq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InstanceRuntime, *InstanceRuntimeQuery]()
	return withInterceptors[[]*InstanceRuntime](ctx, irq, qr, irq.inters)
}

// AllX is like All, but panics if an error occurs.
func (irq *InstanceRuntimeQuery) AllX(ctx context.Context) []*InstanceRuntime {
	nodes, err := irq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InstanceRuntime IDs.
func (irq *InstanceRuntimeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if irq.ctx.Unique == nil && irq.path != nil {
		irq.Unique(true)
	}
	ctx = setContextOp(ctx, irq.ctx, "IDs")
	if err = irq.Select(instanceruntime.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (irq *InstanceRuntimeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := irq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (irq *InstanceRuntimeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, irq.ctx, "Count")
	if err := irq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, irq, querierCount[*InstanceRuntimeQuery](), irq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (irq *InstanceRuntimeQuery) CountX(ctx context.Context) int {
	count, err := irq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (irq *InstanceRuntimeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, irq.ctx, "Exist")
	switch _, err := irq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (irq *InstanceRuntimeQuery) ExistX(ctx context.Context) bool {
	exist, err := irq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InstanceRuntimeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (irq *InstanceRuntimeQuery) Clone() *InstanceRuntimeQuery {
	if irq == nil {
		return nil
	}
	return &InstanceRuntimeQuery{
		config:       irq.config,
		ctx:          irq.ctx.Clone(),
		order:        append([]OrderFunc{}, irq.order...),
		inters:       append([]Interceptor{}, irq.inters...),
		predicates:   append([]predicate.InstanceRuntime{}, irq.predicates...),
		withInstance: irq.withInstance.Clone(),
		withCaller:   irq.withCaller.Clone(),
		// clone intermediate query.
		sql:  irq.sql.Clone(),
		path: irq.path,
	}
}

// WithInstance tells the query-builder to eager-load the nodes that are connected to
// the "instance" edge. The optional arguments are used to configure the query builder of the edge.
func (irq *InstanceRuntimeQuery) WithInstance(opts ...func(*InstanceQuery)) *InstanceRuntimeQuery {
	query := (&InstanceClient{config: irq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	irq.withInstance = query
	return irq
}

// WithCaller tells the query-builder to eager-load the nodes that are connected to
// the "caller" edge. The optional arguments are used to configure the query builder of the edge.
func (irq *InstanceRuntimeQuery) WithCaller(opts ...func(*InstanceQuery)) *InstanceRuntimeQuery {
	query := (&InstanceClient{config: irq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	irq.withCaller = query
	return irq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Input []byte `json:"input,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InstanceRuntime.Query().
//		GroupBy(instanceruntime.FieldInput).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (irq *InstanceRuntimeQuery) GroupBy(field string, fields ...string) *InstanceRuntimeGroupBy {
	irq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InstanceRuntimeGroupBy{build: irq}
	grbuild.flds = &irq.ctx.Fields
	grbuild.label = instanceruntime.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Input []byte `json:"input,omitempty"`
//	}
//
//	client.InstanceRuntime.Query().
//		Select(instanceruntime.FieldInput).
//		Scan(ctx, &v)
func (irq *InstanceRuntimeQuery) Select(fields ...string) *InstanceRuntimeSelect {
	irq.ctx.Fields = append(irq.ctx.Fields, fields...)
	sbuild := &InstanceRuntimeSelect{InstanceRuntimeQuery: irq}
	sbuild.label = instanceruntime.Label
	sbuild.flds, sbuild.scan = &irq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InstanceRuntimeSelect configured with the given aggregations.
func (irq *InstanceRuntimeQuery) Aggregate(fns ...AggregateFunc) *InstanceRuntimeSelect {
	return irq.Select().Aggregate(fns...)
}

func (irq *InstanceRuntimeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range irq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, irq); err != nil {
				return err
			}
		}
	}
	for _, f := range irq.ctx.Fields {
		if !instanceruntime.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if irq.path != nil {
		prev, err := irq.path(ctx)
		if err != nil {
			return err
		}
		irq.sql = prev
	}
	return nil
}

func (irq *InstanceRuntimeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InstanceRuntime, error) {
	var (
		nodes       = []*InstanceRuntime{}
		withFKs     = irq.withFKs
		_spec       = irq.querySpec()
		loadedTypes = [2]bool{
			irq.withInstance != nil,
			irq.withCaller != nil,
		}
	)
	if irq.withInstance != nil || irq.withCaller != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, instanceruntime.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InstanceRuntime).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InstanceRuntime{config: irq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(irq.modifiers) > 0 {
		_spec.Modifiers = irq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, irq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := irq.withInstance; query != nil {
		if err := irq.loadInstance(ctx, query, nodes, nil,
			func(n *InstanceRuntime, e *Instance) { n.Edges.Instance = e }); err != nil {
			return nil, err
		}
	}
	if query := irq.withCaller; query != nil {
		if err := irq.loadCaller(ctx, query, nodes, nil,
			func(n *InstanceRuntime, e *Instance) { n.Edges.Caller = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (irq *InstanceRuntimeQuery) loadInstance(ctx context.Context, query *InstanceQuery, nodes []*InstanceRuntime, init func(*InstanceRuntime), assign func(*InstanceRuntime, *Instance)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*InstanceRuntime)
	for i := range nodes {
		if nodes[i].instance_runtime == nil {
			continue
		}
		fk := *nodes[i].instance_runtime
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(instance.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "instance_runtime" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (irq *InstanceRuntimeQuery) loadCaller(ctx context.Context, query *InstanceQuery, nodes []*InstanceRuntime, init func(*InstanceRuntime), assign func(*InstanceRuntime, *Instance)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*InstanceRuntime)
	for i := range nodes {
		if nodes[i].instance_children == nil {
			continue
		}
		fk := *nodes[i].instance_children
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(instance.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "instance_children" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (irq *InstanceRuntimeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := irq.querySpec()
	if len(irq.modifiers) > 0 {
		_spec.Modifiers = irq.modifiers
	}
	_spec.Node.Columns = irq.ctx.Fields
	if len(irq.ctx.Fields) > 0 {
		_spec.Unique = irq.ctx.Unique != nil && *irq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, irq.driver, _spec)
}

func (irq *InstanceRuntimeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(instanceruntime.Table, instanceruntime.Columns, sqlgraph.NewFieldSpec(instanceruntime.FieldID, field.TypeUUID))
	_spec.From = irq.sql
	if unique := irq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if irq.path != nil {
		_spec.Unique = true
	}
	if fields := irq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instanceruntime.FieldID)
		for i := range fields {
			if fields[i] != instanceruntime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := irq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := irq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := irq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := irq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (irq *InstanceRuntimeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(irq.driver.Dialect())
	t1 := builder.Table(instanceruntime.Table)
	columns := irq.ctx.Fields
	if len(columns) == 0 {
		columns = instanceruntime.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if irq.sql != nil {
		selector = irq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if irq.ctx.Unique != nil && *irq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range irq.modifiers {
		m(selector)
	}
	for _, p := range irq.predicates {
		p(selector)
	}
	for _, p := range irq.order {
		p(selector)
	}
	if offset := irq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := irq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (irq *InstanceRuntimeQuery) ForUpdate(opts ...sql.LockOption) *InstanceRuntimeQuery {
	if irq.driver.Dialect() == dialect.Postgres {
		irq.Unique(false)
	}
	irq.modifiers = append(irq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return irq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (irq *InstanceRuntimeQuery) ForShare(opts ...sql.LockOption) *InstanceRuntimeQuery {
	if irq.driver.Dialect() == dialect.Postgres {
		irq.Unique(false)
	}
	irq.modifiers = append(irq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return irq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (irq *InstanceRuntimeQuery) Modify(modifiers ...func(s *sql.Selector)) *InstanceRuntimeSelect {
	irq.modifiers = append(irq.modifiers, modifiers...)
	return irq.Select()
}

// InstanceRuntimeGroupBy is the group-by builder for InstanceRuntime entities.
type InstanceRuntimeGroupBy struct {
	selector
	build *InstanceRuntimeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (irgb *InstanceRuntimeGroupBy) Aggregate(fns ...AggregateFunc) *InstanceRuntimeGroupBy {
	irgb.fns = append(irgb.fns, fns...)
	return irgb
}

// Scan applies the selector query and scans the result into the given value.
func (irgb *InstanceRuntimeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, irgb.build.ctx, "GroupBy")
	if err := irgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InstanceRuntimeQuery, *InstanceRuntimeGroupBy](ctx, irgb.build, irgb, irgb.build.inters, v)
}

func (irgb *InstanceRuntimeGroupBy) sqlScan(ctx context.Context, root *InstanceRuntimeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(irgb.fns))
	for _, fn := range irgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*irgb.flds)+len(irgb.fns))
		for _, f := range *irgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*irgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := irgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InstanceRuntimeSelect is the builder for selecting fields of InstanceRuntime entities.
type InstanceRuntimeSelect struct {
	*InstanceRuntimeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (irs *InstanceRuntimeSelect) Aggregate(fns ...AggregateFunc) *InstanceRuntimeSelect {
	irs.fns = append(irs.fns, fns...)
	return irs
}

// Scan applies the selector query and scans the result into the given value.
func (irs *InstanceRuntimeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, irs.ctx, "Select")
	if err := irs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InstanceRuntimeQuery, *InstanceRuntimeSelect](ctx, irs.InstanceRuntimeQuery, irs, irs.inters, v)
}

func (irs *InstanceRuntimeSelect) sqlScan(ctx context.Context, root *InstanceRuntimeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(irs.fns))
	for _, fn := range irs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*irs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := irs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (irs *InstanceRuntimeSelect) Modify(modifiers ...func(s *sql.Selector)) *InstanceRuntimeSelect {
	irs.modifiers = append(irs.modifiers, modifiers...)
	return irs
}
