// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_models/ent/predicate"
	"_models/ent/queue"
	"_models/ent/user"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// QueueQuery is the builder for querying Queue entities.
type QueueQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Queue
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QueueQuery builder.
func (qq *QueueQuery) Where(ps ...predicate.Queue) *QueueQuery {
	qq.predicates = append(qq.predicates, ps...)
	return qq
}

// Limit adds a limit step to the query.
func (qq *QueueQuery) Limit(limit int) *QueueQuery {
	qq.limit = &limit
	return qq
}

// Offset adds an offset step to the query.
func (qq *QueueQuery) Offset(offset int) *QueueQuery {
	qq.offset = &offset
	return qq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qq *QueueQuery) Unique(unique bool) *QueueQuery {
	qq.unique = &unique
	return qq
}

// Order adds an order step to the query.
func (qq *QueueQuery) Order(o ...OrderFunc) *QueueQuery {
	qq.order = append(qq.order, o...)
	return qq
}

// QueryUser chains the current query on the "user" edge.
func (qq *QueueQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: qq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(queue.Table, queue.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, queue.UserTable, queue.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Queue entity from the query.
// Returns a *NotFoundError when no Queue was found.
func (qq *QueueQuery) First(ctx context.Context) (*Queue, error) {
	nodes, err := qq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{queue.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qq *QueueQuery) FirstX(ctx context.Context) *Queue {
	node, err := qq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Queue ID from the query.
// Returns a *NotFoundError when no Queue ID was found.
func (qq *QueueQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = qq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{queue.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qq *QueueQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := qq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Queue entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Queue entity is found.
// Returns a *NotFoundError when no Queue entities are found.
func (qq *QueueQuery) Only(ctx context.Context) (*Queue, error) {
	nodes, err := qq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{queue.Label}
	default:
		return nil, &NotSingularError{queue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qq *QueueQuery) OnlyX(ctx context.Context) *Queue {
	node, err := qq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Queue ID in the query.
// Returns a *NotSingularError when more than one Queue ID is found.
// Returns a *NotFoundError when no entities are found.
func (qq *QueueQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = qq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{queue.Label}
	default:
		err = &NotSingularError{queue.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qq *QueueQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := qq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Queues.
func (qq *QueueQuery) All(ctx context.Context) ([]*Queue, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return qq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (qq *QueueQuery) AllX(ctx context.Context) []*Queue {
	nodes, err := qq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Queue IDs.
func (qq *QueueQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := qq.Select(queue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qq *QueueQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := qq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qq *QueueQuery) Count(ctx context.Context) (int, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return qq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (qq *QueueQuery) CountX(ctx context.Context) int {
	count, err := qq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qq *QueueQuery) Exist(ctx context.Context) (bool, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return qq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (qq *QueueQuery) ExistX(ctx context.Context) bool {
	exist, err := qq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QueueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qq *QueueQuery) Clone() *QueueQuery {
	if qq == nil {
		return nil
	}
	return &QueueQuery{
		config:     qq.config,
		limit:      qq.limit,
		offset:     qq.offset,
		order:      append([]OrderFunc{}, qq.order...),
		predicates: append([]predicate.Queue{}, qq.predicates...),
		withUser:   qq.withUser.Clone(),
		// clone intermediate query.
		sql:    qq.sql.Clone(),
		path:   qq.path,
		unique: qq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (qq *QueueQuery) WithUser(opts ...func(*UserQuery)) *QueueQuery {
	query := &UserQuery{config: qq.config}
	for _, opt := range opts {
		opt(query)
	}
	qq.withUser = query
	return qq
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
//	client.Queue.Query().
//		GroupBy(queue.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (qq *QueueQuery) GroupBy(field string, fields ...string) *QueueGroupBy {
	grbuild := &QueueGroupBy{config: qq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qq.sqlQuery(ctx), nil
	}
	grbuild.label = queue.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.Queue.Query().
//		Select(queue.FieldName).
//		Scan(ctx, &v)
func (qq *QueueQuery) Select(fields ...string) *QueueSelect {
	qq.fields = append(qq.fields, fields...)
	selbuild := &QueueSelect{QueueQuery: qq}
	selbuild.label = queue.Label
	selbuild.flds, selbuild.scan = &qq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a QueueSelect configured with the given aggregations.
func (qq *QueueQuery) Aggregate(fns ...AggregateFunc) *QueueSelect {
	return qq.Select().Aggregate(fns...)
}

func (qq *QueueQuery) prepareQuery(ctx context.Context) error {
	for _, f := range qq.fields {
		if !queue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qq.path != nil {
		prev, err := qq.path(ctx)
		if err != nil {
			return err
		}
		qq.sql = prev
	}
	return nil
}

func (qq *QueueQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Queue, error) {
	var (
		nodes       = []*Queue{}
		_spec       = qq.querySpec()
		loadedTypes = [1]bool{
			qq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Queue).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Queue{config: qq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := qq.withUser; query != nil {
		if err := qq.loadUser(ctx, query, nodes,
			func(n *Queue) { n.Edges.User = []*User{} },
			func(n *Queue, e *User) { n.Edges.User = append(n.Edges.User, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (qq *QueueQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Queue, init func(*Queue), assign func(*Queue, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Queue)
	nids := make(map[uuid.UUID]map[*Queue]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(queue.UserTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(queue.UserPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(queue.UserPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(queue.UserPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := *values[1].(*uuid.UUID)
			if nids[inValue] == nil {
				nids[inValue] = map[*Queue]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (qq *QueueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qq.querySpec()
	_spec.Node.Columns = qq.fields
	if len(qq.fields) > 0 {
		_spec.Unique = qq.unique != nil && *qq.unique
	}
	return sqlgraph.CountNodes(ctx, qq.driver, _spec)
}

func (qq *QueueQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := qq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (qq *QueueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   queue.Table,
			Columns: queue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: queue.FieldID,
			},
		},
		From:   qq.sql,
		Unique: true,
	}
	if unique := qq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := qq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, queue.FieldID)
		for i := range fields {
			if fields[i] != queue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qq *QueueQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qq.driver.Dialect())
	t1 := builder.Table(queue.Table)
	columns := qq.fields
	if len(columns) == 0 {
		columns = queue.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qq.sql != nil {
		selector = qq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qq.unique != nil && *qq.unique {
		selector.Distinct()
	}
	for _, p := range qq.predicates {
		p(selector)
	}
	for _, p := range qq.order {
		p(selector)
	}
	if offset := qq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QueueGroupBy is the group-by builder for Queue entities.
type QueueGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qgb *QueueGroupBy) Aggregate(fns ...AggregateFunc) *QueueGroupBy {
	qgb.fns = append(qgb.fns, fns...)
	return qgb
}

// Scan applies the group-by query and scans the result into the given value.
func (qgb *QueueGroupBy) Scan(ctx context.Context, v any) error {
	query, err := qgb.path(ctx)
	if err != nil {
		return err
	}
	qgb.sql = query
	return qgb.sqlScan(ctx, v)
}

func (qgb *QueueGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range qgb.fields {
		if !queue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := qgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qgb *QueueGroupBy) sqlQuery() *sql.Selector {
	selector := qgb.sql.Select()
	aggregation := make([]string, 0, len(qgb.fns))
	for _, fn := range qgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(qgb.fields)+len(qgb.fns))
		for _, f := range qgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(qgb.fields...)...)
}

// QueueSelect is the builder for selecting fields of Queue entities.
type QueueSelect struct {
	*QueueQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qs *QueueSelect) Aggregate(fns ...AggregateFunc) *QueueSelect {
	qs.fns = append(qs.fns, fns...)
	return qs
}

// Scan applies the selector query and scans the result into the given value.
func (qs *QueueSelect) Scan(ctx context.Context, v any) error {
	if err := qs.prepareQuery(ctx); err != nil {
		return err
	}
	qs.sql = qs.QueueQuery.sqlQuery(ctx)
	return qs.sqlScan(ctx, v)
}

func (qs *QueueSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(qs.fns))
	for _, fn := range qs.fns {
		aggregation = append(aggregation, fn(qs.sql))
	}
	switch n := len(*qs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		qs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		qs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := qs.sql.Query()
	if err := qs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}