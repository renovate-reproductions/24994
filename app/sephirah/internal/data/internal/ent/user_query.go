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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx              *QueryContext
	order            []OrderFunc
	inters           []Interceptor
	predicates       []predicate.User
	withBindAccount  *AccountQuery
	withPurchasedApp *AppQuery
	withAppPackage   *AppPackageQuery
	withFeedConfig   *FeedConfigQuery
	withNotifyTarget *NotifyTargetQuery
	withNotifyFlow   *NotifyFlowQuery
	withCreator      *UserQuery
	withCreatedUser  *UserQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...OrderFunc) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryBindAccount chains the current query on the "bind_account" edge.
func (uq *UserQuery) QueryBindAccount() *AccountQuery {
	query := (&AccountClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.BindAccountTable, user.BindAccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPurchasedApp chains the current query on the "purchased_app" edge.
func (uq *UserQuery) QueryPurchasedApp() *AppQuery {
	query := (&AppClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.PurchasedAppTable, user.PurchasedAppPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAppPackage chains the current query on the "app_package" edge.
func (uq *UserQuery) QueryAppPackage() *AppPackageQuery {
	query := (&AppPackageClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(apppackage.Table, apppackage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.AppPackageTable, user.AppPackageColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeedConfig chains the current query on the "feed_config" edge.
func (uq *UserQuery) QueryFeedConfig() *FeedConfigQuery {
	query := (&FeedConfigClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(feedconfig.Table, feedconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FeedConfigTable, user.FeedConfigColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyTarget chains the current query on the "notify_target" edge.
func (uq *UserQuery) QueryNotifyTarget() *NotifyTargetQuery {
	query := (&NotifyTargetClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(notifytarget.Table, notifytarget.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.NotifyTargetTable, user.NotifyTargetColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyFlow chains the current query on the "notify_flow" edge.
func (uq *UserQuery) QueryNotifyFlow() *NotifyFlowQuery {
	query := (&NotifyFlowClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(notifyflow.Table, notifyflow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.NotifyFlowTable, user.NotifyFlowColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreator chains the current query on the "creator" edge.
func (uq *UserQuery) QueryCreator() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.CreatorTable, user.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedUser chains the current query on the "created_user" edge.
func (uq *UserQuery) QueryCreatedUser() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CreatedUserTable, user.CreatedUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, uq.ctx, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, "IDs")
	if err = uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:           uq.config,
		ctx:              uq.ctx.Clone(),
		order:            append([]OrderFunc{}, uq.order...),
		inters:           append([]Interceptor{}, uq.inters...),
		predicates:       append([]predicate.User{}, uq.predicates...),
		withBindAccount:  uq.withBindAccount.Clone(),
		withPurchasedApp: uq.withPurchasedApp.Clone(),
		withAppPackage:   uq.withAppPackage.Clone(),
		withFeedConfig:   uq.withFeedConfig.Clone(),
		withNotifyTarget: uq.withNotifyTarget.Clone(),
		withNotifyFlow:   uq.withNotifyFlow.Clone(),
		withCreator:      uq.withCreator.Clone(),
		withCreatedUser:  uq.withCreatedUser.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithBindAccount tells the query-builder to eager-load the nodes that are connected to
// the "bind_account" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithBindAccount(opts ...func(*AccountQuery)) *UserQuery {
	query := (&AccountClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withBindAccount = query
	return uq
}

// WithPurchasedApp tells the query-builder to eager-load the nodes that are connected to
// the "purchased_app" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithPurchasedApp(opts ...func(*AppQuery)) *UserQuery {
	query := (&AppClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withPurchasedApp = query
	return uq
}

// WithAppPackage tells the query-builder to eager-load the nodes that are connected to
// the "app_package" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithAppPackage(opts ...func(*AppPackageQuery)) *UserQuery {
	query := (&AppPackageClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withAppPackage = query
	return uq
}

// WithFeedConfig tells the query-builder to eager-load the nodes that are connected to
// the "feed_config" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithFeedConfig(opts ...func(*FeedConfigQuery)) *UserQuery {
	query := (&FeedConfigClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withFeedConfig = query
	return uq
}

// WithNotifyTarget tells the query-builder to eager-load the nodes that are connected to
// the "notify_target" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNotifyTarget(opts ...func(*NotifyTargetQuery)) *UserQuery {
	query := (&NotifyTargetClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withNotifyTarget = query
	return uq
}

// WithNotifyFlow tells the query-builder to eager-load the nodes that are connected to
// the "notify_flow" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNotifyFlow(opts ...func(*NotifyFlowQuery)) *UserQuery {
	query := (&NotifyFlowClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withNotifyFlow = query
	return uq
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithCreator(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withCreator = query
	return uq
}

// WithCreatedUser tells the query-builder to eager-load the nodes that are connected to
// the "created_user" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithCreatedUser(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withCreatedUser = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldUsername).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		withFKs     = uq.withFKs
		_spec       = uq.querySpec()
		loadedTypes = [8]bool{
			uq.withBindAccount != nil,
			uq.withPurchasedApp != nil,
			uq.withAppPackage != nil,
			uq.withFeedConfig != nil,
			uq.withNotifyTarget != nil,
			uq.withNotifyFlow != nil,
			uq.withCreator != nil,
			uq.withCreatedUser != nil,
		}
	)
	if uq.withCreator != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, user.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withBindAccount; query != nil {
		if err := uq.loadBindAccount(ctx, query, nodes,
			func(n *User) { n.Edges.BindAccount = []*Account{} },
			func(n *User, e *Account) { n.Edges.BindAccount = append(n.Edges.BindAccount, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withPurchasedApp; query != nil {
		if err := uq.loadPurchasedApp(ctx, query, nodes,
			func(n *User) { n.Edges.PurchasedApp = []*App{} },
			func(n *User, e *App) { n.Edges.PurchasedApp = append(n.Edges.PurchasedApp, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withAppPackage; query != nil {
		if err := uq.loadAppPackage(ctx, query, nodes,
			func(n *User) { n.Edges.AppPackage = []*AppPackage{} },
			func(n *User, e *AppPackage) { n.Edges.AppPackage = append(n.Edges.AppPackage, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withFeedConfig; query != nil {
		if err := uq.loadFeedConfig(ctx, query, nodes,
			func(n *User) { n.Edges.FeedConfig = []*FeedConfig{} },
			func(n *User, e *FeedConfig) { n.Edges.FeedConfig = append(n.Edges.FeedConfig, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withNotifyTarget; query != nil {
		if err := uq.loadNotifyTarget(ctx, query, nodes,
			func(n *User) { n.Edges.NotifyTarget = []*NotifyTarget{} },
			func(n *User, e *NotifyTarget) { n.Edges.NotifyTarget = append(n.Edges.NotifyTarget, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withNotifyFlow; query != nil {
		if err := uq.loadNotifyFlow(ctx, query, nodes,
			func(n *User) { n.Edges.NotifyFlow = []*NotifyFlow{} },
			func(n *User, e *NotifyFlow) { n.Edges.NotifyFlow = append(n.Edges.NotifyFlow, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withCreator; query != nil {
		if err := uq.loadCreator(ctx, query, nodes, nil,
			func(n *User, e *User) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withCreatedUser; query != nil {
		if err := uq.loadCreatedUser(ctx, query, nodes,
			func(n *User) { n.Edges.CreatedUser = []*User{} },
			func(n *User, e *User) { n.Edges.CreatedUser = append(n.Edges.CreatedUser, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadBindAccount(ctx context.Context, query *AccountQuery, nodes []*User, init func(*User), assign func(*User, *Account)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Account(func(s *sql.Selector) {
		s.Where(sql.InValues(user.BindAccountColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_bind_account
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_bind_account" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_bind_account" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadPurchasedApp(ctx context.Context, query *AppQuery, nodes []*User, init func(*User), assign func(*User, *App)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[model.InternalID]*User)
	nids := make(map[model.InternalID]map[*User]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(user.PurchasedAppTable)
		s.Join(joinT).On(s.C(app.FieldID), joinT.C(user.PurchasedAppPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(user.PurchasedAppPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(user.PurchasedAppPrimaryKey[0]))
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
				outValue := model.InternalID(values[0].(*sql.NullInt64).Int64)
				inValue := model.InternalID(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*User]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*App](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "purchased_app" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (uq *UserQuery) loadAppPackage(ctx context.Context, query *AppPackageQuery, nodes []*User, init func(*User), assign func(*User, *AppPackage)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.InValues(user.AppPackageColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_app_package
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_app_package" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_app_package" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadFeedConfig(ctx context.Context, query *FeedConfigQuery, nodes []*User, init func(*User), assign func(*User, *FeedConfig)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.FeedConfig(func(s *sql.Selector) {
		s.Where(sql.InValues(user.FeedConfigColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_feed_config
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_feed_config" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_feed_config" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadNotifyTarget(ctx context.Context, query *NotifyTargetQuery, nodes []*User, init func(*User), assign func(*User, *NotifyTarget)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.NotifyTarget(func(s *sql.Selector) {
		s.Where(sql.InValues(user.NotifyTargetColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_notify_target
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_notify_target" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_notify_target" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadNotifyFlow(ctx context.Context, query *NotifyFlowQuery, nodes []*User, init func(*User), assign func(*User, *NotifyFlow)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.NotifyFlow(func(s *sql.Selector) {
		s.Where(sql.InValues(user.NotifyFlowColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_notify_flow
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_notify_flow" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_notify_flow" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadCreator(ctx context.Context, query *UserQuery, nodes []*User, init func(*User), assign func(*User, *User)) error {
	ids := make([]model.InternalID, 0, len(nodes))
	nodeids := make(map[model.InternalID][]*User)
	for i := range nodes {
		if nodes[i].user_created_user == nil {
			continue
		}
		fk := *nodes[i].user_created_user
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
			return fmt.Errorf(`unexpected foreign-key "user_created_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uq *UserQuery) loadCreatedUser(ctx context.Context, query *UserQuery, nodes []*User, init func(*User), assign func(*User, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(user.CreatedUserColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_created_user
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_created_user" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_created_user" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
