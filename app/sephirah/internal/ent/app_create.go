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
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
)

// AppCreate is the builder for creating a App entity.
type AppCreate struct {
	config
	mutation *AppMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetInternalID sets the "internal_id" field.
func (ac *AppCreate) SetInternalID(i int64) *AppCreate {
	ac.mutation.SetInternalID(i)
	return ac
}

// SetSource sets the "source" field.
func (ac *AppCreate) SetSource(a app.Source) *AppCreate {
	ac.mutation.SetSource(a)
	return ac
}

// SetSourceAppID sets the "source_app_id" field.
func (ac *AppCreate) SetSourceAppID(s string) *AppCreate {
	ac.mutation.SetSourceAppID(s)
	return ac
}

// SetSourceURL sets the "source_url" field.
func (ac *AppCreate) SetSourceURL(s string) *AppCreate {
	ac.mutation.SetSourceURL(s)
	return ac
}

// SetName sets the "name" field.
func (ac *AppCreate) SetName(s string) *AppCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetType sets the "type" field.
func (ac *AppCreate) SetType(a app.Type) *AppCreate {
	ac.mutation.SetType(a)
	return ac
}

// SetShortDescription sets the "short_description" field.
func (ac *AppCreate) SetShortDescription(s string) *AppCreate {
	ac.mutation.SetShortDescription(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AppCreate) SetDescription(s string) *AppCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetImageURL sets the "image_url" field.
func (ac *AppCreate) SetImageURL(s string) *AppCreate {
	ac.mutation.SetImageURL(s)
	return ac
}

// SetReleaseDate sets the "release_date" field.
func (ac *AppCreate) SetReleaseDate(s string) *AppCreate {
	ac.mutation.SetReleaseDate(s)
	return ac
}

// SetDeveloper sets the "developer" field.
func (ac *AppCreate) SetDeveloper(s string) *AppCreate {
	ac.mutation.SetDeveloper(s)
	return ac
}

// SetPublisher sets the "publisher" field.
func (ac *AppCreate) SetPublisher(s string) *AppCreate {
	ac.mutation.SetPublisher(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AppCreate) SetCreatedAt(t time.Time) *AppCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableCreatedAt(t *time.Time) *AppCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// Mutation returns the AppMutation object of the builder.
func (ac *AppCreate) Mutation() *AppMutation {
	return ac.mutation
}

// Save creates the App in the database.
func (ac *AppCreate) Save(ctx context.Context) (*App, error) {
	var (
		err  error
		node *App
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*App)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AppCreate) SaveX(ctx context.Context) *App {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AppCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AppCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AppCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := app.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AppCreate) check() error {
	if _, ok := ac.mutation.InternalID(); !ok {
		return &ValidationError{Name: "internal_id", err: errors.New(`ent: missing required field "App.internal_id"`)}
	}
	if _, ok := ac.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required field "App.source"`)}
	}
	if v, ok := ac.mutation.Source(); ok {
		if err := app.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "App.source": %w`, err)}
		}
	}
	if _, ok := ac.mutation.SourceAppID(); !ok {
		return &ValidationError{Name: "source_app_id", err: errors.New(`ent: missing required field "App.source_app_id"`)}
	}
	if _, ok := ac.mutation.SourceURL(); !ok {
		return &ValidationError{Name: "source_url", err: errors.New(`ent: missing required field "App.source_url"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "App.name"`)}
	}
	if _, ok := ac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "App.type"`)}
	}
	if v, ok := ac.mutation.GetType(); ok {
		if err := app.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "App.type": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ShortDescription(); !ok {
		return &ValidationError{Name: "short_description", err: errors.New(`ent: missing required field "App.short_description"`)}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "App.description"`)}
	}
	if _, ok := ac.mutation.ImageURL(); !ok {
		return &ValidationError{Name: "image_url", err: errors.New(`ent: missing required field "App.image_url"`)}
	}
	if _, ok := ac.mutation.ReleaseDate(); !ok {
		return &ValidationError{Name: "release_date", err: errors.New(`ent: missing required field "App.release_date"`)}
	}
	if _, ok := ac.mutation.Developer(); !ok {
		return &ValidationError{Name: "developer", err: errors.New(`ent: missing required field "App.developer"`)}
	}
	if _, ok := ac.mutation.Publisher(); !ok {
		return &ValidationError{Name: "publisher", err: errors.New(`ent: missing required field "App.publisher"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "App.created_at"`)}
	}
	return nil
}

func (ac *AppCreate) sqlSave(ctx context.Context) (*App, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AppCreate) createSpec() (*App, *sqlgraph.CreateSpec) {
	var (
		_node = &App{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: app.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: app.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.InternalID(); ok {
		_spec.SetField(app.FieldInternalID, field.TypeInt64, value)
		_node.InternalID = value
	}
	if value, ok := ac.mutation.Source(); ok {
		_spec.SetField(app.FieldSource, field.TypeEnum, value)
		_node.Source = value
	}
	if value, ok := ac.mutation.SourceAppID(); ok {
		_spec.SetField(app.FieldSourceAppID, field.TypeString, value)
		_node.SourceAppID = value
	}
	if value, ok := ac.mutation.SourceURL(); ok {
		_spec.SetField(app.FieldSourceURL, field.TypeString, value)
		_node.SourceURL = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(app.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.GetType(); ok {
		_spec.SetField(app.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ac.mutation.ShortDescription(); ok {
		_spec.SetField(app.FieldShortDescription, field.TypeString, value)
		_node.ShortDescription = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(app.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.ImageURL(); ok {
		_spec.SetField(app.FieldImageURL, field.TypeString, value)
		_node.ImageURL = value
	}
	if value, ok := ac.mutation.ReleaseDate(); ok {
		_spec.SetField(app.FieldReleaseDate, field.TypeString, value)
		_node.ReleaseDate = value
	}
	if value, ok := ac.mutation.Developer(); ok {
		_spec.SetField(app.FieldDeveloper, field.TypeString, value)
		_node.Developer = value
	}
	if value, ok := ac.mutation.Publisher(); ok {
		_spec.SetField(app.FieldPublisher, field.TypeString, value)
		_node.Publisher = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(app.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.App.Create().
//		SetInternalID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUpsert) {
//			SetInternalID(v+v).
//		}).
//		Exec(ctx)
//
func (ac *AppCreate) OnConflict(opts ...sql.ConflictOption) *AppUpsertOne {
	ac.conflict = opts
	return &AppUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *AppCreate) OnConflictColumns(columns ...string) *AppUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AppUpsertOne{
		create: ac,
	}
}

type (
	// AppUpsertOne is the builder for "upsert"-ing
	//  one App node.
	AppUpsertOne struct {
		create *AppCreate
	}

	// AppUpsert is the "OnConflict" setter.
	AppUpsert struct {
		*sql.UpdateSet
	}
)

// SetInternalID sets the "internal_id" field.
func (u *AppUpsert) SetInternalID(v int64) *AppUpsert {
	u.Set(app.FieldInternalID, v)
	return u
}

// UpdateInternalID sets the "internal_id" field to the value that was provided on create.
func (u *AppUpsert) UpdateInternalID() *AppUpsert {
	u.SetExcluded(app.FieldInternalID)
	return u
}

// AddInternalID adds v to the "internal_id" field.
func (u *AppUpsert) AddInternalID(v int64) *AppUpsert {
	u.Add(app.FieldInternalID, v)
	return u
}

// SetSource sets the "source" field.
func (u *AppUpsert) SetSource(v app.Source) *AppUpsert {
	u.Set(app.FieldSource, v)
	return u
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppUpsert) UpdateSource() *AppUpsert {
	u.SetExcluded(app.FieldSource)
	return u
}

// SetSourceAppID sets the "source_app_id" field.
func (u *AppUpsert) SetSourceAppID(v string) *AppUpsert {
	u.Set(app.FieldSourceAppID, v)
	return u
}

// UpdateSourceAppID sets the "source_app_id" field to the value that was provided on create.
func (u *AppUpsert) UpdateSourceAppID() *AppUpsert {
	u.SetExcluded(app.FieldSourceAppID)
	return u
}

// SetSourceURL sets the "source_url" field.
func (u *AppUpsert) SetSourceURL(v string) *AppUpsert {
	u.Set(app.FieldSourceURL, v)
	return u
}

// UpdateSourceURL sets the "source_url" field to the value that was provided on create.
func (u *AppUpsert) UpdateSourceURL() *AppUpsert {
	u.SetExcluded(app.FieldSourceURL)
	return u
}

// SetName sets the "name" field.
func (u *AppUpsert) SetName(v string) *AppUpsert {
	u.Set(app.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsert) UpdateName() *AppUpsert {
	u.SetExcluded(app.FieldName)
	return u
}

// SetType sets the "type" field.
func (u *AppUpsert) SetType(v app.Type) *AppUpsert {
	u.Set(app.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *AppUpsert) UpdateType() *AppUpsert {
	u.SetExcluded(app.FieldType)
	return u
}

// SetShortDescription sets the "short_description" field.
func (u *AppUpsert) SetShortDescription(v string) *AppUpsert {
	u.Set(app.FieldShortDescription, v)
	return u
}

// UpdateShortDescription sets the "short_description" field to the value that was provided on create.
func (u *AppUpsert) UpdateShortDescription() *AppUpsert {
	u.SetExcluded(app.FieldShortDescription)
	return u
}

// SetDescription sets the "description" field.
func (u *AppUpsert) SetDescription(v string) *AppUpsert {
	u.Set(app.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsert) UpdateDescription() *AppUpsert {
	u.SetExcluded(app.FieldDescription)
	return u
}

// SetImageURL sets the "image_url" field.
func (u *AppUpsert) SetImageURL(v string) *AppUpsert {
	u.Set(app.FieldImageURL, v)
	return u
}

// UpdateImageURL sets the "image_url" field to the value that was provided on create.
func (u *AppUpsert) UpdateImageURL() *AppUpsert {
	u.SetExcluded(app.FieldImageURL)
	return u
}

// SetReleaseDate sets the "release_date" field.
func (u *AppUpsert) SetReleaseDate(v string) *AppUpsert {
	u.Set(app.FieldReleaseDate, v)
	return u
}

// UpdateReleaseDate sets the "release_date" field to the value that was provided on create.
func (u *AppUpsert) UpdateReleaseDate() *AppUpsert {
	u.SetExcluded(app.FieldReleaseDate)
	return u
}

// SetDeveloper sets the "developer" field.
func (u *AppUpsert) SetDeveloper(v string) *AppUpsert {
	u.Set(app.FieldDeveloper, v)
	return u
}

// UpdateDeveloper sets the "developer" field to the value that was provided on create.
func (u *AppUpsert) UpdateDeveloper() *AppUpsert {
	u.SetExcluded(app.FieldDeveloper)
	return u
}

// SetPublisher sets the "publisher" field.
func (u *AppUpsert) SetPublisher(v string) *AppUpsert {
	u.Set(app.FieldPublisher, v)
	return u
}

// UpdatePublisher sets the "publisher" field to the value that was provided on create.
func (u *AppUpsert) UpdatePublisher() *AppUpsert {
	u.SetExcluded(app.FieldPublisher)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsert) SetCreatedAt(v time.Time) *AppUpsert {
	u.Set(app.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsert) UpdateCreatedAt() *AppUpsert {
	u.SetExcluded(app.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *AppUpsertOne) UpdateNewValues() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.App.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppUpsertOne) Ignore() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUpsertOne) DoNothing() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCreate.OnConflict
// documentation for more info.
func (u *AppUpsertOne) Update(set func(*AppUpsert)) *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUpsert{UpdateSet: update})
	}))
	return u
}

// SetInternalID sets the "internal_id" field.
func (u *AppUpsertOne) SetInternalID(v int64) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetInternalID(v)
	})
}

// AddInternalID adds v to the "internal_id" field.
func (u *AppUpsertOne) AddInternalID(v int64) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.AddInternalID(v)
	})
}

// UpdateInternalID sets the "internal_id" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateInternalID() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateInternalID()
	})
}

// SetSource sets the "source" field.
func (u *AppUpsertOne) SetSource(v app.Source) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateSource() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateSource()
	})
}

// SetSourceAppID sets the "source_app_id" field.
func (u *AppUpsertOne) SetSourceAppID(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetSourceAppID(v)
	})
}

// UpdateSourceAppID sets the "source_app_id" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateSourceAppID() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateSourceAppID()
	})
}

// SetSourceURL sets the "source_url" field.
func (u *AppUpsertOne) SetSourceURL(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetSourceURL(v)
	})
}

// UpdateSourceURL sets the "source_url" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateSourceURL() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateSourceURL()
	})
}

// SetName sets the "name" field.
func (u *AppUpsertOne) SetName(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateName() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *AppUpsertOne) SetType(v app.Type) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateType() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateType()
	})
}

// SetShortDescription sets the "short_description" field.
func (u *AppUpsertOne) SetShortDescription(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetShortDescription(v)
	})
}

// UpdateShortDescription sets the "short_description" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateShortDescription() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateShortDescription()
	})
}

// SetDescription sets the "description" field.
func (u *AppUpsertOne) SetDescription(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateDescription() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDescription()
	})
}

// SetImageURL sets the "image_url" field.
func (u *AppUpsertOne) SetImageURL(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetImageURL(v)
	})
}

// UpdateImageURL sets the "image_url" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateImageURL() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateImageURL()
	})
}

// SetReleaseDate sets the "release_date" field.
func (u *AppUpsertOne) SetReleaseDate(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetReleaseDate(v)
	})
}

// UpdateReleaseDate sets the "release_date" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateReleaseDate() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateReleaseDate()
	})
}

// SetDeveloper sets the "developer" field.
func (u *AppUpsertOne) SetDeveloper(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetDeveloper(v)
	})
}

// UpdateDeveloper sets the "developer" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateDeveloper() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDeveloper()
	})
}

// SetPublisher sets the "publisher" field.
func (u *AppUpsertOne) SetPublisher(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetPublisher(v)
	})
}

// UpdatePublisher sets the "publisher" field to the value that was provided on create.
func (u *AppUpsertOne) UpdatePublisher() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdatePublisher()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsertOne) SetCreatedAt(v time.Time) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateCreatedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppCreateBulk is the builder for creating many App entities in bulk.
type AppCreateBulk struct {
	config
	builders []*AppCreate
	conflict []sql.ConflictOption
}

// Save creates the App entities in the database.
func (acb *AppCreateBulk) Save(ctx context.Context) ([]*App, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*App, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AppCreateBulk) SaveX(ctx context.Context) []*App {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AppCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AppCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.App.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUpsert) {
//			SetInternalID(v+v).
//		}).
//		Exec(ctx)
//
func (acb *AppCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUpsertBulk {
	acb.conflict = opts
	return &AppUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *AppCreateBulk) OnConflictColumns(columns ...string) *AppUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AppUpsertBulk{
		create: acb,
	}
}

// AppUpsertBulk is the builder for "upsert"-ing
// a bulk of App nodes.
type AppUpsertBulk struct {
	create *AppCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *AppUpsertBulk) UpdateNewValues() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppUpsertBulk) Ignore() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUpsertBulk) DoNothing() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCreateBulk.OnConflict
// documentation for more info.
func (u *AppUpsertBulk) Update(set func(*AppUpsert)) *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUpsert{UpdateSet: update})
	}))
	return u
}

// SetInternalID sets the "internal_id" field.
func (u *AppUpsertBulk) SetInternalID(v int64) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetInternalID(v)
	})
}

// AddInternalID adds v to the "internal_id" field.
func (u *AppUpsertBulk) AddInternalID(v int64) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.AddInternalID(v)
	})
}

// UpdateInternalID sets the "internal_id" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateInternalID() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateInternalID()
	})
}

// SetSource sets the "source" field.
func (u *AppUpsertBulk) SetSource(v app.Source) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateSource() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateSource()
	})
}

// SetSourceAppID sets the "source_app_id" field.
func (u *AppUpsertBulk) SetSourceAppID(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetSourceAppID(v)
	})
}

// UpdateSourceAppID sets the "source_app_id" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateSourceAppID() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateSourceAppID()
	})
}

// SetSourceURL sets the "source_url" field.
func (u *AppUpsertBulk) SetSourceURL(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetSourceURL(v)
	})
}

// UpdateSourceURL sets the "source_url" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateSourceURL() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateSourceURL()
	})
}

// SetName sets the "name" field.
func (u *AppUpsertBulk) SetName(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateName() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *AppUpsertBulk) SetType(v app.Type) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateType() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateType()
	})
}

// SetShortDescription sets the "short_description" field.
func (u *AppUpsertBulk) SetShortDescription(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetShortDescription(v)
	})
}

// UpdateShortDescription sets the "short_description" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateShortDescription() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateShortDescription()
	})
}

// SetDescription sets the "description" field.
func (u *AppUpsertBulk) SetDescription(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateDescription() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDescription()
	})
}

// SetImageURL sets the "image_url" field.
func (u *AppUpsertBulk) SetImageURL(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetImageURL(v)
	})
}

// UpdateImageURL sets the "image_url" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateImageURL() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateImageURL()
	})
}

// SetReleaseDate sets the "release_date" field.
func (u *AppUpsertBulk) SetReleaseDate(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetReleaseDate(v)
	})
}

// UpdateReleaseDate sets the "release_date" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateReleaseDate() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateReleaseDate()
	})
}

// SetDeveloper sets the "developer" field.
func (u *AppUpsertBulk) SetDeveloper(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetDeveloper(v)
	})
}

// UpdateDeveloper sets the "developer" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateDeveloper() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDeveloper()
	})
}

// SetPublisher sets the "publisher" field.
func (u *AppUpsertBulk) SetPublisher(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetPublisher(v)
	})
}

// UpdatePublisher sets the "publisher" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdatePublisher() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdatePublisher()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsertBulk) SetCreatedAt(v time.Time) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateCreatedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
