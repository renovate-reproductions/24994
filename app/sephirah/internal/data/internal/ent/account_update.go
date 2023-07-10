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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetPlatform sets the "platform" field.
func (au *AccountUpdate) SetPlatform(a account.Platform) *AccountUpdate {
	au.mutation.SetPlatform(a)
	return au
}

// SetPlatformAccountID sets the "platform_account_id" field.
func (au *AccountUpdate) SetPlatformAccountID(s string) *AccountUpdate {
	au.mutation.SetPlatformAccountID(s)
	return au
}

// SetName sets the "name" field.
func (au *AccountUpdate) SetName(s string) *AccountUpdate {
	au.mutation.SetName(s)
	return au
}

// SetProfileURL sets the "profile_url" field.
func (au *AccountUpdate) SetProfileURL(s string) *AccountUpdate {
	au.mutation.SetProfileURL(s)
	return au
}

// SetAvatarURL sets the "avatar_url" field.
func (au *AccountUpdate) SetAvatarURL(s string) *AccountUpdate {
	au.mutation.SetAvatarURL(s)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AccountUpdate) SetCreatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// AddPurchasedAppIDs adds the "purchased_app" edge to the App entity by IDs.
func (au *AccountUpdate) AddPurchasedAppIDs(ids ...model.InternalID) *AccountUpdate {
	au.mutation.AddPurchasedAppIDs(ids...)
	return au
}

// AddPurchasedApp adds the "purchased_app" edges to the App entity.
func (au *AccountUpdate) AddPurchasedApp(a ...*App) *AccountUpdate {
	ids := make([]model.InternalID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddPurchasedAppIDs(ids...)
}

// SetBindUserID sets the "bind_user" edge to the User entity by ID.
func (au *AccountUpdate) SetBindUserID(id model.InternalID) *AccountUpdate {
	au.mutation.SetBindUserID(id)
	return au
}

// SetNillableBindUserID sets the "bind_user" edge to the User entity by ID if the given value is not nil.
func (au *AccountUpdate) SetNillableBindUserID(id *model.InternalID) *AccountUpdate {
	if id != nil {
		au = au.SetBindUserID(*id)
	}
	return au
}

// SetBindUser sets the "bind_user" edge to the User entity.
func (au *AccountUpdate) SetBindUser(u *User) *AccountUpdate {
	return au.SetBindUserID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearPurchasedApp clears all "purchased_app" edges to the App entity.
func (au *AccountUpdate) ClearPurchasedApp() *AccountUpdate {
	au.mutation.ClearPurchasedApp()
	return au
}

// RemovePurchasedAppIDs removes the "purchased_app" edge to App entities by IDs.
func (au *AccountUpdate) RemovePurchasedAppIDs(ids ...model.InternalID) *AccountUpdate {
	au.mutation.RemovePurchasedAppIDs(ids...)
	return au
}

// RemovePurchasedApp removes "purchased_app" edges to App entities.
func (au *AccountUpdate) RemovePurchasedApp(a ...*App) *AccountUpdate {
	ids := make([]model.InternalID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemovePurchasedAppIDs(ids...)
}

// ClearBindUser clears the "bind_user" edge to the User entity.
func (au *AccountUpdate) ClearBindUser() *AccountUpdate {
	au.mutation.ClearBindUser()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Platform(); ok {
		if err := account.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf(`ent: validator failed for field "Account.platform": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Platform(); ok {
		_spec.SetField(account.FieldPlatform, field.TypeEnum, value)
	}
	if value, ok := au.mutation.PlatformAccountID(); ok {
		_spec.SetField(account.FieldPlatformAccountID, field.TypeString, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.ProfileURL(); ok {
		_spec.SetField(account.FieldProfileURL, field.TypeString, value)
	}
	if value, ok := au.mutation.AvatarURL(); ok {
		_spec.SetField(account.FieldAvatarURL, field.TypeString, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.PurchasedAppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.PurchasedAppTable,
			Columns: account.PurchasedAppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedPurchasedAppIDs(); len(nodes) > 0 && !au.mutation.PurchasedAppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.PurchasedAppTable,
			Columns: account.PurchasedAppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.PurchasedAppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.PurchasedAppTable,
			Columns: account.PurchasedAppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.BindUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.BindUserTable,
			Columns: []string{account.BindUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.BindUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.BindUserTable,
			Columns: []string{account.BindUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetPlatform sets the "platform" field.
func (auo *AccountUpdateOne) SetPlatform(a account.Platform) *AccountUpdateOne {
	auo.mutation.SetPlatform(a)
	return auo
}

// SetPlatformAccountID sets the "platform_account_id" field.
func (auo *AccountUpdateOne) SetPlatformAccountID(s string) *AccountUpdateOne {
	auo.mutation.SetPlatformAccountID(s)
	return auo
}

// SetName sets the "name" field.
func (auo *AccountUpdateOne) SetName(s string) *AccountUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetProfileURL sets the "profile_url" field.
func (auo *AccountUpdateOne) SetProfileURL(s string) *AccountUpdateOne {
	auo.mutation.SetProfileURL(s)
	return auo
}

// SetAvatarURL sets the "avatar_url" field.
func (auo *AccountUpdateOne) SetAvatarURL(s string) *AccountUpdateOne {
	auo.mutation.SetAvatarURL(s)
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AccountUpdateOne) SetCreatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// AddPurchasedAppIDs adds the "purchased_app" edge to the App entity by IDs.
func (auo *AccountUpdateOne) AddPurchasedAppIDs(ids ...model.InternalID) *AccountUpdateOne {
	auo.mutation.AddPurchasedAppIDs(ids...)
	return auo
}

// AddPurchasedApp adds the "purchased_app" edges to the App entity.
func (auo *AccountUpdateOne) AddPurchasedApp(a ...*App) *AccountUpdateOne {
	ids := make([]model.InternalID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddPurchasedAppIDs(ids...)
}

// SetBindUserID sets the "bind_user" edge to the User entity by ID.
func (auo *AccountUpdateOne) SetBindUserID(id model.InternalID) *AccountUpdateOne {
	auo.mutation.SetBindUserID(id)
	return auo
}

// SetNillableBindUserID sets the "bind_user" edge to the User entity by ID if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableBindUserID(id *model.InternalID) *AccountUpdateOne {
	if id != nil {
		auo = auo.SetBindUserID(*id)
	}
	return auo
}

// SetBindUser sets the "bind_user" edge to the User entity.
func (auo *AccountUpdateOne) SetBindUser(u *User) *AccountUpdateOne {
	return auo.SetBindUserID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearPurchasedApp clears all "purchased_app" edges to the App entity.
func (auo *AccountUpdateOne) ClearPurchasedApp() *AccountUpdateOne {
	auo.mutation.ClearPurchasedApp()
	return auo
}

// RemovePurchasedAppIDs removes the "purchased_app" edge to App entities by IDs.
func (auo *AccountUpdateOne) RemovePurchasedAppIDs(ids ...model.InternalID) *AccountUpdateOne {
	auo.mutation.RemovePurchasedAppIDs(ids...)
	return auo
}

// RemovePurchasedApp removes "purchased_app" edges to App entities.
func (auo *AccountUpdateOne) RemovePurchasedApp(a ...*App) *AccountUpdateOne {
	ids := make([]model.InternalID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemovePurchasedAppIDs(ids...)
}

// ClearBindUser clears the "bind_user" edge to the User entity.
func (auo *AccountUpdateOne) ClearBindUser() *AccountUpdateOne {
	auo.mutation.ClearBindUser()
	return auo
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Platform(); ok {
		if err := account.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf(`ent: validator failed for field "Account.platform": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Platform(); ok {
		_spec.SetField(account.FieldPlatform, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.PlatformAccountID(); ok {
		_spec.SetField(account.FieldPlatformAccountID, field.TypeString, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.ProfileURL(); ok {
		_spec.SetField(account.FieldProfileURL, field.TypeString, value)
	}
	if value, ok := auo.mutation.AvatarURL(); ok {
		_spec.SetField(account.FieldAvatarURL, field.TypeString, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.PurchasedAppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.PurchasedAppTable,
			Columns: account.PurchasedAppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedPurchasedAppIDs(); len(nodes) > 0 && !auo.mutation.PurchasedAppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.PurchasedAppTable,
			Columns: account.PurchasedAppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.PurchasedAppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.PurchasedAppTable,
			Columns: account.PurchasedAppPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.BindUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.BindUserTable,
			Columns: []string{account.BindUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.BindUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.BindUserTable,
			Columns: []string{account.BindUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
