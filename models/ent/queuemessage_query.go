// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_models/ent/predicate"
	"_models/ent/queue"
	"_models/ent/queuemessage"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// QueueMessageQuery is the builder for querying QueueMessage entities.
type QueueMessageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.QueueMessage
	withQueue  *QueueQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QueueMessageQuery builder.
func (qmq *QueueMessageQuery) Where(ps ...predicate.QueueMessage) *QueueMessageQuery {
	qmq.predicates = append(qmq.predicates, ps...)
	return qmq
}

// Limit adds a limit step to the query.
func (qmq *QueueMessageQuery) Limit(limit int) *QueueMessageQuery {
	qmq.limit = &limit
	return qmq
}

// Offset adds an offset step to the query.
func (qmq *QueueMessageQuery) Offset(offset int) *QueueMessageQuery {
	qmq.offset = &offset
	return qmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qmq *QueueMessageQuery) Unique(unique bool) *QueueMessageQuery {
	qmq.unique = &unique
	return qmq
}

// Order adds an order step to the query.
func (qmq *QueueMessageQuery) Order(o ...OrderFunc) *QueueMessageQuery {
	qmq.order = append(qmq.order, o...)
	return qmq
}

// QueryQueue chains the current query on the "queue" edge.
func (qmq *QueueMessageQuery) QueryQueue() *QueueQuery {
	query := &QueueQuery{config: qmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(queuemessage.Table, queuemessage.FieldID, selector),
			sqlgraph.To(queue.Table, queue.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, queuemessage.QueueTable, queuemessage.QueuePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(qmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first QueueMessage entity from the query.
// Returns a *NotFoundError when no QueueMessage was found.
func (qmq *QueueMessageQuery) First(ctx context.Context) (*QueueMessage, error) {
	nodes, err := qmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{queuemessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qmq *QueueMessageQuery) FirstX(ctx context.Context) *QueueMessage {
	node, err := qmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first QueueMessage ID from the query.
// Returns a *NotFoundError when no QueueMessage ID was found.
func (qmq *QueueMessageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = qmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{queuemessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qmq *QueueMessageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := qmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single QueueMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one QueueMessage entity is found.
// Returns a *NotFoundError when no QueueMessage entities are found.
func (qmq *QueueMessageQuery) Only(ctx context.Context) (*QueueMessage, error) {
	nodes, err := qmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{queuemessage.Label}
	default:
		return nil, &NotSingularError{queuemessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qmq *QueueMessageQuery) OnlyX(ctx context.Context) *QueueMessage {
	node, err := qmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only QueueMessage ID in the query.
// Returns a *NotSingularError when more than one QueueMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (qmq *QueueMessageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = qmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{queuemessage.Label}
	default:
		err = &NotSingularError{queuemessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qmq *QueueMessageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := qmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of QueueMessages.
func (qmq *QueueMessageQuery) All(ctx context.Context) ([]*QueueMessage, error) {
	if err := qmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return qmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (qmq *QueueMessageQuery) AllX(ctx context.Context) []*QueueMessage {
	nodes, err := qmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of QueueMessage IDs.
func (qmq *QueueMessageQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := qmq.Select(queuemessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qmq *QueueMessageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := qmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qmq *QueueMessageQuery) Count(ctx context.Context) (int, error) {
	if err := qmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return qmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (qmq *QueueMessageQuery) CountX(ctx context.Context) int {
	count, err := qmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qmq *QueueMessageQuery) Exist(ctx context.Context) (bool, error) {
	if err := qmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return qmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (qmq *QueueMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := qmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QueueMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qmq *QueueMessageQuery) Clone() *QueueMessageQuery {
	if qmq == nil {
		return nil
	}
	return &QueueMessageQuery{
		config:     qmq.config,
		limit:      qmq.limit,
		offset:     qmq.offset,
		order:      append([]OrderFunc{}, qmq.order...),
		predicates: append([]predicate.QueueMessage{}, qmq.predicates...),
		withQueue:  qmq.withQueue.Clone(),
		// clone intermediate query.
		sql:    qmq.sql.Clone(),
		path:   qmq.path,
		unique: qmq.unique,
	}
}

// WithQueue tells the query-builder to eager-load the nodes that are connected to
// the "queue" edge. The optional arguments are used to configure the query builder of the edge.
func (qmq *QueueMessageQuery) WithQueue(opts ...func(*QueueQuery)) *QueueMessageQuery {
	query := &QueueQuery{config: qmq.config}
	for _, opt := range opts {
		opt(query)
	}
	qmq.withQueue = query
	return qmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Body string `json:"body,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.QueueMessage.Query().
//		GroupBy(queuemessage.FieldBody).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (qmq *QueueMessageQuery) GroupBy(field string, fields ...string) *QueueMessageGroupBy {
	grbuild := &QueueMessageGroupBy{config: qmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qmq.sqlQuery(ctx), nil
	}
	grbuild.label = queuemessage.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Body string `json:"body,omitempty"`
//	}
//
//	client.QueueMessage.Query().
//		Select(queuemessage.FieldBody).
//		Scan(ctx, &v)
func (qmq *QueueMessageQuery) Select(fields ...string) *QueueMessageSelect {
	qmq.fields = append(qmq.fields, fields...)
	selbuild := &QueueMessageSelect{QueueMessageQuery: qmq}
	selbuild.label = queuemessage.Label
	selbuild.flds, selbuild.scan = &qmq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a QueueMessageSelect configured with the given aggregations.
func (qmq *QueueMessageQuery) Aggregate(fns ...AggregateFunc) *QueueMessageSelect {
	return qmq.Select().Aggregate(fns...)
}

func (qmq *QueueMessageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range qmq.fields {
		if !queuemessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qmq.path != nil {
		prev, err := qmq.path(ctx)
		if err != nil {
			return err
		}
		qmq.sql = prev
	}
	return nil
}

func (qmq *QueueMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*QueueMessage, error) {
	var (
		nodes       = []*QueueMessage{}
		_spec       = qmq.querySpec()
		loadedTypes = [1]bool{
			qmq.withQueue != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*QueueMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &QueueMessage{config: qmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := qmq.withQueue; query != nil {
		if err := qmq.loadQueue(ctx, query, nodes,
			func(n *QueueMessage) { n.Edges.Queue = []*Queue{} },
			func(n *QueueMessage, e *Queue) { n.Edges.Queue = append(n.Edges.Queue, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (qmq *QueueMessageQuery) loadQueue(ctx context.Context, query *QueueQuery, nodes []*QueueMessage, init func(*QueueMessage), assign func(*QueueMessage, *Queue)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*QueueMessage)
	nids := make(map[uuid.UUID]map[*QueueMessage]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(queuemessage.QueueTable)
		s.Join(joinT).On(s.C(queue.FieldID), joinT.C(queuemessage.QueuePrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(queuemessage.QueuePrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(queuemessage.QueuePrimaryKey[1]))
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
				nids[inValue] = map[*QueueMessage]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "queue" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (qmq *QueueMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qmq.querySpec()
	_spec.Node.Columns = qmq.fields
	if len(qmq.fields) > 0 {
		_spec.Unique = qmq.unique != nil && *qmq.unique
	}
	return sqlgraph.CountNodes(ctx, qmq.driver, _spec)
}

func (qmq *QueueMessageQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := qmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (qmq *QueueMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   queuemessage.Table,
			Columns: queuemessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: queuemessage.FieldID,
			},
		},
		From:   qmq.sql,
		Unique: true,
	}
	if unique := qmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := qmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, queuemessage.FieldID)
		for i := range fields {
			if fields[i] != queuemessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qmq *QueueMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qmq.driver.Dialect())
	t1 := builder.Table(queuemessage.Table)
	columns := qmq.fields
	if len(columns) == 0 {
		columns = queuemessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qmq.sql != nil {
		selector = qmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qmq.unique != nil && *qmq.unique {
		selector.Distinct()
	}
	for _, p := range qmq.predicates {
		p(selector)
	}
	for _, p := range qmq.order {
		p(selector)
	}
	if offset := qmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QueueMessageGroupBy is the group-by builder for QueueMessage entities.
type QueueMessageGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qmgb *QueueMessageGroupBy) Aggregate(fns ...AggregateFunc) *QueueMessageGroupBy {
	qmgb.fns = append(qmgb.fns, fns...)
	return qmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (qmgb *QueueMessageGroupBy) Scan(ctx context.Context, v any) error {
	query, err := qmgb.path(ctx)
	if err != nil {
		return err
	}
	qmgb.sql = query
	return qmgb.sqlScan(ctx, v)
}

func (qmgb *QueueMessageGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range qmgb.fields {
		if !queuemessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := qmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qmgb *QueueMessageGroupBy) sqlQuery() *sql.Selector {
	selector := qmgb.sql.Select()
	aggregation := make([]string, 0, len(qmgb.fns))
	for _, fn := range qmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(qmgb.fields)+len(qmgb.fns))
		for _, f := range qmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(qmgb.fields...)...)
}

// QueueMessageSelect is the builder for selecting fields of QueueMessage entities.
type QueueMessageSelect struct {
	*QueueMessageQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qms *QueueMessageSelect) Aggregate(fns ...AggregateFunc) *QueueMessageSelect {
	qms.fns = append(qms.fns, fns...)
	return qms
}

// Scan applies the selector query and scans the result into the given value.
func (qms *QueueMessageSelect) Scan(ctx context.Context, v any) error {
	if err := qms.prepareQuery(ctx); err != nil {
		return err
	}
	qms.sql = qms.QueueMessageQuery.sqlQuery(ctx)
	return qms.sqlScan(ctx, v)
}

func (qms *QueueMessageSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(qms.fns))
	for _, fn := range qms.fns {
		aggregation = append(aggregation, fn(qms.sql))
	}
	switch n := len(*qms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		qms.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		qms.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := qms.sql.Query()
	if err := qms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
