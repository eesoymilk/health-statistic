// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eesoymilk/health-statistic-api/ent/mycard"
	"github.com/eesoymilk/health-statistic-api/ent/notification"
	"github.com/eesoymilk/health-statistic-api/ent/predicate"
	"github.com/eesoymilk/health-statistic-api/ent/user"
)

// MyCardQuery is the builder for querying MyCard entities.
type MyCardQuery struct {
	config
	ctx               *QueryContext
	order             []mycard.OrderOption
	inters            []Interceptor
	predicates        []predicate.MyCard
	withRecipient     *UserQuery
	withNotifications *NotificationQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MyCardQuery builder.
func (mcq *MyCardQuery) Where(ps ...predicate.MyCard) *MyCardQuery {
	mcq.predicates = append(mcq.predicates, ps...)
	return mcq
}

// Limit the number of records to be returned by this query.
func (mcq *MyCardQuery) Limit(limit int) *MyCardQuery {
	mcq.ctx.Limit = &limit
	return mcq
}

// Offset to start from.
func (mcq *MyCardQuery) Offset(offset int) *MyCardQuery {
	mcq.ctx.Offset = &offset
	return mcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mcq *MyCardQuery) Unique(unique bool) *MyCardQuery {
	mcq.ctx.Unique = &unique
	return mcq
}

// Order specifies how the records should be ordered.
func (mcq *MyCardQuery) Order(o ...mycard.OrderOption) *MyCardQuery {
	mcq.order = append(mcq.order, o...)
	return mcq
}

// QueryRecipient chains the current query on the "recipient" edge.
func (mcq *MyCardQuery) QueryRecipient() *UserQuery {
	query := (&UserClient{config: mcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mycard.Table, mycard.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, mycard.RecipientTable, mycard.RecipientColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifications chains the current query on the "notifications" edge.
func (mcq *MyCardQuery) QueryNotifications() *NotificationQuery {
	query := (&NotificationClient{config: mcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mycard.Table, mycard.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mycard.NotificationsTable, mycard.NotificationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MyCard entity from the query.
// Returns a *NotFoundError when no MyCard was found.
func (mcq *MyCardQuery) First(ctx context.Context) (*MyCard, error) {
	nodes, err := mcq.Limit(1).All(setContextOp(ctx, mcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mycard.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mcq *MyCardQuery) FirstX(ctx context.Context) *MyCard {
	node, err := mcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MyCard ID from the query.
// Returns a *NotFoundError when no MyCard ID was found.
func (mcq *MyCardQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mcq.Limit(1).IDs(setContextOp(ctx, mcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mycard.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mcq *MyCardQuery) FirstIDX(ctx context.Context) string {
	id, err := mcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MyCard entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MyCard entity is found.
// Returns a *NotFoundError when no MyCard entities are found.
func (mcq *MyCardQuery) Only(ctx context.Context) (*MyCard, error) {
	nodes, err := mcq.Limit(2).All(setContextOp(ctx, mcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mycard.Label}
	default:
		return nil, &NotSingularError{mycard.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mcq *MyCardQuery) OnlyX(ctx context.Context) *MyCard {
	node, err := mcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MyCard ID in the query.
// Returns a *NotSingularError when more than one MyCard ID is found.
// Returns a *NotFoundError when no entities are found.
func (mcq *MyCardQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mcq.Limit(2).IDs(setContextOp(ctx, mcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mycard.Label}
	default:
		err = &NotSingularError{mycard.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mcq *MyCardQuery) OnlyIDX(ctx context.Context) string {
	id, err := mcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MyCards.
func (mcq *MyCardQuery) All(ctx context.Context) ([]*MyCard, error) {
	ctx = setContextOp(ctx, mcq.ctx, "All")
	if err := mcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MyCard, *MyCardQuery]()
	return withInterceptors[[]*MyCard](ctx, mcq, qr, mcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mcq *MyCardQuery) AllX(ctx context.Context) []*MyCard {
	nodes, err := mcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MyCard IDs.
func (mcq *MyCardQuery) IDs(ctx context.Context) (ids []string, err error) {
	if mcq.ctx.Unique == nil && mcq.path != nil {
		mcq.Unique(true)
	}
	ctx = setContextOp(ctx, mcq.ctx, "IDs")
	if err = mcq.Select(mycard.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mcq *MyCardQuery) IDsX(ctx context.Context) []string {
	ids, err := mcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mcq *MyCardQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mcq.ctx, "Count")
	if err := mcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mcq, querierCount[*MyCardQuery](), mcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mcq *MyCardQuery) CountX(ctx context.Context) int {
	count, err := mcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mcq *MyCardQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mcq.ctx, "Exist")
	switch _, err := mcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mcq *MyCardQuery) ExistX(ctx context.Context) bool {
	exist, err := mcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MyCardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mcq *MyCardQuery) Clone() *MyCardQuery {
	if mcq == nil {
		return nil
	}
	return &MyCardQuery{
		config:            mcq.config,
		ctx:               mcq.ctx.Clone(),
		order:             append([]mycard.OrderOption{}, mcq.order...),
		inters:            append([]Interceptor{}, mcq.inters...),
		predicates:        append([]predicate.MyCard{}, mcq.predicates...),
		withRecipient:     mcq.withRecipient.Clone(),
		withNotifications: mcq.withNotifications.Clone(),
		// clone intermediate query.
		sql:  mcq.sql.Clone(),
		path: mcq.path,
	}
}

// WithRecipient tells the query-builder to eager-load the nodes that are connected to
// the "recipient" edge. The optional arguments are used to configure the query builder of the edge.
func (mcq *MyCardQuery) WithRecipient(opts ...func(*UserQuery)) *MyCardQuery {
	query := (&UserClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcq.withRecipient = query
	return mcq
}

// WithNotifications tells the query-builder to eager-load the nodes that are connected to
// the "notifications" edge. The optional arguments are used to configure the query builder of the edge.
func (mcq *MyCardQuery) WithNotifications(opts ...func(*NotificationQuery)) *MyCardQuery {
	query := (&NotificationClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcq.withNotifications = query
	return mcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CardPassword string `json:"card_password,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MyCard.Query().
//		GroupBy(mycard.FieldCardPassword).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mcq *MyCardQuery) GroupBy(field string, fields ...string) *MyCardGroupBy {
	mcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MyCardGroupBy{build: mcq}
	grbuild.flds = &mcq.ctx.Fields
	grbuild.label = mycard.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CardPassword string `json:"card_password,omitempty"`
//	}
//
//	client.MyCard.Query().
//		Select(mycard.FieldCardPassword).
//		Scan(ctx, &v)
func (mcq *MyCardQuery) Select(fields ...string) *MyCardSelect {
	mcq.ctx.Fields = append(mcq.ctx.Fields, fields...)
	sbuild := &MyCardSelect{MyCardQuery: mcq}
	sbuild.label = mycard.Label
	sbuild.flds, sbuild.scan = &mcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MyCardSelect configured with the given aggregations.
func (mcq *MyCardQuery) Aggregate(fns ...AggregateFunc) *MyCardSelect {
	return mcq.Select().Aggregate(fns...)
}

func (mcq *MyCardQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mcq); err != nil {
				return err
			}
		}
	}
	for _, f := range mcq.ctx.Fields {
		if !mycard.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mcq.path != nil {
		prev, err := mcq.path(ctx)
		if err != nil {
			return err
		}
		mcq.sql = prev
	}
	return nil
}

func (mcq *MyCardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MyCard, error) {
	var (
		nodes       = []*MyCard{}
		withFKs     = mcq.withFKs
		_spec       = mcq.querySpec()
		loadedTypes = [2]bool{
			mcq.withRecipient != nil,
			mcq.withNotifications != nil,
		}
	)
	if mcq.withRecipient != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, mycard.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MyCard).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MyCard{config: mcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mcq.withRecipient; query != nil {
		if err := mcq.loadRecipient(ctx, query, nodes, nil,
			func(n *MyCard, e *User) { n.Edges.Recipient = e }); err != nil {
			return nil, err
		}
	}
	if query := mcq.withNotifications; query != nil {
		if err := mcq.loadNotifications(ctx, query, nodes,
			func(n *MyCard) { n.Edges.Notifications = []*Notification{} },
			func(n *MyCard, e *Notification) { n.Edges.Notifications = append(n.Edges.Notifications, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mcq *MyCardQuery) loadRecipient(ctx context.Context, query *UserQuery, nodes []*MyCard, init func(*MyCard), assign func(*MyCard, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*MyCard)
	for i := range nodes {
		if nodes[i].my_card_recipient == nil {
			continue
		}
		fk := *nodes[i].my_card_recipient
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "my_card_recipient" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mcq *MyCardQuery) loadNotifications(ctx context.Context, query *NotificationQuery, nodes []*MyCard, init func(*MyCard), assign func(*MyCard, *Notification)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*MyCard)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(mycard.NotificationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.my_card_notifications
		if fk == nil {
			return fmt.Errorf(`foreign-key "my_card_notifications" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "my_card_notifications" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mcq *MyCardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mcq.querySpec()
	_spec.Node.Columns = mcq.ctx.Fields
	if len(mcq.ctx.Fields) > 0 {
		_spec.Unique = mcq.ctx.Unique != nil && *mcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mcq.driver, _spec)
}

func (mcq *MyCardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(mycard.Table, mycard.Columns, sqlgraph.NewFieldSpec(mycard.FieldID, field.TypeString))
	_spec.From = mcq.sql
	if unique := mcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mcq.path != nil {
		_spec.Unique = true
	}
	if fields := mcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mycard.FieldID)
		for i := range fields {
			if fields[i] != mycard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mcq *MyCardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mcq.driver.Dialect())
	t1 := builder.Table(mycard.Table)
	columns := mcq.ctx.Fields
	if len(columns) == 0 {
		columns = mycard.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mcq.sql != nil {
		selector = mcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mcq.ctx.Unique != nil && *mcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mcq.predicates {
		p(selector)
	}
	for _, p := range mcq.order {
		p(selector)
	}
	if offset := mcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MyCardGroupBy is the group-by builder for MyCard entities.
type MyCardGroupBy struct {
	selector
	build *MyCardQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mcgb *MyCardGroupBy) Aggregate(fns ...AggregateFunc) *MyCardGroupBy {
	mcgb.fns = append(mcgb.fns, fns...)
	return mcgb
}

// Scan applies the selector query and scans the result into the given value.
func (mcgb *MyCardGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mcgb.build.ctx, "GroupBy")
	if err := mcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MyCardQuery, *MyCardGroupBy](ctx, mcgb.build, mcgb, mcgb.build.inters, v)
}

func (mcgb *MyCardGroupBy) sqlScan(ctx context.Context, root *MyCardQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mcgb.fns))
	for _, fn := range mcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mcgb.flds)+len(mcgb.fns))
		for _, f := range *mcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MyCardSelect is the builder for selecting fields of MyCard entities.
type MyCardSelect struct {
	*MyCardQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mcs *MyCardSelect) Aggregate(fns ...AggregateFunc) *MyCardSelect {
	mcs.fns = append(mcs.fns, fns...)
	return mcs
}

// Scan applies the selector query and scans the result into the given value.
func (mcs *MyCardSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mcs.ctx, "Select")
	if err := mcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MyCardQuery, *MyCardSelect](ctx, mcs.MyCardQuery, mcs, mcs.inters, v)
}

func (mcs *MyCardSelect) sqlScan(ctx context.Context, root *MyCardQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mcs.fns))
	for _, fn := range mcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
