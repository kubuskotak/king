// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/kubuskotak/king/pkg/persist/crud/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/kubuskotak/king/pkg/persist/crud/ent/article"
	"github.com/kubuskotak/king/pkg/persist/crud/ent/ymir"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Article is the client for interacting with the Article builders.
	Article *ArticleClient
	// Ymir is the client for interacting with the Ymir builders.
	Ymir *YmirClient
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
	c.Article = NewArticleClient(c.config)
	c.Ymir = NewYmirClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		ctx:     ctx,
		config:  cfg,
		Article: NewArticleClient(cfg),
		Ymir:    NewYmirClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Article: NewArticleClient(cfg),
		Ymir:    NewYmirClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Article.
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
	c.Article.Use(hooks...)
	c.Ymir.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Article.Intercept(interceptors...)
	c.Ymir.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ArticleMutation:
		return c.Article.mutate(ctx, m)
	case *YmirMutation:
		return c.Ymir.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ArticleClient is a client for the Article schema.
type ArticleClient struct {
	config
}

// NewArticleClient returns a client for the Article from the given config.
func NewArticleClient(c config) *ArticleClient {
	return &ArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `article.Hooks(f(g(h())))`.
func (c *ArticleClient) Use(hooks ...Hook) {
	c.hooks.Article = append(c.hooks.Article, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `article.Intercept(f(g(h())))`.
func (c *ArticleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Article = append(c.inters.Article, interceptors...)
}

// Create returns a builder for creating a Article entity.
func (c *ArticleClient) Create() *ArticleCreate {
	mutation := newArticleMutation(c.config, OpCreate)
	return &ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Article entities.
func (c *ArticleClient) CreateBulk(builders ...*ArticleCreate) *ArticleCreateBulk {
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Article.
func (c *ArticleClient) Update() *ArticleUpdate {
	mutation := newArticleMutation(c.config, OpUpdate)
	return &ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleClient) UpdateOne(a *Article) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticle(a))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleClient) UpdateOneID(id int) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticleID(id))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Article.
func (c *ArticleClient) Delete() *ArticleDelete {
	mutation := newArticleMutation(c.config, OpDelete)
	return &ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArticleClient) DeleteOne(a *Article) *ArticleDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ArticleClient) DeleteOneID(id int) *ArticleDeleteOne {
	builder := c.Delete().Where(article.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleDeleteOne{builder}
}

// Query returns a query builder for Article.
func (c *ArticleClient) Query() *ArticleQuery {
	return &ArticleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeArticle},
		inters: c.Interceptors(),
	}
}

// Get returns a Article entity by its id.
func (c *ArticleClient) Get(ctx context.Context, id int) (*Article, error) {
	return c.Query().Where(article.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleClient) GetX(ctx context.Context, id int) *Article {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ArticleClient) Hooks() []Hook {
	return c.hooks.Article
}

// Interceptors returns the client interceptors.
func (c *ArticleClient) Interceptors() []Interceptor {
	return c.inters.Article
}

func (c *ArticleClient) mutate(ctx context.Context, m *ArticleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Article mutation op: %q", m.Op())
	}
}

// YmirClient is a client for the Ymir schema.
type YmirClient struct {
	config
}

// NewYmirClient returns a client for the Ymir from the given config.
func NewYmirClient(c config) *YmirClient {
	return &YmirClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ymir.Hooks(f(g(h())))`.
func (c *YmirClient) Use(hooks ...Hook) {
	c.hooks.Ymir = append(c.hooks.Ymir, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `ymir.Intercept(f(g(h())))`.
func (c *YmirClient) Intercept(interceptors ...Interceptor) {
	c.inters.Ymir = append(c.inters.Ymir, interceptors...)
}

// Create returns a builder for creating a Ymir entity.
func (c *YmirClient) Create() *YmirCreate {
	mutation := newYmirMutation(c.config, OpCreate)
	return &YmirCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ymir entities.
func (c *YmirClient) CreateBulk(builders ...*YmirCreate) *YmirCreateBulk {
	return &YmirCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ymir.
func (c *YmirClient) Update() *YmirUpdate {
	mutation := newYmirMutation(c.config, OpUpdate)
	return &YmirUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *YmirClient) UpdateOne(y *Ymir) *YmirUpdateOne {
	mutation := newYmirMutation(c.config, OpUpdateOne, withYmir(y))
	return &YmirUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *YmirClient) UpdateOneID(id int) *YmirUpdateOne {
	mutation := newYmirMutation(c.config, OpUpdateOne, withYmirID(id))
	return &YmirUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ymir.
func (c *YmirClient) Delete() *YmirDelete {
	mutation := newYmirMutation(c.config, OpDelete)
	return &YmirDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *YmirClient) DeleteOne(y *Ymir) *YmirDeleteOne {
	return c.DeleteOneID(y.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *YmirClient) DeleteOneID(id int) *YmirDeleteOne {
	builder := c.Delete().Where(ymir.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &YmirDeleteOne{builder}
}

// Query returns a query builder for Ymir.
func (c *YmirClient) Query() *YmirQuery {
	return &YmirQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeYmir},
		inters: c.Interceptors(),
	}
}

// Get returns a Ymir entity by its id.
func (c *YmirClient) Get(ctx context.Context, id int) (*Ymir, error) {
	return c.Query().Where(ymir.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *YmirClient) GetX(ctx context.Context, id int) *Ymir {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *YmirClient) Hooks() []Hook {
	return c.hooks.Ymir
}

// Interceptors returns the client interceptors.
func (c *YmirClient) Interceptors() []Interceptor {
	return c.inters.Ymir
}

func (c *YmirClient) mutate(ctx context.Context, m *YmirMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&YmirCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&YmirUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&YmirUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&YmirDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Ymir mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Article, Ymir []ent.Hook
	}
	inters struct {
		Article, Ymir []ent.Interceptor
	}
)
