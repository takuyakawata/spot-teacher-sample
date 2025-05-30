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
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplan"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplansubject"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/subject"
)

// LessonPlanSubjectQuery is the builder for querying LessonPlanSubject entities.
type LessonPlanSubjectQuery struct {
	config
	ctx            *QueryContext
	order          []lessonplansubject.OrderOption
	inters         []Interceptor
	predicates     []predicate.LessonPlanSubject
	withLessonPlan *LessonPlanQuery
	withSubject    *SubjectQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LessonPlanSubjectQuery builder.
func (lpsq *LessonPlanSubjectQuery) Where(ps ...predicate.LessonPlanSubject) *LessonPlanSubjectQuery {
	lpsq.predicates = append(lpsq.predicates, ps...)
	return lpsq
}

// Limit the number of records to be returned by this query.
func (lpsq *LessonPlanSubjectQuery) Limit(limit int) *LessonPlanSubjectQuery {
	lpsq.ctx.Limit = &limit
	return lpsq
}

// Offset to start from.
func (lpsq *LessonPlanSubjectQuery) Offset(offset int) *LessonPlanSubjectQuery {
	lpsq.ctx.Offset = &offset
	return lpsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lpsq *LessonPlanSubjectQuery) Unique(unique bool) *LessonPlanSubjectQuery {
	lpsq.ctx.Unique = &unique
	return lpsq
}

// Order specifies how the records should be ordered.
func (lpsq *LessonPlanSubjectQuery) Order(o ...lessonplansubject.OrderOption) *LessonPlanSubjectQuery {
	lpsq.order = append(lpsq.order, o...)
	return lpsq
}

// QueryLessonPlan chains the current query on the "lesson_plan" edge.
func (lpsq *LessonPlanSubjectQuery) QueryLessonPlan() *LessonPlanQuery {
	query := (&LessonPlanClient{config: lpsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lpsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lpsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lessonplansubject.Table, lessonplansubject.FieldID, selector),
			sqlgraph.To(lessonplan.Table, lessonplan.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, lessonplansubject.LessonPlanTable, lessonplansubject.LessonPlanColumn),
		)
		fromU = sqlgraph.SetNeighbors(lpsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubject chains the current query on the "subject" edge.
func (lpsq *LessonPlanSubjectQuery) QuerySubject() *SubjectQuery {
	query := (&SubjectClient{config: lpsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lpsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lpsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lessonplansubject.Table, lessonplansubject.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, lessonplansubject.SubjectTable, lessonplansubject.SubjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(lpsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LessonPlanSubject entity from the query.
// Returns a *NotFoundError when no LessonPlanSubject was found.
func (lpsq *LessonPlanSubjectQuery) First(ctx context.Context) (*LessonPlanSubject, error) {
	nodes, err := lpsq.Limit(1).All(setContextOp(ctx, lpsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lessonplansubject.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lpsq *LessonPlanSubjectQuery) FirstX(ctx context.Context) *LessonPlanSubject {
	node, err := lpsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LessonPlanSubject ID from the query.
// Returns a *NotFoundError when no LessonPlanSubject ID was found.
func (lpsq *LessonPlanSubjectQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lpsq.Limit(1).IDs(setContextOp(ctx, lpsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lessonplansubject.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lpsq *LessonPlanSubjectQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lpsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LessonPlanSubject entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LessonPlanSubject entity is found.
// Returns a *NotFoundError when no LessonPlanSubject entities are found.
func (lpsq *LessonPlanSubjectQuery) Only(ctx context.Context) (*LessonPlanSubject, error) {
	nodes, err := lpsq.Limit(2).All(setContextOp(ctx, lpsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lessonplansubject.Label}
	default:
		return nil, &NotSingularError{lessonplansubject.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lpsq *LessonPlanSubjectQuery) OnlyX(ctx context.Context) *LessonPlanSubject {
	node, err := lpsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LessonPlanSubject ID in the query.
// Returns a *NotSingularError when more than one LessonPlanSubject ID is found.
// Returns a *NotFoundError when no entities are found.
func (lpsq *LessonPlanSubjectQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lpsq.Limit(2).IDs(setContextOp(ctx, lpsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lessonplansubject.Label}
	default:
		err = &NotSingularError{lessonplansubject.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lpsq *LessonPlanSubjectQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lpsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LessonPlanSubjects.
func (lpsq *LessonPlanSubjectQuery) All(ctx context.Context) ([]*LessonPlanSubject, error) {
	ctx = setContextOp(ctx, lpsq.ctx, ent.OpQueryAll)
	if err := lpsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LessonPlanSubject, *LessonPlanSubjectQuery]()
	return withInterceptors[[]*LessonPlanSubject](ctx, lpsq, qr, lpsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lpsq *LessonPlanSubjectQuery) AllX(ctx context.Context) []*LessonPlanSubject {
	nodes, err := lpsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LessonPlanSubject IDs.
func (lpsq *LessonPlanSubjectQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lpsq.ctx.Unique == nil && lpsq.path != nil {
		lpsq.Unique(true)
	}
	ctx = setContextOp(ctx, lpsq.ctx, ent.OpQueryIDs)
	if err = lpsq.Select(lessonplansubject.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lpsq *LessonPlanSubjectQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lpsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lpsq *LessonPlanSubjectQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lpsq.ctx, ent.OpQueryCount)
	if err := lpsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lpsq, querierCount[*LessonPlanSubjectQuery](), lpsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lpsq *LessonPlanSubjectQuery) CountX(ctx context.Context) int {
	count, err := lpsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lpsq *LessonPlanSubjectQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lpsq.ctx, ent.OpQueryExist)
	switch _, err := lpsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lpsq *LessonPlanSubjectQuery) ExistX(ctx context.Context) bool {
	exist, err := lpsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LessonPlanSubjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lpsq *LessonPlanSubjectQuery) Clone() *LessonPlanSubjectQuery {
	if lpsq == nil {
		return nil
	}
	return &LessonPlanSubjectQuery{
		config:         lpsq.config,
		ctx:            lpsq.ctx.Clone(),
		order:          append([]lessonplansubject.OrderOption{}, lpsq.order...),
		inters:         append([]Interceptor{}, lpsq.inters...),
		predicates:     append([]predicate.LessonPlanSubject{}, lpsq.predicates...),
		withLessonPlan: lpsq.withLessonPlan.Clone(),
		withSubject:    lpsq.withSubject.Clone(),
		// clone intermediate query.
		sql:  lpsq.sql.Clone(),
		path: lpsq.path,
	}
}

// WithLessonPlan tells the query-builder to eager-load the nodes that are connected to
// the "lesson_plan" edge. The optional arguments are used to configure the query builder of the edge.
func (lpsq *LessonPlanSubjectQuery) WithLessonPlan(opts ...func(*LessonPlanQuery)) *LessonPlanSubjectQuery {
	query := (&LessonPlanClient{config: lpsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lpsq.withLessonPlan = query
	return lpsq
}

// WithSubject tells the query-builder to eager-load the nodes that are connected to
// the "subject" edge. The optional arguments are used to configure the query builder of the edge.
func (lpsq *LessonPlanSubjectQuery) WithSubject(opts ...func(*SubjectQuery)) *LessonPlanSubjectQuery {
	query := (&SubjectClient{config: lpsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lpsq.withSubject = query
	return lpsq
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
//	client.LessonPlanSubject.Query().
//		GroupBy(lessonplansubject.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lpsq *LessonPlanSubjectQuery) GroupBy(field string, fields ...string) *LessonPlanSubjectGroupBy {
	lpsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LessonPlanSubjectGroupBy{build: lpsq}
	grbuild.flds = &lpsq.ctx.Fields
	grbuild.label = lessonplansubject.Label
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
//	client.LessonPlanSubject.Query().
//		Select(lessonplansubject.FieldCreatedAt).
//		Scan(ctx, &v)
func (lpsq *LessonPlanSubjectQuery) Select(fields ...string) *LessonPlanSubjectSelect {
	lpsq.ctx.Fields = append(lpsq.ctx.Fields, fields...)
	sbuild := &LessonPlanSubjectSelect{LessonPlanSubjectQuery: lpsq}
	sbuild.label = lessonplansubject.Label
	sbuild.flds, sbuild.scan = &lpsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LessonPlanSubjectSelect configured with the given aggregations.
func (lpsq *LessonPlanSubjectQuery) Aggregate(fns ...AggregateFunc) *LessonPlanSubjectSelect {
	return lpsq.Select().Aggregate(fns...)
}

func (lpsq *LessonPlanSubjectQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lpsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lpsq); err != nil {
				return err
			}
		}
	}
	for _, f := range lpsq.ctx.Fields {
		if !lessonplansubject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lpsq.path != nil {
		prev, err := lpsq.path(ctx)
		if err != nil {
			return err
		}
		lpsq.sql = prev
	}
	return nil
}

func (lpsq *LessonPlanSubjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LessonPlanSubject, error) {
	var (
		nodes       = []*LessonPlanSubject{}
		_spec       = lpsq.querySpec()
		loadedTypes = [2]bool{
			lpsq.withLessonPlan != nil,
			lpsq.withSubject != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LessonPlanSubject).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LessonPlanSubject{config: lpsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lpsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lpsq.withLessonPlan; query != nil {
		if err := lpsq.loadLessonPlan(ctx, query, nodes, nil,
			func(n *LessonPlanSubject, e *LessonPlan) { n.Edges.LessonPlan = e }); err != nil {
			return nil, err
		}
	}
	if query := lpsq.withSubject; query != nil {
		if err := lpsq.loadSubject(ctx, query, nodes, nil,
			func(n *LessonPlanSubject, e *Subject) { n.Edges.Subject = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lpsq *LessonPlanSubjectQuery) loadLessonPlan(ctx context.Context, query *LessonPlanQuery, nodes []*LessonPlanSubject, init func(*LessonPlanSubject), assign func(*LessonPlanSubject, *LessonPlan)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LessonPlanSubject)
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
func (lpsq *LessonPlanSubjectQuery) loadSubject(ctx context.Context, query *SubjectQuery, nodes []*LessonPlanSubject, init func(*LessonPlanSubject), assign func(*LessonPlanSubject, *Subject)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*LessonPlanSubject)
	for i := range nodes {
		fk := nodes[i].SubjectID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(subject.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "subject_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lpsq *LessonPlanSubjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lpsq.querySpec()
	_spec.Node.Columns = lpsq.ctx.Fields
	if len(lpsq.ctx.Fields) > 0 {
		_spec.Unique = lpsq.ctx.Unique != nil && *lpsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lpsq.driver, _spec)
}

func (lpsq *LessonPlanSubjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lessonplansubject.Table, lessonplansubject.Columns, sqlgraph.NewFieldSpec(lessonplansubject.FieldID, field.TypeInt64))
	_spec.From = lpsq.sql
	if unique := lpsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lpsq.path != nil {
		_spec.Unique = true
	}
	if fields := lpsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lessonplansubject.FieldID)
		for i := range fields {
			if fields[i] != lessonplansubject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if lpsq.withLessonPlan != nil {
			_spec.Node.AddColumnOnce(lessonplansubject.FieldLessonPlanID)
		}
		if lpsq.withSubject != nil {
			_spec.Node.AddColumnOnce(lessonplansubject.FieldSubjectID)
		}
	}
	if ps := lpsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lpsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lpsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lpsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lpsq *LessonPlanSubjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lpsq.driver.Dialect())
	t1 := builder.Table(lessonplansubject.Table)
	columns := lpsq.ctx.Fields
	if len(columns) == 0 {
		columns = lessonplansubject.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lpsq.sql != nil {
		selector = lpsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lpsq.ctx.Unique != nil && *lpsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lpsq.predicates {
		p(selector)
	}
	for _, p := range lpsq.order {
		p(selector)
	}
	if offset := lpsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lpsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LessonPlanSubjectGroupBy is the group-by builder for LessonPlanSubject entities.
type LessonPlanSubjectGroupBy struct {
	selector
	build *LessonPlanSubjectQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lpsgb *LessonPlanSubjectGroupBy) Aggregate(fns ...AggregateFunc) *LessonPlanSubjectGroupBy {
	lpsgb.fns = append(lpsgb.fns, fns...)
	return lpsgb
}

// Scan applies the selector query and scans the result into the given value.
func (lpsgb *LessonPlanSubjectGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lpsgb.build.ctx, ent.OpQueryGroupBy)
	if err := lpsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LessonPlanSubjectQuery, *LessonPlanSubjectGroupBy](ctx, lpsgb.build, lpsgb, lpsgb.build.inters, v)
}

func (lpsgb *LessonPlanSubjectGroupBy) sqlScan(ctx context.Context, root *LessonPlanSubjectQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lpsgb.fns))
	for _, fn := range lpsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lpsgb.flds)+len(lpsgb.fns))
		for _, f := range *lpsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lpsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lpsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LessonPlanSubjectSelect is the builder for selecting fields of LessonPlanSubject entities.
type LessonPlanSubjectSelect struct {
	*LessonPlanSubjectQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lpss *LessonPlanSubjectSelect) Aggregate(fns ...AggregateFunc) *LessonPlanSubjectSelect {
	lpss.fns = append(lpss.fns, fns...)
	return lpss
}

// Scan applies the selector query and scans the result into the given value.
func (lpss *LessonPlanSubjectSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lpss.ctx, ent.OpQuerySelect)
	if err := lpss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LessonPlanSubjectQuery, *LessonPlanSubjectSelect](ctx, lpss.LessonPlanSubjectQuery, lpss, lpss.inters, v)
}

func (lpss *LessonPlanSubjectSelect) sqlScan(ctx context.Context, root *LessonPlanSubjectQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lpss.fns))
	for _, fn := range lpss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lpss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lpss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
