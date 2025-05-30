// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplan"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplanuploadfile"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/uploadfile"
)

// UploadFileQuery is the builder for querying UploadFile entities.
type UploadFileQuery struct {
	config
	ctx                       *QueryContext
	order                     []uploadfile.OrderOption
	inters                    []Interceptor
	predicates                []predicate.UploadFile
	withLessonPlan            *LessonPlanQuery
	withLessonPlanUploadFiles *LessonPlanUploadFileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UploadFileQuery builder.
func (ufq *UploadFileQuery) Where(ps ...predicate.UploadFile) *UploadFileQuery {
	ufq.predicates = append(ufq.predicates, ps...)
	return ufq
}

// Limit the number of records to be returned by this query.
func (ufq *UploadFileQuery) Limit(limit int) *UploadFileQuery {
	ufq.ctx.Limit = &limit
	return ufq
}

// Offset to start from.
func (ufq *UploadFileQuery) Offset(offset int) *UploadFileQuery {
	ufq.ctx.Offset = &offset
	return ufq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ufq *UploadFileQuery) Unique(unique bool) *UploadFileQuery {
	ufq.ctx.Unique = &unique
	return ufq
}

// Order specifies how the records should be ordered.
func (ufq *UploadFileQuery) Order(o ...uploadfile.OrderOption) *UploadFileQuery {
	ufq.order = append(ufq.order, o...)
	return ufq
}

// QueryLessonPlan chains the current query on the "LessonPlan" edge.
func (ufq *UploadFileQuery) QueryLessonPlan() *LessonPlanQuery {
	query := (&LessonPlanClient{config: ufq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(uploadfile.Table, uploadfile.FieldID, selector),
			sqlgraph.To(lessonplan.Table, lessonplan.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, uploadfile.LessonPlanTable, uploadfile.LessonPlanPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLessonPlanUploadFiles chains the current query on the "lesson_plan_upload_files" edge.
func (ufq *UploadFileQuery) QueryLessonPlanUploadFiles() *LessonPlanUploadFileQuery {
	query := (&LessonPlanUploadFileClient{config: ufq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(uploadfile.Table, uploadfile.FieldID, selector),
			sqlgraph.To(lessonplanuploadfile.Table, lessonplanuploadfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, uploadfile.LessonPlanUploadFilesTable, uploadfile.LessonPlanUploadFilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UploadFile entity from the query.
// Returns a *NotFoundError when no UploadFile was found.
func (ufq *UploadFileQuery) First(ctx context.Context) (*UploadFile, error) {
	nodes, err := ufq.Limit(1).All(setContextOp(ctx, ufq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{uploadfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ufq *UploadFileQuery) FirstX(ctx context.Context) *UploadFile {
	node, err := ufq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UploadFile ID from the query.
// Returns a *NotFoundError when no UploadFile ID was found.
func (ufq *UploadFileQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ufq.Limit(1).IDs(setContextOp(ctx, ufq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{uploadfile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ufq *UploadFileQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ufq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UploadFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UploadFile entity is found.
// Returns a *NotFoundError when no UploadFile entities are found.
func (ufq *UploadFileQuery) Only(ctx context.Context) (*UploadFile, error) {
	nodes, err := ufq.Limit(2).All(setContextOp(ctx, ufq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{uploadfile.Label}
	default:
		return nil, &NotSingularError{uploadfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ufq *UploadFileQuery) OnlyX(ctx context.Context) *UploadFile {
	node, err := ufq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UploadFile ID in the query.
// Returns a *NotSingularError when more than one UploadFile ID is found.
// Returns a *NotFoundError when no entities are found.
func (ufq *UploadFileQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ufq.Limit(2).IDs(setContextOp(ctx, ufq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{uploadfile.Label}
	default:
		err = &NotSingularError{uploadfile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ufq *UploadFileQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ufq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UploadFiles.
func (ufq *UploadFileQuery) All(ctx context.Context) ([]*UploadFile, error) {
	ctx = setContextOp(ctx, ufq.ctx, ent.OpQueryAll)
	if err := ufq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UploadFile, *UploadFileQuery]()
	return withInterceptors[[]*UploadFile](ctx, ufq, qr, ufq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ufq *UploadFileQuery) AllX(ctx context.Context) []*UploadFile {
	nodes, err := ufq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UploadFile IDs.
func (ufq *UploadFileQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if ufq.ctx.Unique == nil && ufq.path != nil {
		ufq.Unique(true)
	}
	ctx = setContextOp(ctx, ufq.ctx, ent.OpQueryIDs)
	if err = ufq.Select(uploadfile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ufq *UploadFileQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ufq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ufq *UploadFileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ufq.ctx, ent.OpQueryCount)
	if err := ufq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ufq, querierCount[*UploadFileQuery](), ufq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ufq *UploadFileQuery) CountX(ctx context.Context) int {
	count, err := ufq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ufq *UploadFileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ufq.ctx, ent.OpQueryExist)
	switch _, err := ufq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ufq *UploadFileQuery) ExistX(ctx context.Context) bool {
	exist, err := ufq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UploadFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ufq *UploadFileQuery) Clone() *UploadFileQuery {
	if ufq == nil {
		return nil
	}
	return &UploadFileQuery{
		config:                    ufq.config,
		ctx:                       ufq.ctx.Clone(),
		order:                     append([]uploadfile.OrderOption{}, ufq.order...),
		inters:                    append([]Interceptor{}, ufq.inters...),
		predicates:                append([]predicate.UploadFile{}, ufq.predicates...),
		withLessonPlan:            ufq.withLessonPlan.Clone(),
		withLessonPlanUploadFiles: ufq.withLessonPlanUploadFiles.Clone(),
		// clone intermediate query.
		sql:  ufq.sql.Clone(),
		path: ufq.path,
	}
}

// WithLessonPlan tells the query-builder to eager-load the nodes that are connected to
// the "LessonPlan" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UploadFileQuery) WithLessonPlan(opts ...func(*LessonPlanQuery)) *UploadFileQuery {
	query := (&LessonPlanClient{config: ufq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufq.withLessonPlan = query
	return ufq
}

// WithLessonPlanUploadFiles tells the query-builder to eager-load the nodes that are connected to
// the "lesson_plan_upload_files" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UploadFileQuery) WithLessonPlanUploadFiles(opts ...func(*LessonPlanUploadFileQuery)) *UploadFileQuery {
	query := (&LessonPlanUploadFileClient{config: ufq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufq.withLessonPlanUploadFiles = query
	return ufq
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
//	client.UploadFile.Query().
//		GroupBy(uploadfile.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ufq *UploadFileQuery) GroupBy(field string, fields ...string) *UploadFileGroupBy {
	ufq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UploadFileGroupBy{build: ufq}
	grbuild.flds = &ufq.ctx.Fields
	grbuild.label = uploadfile.Label
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
//	client.UploadFile.Query().
//		Select(uploadfile.FieldCreatedAt).
//		Scan(ctx, &v)
func (ufq *UploadFileQuery) Select(fields ...string) *UploadFileSelect {
	ufq.ctx.Fields = append(ufq.ctx.Fields, fields...)
	sbuild := &UploadFileSelect{UploadFileQuery: ufq}
	sbuild.label = uploadfile.Label
	sbuild.flds, sbuild.scan = &ufq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UploadFileSelect configured with the given aggregations.
func (ufq *UploadFileQuery) Aggregate(fns ...AggregateFunc) *UploadFileSelect {
	return ufq.Select().Aggregate(fns...)
}

func (ufq *UploadFileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ufq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ufq); err != nil {
				return err
			}
		}
	}
	for _, f := range ufq.ctx.Fields {
		if !uploadfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ufq.path != nil {
		prev, err := ufq.path(ctx)
		if err != nil {
			return err
		}
		ufq.sql = prev
	}
	return nil
}

func (ufq *UploadFileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UploadFile, error) {
	var (
		nodes       = []*UploadFile{}
		_spec       = ufq.querySpec()
		loadedTypes = [2]bool{
			ufq.withLessonPlan != nil,
			ufq.withLessonPlanUploadFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UploadFile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UploadFile{config: ufq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ufq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ufq.withLessonPlan; query != nil {
		if err := ufq.loadLessonPlan(ctx, query, nodes,
			func(n *UploadFile) { n.Edges.LessonPlan = []*LessonPlan{} },
			func(n *UploadFile, e *LessonPlan) { n.Edges.LessonPlan = append(n.Edges.LessonPlan, e) }); err != nil {
			return nil, err
		}
	}
	if query := ufq.withLessonPlanUploadFiles; query != nil {
		if err := ufq.loadLessonPlanUploadFiles(ctx, query, nodes,
			func(n *UploadFile) { n.Edges.LessonPlanUploadFiles = []*LessonPlanUploadFile{} },
			func(n *UploadFile, e *LessonPlanUploadFile) {
				n.Edges.LessonPlanUploadFiles = append(n.Edges.LessonPlanUploadFiles, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ufq *UploadFileQuery) loadLessonPlan(ctx context.Context, query *LessonPlanQuery, nodes []*UploadFile, init func(*UploadFile), assign func(*UploadFile, *LessonPlan)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*UploadFile)
	nids := make(map[int64]map[*UploadFile]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(uploadfile.LessonPlanTable)
		s.Join(joinT).On(s.C(lessonplan.FieldID), joinT.C(uploadfile.LessonPlanPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(uploadfile.LessonPlanPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(uploadfile.LessonPlanPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*UploadFile]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*LessonPlan](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "LessonPlan" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (ufq *UploadFileQuery) loadLessonPlanUploadFiles(ctx context.Context, query *LessonPlanUploadFileQuery, nodes []*UploadFile, init func(*UploadFile), assign func(*UploadFile, *LessonPlanUploadFile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*UploadFile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(lessonplanuploadfile.FieldUploadFileID)
	}
	query.Where(predicate.LessonPlanUploadFile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(uploadfile.LessonPlanUploadFilesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UploadFileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "upload_file_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ufq *UploadFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ufq.querySpec()
	_spec.Node.Columns = ufq.ctx.Fields
	if len(ufq.ctx.Fields) > 0 {
		_spec.Unique = ufq.ctx.Unique != nil && *ufq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ufq.driver, _spec)
}

func (ufq *UploadFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(uploadfile.Table, uploadfile.Columns, sqlgraph.NewFieldSpec(uploadfile.FieldID, field.TypeInt64))
	_spec.From = ufq.sql
	if unique := ufq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ufq.path != nil {
		_spec.Unique = true
	}
	if fields := ufq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uploadfile.FieldID)
		for i := range fields {
			if fields[i] != uploadfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ufq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ufq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ufq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ufq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ufq *UploadFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ufq.driver.Dialect())
	t1 := builder.Table(uploadfile.Table)
	columns := ufq.ctx.Fields
	if len(columns) == 0 {
		columns = uploadfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ufq.sql != nil {
		selector = ufq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ufq.ctx.Unique != nil && *ufq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ufq.predicates {
		p(selector)
	}
	for _, p := range ufq.order {
		p(selector)
	}
	if offset := ufq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ufq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UploadFileGroupBy is the group-by builder for UploadFile entities.
type UploadFileGroupBy struct {
	selector
	build *UploadFileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ufgb *UploadFileGroupBy) Aggregate(fns ...AggregateFunc) *UploadFileGroupBy {
	ufgb.fns = append(ufgb.fns, fns...)
	return ufgb
}

// Scan applies the selector query and scans the result into the given value.
func (ufgb *UploadFileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufgb.build.ctx, ent.OpQueryGroupBy)
	if err := ufgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UploadFileQuery, *UploadFileGroupBy](ctx, ufgb.build, ufgb, ufgb.build.inters, v)
}

func (ufgb *UploadFileGroupBy) sqlScan(ctx context.Context, root *UploadFileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ufgb.fns))
	for _, fn := range ufgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ufgb.flds)+len(ufgb.fns))
		for _, f := range *ufgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ufgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UploadFileSelect is the builder for selecting fields of UploadFile entities.
type UploadFileSelect struct {
	*UploadFileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ufs *UploadFileSelect) Aggregate(fns ...AggregateFunc) *UploadFileSelect {
	ufs.fns = append(ufs.fns, fns...)
	return ufs
}

// Scan applies the selector query and scans the result into the given value.
func (ufs *UploadFileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufs.ctx, ent.OpQuerySelect)
	if err := ufs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UploadFileQuery, *UploadFileSelect](ctx, ufs.UploadFileQuery, ufs, ufs.inters, v)
}

func (ufs *UploadFileSelect) sqlScan(ctx context.Context, root *UploadFileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ufs.fns))
	for _, fn := range ufs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ufs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
