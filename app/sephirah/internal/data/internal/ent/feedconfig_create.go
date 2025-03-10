// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// FeedConfigCreate is the builder for creating a FeedConfig entity.
type FeedConfigCreate struct {
	config
	mutation *FeedConfigMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserFeedConfig sets the "user_feed_config" field.
func (fcc *FeedConfigCreate) SetUserFeedConfig(mi model.InternalID) *FeedConfigCreate {
	fcc.mutation.SetUserFeedConfig(mi)
	return fcc
}

// SetName sets the "name" field.
func (fcc *FeedConfigCreate) SetName(s string) *FeedConfigCreate {
	fcc.mutation.SetName(s)
	return fcc
}

// SetFeedURL sets the "feed_url" field.
func (fcc *FeedConfigCreate) SetFeedURL(s string) *FeedConfigCreate {
	fcc.mutation.SetFeedURL(s)
	return fcc
}

// SetAuthorAccount sets the "author_account" field.
func (fcc *FeedConfigCreate) SetAuthorAccount(mi model.InternalID) *FeedConfigCreate {
	fcc.mutation.SetAuthorAccount(mi)
	return fcc
}

// SetSource sets the "source" field.
func (fcc *FeedConfigCreate) SetSource(f feedconfig.Source) *FeedConfigCreate {
	fcc.mutation.SetSource(f)
	return fcc
}

// SetStatus sets the "status" field.
func (fcc *FeedConfigCreate) SetStatus(f feedconfig.Status) *FeedConfigCreate {
	fcc.mutation.SetStatus(f)
	return fcc
}

// SetCategory sets the "category" field.
func (fcc *FeedConfigCreate) SetCategory(s string) *FeedConfigCreate {
	fcc.mutation.SetCategory(s)
	return fcc
}

// SetPullInterval sets the "pull_interval" field.
func (fcc *FeedConfigCreate) SetPullInterval(t time.Duration) *FeedConfigCreate {
	fcc.mutation.SetPullInterval(t)
	return fcc
}

// SetLatestPullAt sets the "latest_pull_at" field.
func (fcc *FeedConfigCreate) SetLatestPullAt(t time.Time) *FeedConfigCreate {
	fcc.mutation.SetLatestPullAt(t)
	return fcc
}

// SetNillableLatestPullAt sets the "latest_pull_at" field if the given value is not nil.
func (fcc *FeedConfigCreate) SetNillableLatestPullAt(t *time.Time) *FeedConfigCreate {
	if t != nil {
		fcc.SetLatestPullAt(*t)
	}
	return fcc
}

// SetNextPullBeginAt sets the "next_pull_begin_at" field.
func (fcc *FeedConfigCreate) SetNextPullBeginAt(t time.Time) *FeedConfigCreate {
	fcc.mutation.SetNextPullBeginAt(t)
	return fcc
}

// SetNillableNextPullBeginAt sets the "next_pull_begin_at" field if the given value is not nil.
func (fcc *FeedConfigCreate) SetNillableNextPullBeginAt(t *time.Time) *FeedConfigCreate {
	if t != nil {
		fcc.SetNextPullBeginAt(*t)
	}
	return fcc
}

// SetUpdatedAt sets the "updated_at" field.
func (fcc *FeedConfigCreate) SetUpdatedAt(t time.Time) *FeedConfigCreate {
	fcc.mutation.SetUpdatedAt(t)
	return fcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fcc *FeedConfigCreate) SetNillableUpdatedAt(t *time.Time) *FeedConfigCreate {
	if t != nil {
		fcc.SetUpdatedAt(*t)
	}
	return fcc
}

// SetCreatedAt sets the "created_at" field.
func (fcc *FeedConfigCreate) SetCreatedAt(t time.Time) *FeedConfigCreate {
	fcc.mutation.SetCreatedAt(t)
	return fcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fcc *FeedConfigCreate) SetNillableCreatedAt(t *time.Time) *FeedConfigCreate {
	if t != nil {
		fcc.SetCreatedAt(*t)
	}
	return fcc
}

// SetID sets the "id" field.
func (fcc *FeedConfigCreate) SetID(mi model.InternalID) *FeedConfigCreate {
	fcc.mutation.SetID(mi)
	return fcc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fcc *FeedConfigCreate) SetOwnerID(id model.InternalID) *FeedConfigCreate {
	fcc.mutation.SetOwnerID(id)
	return fcc
}

// SetOwner sets the "owner" edge to the User entity.
func (fcc *FeedConfigCreate) SetOwner(u *User) *FeedConfigCreate {
	return fcc.SetOwnerID(u.ID)
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (fcc *FeedConfigCreate) SetFeedID(id model.InternalID) *FeedConfigCreate {
	fcc.mutation.SetFeedID(id)
	return fcc
}

// SetNillableFeedID sets the "feed" edge to the Feed entity by ID if the given value is not nil.
func (fcc *FeedConfigCreate) SetNillableFeedID(id *model.InternalID) *FeedConfigCreate {
	if id != nil {
		fcc = fcc.SetFeedID(*id)
	}
	return fcc
}

// SetFeed sets the "feed" edge to the Feed entity.
func (fcc *FeedConfigCreate) SetFeed(f *Feed) *FeedConfigCreate {
	return fcc.SetFeedID(f.ID)
}

// AddNotifyFlowIDs adds the "notify_flow" edge to the NotifyFlow entity by IDs.
func (fcc *FeedConfigCreate) AddNotifyFlowIDs(ids ...model.InternalID) *FeedConfigCreate {
	fcc.mutation.AddNotifyFlowIDs(ids...)
	return fcc
}

// AddNotifyFlow adds the "notify_flow" edges to the NotifyFlow entity.
func (fcc *FeedConfigCreate) AddNotifyFlow(n ...*NotifyFlow) *FeedConfigCreate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return fcc.AddNotifyFlowIDs(ids...)
}

// Mutation returns the FeedConfigMutation object of the builder.
func (fcc *FeedConfigCreate) Mutation() *FeedConfigMutation {
	return fcc.mutation
}

// Save creates the FeedConfig in the database.
func (fcc *FeedConfigCreate) Save(ctx context.Context) (*FeedConfig, error) {
	fcc.defaults()
	return withHooks(ctx, fcc.sqlSave, fcc.mutation, fcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fcc *FeedConfigCreate) SaveX(ctx context.Context) *FeedConfig {
	v, err := fcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcc *FeedConfigCreate) Exec(ctx context.Context) error {
	_, err := fcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcc *FeedConfigCreate) ExecX(ctx context.Context) {
	if err := fcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcc *FeedConfigCreate) defaults() {
	if _, ok := fcc.mutation.LatestPullAt(); !ok {
		v := feedconfig.DefaultLatestPullAt
		fcc.mutation.SetLatestPullAt(v)
	}
	if _, ok := fcc.mutation.NextPullBeginAt(); !ok {
		v := feedconfig.DefaultNextPullBeginAt
		fcc.mutation.SetNextPullBeginAt(v)
	}
	if _, ok := fcc.mutation.UpdatedAt(); !ok {
		v := feedconfig.DefaultUpdatedAt()
		fcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fcc.mutation.CreatedAt(); !ok {
		v := feedconfig.DefaultCreatedAt()
		fcc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fcc *FeedConfigCreate) check() error {
	if _, ok := fcc.mutation.UserFeedConfig(); !ok {
		return &ValidationError{Name: "user_feed_config", err: errors.New(`ent: missing required field "FeedConfig.user_feed_config"`)}
	}
	if _, ok := fcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FeedConfig.name"`)}
	}
	if _, ok := fcc.mutation.FeedURL(); !ok {
		return &ValidationError{Name: "feed_url", err: errors.New(`ent: missing required field "FeedConfig.feed_url"`)}
	}
	if _, ok := fcc.mutation.AuthorAccount(); !ok {
		return &ValidationError{Name: "author_account", err: errors.New(`ent: missing required field "FeedConfig.author_account"`)}
	}
	if _, ok := fcc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required field "FeedConfig.source"`)}
	}
	if v, ok := fcc.mutation.Source(); ok {
		if err := feedconfig.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "FeedConfig.source": %w`, err)}
		}
	}
	if _, ok := fcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "FeedConfig.status"`)}
	}
	if v, ok := fcc.mutation.Status(); ok {
		if err := feedconfig.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "FeedConfig.status": %w`, err)}
		}
	}
	if _, ok := fcc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "FeedConfig.category"`)}
	}
	if _, ok := fcc.mutation.PullInterval(); !ok {
		return &ValidationError{Name: "pull_interval", err: errors.New(`ent: missing required field "FeedConfig.pull_interval"`)}
	}
	if _, ok := fcc.mutation.LatestPullAt(); !ok {
		return &ValidationError{Name: "latest_pull_at", err: errors.New(`ent: missing required field "FeedConfig.latest_pull_at"`)}
	}
	if _, ok := fcc.mutation.NextPullBeginAt(); !ok {
		return &ValidationError{Name: "next_pull_begin_at", err: errors.New(`ent: missing required field "FeedConfig.next_pull_begin_at"`)}
	}
	if _, ok := fcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FeedConfig.updated_at"`)}
	}
	if _, ok := fcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FeedConfig.created_at"`)}
	}
	if _, ok := fcc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "FeedConfig.owner"`)}
	}
	return nil
}

func (fcc *FeedConfigCreate) sqlSave(ctx context.Context) (*FeedConfig, error) {
	if err := fcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	fcc.mutation.id = &_node.ID
	fcc.mutation.done = true
	return _node, nil
}

func (fcc *FeedConfigCreate) createSpec() (*FeedConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &FeedConfig{config: fcc.config}
		_spec = sqlgraph.NewCreateSpec(feedconfig.Table, sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = fcc.conflict
	if id, ok := fcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fcc.mutation.Name(); ok {
		_spec.SetField(feedconfig.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fcc.mutation.FeedURL(); ok {
		_spec.SetField(feedconfig.FieldFeedURL, field.TypeString, value)
		_node.FeedURL = value
	}
	if value, ok := fcc.mutation.AuthorAccount(); ok {
		_spec.SetField(feedconfig.FieldAuthorAccount, field.TypeInt64, value)
		_node.AuthorAccount = value
	}
	if value, ok := fcc.mutation.Source(); ok {
		_spec.SetField(feedconfig.FieldSource, field.TypeEnum, value)
		_node.Source = value
	}
	if value, ok := fcc.mutation.Status(); ok {
		_spec.SetField(feedconfig.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := fcc.mutation.Category(); ok {
		_spec.SetField(feedconfig.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := fcc.mutation.PullInterval(); ok {
		_spec.SetField(feedconfig.FieldPullInterval, field.TypeInt64, value)
		_node.PullInterval = value
	}
	if value, ok := fcc.mutation.LatestPullAt(); ok {
		_spec.SetField(feedconfig.FieldLatestPullAt, field.TypeTime, value)
		_node.LatestPullAt = value
	}
	if value, ok := fcc.mutation.NextPullBeginAt(); ok {
		_spec.SetField(feedconfig.FieldNextPullBeginAt, field.TypeTime, value)
		_node.NextPullBeginAt = value
	}
	if value, ok := fcc.mutation.UpdatedAt(); ok {
		_spec.SetField(feedconfig.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fcc.mutation.CreatedAt(); ok {
		_spec.SetField(feedconfig.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := fcc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feedconfig.OwnerTable,
			Columns: []string{feedconfig.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserFeedConfig = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fcc.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   feedconfig.FeedTable,
			Columns: []string{feedconfig.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feed.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fcc.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feedconfig.NotifyFlowTable,
			Columns: feedconfig.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedConfig.Create().
//		SetUserFeedConfig(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedConfigUpsert) {
//			SetUserFeedConfig(v+v).
//		}).
//		Exec(ctx)
func (fcc *FeedConfigCreate) OnConflict(opts ...sql.ConflictOption) *FeedConfigUpsertOne {
	fcc.conflict = opts
	return &FeedConfigUpsertOne{
		create: fcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcc *FeedConfigCreate) OnConflictColumns(columns ...string) *FeedConfigUpsertOne {
	fcc.conflict = append(fcc.conflict, sql.ConflictColumns(columns...))
	return &FeedConfigUpsertOne{
		create: fcc,
	}
}

type (
	// FeedConfigUpsertOne is the builder for "upsert"-ing
	//  one FeedConfig node.
	FeedConfigUpsertOne struct {
		create *FeedConfigCreate
	}

	// FeedConfigUpsert is the "OnConflict" setter.
	FeedConfigUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserFeedConfig sets the "user_feed_config" field.
func (u *FeedConfigUpsert) SetUserFeedConfig(v model.InternalID) *FeedConfigUpsert {
	u.Set(feedconfig.FieldUserFeedConfig, v)
	return u
}

// UpdateUserFeedConfig sets the "user_feed_config" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateUserFeedConfig() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldUserFeedConfig)
	return u
}

// SetName sets the "name" field.
func (u *FeedConfigUpsert) SetName(v string) *FeedConfigUpsert {
	u.Set(feedconfig.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateName() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldName)
	return u
}

// SetFeedURL sets the "feed_url" field.
func (u *FeedConfigUpsert) SetFeedURL(v string) *FeedConfigUpsert {
	u.Set(feedconfig.FieldFeedURL, v)
	return u
}

// UpdateFeedURL sets the "feed_url" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateFeedURL() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldFeedURL)
	return u
}

// SetAuthorAccount sets the "author_account" field.
func (u *FeedConfigUpsert) SetAuthorAccount(v model.InternalID) *FeedConfigUpsert {
	u.Set(feedconfig.FieldAuthorAccount, v)
	return u
}

// UpdateAuthorAccount sets the "author_account" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateAuthorAccount() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldAuthorAccount)
	return u
}

// AddAuthorAccount adds v to the "author_account" field.
func (u *FeedConfigUpsert) AddAuthorAccount(v model.InternalID) *FeedConfigUpsert {
	u.Add(feedconfig.FieldAuthorAccount, v)
	return u
}

// SetSource sets the "source" field.
func (u *FeedConfigUpsert) SetSource(v feedconfig.Source) *FeedConfigUpsert {
	u.Set(feedconfig.FieldSource, v)
	return u
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateSource() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldSource)
	return u
}

// SetStatus sets the "status" field.
func (u *FeedConfigUpsert) SetStatus(v feedconfig.Status) *FeedConfigUpsert {
	u.Set(feedconfig.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateStatus() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldStatus)
	return u
}

// SetCategory sets the "category" field.
func (u *FeedConfigUpsert) SetCategory(v string) *FeedConfigUpsert {
	u.Set(feedconfig.FieldCategory, v)
	return u
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateCategory() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldCategory)
	return u
}

// SetPullInterval sets the "pull_interval" field.
func (u *FeedConfigUpsert) SetPullInterval(v time.Duration) *FeedConfigUpsert {
	u.Set(feedconfig.FieldPullInterval, v)
	return u
}

// UpdatePullInterval sets the "pull_interval" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdatePullInterval() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldPullInterval)
	return u
}

// AddPullInterval adds v to the "pull_interval" field.
func (u *FeedConfigUpsert) AddPullInterval(v time.Duration) *FeedConfigUpsert {
	u.Add(feedconfig.FieldPullInterval, v)
	return u
}

// SetLatestPullAt sets the "latest_pull_at" field.
func (u *FeedConfigUpsert) SetLatestPullAt(v time.Time) *FeedConfigUpsert {
	u.Set(feedconfig.FieldLatestPullAt, v)
	return u
}

// UpdateLatestPullAt sets the "latest_pull_at" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateLatestPullAt() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldLatestPullAt)
	return u
}

// SetNextPullBeginAt sets the "next_pull_begin_at" field.
func (u *FeedConfigUpsert) SetNextPullBeginAt(v time.Time) *FeedConfigUpsert {
	u.Set(feedconfig.FieldNextPullBeginAt, v)
	return u
}

// UpdateNextPullBeginAt sets the "next_pull_begin_at" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateNextPullBeginAt() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldNextPullBeginAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedConfigUpsert) SetUpdatedAt(v time.Time) *FeedConfigUpsert {
	u.Set(feedconfig.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateUpdatedAt() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedConfigUpsert) SetCreatedAt(v time.Time) *FeedConfigUpsert {
	u.Set(feedconfig.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedConfigUpsert) UpdateCreatedAt() *FeedConfigUpsert {
	u.SetExcluded(feedconfig.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FeedConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feedconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedConfigUpsertOne) UpdateNewValues() *FeedConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(feedconfig.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedConfig.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FeedConfigUpsertOne) Ignore() *FeedConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedConfigUpsertOne) DoNothing() *FeedConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedConfigCreate.OnConflict
// documentation for more info.
func (u *FeedConfigUpsertOne) Update(set func(*FeedConfigUpsert)) *FeedConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserFeedConfig sets the "user_feed_config" field.
func (u *FeedConfigUpsertOne) SetUserFeedConfig(v model.InternalID) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetUserFeedConfig(v)
	})
}

// UpdateUserFeedConfig sets the "user_feed_config" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateUserFeedConfig() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateUserFeedConfig()
	})
}

// SetName sets the "name" field.
func (u *FeedConfigUpsertOne) SetName(v string) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateName() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateName()
	})
}

// SetFeedURL sets the "feed_url" field.
func (u *FeedConfigUpsertOne) SetFeedURL(v string) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetFeedURL(v)
	})
}

// UpdateFeedURL sets the "feed_url" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateFeedURL() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateFeedURL()
	})
}

// SetAuthorAccount sets the "author_account" field.
func (u *FeedConfigUpsertOne) SetAuthorAccount(v model.InternalID) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetAuthorAccount(v)
	})
}

// AddAuthorAccount adds v to the "author_account" field.
func (u *FeedConfigUpsertOne) AddAuthorAccount(v model.InternalID) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.AddAuthorAccount(v)
	})
}

// UpdateAuthorAccount sets the "author_account" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateAuthorAccount() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateAuthorAccount()
	})
}

// SetSource sets the "source" field.
func (u *FeedConfigUpsertOne) SetSource(v feedconfig.Source) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateSource() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateSource()
	})
}

// SetStatus sets the "status" field.
func (u *FeedConfigUpsertOne) SetStatus(v feedconfig.Status) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateStatus() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateStatus()
	})
}

// SetCategory sets the "category" field.
func (u *FeedConfigUpsertOne) SetCategory(v string) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateCategory() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateCategory()
	})
}

// SetPullInterval sets the "pull_interval" field.
func (u *FeedConfigUpsertOne) SetPullInterval(v time.Duration) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetPullInterval(v)
	})
}

// AddPullInterval adds v to the "pull_interval" field.
func (u *FeedConfigUpsertOne) AddPullInterval(v time.Duration) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.AddPullInterval(v)
	})
}

// UpdatePullInterval sets the "pull_interval" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdatePullInterval() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdatePullInterval()
	})
}

// SetLatestPullAt sets the "latest_pull_at" field.
func (u *FeedConfigUpsertOne) SetLatestPullAt(v time.Time) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetLatestPullAt(v)
	})
}

// UpdateLatestPullAt sets the "latest_pull_at" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateLatestPullAt() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateLatestPullAt()
	})
}

// SetNextPullBeginAt sets the "next_pull_begin_at" field.
func (u *FeedConfigUpsertOne) SetNextPullBeginAt(v time.Time) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetNextPullBeginAt(v)
	})
}

// UpdateNextPullBeginAt sets the "next_pull_begin_at" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateNextPullBeginAt() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateNextPullBeginAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedConfigUpsertOne) SetUpdatedAt(v time.Time) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateUpdatedAt() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedConfigUpsertOne) SetCreatedAt(v time.Time) *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedConfigUpsertOne) UpdateCreatedAt() *FeedConfigUpsertOne {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FeedConfigUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedConfigCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedConfigUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FeedConfigUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FeedConfigUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FeedConfigCreateBulk is the builder for creating many FeedConfig entities in bulk.
type FeedConfigCreateBulk struct {
	config
	builders []*FeedConfigCreate
	conflict []sql.ConflictOption
}

// Save creates the FeedConfig entities in the database.
func (fccb *FeedConfigCreateBulk) Save(ctx context.Context) ([]*FeedConfig, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fccb.builders))
	nodes := make([]*FeedConfig, len(fccb.builders))
	mutators := make([]Mutator, len(fccb.builders))
	for i := range fccb.builders {
		func(i int, root context.Context) {
			builder := fccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = model.InternalID(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fccb *FeedConfigCreateBulk) SaveX(ctx context.Context) []*FeedConfig {
	v, err := fccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fccb *FeedConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := fccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fccb *FeedConfigCreateBulk) ExecX(ctx context.Context) {
	if err := fccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedConfig.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedConfigUpsert) {
//			SetUserFeedConfig(v+v).
//		}).
//		Exec(ctx)
func (fccb *FeedConfigCreateBulk) OnConflict(opts ...sql.ConflictOption) *FeedConfigUpsertBulk {
	fccb.conflict = opts
	return &FeedConfigUpsertBulk{
		create: fccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fccb *FeedConfigCreateBulk) OnConflictColumns(columns ...string) *FeedConfigUpsertBulk {
	fccb.conflict = append(fccb.conflict, sql.ConflictColumns(columns...))
	return &FeedConfigUpsertBulk{
		create: fccb,
	}
}

// FeedConfigUpsertBulk is the builder for "upsert"-ing
// a bulk of FeedConfig nodes.
type FeedConfigUpsertBulk struct {
	create *FeedConfigCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FeedConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feedconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedConfigUpsertBulk) UpdateNewValues() *FeedConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(feedconfig.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedConfig.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FeedConfigUpsertBulk) Ignore() *FeedConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedConfigUpsertBulk) DoNothing() *FeedConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedConfigCreateBulk.OnConflict
// documentation for more info.
func (u *FeedConfigUpsertBulk) Update(set func(*FeedConfigUpsert)) *FeedConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserFeedConfig sets the "user_feed_config" field.
func (u *FeedConfigUpsertBulk) SetUserFeedConfig(v model.InternalID) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetUserFeedConfig(v)
	})
}

// UpdateUserFeedConfig sets the "user_feed_config" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateUserFeedConfig() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateUserFeedConfig()
	})
}

// SetName sets the "name" field.
func (u *FeedConfigUpsertBulk) SetName(v string) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateName() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateName()
	})
}

// SetFeedURL sets the "feed_url" field.
func (u *FeedConfigUpsertBulk) SetFeedURL(v string) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetFeedURL(v)
	})
}

// UpdateFeedURL sets the "feed_url" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateFeedURL() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateFeedURL()
	})
}

// SetAuthorAccount sets the "author_account" field.
func (u *FeedConfigUpsertBulk) SetAuthorAccount(v model.InternalID) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetAuthorAccount(v)
	})
}

// AddAuthorAccount adds v to the "author_account" field.
func (u *FeedConfigUpsertBulk) AddAuthorAccount(v model.InternalID) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.AddAuthorAccount(v)
	})
}

// UpdateAuthorAccount sets the "author_account" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateAuthorAccount() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateAuthorAccount()
	})
}

// SetSource sets the "source" field.
func (u *FeedConfigUpsertBulk) SetSource(v feedconfig.Source) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateSource() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateSource()
	})
}

// SetStatus sets the "status" field.
func (u *FeedConfigUpsertBulk) SetStatus(v feedconfig.Status) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateStatus() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateStatus()
	})
}

// SetCategory sets the "category" field.
func (u *FeedConfigUpsertBulk) SetCategory(v string) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateCategory() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateCategory()
	})
}

// SetPullInterval sets the "pull_interval" field.
func (u *FeedConfigUpsertBulk) SetPullInterval(v time.Duration) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetPullInterval(v)
	})
}

// AddPullInterval adds v to the "pull_interval" field.
func (u *FeedConfigUpsertBulk) AddPullInterval(v time.Duration) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.AddPullInterval(v)
	})
}

// UpdatePullInterval sets the "pull_interval" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdatePullInterval() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdatePullInterval()
	})
}

// SetLatestPullAt sets the "latest_pull_at" field.
func (u *FeedConfigUpsertBulk) SetLatestPullAt(v time.Time) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetLatestPullAt(v)
	})
}

// UpdateLatestPullAt sets the "latest_pull_at" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateLatestPullAt() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateLatestPullAt()
	})
}

// SetNextPullBeginAt sets the "next_pull_begin_at" field.
func (u *FeedConfigUpsertBulk) SetNextPullBeginAt(v time.Time) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetNextPullBeginAt(v)
	})
}

// UpdateNextPullBeginAt sets the "next_pull_begin_at" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateNextPullBeginAt() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateNextPullBeginAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedConfigUpsertBulk) SetUpdatedAt(v time.Time) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateUpdatedAt() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedConfigUpsertBulk) SetCreatedAt(v time.Time) *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedConfigUpsertBulk) UpdateCreatedAt() *FeedConfigUpsertBulk {
	return u.Update(func(s *FeedConfigUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FeedConfigUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FeedConfigCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedConfigCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedConfigUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
