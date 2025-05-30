// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/grade"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplan"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplangrade"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
)

// LessonPlanGradeQuery is the builder for querying LessonPlanGrade entities.
type LessonPlanGradeQuery struct {
	config
	ctx            *QueryContext
	order          []lessonplangrade.OrderOption
	inters         []Interceptor
	predicates     []predicate.LessonPlanGrade
	withLessonPlan *LessonPlanQuery
	withGrade      *GradeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LessonPlanGradeQuery builder.
func (lpgq *LessonPlanGradeQuery) Where(ps ...predicate.LessonPlanGrade) *LessonPlanGradeQuery {
	lpgq.predicates = append(lpgq.predicates, ps...)
	return lpgq
}

// Limit the number of records to be returned by this query.
func (lpgq *LessonPlanGradeQuery) Limit(limit int) *LessonPlanGradeQuery {
	lpgq.ctx.Limit = &limit
	return lpgq
}

// Offset to start from.
func (lpgq *LessonPlanGradeQuery) Offset(offset int) *LessonPlanGradeQuery {
	lpgq.ctx.Offset = &offset
	return lpgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lpgq *LessonPlanGradeQuery) Unique(unique bool) *LessonPlanGradeQuery {
	lpgq.ctx.Unique = &unique
	return lpgq
}

// Order specifies how the records should be ordered.
func (lpgq *LessonPlanGradeQuery) Order(o ...lessonplangrade.OrderOption) *LessonPlanGradeQuery {
	lpgq.order = append(lpgq.order, o...)
	return lpgq
}

// QueryLessonPlan chains the current query on the "lesson_plan" edge.
func (lpgq *LessonPlanGradeQuery) QueryLessonPlan() *LessonPlanQuery {
	query := (&LessonPlanClient{config: lpgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lpgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lpgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lessonplangrade.Table, lessonplangrade.FieldID, selector),
			sqlgraph.To(lessonplan.Table, lessonplan.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, lessonplangrade.LessonPlanTable, lessonplangrade.LessonPlanColumn),
		)
		fromU = sqlgraph.SetNeighbors(lpgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGrade chains the current query on the "grade" edge.
func (lpgq *LessonPlanGradeQuery) QueryGrade() *GradeQuery {
	query := (&GradeClient{config: lpgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lpgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lpgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lessonplangrade.Table, lessonplangrade.FieldID, selector),
			sqlgraph.To(grade.Table, grade.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, lessonplangrade.GradeTable, lessonplangrade.GradeColumn),
		)
		fromU = sqlgraph.SetNeighbors(lpgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LessonPlanGrade entity from the query.
// Returns a *NotFoundError when no LessonPlanGrade was found.
func (lpgq *LessonPlanGradeQuery) First(ctx context.Context) (*LessonPlanGrade, error) {
	nodes, err := lpgq.Limit(1).All(setContextOp(ctx, lpgq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lessonplangrade.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lpgq *LessonPlanGradeQuery) FirstX(ctx context.Context) *LessonPlanGrade {
	node, err := lpgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LessonPlanGrade ID from the query.
// Returns a *NotFoundError when no LessonPlanGrade ID was found.
func (lpgq *LessonPlanGradeQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lpgq.Limit(1).IDs(setContextOp(ctx, lpgq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lessonplangrade.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lpgq *LessonPlanGradeQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lpgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LessonPlanGrade entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LessonPlanGrade entity is found.
// Returns a *NotFoundError when no LessonPlanGrade entities are found.
func (lpgq *LessonPlanGradeQuery) Only(ctx context.Context) (*LessonPlanGrade, error) {
	nodes, err := lpgq.Limit(2).All(setContextOp(ctx, lpgq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lessonplangrade.Label}
	default:
		return nil, &NotSingularError{lessonplangrade.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lpgq *LessonPlanGradeQuery) OnlyX(ctx context.Context) *LessonPlanGrade {
	node, err := lpgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LessonPlanGrade ID in the query.
// Returns a *NotSingularError when more than one LessonPlanGrade ID is found.
// Returns a *NotFoundError when no entities are found.
func (lpgq *LessonPlanGradeQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lpgq.Limit(2).IDs(setContextOp(ctx, lpgq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lessonplangrade.Label}
	default:
		err = &NotSingularError{lessonplangrade.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lpgq *LessonPlanGradeQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lpgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LessonPlanGrades.
func (lpgq *LessonPlanGradeQuery) All(ctx context.Context) ([]*LessonPlanGrade, error) {
	ctx = setContextOp(ctx, lpgq.ctx, ent.OpQueryAll)
	if err := lpgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LessonPlanGrade, *LessonPlanGradeQuery]()
	return withInterceptors[[]*LessonPlanGrade](ctx, lpgq, qr, lpgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lpgq *LessonPlanGradeQuery) AllX(ctx context.Context) []*LessonPlanGrade {
	nodes, err := lpgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LessonPlanGrade IDs.
func (lpgq *LessonPlanGradeQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lpgq.ctx.Unique == nil && lpgq.path != nil {
		lpgq.Unique(true)
	}
	ctx = setContextOp(ctx, lpgq.ctx, ent.OpQueryIDs)
	if err = lpgq.Select(lessonplangrade.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lpgq *LessonPlanGradeQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lpgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lpgq *LessonPlanGradeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lpgq.ctx, ent.OpQueryCount)
	if err := lpgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lpgq, querierCount[*LessonPlanGradeQuery](), lpgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lpgq *LessonPlanGradeQuery) CountX(ctx context.Context) int {
	count, err := lpgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lpgq *LessonPlanGradeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lpgq.ctx, ent.OpQueryExist)
	switch _, err := lpgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lpgq *LessonPlanGradeQuery) ExistX(ctx context.Context) bool {
	exist, err := lpgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LessonPlanGradeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lpgq *LessonPlanGradeQuery) Clone() *LessonPlanGradeQuery {
	if lpgq == nil {
		return nil
	}
	return &LessonPlanGradeQuery{
		config:         lpgq.config,
		ctx:            lpgq.ctx.Clone(),
		order:          append([]lessonplangrade.OrderOption{}, lpgq.order...),
		inters:         append([]Interceptor{}, lpgq.inters...),
		predicates:     append([]predicate.LessonPlanGrade{}, lpgq.predicates...),
		withLessonPlan: lpgq.withLessonPlan.Clone(),
		withGrade:      lpgq.withGrade.Clone(),
		// clone intermediate query.
		sql:  lpgq.sql.Clone(),
		path: lpgq.path,
	}
}

// WithLessonPlan tells the query-builder to eager-load the nodes that are connected to
// the "lesson_plan" edge. The optional arguments are used to configure the query builder of the edge.
func (lpgq *LessonPlanGradeQuery) WithLessonPlan(opts ...func(*LessonPlanQuery)) *LessonPlanGradeQuery {
	query := (&LessonPlanClient{config: lpgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lpgq.withLessonPlan = query
	return lpgq
}

// WithGrade tells the query-builder to eager-load the nodes that are connected to
// the "grade" edge. The optional arguments are used to configure the query builder of the edge.
func (lpgq *LessonPlanGradeQuery) WithGrade(opts ...func(*GradeQuery)) *LessonPlanGradeQuery {
	query := (&GradeClient{config: lpgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lpgq.withGrade = query
	return lpgq
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
//	client.LessonPlanGrade.Query().
//		GroupBy(lessonplangrade.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lpgq *LessonPlanGradeQuery) GroupBy(field string, fields ...string) *LessonPlanGradeGroupBy {
	lpgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LessonPlanGradeGroupBy{build: lpgq}
	grbuild.flds = &lpgq.ctx.Fields
	grbuild.label = lessonplangrade.Label
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
//	client.LessonPlanGrade.Query().
//		Select(lessonplangrade.FieldCreatedAt).
//		Scan(ctx, &v)
func (lpgq *LessonPlanGradeQuery) Select(fields ...string) *LessonPlanGradeSelect {
	lpgq.ctx.Fields = append(lpgq.ctx.Fields, fields...)
	sbuild := &LessonPlanGradeSelect{LessonPlanGradeQuery: lpgq}
	sbuild.label = lessonplangrade.Label
	sbuild.flds, sbuild.scan = &lpgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LessonPlanGradeSelect configured with the given aggregations.
func (lpgq *LessonPlanGradeQuery) Aggregate(fns ...AggregateFunc) *LessonPlanGradeSelect {
	return lpgq.Select().Aggregate(fns...)
}

func (lpgq *LessonPlanGradeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lpgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lpgq); err != nil {
				return err
			}
		}
	}
	for _, f := range lpgq.ctx.Fields {
		if !lessonplangrade.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lpgq.path != nil {
		prev, err := lpgq.path(ctx)
		if err != nil {
			return err
		}
		lpgq.sql = prev
	}
	return nil
}

func (lpgq *LessonPlanGradeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LessonPlanGrade, error) {
	var (
		nodes       = []*LessonPlanGrade{}
		_spec       = lpgq.querySpec()
		loadedTypes = [2]bool{
			lpgq.withLessonPlan != nil,
			lpgq.withGrade != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LessonPlanGrade).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LessonPlanGrade{config: lpgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lpgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lpgq.withLessonPlan; query != nil {
		if err := lpgq.loadLessonPlan(ctx, query, nodes, nil,
			func(n *LessonPlanGrade, e *LessonPlan) { n.Edges.LessonPlan = e }); err != nil {
			return nil, err
		}
	}
	if query := lpgq.withGrade; query != nil {
		if err := lpgq.loadGrade(ctx, query, nodes, nil,
			func(n *LessonPlanGrade, e *Grade) { n.Edges.Grade = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lpgq *LessonPlanGradeQuery) loadLessonPlan(ctx context.Context, query *LessonPlanQuery, nodes []*LessonPlanGrade, init func(*LessonPlanGrade), assign func(*LessonPlanGrade, *LessonPlan)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LessonPlanGrade)
	for i := range nodes {
		fk := nodes[i].LessonPlanID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(lessonplan.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "lesson_plan_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lpgq *LessonPlanGradeQuery) loadGrade(ctx context.Context, query *GradeQuery, nodes []*LessonPlanGrade, init func(*LessonPlanGrade), assign func(*LessonPlanGrade, *Grade)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LessonPlanGrade)
	for i := range nodes {
		fk := nodes[i].GradeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(grade.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "grade_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lpgq *LessonPlanGradeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lpgq.querySpec()
	_spec.Node.Columns = lpgq.ctx.Fields
	if len(lpgq.ctx.Fields) > 0 {
		_spec.Unique = lpgq.ctx.Unique != nil && *lpgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lpgq.driver, _spec)
}

func (lpgq *LessonPlanGradeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lessonplangrade.Table, lessonplangrade.Columns, sqlgraph.NewFieldSpec(lessonplangrade.FieldID, field.TypeInt64))
	_spec.From = lpgq.sql
	if unique := lpgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lpgq.path != nil {
		_spec.Unique = true
	}
	if fields := lpgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lessonplangrade.FieldID)
		for i := range fields {
			if fields[i] != lessonplangrade.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if lpgq.withLessonPlan != nil {
			_spec.Node.AddColumnOnce(lessonplangrade.FieldLessonPlanID)
		}
		if lpgq.withGrade != nil {
			_spec.Node.AddColumnOnce(lessonplangrade.FieldGradeID)
		}
	}
	if ps := lpgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lpgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lpgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lpgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lpgq *LessonPlanGradeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lpgq.driver.Dialect())
	t1 := builder.Table(lessonplangrade.Table)
	columns := lpgq.ctx.Fields
	if len(columns) == 0 {
		columns = lessonplangrade.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lpgq.sql != nil {
		selector = lpgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lpgq.ctx.Unique != nil && *lpgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lpgq.predicates {
		p(selector)
	}
	for _, p := range lpgq.order {
		p(selector)
	}
	if offset := lpgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lpgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LessonPlanGradeGroupBy is the group-by builder for LessonPlanGrade entities.
type LessonPlanGradeGroupBy struct {
	selector
	build *LessonPlanGradeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lpggb *LessonPlanGradeGroupBy) Aggregate(fns ...AggregateFunc) *LessonPlanGradeGroupBy {
	lpggb.fns = append(lpggb.fns, fns...)
	return lpggb
}

// Scan applies the selector query and scans the result into the given value.
func (lpggb *LessonPlanGradeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lpggb.build.ctx, ent.OpQueryGroupBy)
	if err := lpggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LessonPlanGradeQuery, *LessonPlanGradeGroupBy](ctx, lpggb.build, lpggb, lpggb.build.inters, v)
}

func (lpggb *LessonPlanGradeGroupBy) sqlScan(ctx context.Context, root *LessonPlanGradeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lpggb.fns))
	for _, fn := range lpggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lpggb.flds)+len(lpggb.fns))
		for _, f := range *lpggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lpggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lpggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LessonPlanGradeSelect is the builder for selecting fields of LessonPlanGrade entities.
type LessonPlanGradeSelect struct {
	*LessonPlanGradeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lpgs *LessonPlanGradeSelect) Aggregate(fns ...AggregateFunc) *LessonPlanGradeSelect {
	lpgs.fns = append(lpgs.fns, fns...)
	return lpgs
}

// Scan applies the selector query and scans the result into the given value.
func (lpgs *LessonPlanGradeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lpgs.ctx, ent.OpQuerySelect)
	if err := lpgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LessonPlanGradeQuery, *LessonPlanGradeSelect](ctx, lpgs.LessonPlanGradeQuery, lpgs, lpgs.inters, v)
}

func (lpgs *LessonPlanGradeSelect) sqlScan(ctx context.Context, root *LessonPlanGradeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lpgs.fns))
	for _, fn := range lpgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lpgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lpgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
