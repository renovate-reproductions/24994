// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/tuihub/librarian/app/sephirah/internal/ent/migrate"

	"github.com/tuihub/librarian/app/sephirah/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// App is the client for interacting with the App builders.
	App *AppClient
	// AppPackage is the client for interacting with the AppPackage builders.
	AppPackage *AppPackageClient
	// Feed is the client for interacting with the Feed builders.
	Feed *FeedClient
	// FeedConfig is the client for interacting with the FeedConfig builders.
	FeedConfig *FeedConfigClient
	// FeedItem is the client for interacting with the FeedItem builders.
	FeedItem *FeedItemClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.App = NewAppClient(c.config)
	c.AppPackage = NewAppPackageClient(c.config)
	c.Feed = NewFeedClient(c.config)
	c.FeedConfig = NewFeedConfigClient(c.config)
	c.FeedItem = NewFeedItemClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Account:    NewAccountClient(cfg),
		App:        NewAppClient(cfg),
		AppPackage: NewAppPackageClient(cfg),
		Feed:       NewFeedClient(cfg),
		FeedConfig: NewFeedConfigClient(cfg),
		FeedItem:   NewFeedItemClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Account:    NewAccountClient(cfg),
		App:        NewAppClient(cfg),
		AppPackage: NewAppPackageClient(cfg),
		Feed:       NewFeedClient(cfg),
		FeedConfig: NewFeedConfigClient(cfg),
		FeedItem:   NewFeedItemClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.App.Use(hooks...)
	c.AppPackage.Use(hooks...)
	c.Feed.Use(hooks...)
	c.FeedConfig.Use(hooks...)
	c.FeedItem.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Account.Intercept(interceptors...)
	c.App.Intercept(interceptors...)
	c.AppPackage.Intercept(interceptors...)
	c.Feed.Intercept(interceptors...)
	c.FeedConfig.Intercept(interceptors...)
	c.FeedItem.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AccountMutation:
		return c.Account.mutate(ctx, m)
	case *AppMutation:
		return c.App.mutate(ctx, m)
	case *AppPackageMutation:
		return c.AppPackage.mutate(ctx, m)
	case *FeedMutation:
		return c.Feed.mutate(ctx, m)
	case *FeedConfigMutation:
		return c.FeedConfig.mutate(ctx, m)
	case *FeedItemMutation:
		return c.FeedItem.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `account.Intercept(f(g(h())))`.
func (c *AccountClient) Intercept(interceptors ...Interceptor) {
	c.inters.Account = append(c.inters.Account, interceptors...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id int) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id int) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAccount},
		inters: c.Interceptors(),
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id int) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id int) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// Interceptors returns the client interceptors.
func (c *AccountClient) Interceptors() []Interceptor {
	return c.inters.Account
}

func (c *AccountClient) mutate(ctx context.Context, m *AccountMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AccountCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AccountDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Account mutation op: %q", m.Op())
	}
}

// AppClient is a client for the App schema.
type AppClient struct {
	config
}

// NewAppClient returns a client for the App from the given config.
func NewAppClient(c config) *AppClient {
	return &AppClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `app.Hooks(f(g(h())))`.
func (c *AppClient) Use(hooks ...Hook) {
	c.hooks.App = append(c.hooks.App, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `app.Intercept(f(g(h())))`.
func (c *AppClient) Intercept(interceptors ...Interceptor) {
	c.inters.App = append(c.inters.App, interceptors...)
}

// Create returns a builder for creating a App entity.
func (c *AppClient) Create() *AppCreate {
	mutation := newAppMutation(c.config, OpCreate)
	return &AppCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of App entities.
func (c *AppClient) CreateBulk(builders ...*AppCreate) *AppCreateBulk {
	return &AppCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for App.
func (c *AppClient) Update() *AppUpdate {
	mutation := newAppMutation(c.config, OpUpdate)
	return &AppUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppClient) UpdateOne(a *App) *AppUpdateOne {
	mutation := newAppMutation(c.config, OpUpdateOne, withApp(a))
	return &AppUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppClient) UpdateOneID(id int) *AppUpdateOne {
	mutation := newAppMutation(c.config, OpUpdateOne, withAppID(id))
	return &AppUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for App.
func (c *AppClient) Delete() *AppDelete {
	mutation := newAppMutation(c.config, OpDelete)
	return &AppDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AppClient) DeleteOne(a *App) *AppDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AppClient) DeleteOneID(id int) *AppDeleteOne {
	builder := c.Delete().Where(app.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppDeleteOne{builder}
}

// Query returns a query builder for App.
func (c *AppClient) Query() *AppQuery {
	return &AppQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeApp},
		inters: c.Interceptors(),
	}
}

// Get returns a App entity by its id.
func (c *AppClient) Get(ctx context.Context, id int) (*App, error) {
	return c.Query().Where(app.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppClient) GetX(ctx context.Context, id int) *App {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppClient) Hooks() []Hook {
	return c.hooks.App
}

// Interceptors returns the client interceptors.
func (c *AppClient) Interceptors() []Interceptor {
	return c.inters.App
}

func (c *AppClient) mutate(ctx context.Context, m *AppMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AppCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AppUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AppUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AppDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown App mutation op: %q", m.Op())
	}
}

// AppPackageClient is a client for the AppPackage schema.
type AppPackageClient struct {
	config
}

// NewAppPackageClient returns a client for the AppPackage from the given config.
func NewAppPackageClient(c config) *AppPackageClient {
	return &AppPackageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `apppackage.Hooks(f(g(h())))`.
func (c *AppPackageClient) Use(hooks ...Hook) {
	c.hooks.AppPackage = append(c.hooks.AppPackage, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `apppackage.Intercept(f(g(h())))`.
func (c *AppPackageClient) Intercept(interceptors ...Interceptor) {
	c.inters.AppPackage = append(c.inters.AppPackage, interceptors...)
}

// Create returns a builder for creating a AppPackage entity.
func (c *AppPackageClient) Create() *AppPackageCreate {
	mutation := newAppPackageMutation(c.config, OpCreate)
	return &AppPackageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AppPackage entities.
func (c *AppPackageClient) CreateBulk(builders ...*AppPackageCreate) *AppPackageCreateBulk {
	return &AppPackageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AppPackage.
func (c *AppPackageClient) Update() *AppPackageUpdate {
	mutation := newAppPackageMutation(c.config, OpUpdate)
	return &AppPackageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppPackageClient) UpdateOne(ap *AppPackage) *AppPackageUpdateOne {
	mutation := newAppPackageMutation(c.config, OpUpdateOne, withAppPackage(ap))
	return &AppPackageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppPackageClient) UpdateOneID(id int) *AppPackageUpdateOne {
	mutation := newAppPackageMutation(c.config, OpUpdateOne, withAppPackageID(id))
	return &AppPackageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AppPackage.
func (c *AppPackageClient) Delete() *AppPackageDelete {
	mutation := newAppPackageMutation(c.config, OpDelete)
	return &AppPackageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AppPackageClient) DeleteOne(ap *AppPackage) *AppPackageDeleteOne {
	return c.DeleteOneID(ap.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AppPackageClient) DeleteOneID(id int) *AppPackageDeleteOne {
	builder := c.Delete().Where(apppackage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppPackageDeleteOne{builder}
}

// Query returns a query builder for AppPackage.
func (c *AppPackageClient) Query() *AppPackageQuery {
	return &AppPackageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAppPackage},
		inters: c.Interceptors(),
	}
}

// Get returns a AppPackage entity by its id.
func (c *AppPackageClient) Get(ctx context.Context, id int) (*AppPackage, error) {
	return c.Query().Where(apppackage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppPackageClient) GetX(ctx context.Context, id int) *AppPackage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppPackageClient) Hooks() []Hook {
	return c.hooks.AppPackage
}

// Interceptors returns the client interceptors.
func (c *AppPackageClient) Interceptors() []Interceptor {
	return c.inters.AppPackage
}

func (c *AppPackageClient) mutate(ctx context.Context, m *AppPackageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AppPackageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AppPackageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AppPackageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AppPackageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown AppPackage mutation op: %q", m.Op())
	}
}

// FeedClient is a client for the Feed schema.
type FeedClient struct {
	config
}

// NewFeedClient returns a client for the Feed from the given config.
func NewFeedClient(c config) *FeedClient {
	return &FeedClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `feed.Hooks(f(g(h())))`.
func (c *FeedClient) Use(hooks ...Hook) {
	c.hooks.Feed = append(c.hooks.Feed, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `feed.Intercept(f(g(h())))`.
func (c *FeedClient) Intercept(interceptors ...Interceptor) {
	c.inters.Feed = append(c.inters.Feed, interceptors...)
}

// Create returns a builder for creating a Feed entity.
func (c *FeedClient) Create() *FeedCreate {
	mutation := newFeedMutation(c.config, OpCreate)
	return &FeedCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Feed entities.
func (c *FeedClient) CreateBulk(builders ...*FeedCreate) *FeedCreateBulk {
	return &FeedCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Feed.
func (c *FeedClient) Update() *FeedUpdate {
	mutation := newFeedMutation(c.config, OpUpdate)
	return &FeedUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FeedClient) UpdateOne(f *Feed) *FeedUpdateOne {
	mutation := newFeedMutation(c.config, OpUpdateOne, withFeed(f))
	return &FeedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FeedClient) UpdateOneID(id int) *FeedUpdateOne {
	mutation := newFeedMutation(c.config, OpUpdateOne, withFeedID(id))
	return &FeedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Feed.
func (c *FeedClient) Delete() *FeedDelete {
	mutation := newFeedMutation(c.config, OpDelete)
	return &FeedDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FeedClient) DeleteOne(f *Feed) *FeedDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FeedClient) DeleteOneID(id int) *FeedDeleteOne {
	builder := c.Delete().Where(feed.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FeedDeleteOne{builder}
}

// Query returns a query builder for Feed.
func (c *FeedClient) Query() *FeedQuery {
	return &FeedQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFeed},
		inters: c.Interceptors(),
	}
}

// Get returns a Feed entity by its id.
func (c *FeedClient) Get(ctx context.Context, id int) (*Feed, error) {
	return c.Query().Where(feed.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FeedClient) GetX(ctx context.Context, id int) *Feed {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FeedClient) Hooks() []Hook {
	return c.hooks.Feed
}

// Interceptors returns the client interceptors.
func (c *FeedClient) Interceptors() []Interceptor {
	return c.inters.Feed
}

func (c *FeedClient) mutate(ctx context.Context, m *FeedMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FeedCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FeedUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FeedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FeedDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Feed mutation op: %q", m.Op())
	}
}

// FeedConfigClient is a client for the FeedConfig schema.
type FeedConfigClient struct {
	config
}

// NewFeedConfigClient returns a client for the FeedConfig from the given config.
func NewFeedConfigClient(c config) *FeedConfigClient {
	return &FeedConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `feedconfig.Hooks(f(g(h())))`.
func (c *FeedConfigClient) Use(hooks ...Hook) {
	c.hooks.FeedConfig = append(c.hooks.FeedConfig, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `feedconfig.Intercept(f(g(h())))`.
func (c *FeedConfigClient) Intercept(interceptors ...Interceptor) {
	c.inters.FeedConfig = append(c.inters.FeedConfig, interceptors...)
}

// Create returns a builder for creating a FeedConfig entity.
func (c *FeedConfigClient) Create() *FeedConfigCreate {
	mutation := newFeedConfigMutation(c.config, OpCreate)
	return &FeedConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FeedConfig entities.
func (c *FeedConfigClient) CreateBulk(builders ...*FeedConfigCreate) *FeedConfigCreateBulk {
	return &FeedConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FeedConfig.
func (c *FeedConfigClient) Update() *FeedConfigUpdate {
	mutation := newFeedConfigMutation(c.config, OpUpdate)
	return &FeedConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FeedConfigClient) UpdateOne(fc *FeedConfig) *FeedConfigUpdateOne {
	mutation := newFeedConfigMutation(c.config, OpUpdateOne, withFeedConfig(fc))
	return &FeedConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FeedConfigClient) UpdateOneID(id int) *FeedConfigUpdateOne {
	mutation := newFeedConfigMutation(c.config, OpUpdateOne, withFeedConfigID(id))
	return &FeedConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FeedConfig.
func (c *FeedConfigClient) Delete() *FeedConfigDelete {
	mutation := newFeedConfigMutation(c.config, OpDelete)
	return &FeedConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FeedConfigClient) DeleteOne(fc *FeedConfig) *FeedConfigDeleteOne {
	return c.DeleteOneID(fc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FeedConfigClient) DeleteOneID(id int) *FeedConfigDeleteOne {
	builder := c.Delete().Where(feedconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FeedConfigDeleteOne{builder}
}

// Query returns a query builder for FeedConfig.
func (c *FeedConfigClient) Query() *FeedConfigQuery {
	return &FeedConfigQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFeedConfig},
		inters: c.Interceptors(),
	}
}

// Get returns a FeedConfig entity by its id.
func (c *FeedConfigClient) Get(ctx context.Context, id int) (*FeedConfig, error) {
	return c.Query().Where(feedconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FeedConfigClient) GetX(ctx context.Context, id int) *FeedConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FeedConfigClient) Hooks() []Hook {
	return c.hooks.FeedConfig
}

// Interceptors returns the client interceptors.
func (c *FeedConfigClient) Interceptors() []Interceptor {
	return c.inters.FeedConfig
}

func (c *FeedConfigClient) mutate(ctx context.Context, m *FeedConfigMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FeedConfigCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FeedConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FeedConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FeedConfigDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown FeedConfig mutation op: %q", m.Op())
	}
}

// FeedItemClient is a client for the FeedItem schema.
type FeedItemClient struct {
	config
}

// NewFeedItemClient returns a client for the FeedItem from the given config.
func NewFeedItemClient(c config) *FeedItemClient {
	return &FeedItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `feeditem.Hooks(f(g(h())))`.
func (c *FeedItemClient) Use(hooks ...Hook) {
	c.hooks.FeedItem = append(c.hooks.FeedItem, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `feeditem.Intercept(f(g(h())))`.
func (c *FeedItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.FeedItem = append(c.inters.FeedItem, interceptors...)
}

// Create returns a builder for creating a FeedItem entity.
func (c *FeedItemClient) Create() *FeedItemCreate {
	mutation := newFeedItemMutation(c.config, OpCreate)
	return &FeedItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FeedItem entities.
func (c *FeedItemClient) CreateBulk(builders ...*FeedItemCreate) *FeedItemCreateBulk {
	return &FeedItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FeedItem.
func (c *FeedItemClient) Update() *FeedItemUpdate {
	mutation := newFeedItemMutation(c.config, OpUpdate)
	return &FeedItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FeedItemClient) UpdateOne(fi *FeedItem) *FeedItemUpdateOne {
	mutation := newFeedItemMutation(c.config, OpUpdateOne, withFeedItem(fi))
	return &FeedItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FeedItemClient) UpdateOneID(id int) *FeedItemUpdateOne {
	mutation := newFeedItemMutation(c.config, OpUpdateOne, withFeedItemID(id))
	return &FeedItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FeedItem.
func (c *FeedItemClient) Delete() *FeedItemDelete {
	mutation := newFeedItemMutation(c.config, OpDelete)
	return &FeedItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FeedItemClient) DeleteOne(fi *FeedItem) *FeedItemDeleteOne {
	return c.DeleteOneID(fi.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FeedItemClient) DeleteOneID(id int) *FeedItemDeleteOne {
	builder := c.Delete().Where(feeditem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FeedItemDeleteOne{builder}
}

// Query returns a query builder for FeedItem.
func (c *FeedItemClient) Query() *FeedItemQuery {
	return &FeedItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFeedItem},
		inters: c.Interceptors(),
	}
}

// Get returns a FeedItem entity by its id.
func (c *FeedItemClient) Get(ctx context.Context, id int) (*FeedItem, error) {
	return c.Query().Where(feeditem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FeedItemClient) GetX(ctx context.Context, id int) *FeedItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FeedItemClient) Hooks() []Hook {
	return c.hooks.FeedItem
}

// Interceptors returns the client interceptors.
func (c *FeedItemClient) Interceptors() []Interceptor {
	return c.inters.FeedItem
}

func (c *FeedItemClient) mutate(ctx context.Context, m *FeedItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FeedItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FeedItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FeedItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FeedItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown FeedItem mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}
