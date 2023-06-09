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
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// LogMsgQuery is the builder for querying LogMsg entities.
type LogMsgQuery struct {
	config
	ctx           *QueryContext
	order         []OrderFunc
	inters        []Interceptor
	predicates    []predicate.LogMsg
	withNamespace *NamespaceQuery
	withFKs       bool
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LogMsgQuery builder.
func (lmq *LogMsgQuery) Where(ps ...predicate.LogMsg) *LogMsgQuery {
	lmq.predicates = append(lmq.predicates, ps...)
	return lmq
}

// Limit the number of records to be returned by this query.
func (lmq *LogMsgQuery) Limit(limit int) *LogMsgQuery {
	lmq.ctx.Limit = &limit
	return lmq
}

// Offset to start from.
func (lmq *LogMsgQuery) Offset(offset int) *LogMsgQuery {
	lmq.ctx.Offset = &offset
	return lmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lmq *LogMsgQuery) Unique(unique bool) *LogMsgQuery {
	lmq.ctx.Unique = &unique
	return lmq
}

// Order specifies how the records should be ordered.
func (lmq *LogMsgQuery) Order(o ...OrderFunc) *LogMsgQuery {
	lmq.order = append(lmq.order, o...)
	return lmq
}

// QueryNamespace chains the current query on the "namespace" edge.
func (lmq *LogMsgQuery) QueryNamespace() *NamespaceQuery {
	query := (&NamespaceClient{config: lmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(logmsg.Table, logmsg.FieldID, selector),
			sqlgraph.To(namespace.Table, namespace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, logmsg.NamespaceTable, logmsg.NamespaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(lmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LogMsg entity from the query.
// Returns a *NotFoundError when no LogMsg was found.
func (lmq *LogMsgQuery) First(ctx context.Context) (*LogMsg, error) {
	nodes, err := lmq.Limit(1).All(setContextOp(ctx, lmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{logmsg.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lmq *LogMsgQuery) FirstX(ctx context.Context) *LogMsg {
	node, err := lmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LogMsg ID from the query.
// Returns a *NotFoundError when no LogMsg ID was found.
func (lmq *LogMsgQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lmq.Limit(1).IDs(setContextOp(ctx, lmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{logmsg.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lmq *LogMsgQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LogMsg entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LogMsg entity is found.
// Returns a *NotFoundError when no LogMsg entities are found.
func (lmq *LogMsgQuery) Only(ctx context.Context) (*LogMsg, error) {
	nodes, err := lmq.Limit(2).All(setContextOp(ctx, lmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{logmsg.Label}
	default:
		return nil, &NotSingularError{logmsg.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lmq *LogMsgQuery) OnlyX(ctx context.Context) *LogMsg {
	node, err := lmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LogMsg ID in the query.
// Returns a *NotSingularError when more than one LogMsg ID is found.
// Returns a *NotFoundError when no entities are found.
func (lmq *LogMsgQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lmq.Limit(2).IDs(setContextOp(ctx, lmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{logmsg.Label}
	default:
		err = &NotSingularError{logmsg.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lmq *LogMsgQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LogMsgs.
func (lmq *LogMsgQuery) All(ctx context.Context) ([]*LogMsg, error) {
	ctx = setContextOp(ctx, lmq.ctx, "All")
	if err := lmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LogMsg, *LogMsgQuery]()
	return withInterceptors[[]*LogMsg](ctx, lmq, qr, lmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lmq *LogMsgQuery) AllX(ctx context.Context) []*LogMsg {
	nodes, err := lmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LogMsg IDs.
func (lmq *LogMsgQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if lmq.ctx.Unique == nil && lmq.path != nil {
		lmq.Unique(true)
	}
	ctx = setContextOp(ctx, lmq.ctx, "IDs")
	if err = lmq.Select(logmsg.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lmq *LogMsgQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lmq *LogMsgQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lmq.ctx, "Count")
	if err := lmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lmq, querierCount[*LogMsgQuery](), lmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lmq *LogMsgQuery) CountX(ctx context.Context) int {
	count, err := lmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lmq *LogMsgQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lmq.ctx, "Exist")
	switch _, err := lmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lmq *LogMsgQuery) ExistX(ctx context.Context) bool {
	exist, err := lmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LogMsgQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lmq *LogMsgQuery) Clone() *LogMsgQuery {
	if lmq == nil {
		return nil
	}
	return &LogMsgQuery{
		config:        lmq.config,
		ctx:           lmq.ctx.Clone(),
		order:         append([]OrderFunc{}, lmq.order...),
		inters:        append([]Interceptor{}, lmq.inters...),
		predicates:    append([]predicate.LogMsg{}, lmq.predicates...),
		withNamespace: lmq.withNamespace.Clone(),
		// clone intermediate query.
		sql:  lmq.sql.Clone(),
		path: lmq.path,
	}
}

// WithNamespace tells the query-builder to eager-load the nodes that are connected to
// the "namespace" edge. The optional arguments are used to configure the query builder of the edge.
func (lmq *LogMsgQuery) WithNamespace(opts ...func(*NamespaceQuery)) *LogMsgQuery {
	query := (&NamespaceClient{config: lmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lmq.withNamespace = query
	return lmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		T time.Time `json:"t,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LogMsg.Query().
//		GroupBy(logmsg.FieldT).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lmq *LogMsgQuery) GroupBy(field string, fields ...string) *LogMsgGroupBy {
	lmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LogMsgGroupBy{build: lmq}
	grbuild.flds = &lmq.ctx.Fields
	grbuild.label = logmsg.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		T time.Time `json:"t,omitempty"`
//	}
//
//	client.LogMsg.Query().
//		Select(logmsg.FieldT).
//		Scan(ctx, &v)
func (lmq *LogMsgQuery) Select(fields ...string) *LogMsgSelect {
	lmq.ctx.Fields = append(lmq.ctx.Fields, fields...)
	sbuild := &LogMsgSelect{LogMsgQuery: lmq}
	sbuild.label = logmsg.Label
	sbuild.flds, sbuild.scan = &lmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LogMsgSelect configured with the given aggregations.
func (lmq *LogMsgQuery) Aggregate(fns ...AggregateFunc) *LogMsgSelect {
	return lmq.Select().Aggregate(fns...)
}

func (lmq *LogMsgQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lmq); err != nil {
				return err
			}
		}
	}
	for _, f := range lmq.ctx.Fields {
		if !logmsg.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lmq.path != nil {
		prev, err := lmq.path(ctx)
		if err != nil {
			return err
		}
		lmq.sql = prev
	}
	return nil
}

func (lmq *LogMsgQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LogMsg, error) {
	var (
		nodes       = []*LogMsg{}
		withFKs     = lmq.withFKs
		_spec       = lmq.querySpec()
		loadedTypes = [1]bool{
			lmq.withNamespace != nil,
		}
	)
	if lmq.withNamespace != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, logmsg.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LogMsg).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LogMsg{config: lmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lmq.modifiers) > 0 {
		_spec.Modifiers = lmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lmq.withNamespace; query != nil {
		if err := lmq.loadNamespace(ctx, query, nodes, nil,
			func(n *LogMsg, e *Namespace) { n.Edges.Namespace = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lmq *LogMsgQuery) loadNamespace(ctx context.Context, query *NamespaceQuery, nodes []*LogMsg, init func(*LogMsg), assign func(*LogMsg, *Namespace)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*LogMsg)
	for i := range nodes {
		if nodes[i].namespace_logs == nil {
			continue
		}
		fk := *nodes[i].namespace_logs
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
			return fmt.Errorf(`unexpected foreign-key "namespace_logs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lmq *LogMsgQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lmq.querySpec()
	if len(lmq.modifiers) > 0 {
		_spec.Modifiers = lmq.modifiers
	}
	_spec.Node.Columns = lmq.ctx.Fields
	if len(lmq.ctx.Fields) > 0 {
		_spec.Unique = lmq.ctx.Unique != nil && *lmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lmq.driver, _spec)
}

func (lmq *LogMsgQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(logmsg.Table, logmsg.Columns, sqlgraph.NewFieldSpec(logmsg.FieldID, field.TypeUUID))
	_spec.From = lmq.sql
	if unique := lmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lmq.path != nil {
		_spec.Unique = true
	}
	if fields := lmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logmsg.FieldID)
		for i := range fields {
			if fields[i] != logmsg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lmq *LogMsgQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lmq.driver.Dialect())
	t1 := builder.Table(logmsg.Table)
	columns := lmq.ctx.Fields
	if len(columns) == 0 {
		columns = logmsg.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lmq.sql != nil {
		selector = lmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lmq.ctx.Unique != nil && *lmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range lmq.modifiers {
		m(selector)
	}
	for _, p := range lmq.predicates {
		p(selector)
	}
	for _, p := range lmq.order {
		p(selector)
	}
	if offset := lmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (lmq *LogMsgQuery) ForUpdate(opts ...sql.LockOption) *LogMsgQuery {
	if lmq.driver.Dialect() == dialect.Postgres {
		lmq.Unique(false)
	}
	lmq.modifiers = append(lmq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return lmq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (lmq *LogMsgQuery) ForShare(opts ...sql.LockOption) *LogMsgQuery {
	if lmq.driver.Dialect() == dialect.Postgres {
		lmq.Unique(false)
	}
	lmq.modifiers = append(lmq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return lmq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lmq *LogMsgQuery) Modify(modifiers ...func(s *sql.Selector)) *LogMsgSelect {
	lmq.modifiers = append(lmq.modifiers, modifiers...)
	return lmq.Select()
}

// LogMsgGroupBy is the group-by builder for LogMsg entities.
type LogMsgGroupBy struct {
	selector
	build *LogMsgQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lmgb *LogMsgGroupBy) Aggregate(fns ...AggregateFunc) *LogMsgGroupBy {
	lmgb.fns = append(lmgb.fns, fns...)
	return lmgb
}

// Scan applies the selector query and scans the result into the given value.
func (lmgb *LogMsgGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lmgb.build.ctx, "GroupBy")
	if err := lmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogMsgQuery, *LogMsgGroupBy](ctx, lmgb.build, lmgb, lmgb.build.inters, v)
}

func (lmgb *LogMsgGroupBy) sqlScan(ctx context.Context, root *LogMsgQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lmgb.fns))
	for _, fn := range lmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lmgb.flds)+len(lmgb.fns))
		for _, f := range *lmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LogMsgSelect is the builder for selecting fields of LogMsg entities.
type LogMsgSelect struct {
	*LogMsgQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lms *LogMsgSelect) Aggregate(fns ...AggregateFunc) *LogMsgSelect {
	lms.fns = append(lms.fns, fns...)
	return lms
}

// Scan applies the selector query and scans the result into the given value.
func (lms *LogMsgSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lms.ctx, "Select")
	if err := lms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogMsgQuery, *LogMsgSelect](ctx, lms.LogMsgQuery, lms, lms.inters, v)
}

func (lms *LogMsgSelect) sqlScan(ctx context.Context, root *LogMsgQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lms.fns))
	for _, fn := range lms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lms *LogMsgSelect) Modify(modifiers ...func(s *sql.Selector)) *LogMsgSelect {
	lms.modifiers = append(lms.modifiers, modifiers...)
	return lms
}
