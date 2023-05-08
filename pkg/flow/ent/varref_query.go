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
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/vardata"
	"github.com/direktiv/direktiv/pkg/flow/ent/varref"
	"github.com/google/uuid"
)

// VarRefQuery is the builder for querying VarRef entities.
type VarRefQuery struct {
	config
	ctx           *QueryContext
	order         []OrderFunc
	inters        []Interceptor
	predicates    []predicate.VarRef
	withVardata   *VarDataQuery
	withNamespace *NamespaceQuery
	withInstance  *InstanceQuery
	withFKs       bool
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VarRefQuery builder.
func (vrq *VarRefQuery) Where(ps ...predicate.VarRef) *VarRefQuery {
	vrq.predicates = append(vrq.predicates, ps...)
	return vrq
}

// Limit the number of records to be returned by this query.
func (vrq *VarRefQuery) Limit(limit int) *VarRefQuery {
	vrq.ctx.Limit = &limit
	return vrq
}

// Offset to start from.
func (vrq *VarRefQuery) Offset(offset int) *VarRefQuery {
	vrq.ctx.Offset = &offset
	return vrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vrq *VarRefQuery) Unique(unique bool) *VarRefQuery {
	vrq.ctx.Unique = &unique
	return vrq
}

// Order specifies how the records should be ordered.
func (vrq *VarRefQuery) Order(o ...OrderFunc) *VarRefQuery {
	vrq.order = append(vrq.order, o...)
	return vrq
}

// QueryVardata chains the current query on the "vardata" edge.
func (vrq *VarRefQuery) QueryVardata() *VarDataQuery {
	query := (&VarDataClient{config: vrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(varref.Table, varref.FieldID, selector),
			sqlgraph.To(vardata.Table, vardata.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, varref.VardataTable, varref.VardataColumn),
		)
		fromU = sqlgraph.SetNeighbors(vrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNamespace chains the current query on the "namespace" edge.
func (vrq *VarRefQuery) QueryNamespace() *NamespaceQuery {
	query := (&NamespaceClient{config: vrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(varref.Table, varref.FieldID, selector),
			sqlgraph.To(namespace.Table, namespace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, varref.NamespaceTable, varref.NamespaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(vrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInstance chains the current query on the "instance" edge.
func (vrq *VarRefQuery) QueryInstance() *InstanceQuery {
	query := (&InstanceClient{config: vrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(varref.Table, varref.FieldID, selector),
			sqlgraph.To(instance.Table, instance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, varref.InstanceTable, varref.InstanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(vrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VarRef entity from the query.
// Returns a *NotFoundError when no VarRef was found.
func (vrq *VarRefQuery) First(ctx context.Context) (*VarRef, error) {
	nodes, err := vrq.Limit(1).All(setContextOp(ctx, vrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{varref.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vrq *VarRefQuery) FirstX(ctx context.Context) *VarRef {
	node, err := vrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VarRef ID from the query.
// Returns a *NotFoundError when no VarRef ID was found.
func (vrq *VarRefQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vrq.Limit(1).IDs(setContextOp(ctx, vrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{varref.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vrq *VarRefQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VarRef entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VarRef entity is found.
// Returns a *NotFoundError when no VarRef entities are found.
func (vrq *VarRefQuery) Only(ctx context.Context) (*VarRef, error) {
	nodes, err := vrq.Limit(2).All(setContextOp(ctx, vrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{varref.Label}
	default:
		return nil, &NotSingularError{varref.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vrq *VarRefQuery) OnlyX(ctx context.Context) *VarRef {
	node, err := vrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VarRef ID in the query.
// Returns a *NotSingularError when more than one VarRef ID is found.
// Returns a *NotFoundError when no entities are found.
func (vrq *VarRefQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vrq.Limit(2).IDs(setContextOp(ctx, vrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{varref.Label}
	default:
		err = &NotSingularError{varref.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vrq *VarRefQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VarRefs.
func (vrq *VarRefQuery) All(ctx context.Context) ([]*VarRef, error) {
	ctx = setContextOp(ctx, vrq.ctx, "All")
	if err := vrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VarRef, *VarRefQuery]()
	return withInterceptors[[]*VarRef](ctx, vrq, qr, vrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vrq *VarRefQuery) AllX(ctx context.Context) []*VarRef {
	nodes, err := vrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VarRef IDs.
func (vrq *VarRefQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if vrq.ctx.Unique == nil && vrq.path != nil {
		vrq.Unique(true)
	}
	ctx = setContextOp(ctx, vrq.ctx, "IDs")
	if err = vrq.Select(varref.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vrq *VarRefQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vrq *VarRefQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vrq.ctx, "Count")
	if err := vrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vrq, querierCount[*VarRefQuery](), vrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vrq *VarRefQuery) CountX(ctx context.Context) int {
	count, err := vrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vrq *VarRefQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vrq.ctx, "Exist")
	switch _, err := vrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vrq *VarRefQuery) ExistX(ctx context.Context) bool {
	exist, err := vrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VarRefQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vrq *VarRefQuery) Clone() *VarRefQuery {
	if vrq == nil {
		return nil
	}
	return &VarRefQuery{
		config:        vrq.config,
		ctx:           vrq.ctx.Clone(),
		order:         append([]OrderFunc{}, vrq.order...),
		inters:        append([]Interceptor{}, vrq.inters...),
		predicates:    append([]predicate.VarRef{}, vrq.predicates...),
		withVardata:   vrq.withVardata.Clone(),
		withNamespace: vrq.withNamespace.Clone(),
		withInstance:  vrq.withInstance.Clone(),
		// clone intermediate query.
		sql:  vrq.sql.Clone(),
		path: vrq.path,
	}
}

// WithVardata tells the query-builder to eager-load the nodes that are connected to
// the "vardata" edge. The optional arguments are used to configure the query builder of the edge.
func (vrq *VarRefQuery) WithVardata(opts ...func(*VarDataQuery)) *VarRefQuery {
	query := (&VarDataClient{config: vrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vrq.withVardata = query
	return vrq
}

// WithNamespace tells the query-builder to eager-load the nodes that are connected to
// the "namespace" edge. The optional arguments are used to configure the query builder of the edge.
func (vrq *VarRefQuery) WithNamespace(opts ...func(*NamespaceQuery)) *VarRefQuery {
	query := (&NamespaceClient{config: vrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vrq.withNamespace = query
	return vrq
}

// WithInstance tells the query-builder to eager-load the nodes that are connected to
// the "instance" edge. The optional arguments are used to configure the query builder of the edge.
func (vrq *VarRefQuery) WithInstance(opts ...func(*InstanceQuery)) *VarRefQuery {
	query := (&InstanceClient{config: vrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vrq.withInstance = query
	return vrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VarRef.Query().
//		GroupBy(varref.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vrq *VarRefQuery) GroupBy(field string, fields ...string) *VarRefGroupBy {
	vrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VarRefGroupBy{build: vrq}
	grbuild.flds = &vrq.ctx.Fields
	grbuild.label = varref.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.VarRef.Query().
//		Select(varref.FieldName).
//		Scan(ctx, &v)
func (vrq *VarRefQuery) Select(fields ...string) *VarRefSelect {
	vrq.ctx.Fields = append(vrq.ctx.Fields, fields...)
	sbuild := &VarRefSelect{VarRefQuery: vrq}
	sbuild.label = varref.Label
	sbuild.flds, sbuild.scan = &vrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VarRefSelect configured with the given aggregations.
func (vrq *VarRefQuery) Aggregate(fns ...AggregateFunc) *VarRefSelect {
	return vrq.Select().Aggregate(fns...)
}

func (vrq *VarRefQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vrq); err != nil {
				return err
			}
		}
	}
	for _, f := range vrq.ctx.Fields {
		if !varref.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vrq.path != nil {
		prev, err := vrq.path(ctx)
		if err != nil {
			return err
		}
		vrq.sql = prev
	}
	return nil
}

func (vrq *VarRefQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VarRef, error) {
	var (
		nodes       = []*VarRef{}
		withFKs     = vrq.withFKs
		_spec       = vrq.querySpec()
		loadedTypes = [3]bool{
			vrq.withVardata != nil,
			vrq.withNamespace != nil,
			vrq.withInstance != nil,
		}
	)
	if vrq.withVardata != nil || vrq.withNamespace != nil || vrq.withInstance != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, varref.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VarRef).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VarRef{config: vrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vrq.modifiers) > 0 {
		_spec.Modifiers = vrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vrq.withVardata; query != nil {
		if err := vrq.loadVardata(ctx, query, nodes, nil,
			func(n *VarRef, e *VarData) { n.Edges.Vardata = e }); err != nil {
			return nil, err
		}
	}
	if query := vrq.withNamespace; query != nil {
		if err := vrq.loadNamespace(ctx, query, nodes, nil,
			func(n *VarRef, e *Namespace) { n.Edges.Namespace = e }); err != nil {
			return nil, err
		}
	}
	if query := vrq.withInstance; query != nil {
		if err := vrq.loadInstance(ctx, query, nodes, nil,
			func(n *VarRef, e *Instance) { n.Edges.Instance = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vrq *VarRefQuery) loadVardata(ctx context.Context, query *VarDataQuery, nodes []*VarRef, init func(*VarRef), assign func(*VarRef, *VarData)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*VarRef)
	for i := range nodes {
		if nodes[i].var_data_varrefs == nil {
			continue
		}
		fk := *nodes[i].var_data_varrefs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(vardata.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "var_data_varrefs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vrq *VarRefQuery) loadNamespace(ctx context.Context, query *NamespaceQuery, nodes []*VarRef, init func(*VarRef), assign func(*VarRef, *Namespace)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*VarRef)
	for i := range nodes {
		if nodes[i].namespace_vars == nil {
			continue
		}
		fk := *nodes[i].namespace_vars
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
			return fmt.Errorf(`unexpected foreign-key "namespace_vars" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vrq *VarRefQuery) loadInstance(ctx context.Context, query *InstanceQuery, nodes []*VarRef, init func(*VarRef), assign func(*VarRef, *Instance)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*VarRef)
	for i := range nodes {
		if nodes[i].instance_vars == nil {
			continue
		}
		fk := *nodes[i].instance_vars
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
			return fmt.Errorf(`unexpected foreign-key "instance_vars" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vrq *VarRefQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vrq.querySpec()
	if len(vrq.modifiers) > 0 {
		_spec.Modifiers = vrq.modifiers
	}
	_spec.Node.Columns = vrq.ctx.Fields
	if len(vrq.ctx.Fields) > 0 {
		_spec.Unique = vrq.ctx.Unique != nil && *vrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vrq.driver, _spec)
}

func (vrq *VarRefQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(varref.Table, varref.Columns, sqlgraph.NewFieldSpec(varref.FieldID, field.TypeUUID))
	_spec.From = vrq.sql
	if unique := vrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vrq.path != nil {
		_spec.Unique = true
	}
	if fields := vrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, varref.FieldID)
		for i := range fields {
			if fields[i] != varref.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vrq *VarRefQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vrq.driver.Dialect())
	t1 := builder.Table(varref.Table)
	columns := vrq.ctx.Fields
	if len(columns) == 0 {
		columns = varref.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vrq.sql != nil {
		selector = vrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vrq.ctx.Unique != nil && *vrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range vrq.modifiers {
		m(selector)
	}
	for _, p := range vrq.predicates {
		p(selector)
	}
	for _, p := range vrq.order {
		p(selector)
	}
	if offset := vrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (vrq *VarRefQuery) ForUpdate(opts ...sql.LockOption) *VarRefQuery {
	if vrq.driver.Dialect() == dialect.Postgres {
		vrq.Unique(false)
	}
	vrq.modifiers = append(vrq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return vrq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (vrq *VarRefQuery) ForShare(opts ...sql.LockOption) *VarRefQuery {
	if vrq.driver.Dialect() == dialect.Postgres {
		vrq.Unique(false)
	}
	vrq.modifiers = append(vrq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return vrq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vrq *VarRefQuery) Modify(modifiers ...func(s *sql.Selector)) *VarRefSelect {
	vrq.modifiers = append(vrq.modifiers, modifiers...)
	return vrq.Select()
}

// VarRefGroupBy is the group-by builder for VarRef entities.
type VarRefGroupBy struct {
	selector
	build *VarRefQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vrgb *VarRefGroupBy) Aggregate(fns ...AggregateFunc) *VarRefGroupBy {
	vrgb.fns = append(vrgb.fns, fns...)
	return vrgb
}

// Scan applies the selector query and scans the result into the given value.
func (vrgb *VarRefGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vrgb.build.ctx, "GroupBy")
	if err := vrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VarRefQuery, *VarRefGroupBy](ctx, vrgb.build, vrgb, vrgb.build.inters, v)
}

func (vrgb *VarRefGroupBy) sqlScan(ctx context.Context, root *VarRefQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vrgb.fns))
	for _, fn := range vrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vrgb.flds)+len(vrgb.fns))
		for _, f := range *vrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VarRefSelect is the builder for selecting fields of VarRef entities.
type VarRefSelect struct {
	*VarRefQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vrs *VarRefSelect) Aggregate(fns ...AggregateFunc) *VarRefSelect {
	vrs.fns = append(vrs.fns, fns...)
	return vrs
}

// Scan applies the selector query and scans the result into the given value.
func (vrs *VarRefSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vrs.ctx, "Select")
	if err := vrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VarRefQuery, *VarRefSelect](ctx, vrs.VarRefQuery, vrs, vrs.inters, v)
}

func (vrs *VarRefSelect) sqlScan(ctx context.Context, root *VarRefQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vrs.fns))
	for _, fn := range vrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vrs *VarRefSelect) Modify(modifiers ...func(s *sql.Selector)) *VarRefSelect {
	vrs.modifiers = append(vrs.modifiers, modifiers...)
	return vrs
}
