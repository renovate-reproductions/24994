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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// NotifyTargetUpdate is the builder for updating NotifyTarget entities.
type NotifyTargetUpdate struct {
	config
	hooks    []Hook
	mutation *NotifyTargetMutation
}

// Where appends a list predicates to the NotifyTargetUpdate builder.
func (ntu *NotifyTargetUpdate) Where(ps ...predicate.NotifyTarget) *NotifyTargetUpdate {
	ntu.mutation.Where(ps...)
	return ntu
}

// SetToken sets the "token" field.
func (ntu *NotifyTargetUpdate) SetToken(s string) *NotifyTargetUpdate {
	ntu.mutation.SetToken(s)
	return ntu
}

// SetName sets the "name" field.
func (ntu *NotifyTargetUpdate) SetName(s string) *NotifyTargetUpdate {
	ntu.mutation.SetName(s)
	return ntu
}

// SetDescription sets the "description" field.
func (ntu *NotifyTargetUpdate) SetDescription(s string) *NotifyTargetUpdate {
	ntu.mutation.SetDescription(s)
	return ntu
}

// SetType sets the "type" field.
func (ntu *NotifyTargetUpdate) SetType(n notifytarget.Type) *NotifyTargetUpdate {
	ntu.mutation.SetType(n)
	return ntu
}

// SetStatus sets the "status" field.
func (ntu *NotifyTargetUpdate) SetStatus(n notifytarget.Status) *NotifyTargetUpdate {
	ntu.mutation.SetStatus(n)
	return ntu
}

// SetUpdatedAt sets the "updated_at" field.
func (ntu *NotifyTargetUpdate) SetUpdatedAt(t time.Time) *NotifyTargetUpdate {
	ntu.mutation.SetUpdatedAt(t)
	return ntu
}

// SetCreatedAt sets the "created_at" field.
func (ntu *NotifyTargetUpdate) SetCreatedAt(t time.Time) *NotifyTargetUpdate {
	ntu.mutation.SetCreatedAt(t)
	return ntu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ntu *NotifyTargetUpdate) SetNillableCreatedAt(t *time.Time) *NotifyTargetUpdate {
	if t != nil {
		ntu.SetCreatedAt(*t)
	}
	return ntu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ntu *NotifyTargetUpdate) SetOwnerID(id model.InternalID) *NotifyTargetUpdate {
	ntu.mutation.SetOwnerID(id)
	return ntu
}

// SetOwner sets the "owner" edge to the User entity.
func (ntu *NotifyTargetUpdate) SetOwner(u *User) *NotifyTargetUpdate {
	return ntu.SetOwnerID(u.ID)
}

// AddNotifyFlowIDs adds the "notify_flow" edge to the NotifyFlow entity by IDs.
func (ntu *NotifyTargetUpdate) AddNotifyFlowIDs(ids ...model.InternalID) *NotifyTargetUpdate {
	ntu.mutation.AddNotifyFlowIDs(ids...)
	return ntu
}

// AddNotifyFlow adds the "notify_flow" edges to the NotifyFlow entity.
func (ntu *NotifyTargetUpdate) AddNotifyFlow(n ...*NotifyFlow) *NotifyTargetUpdate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntu.AddNotifyFlowIDs(ids...)
}

// AddNotifyFlowTargetIDs adds the "notify_flow_target" edge to the NotifyFlowTarget entity by IDs.
func (ntu *NotifyTargetUpdate) AddNotifyFlowTargetIDs(ids ...model.InternalID) *NotifyTargetUpdate {
	ntu.mutation.AddNotifyFlowTargetIDs(ids...)
	return ntu
}

// AddNotifyFlowTarget adds the "notify_flow_target" edges to the NotifyFlowTarget entity.
func (ntu *NotifyTargetUpdate) AddNotifyFlowTarget(n ...*NotifyFlowTarget) *NotifyTargetUpdate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntu.AddNotifyFlowTargetIDs(ids...)
}

// Mutation returns the NotifyTargetMutation object of the builder.
func (ntu *NotifyTargetUpdate) Mutation() *NotifyTargetMutation {
	return ntu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (ntu *NotifyTargetUpdate) ClearOwner() *NotifyTargetUpdate {
	ntu.mutation.ClearOwner()
	return ntu
}

// ClearNotifyFlow clears all "notify_flow" edges to the NotifyFlow entity.
func (ntu *NotifyTargetUpdate) ClearNotifyFlow() *NotifyTargetUpdate {
	ntu.mutation.ClearNotifyFlow()
	return ntu
}

// RemoveNotifyFlowIDs removes the "notify_flow" edge to NotifyFlow entities by IDs.
func (ntu *NotifyTargetUpdate) RemoveNotifyFlowIDs(ids ...model.InternalID) *NotifyTargetUpdate {
	ntu.mutation.RemoveNotifyFlowIDs(ids...)
	return ntu
}

// RemoveNotifyFlow removes "notify_flow" edges to NotifyFlow entities.
func (ntu *NotifyTargetUpdate) RemoveNotifyFlow(n ...*NotifyFlow) *NotifyTargetUpdate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntu.RemoveNotifyFlowIDs(ids...)
}

// ClearNotifyFlowTarget clears all "notify_flow_target" edges to the NotifyFlowTarget entity.
func (ntu *NotifyTargetUpdate) ClearNotifyFlowTarget() *NotifyTargetUpdate {
	ntu.mutation.ClearNotifyFlowTarget()
	return ntu
}

// RemoveNotifyFlowTargetIDs removes the "notify_flow_target" edge to NotifyFlowTarget entities by IDs.
func (ntu *NotifyTargetUpdate) RemoveNotifyFlowTargetIDs(ids ...model.InternalID) *NotifyTargetUpdate {
	ntu.mutation.RemoveNotifyFlowTargetIDs(ids...)
	return ntu
}

// RemoveNotifyFlowTarget removes "notify_flow_target" edges to NotifyFlowTarget entities.
func (ntu *NotifyTargetUpdate) RemoveNotifyFlowTarget(n ...*NotifyFlowTarget) *NotifyTargetUpdate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntu.RemoveNotifyFlowTargetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ntu *NotifyTargetUpdate) Save(ctx context.Context) (int, error) {
	ntu.defaults()
	return withHooks[int, NotifyTargetMutation](ctx, ntu.sqlSave, ntu.mutation, ntu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ntu *NotifyTargetUpdate) SaveX(ctx context.Context) int {
	affected, err := ntu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ntu *NotifyTargetUpdate) Exec(ctx context.Context) error {
	_, err := ntu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntu *NotifyTargetUpdate) ExecX(ctx context.Context) {
	if err := ntu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ntu *NotifyTargetUpdate) defaults() {
	if _, ok := ntu.mutation.UpdatedAt(); !ok {
		v := notifytarget.UpdateDefaultUpdatedAt()
		ntu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntu *NotifyTargetUpdate) check() error {
	if v, ok := ntu.mutation.GetType(); ok {
		if err := notifytarget.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "NotifyTarget.type": %w`, err)}
		}
	}
	if v, ok := ntu.mutation.Status(); ok {
		if err := notifytarget.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "NotifyTarget.status": %w`, err)}
		}
	}
	if _, ok := ntu.mutation.OwnerID(); ntu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NotifyTarget.owner"`)
	}
	return nil
}

func (ntu *NotifyTargetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ntu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notifytarget.Table, notifytarget.Columns, sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64))
	if ps := ntu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ntu.mutation.Token(); ok {
		_spec.SetField(notifytarget.FieldToken, field.TypeString, value)
	}
	if value, ok := ntu.mutation.Name(); ok {
		_spec.SetField(notifytarget.FieldName, field.TypeString, value)
	}
	if value, ok := ntu.mutation.Description(); ok {
		_spec.SetField(notifytarget.FieldDescription, field.TypeString, value)
	}
	if value, ok := ntu.mutation.GetType(); ok {
		_spec.SetField(notifytarget.FieldType, field.TypeEnum, value)
	}
	if value, ok := ntu.mutation.Status(); ok {
		_spec.SetField(notifytarget.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ntu.mutation.UpdatedAt(); ok {
		_spec.SetField(notifytarget.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ntu.mutation.CreatedAt(); ok {
		_spec.SetField(notifytarget.FieldCreatedAt, field.TypeTime, value)
	}
	if ntu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifytarget.OwnerTable,
			Columns: []string{notifytarget.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifytarget.OwnerTable,
			Columns: []string{notifytarget.OwnerColumn},
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
	if ntu.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTable,
			Columns: notifytarget.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		createE := &NotifyFlowTargetCreate{config: ntu.config, mutation: newNotifyFlowTargetMutation(ntu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntu.mutation.RemovedNotifyFlowIDs(); len(nodes) > 0 && !ntu.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTable,
			Columns: notifytarget.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowTargetCreate{config: ntu.config, mutation: newNotifyFlowTargetMutation(ntu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntu.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTable,
			Columns: notifytarget.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowTargetCreate{config: ntu.config, mutation: newNotifyFlowTargetMutation(ntu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ntu.mutation.NotifyFlowTargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTargetTable,
			Columns: []string{notifytarget.NotifyFlowTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntu.mutation.RemovedNotifyFlowTargetIDs(); len(nodes) > 0 && !ntu.mutation.NotifyFlowTargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTargetTable,
			Columns: []string{notifytarget.NotifyFlowTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntu.mutation.NotifyFlowTargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTargetTable,
			Columns: []string{notifytarget.NotifyFlowTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ntu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notifytarget.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ntu.mutation.done = true
	return n, nil
}

// NotifyTargetUpdateOne is the builder for updating a single NotifyTarget entity.
type NotifyTargetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotifyTargetMutation
}

// SetToken sets the "token" field.
func (ntuo *NotifyTargetUpdateOne) SetToken(s string) *NotifyTargetUpdateOne {
	ntuo.mutation.SetToken(s)
	return ntuo
}

// SetName sets the "name" field.
func (ntuo *NotifyTargetUpdateOne) SetName(s string) *NotifyTargetUpdateOne {
	ntuo.mutation.SetName(s)
	return ntuo
}

// SetDescription sets the "description" field.
func (ntuo *NotifyTargetUpdateOne) SetDescription(s string) *NotifyTargetUpdateOne {
	ntuo.mutation.SetDescription(s)
	return ntuo
}

// SetType sets the "type" field.
func (ntuo *NotifyTargetUpdateOne) SetType(n notifytarget.Type) *NotifyTargetUpdateOne {
	ntuo.mutation.SetType(n)
	return ntuo
}

// SetStatus sets the "status" field.
func (ntuo *NotifyTargetUpdateOne) SetStatus(n notifytarget.Status) *NotifyTargetUpdateOne {
	ntuo.mutation.SetStatus(n)
	return ntuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ntuo *NotifyTargetUpdateOne) SetUpdatedAt(t time.Time) *NotifyTargetUpdateOne {
	ntuo.mutation.SetUpdatedAt(t)
	return ntuo
}

// SetCreatedAt sets the "created_at" field.
func (ntuo *NotifyTargetUpdateOne) SetCreatedAt(t time.Time) *NotifyTargetUpdateOne {
	ntuo.mutation.SetCreatedAt(t)
	return ntuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ntuo *NotifyTargetUpdateOne) SetNillableCreatedAt(t *time.Time) *NotifyTargetUpdateOne {
	if t != nil {
		ntuo.SetCreatedAt(*t)
	}
	return ntuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ntuo *NotifyTargetUpdateOne) SetOwnerID(id model.InternalID) *NotifyTargetUpdateOne {
	ntuo.mutation.SetOwnerID(id)
	return ntuo
}

// SetOwner sets the "owner" edge to the User entity.
func (ntuo *NotifyTargetUpdateOne) SetOwner(u *User) *NotifyTargetUpdateOne {
	return ntuo.SetOwnerID(u.ID)
}

// AddNotifyFlowIDs adds the "notify_flow" edge to the NotifyFlow entity by IDs.
func (ntuo *NotifyTargetUpdateOne) AddNotifyFlowIDs(ids ...model.InternalID) *NotifyTargetUpdateOne {
	ntuo.mutation.AddNotifyFlowIDs(ids...)
	return ntuo
}

// AddNotifyFlow adds the "notify_flow" edges to the NotifyFlow entity.
func (ntuo *NotifyTargetUpdateOne) AddNotifyFlow(n ...*NotifyFlow) *NotifyTargetUpdateOne {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntuo.AddNotifyFlowIDs(ids...)
}

// AddNotifyFlowTargetIDs adds the "notify_flow_target" edge to the NotifyFlowTarget entity by IDs.
func (ntuo *NotifyTargetUpdateOne) AddNotifyFlowTargetIDs(ids ...model.InternalID) *NotifyTargetUpdateOne {
	ntuo.mutation.AddNotifyFlowTargetIDs(ids...)
	return ntuo
}

// AddNotifyFlowTarget adds the "notify_flow_target" edges to the NotifyFlowTarget entity.
func (ntuo *NotifyTargetUpdateOne) AddNotifyFlowTarget(n ...*NotifyFlowTarget) *NotifyTargetUpdateOne {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntuo.AddNotifyFlowTargetIDs(ids...)
}

// Mutation returns the NotifyTargetMutation object of the builder.
func (ntuo *NotifyTargetUpdateOne) Mutation() *NotifyTargetMutation {
	return ntuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (ntuo *NotifyTargetUpdateOne) ClearOwner() *NotifyTargetUpdateOne {
	ntuo.mutation.ClearOwner()
	return ntuo
}

// ClearNotifyFlow clears all "notify_flow" edges to the NotifyFlow entity.
func (ntuo *NotifyTargetUpdateOne) ClearNotifyFlow() *NotifyTargetUpdateOne {
	ntuo.mutation.ClearNotifyFlow()
	return ntuo
}

// RemoveNotifyFlowIDs removes the "notify_flow" edge to NotifyFlow entities by IDs.
func (ntuo *NotifyTargetUpdateOne) RemoveNotifyFlowIDs(ids ...model.InternalID) *NotifyTargetUpdateOne {
	ntuo.mutation.RemoveNotifyFlowIDs(ids...)
	return ntuo
}

// RemoveNotifyFlow removes "notify_flow" edges to NotifyFlow entities.
func (ntuo *NotifyTargetUpdateOne) RemoveNotifyFlow(n ...*NotifyFlow) *NotifyTargetUpdateOne {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntuo.RemoveNotifyFlowIDs(ids...)
}

// ClearNotifyFlowTarget clears all "notify_flow_target" edges to the NotifyFlowTarget entity.
func (ntuo *NotifyTargetUpdateOne) ClearNotifyFlowTarget() *NotifyTargetUpdateOne {
	ntuo.mutation.ClearNotifyFlowTarget()
	return ntuo
}

// RemoveNotifyFlowTargetIDs removes the "notify_flow_target" edge to NotifyFlowTarget entities by IDs.
func (ntuo *NotifyTargetUpdateOne) RemoveNotifyFlowTargetIDs(ids ...model.InternalID) *NotifyTargetUpdateOne {
	ntuo.mutation.RemoveNotifyFlowTargetIDs(ids...)
	return ntuo
}

// RemoveNotifyFlowTarget removes "notify_flow_target" edges to NotifyFlowTarget entities.
func (ntuo *NotifyTargetUpdateOne) RemoveNotifyFlowTarget(n ...*NotifyFlowTarget) *NotifyTargetUpdateOne {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntuo.RemoveNotifyFlowTargetIDs(ids...)
}

// Where appends a list predicates to the NotifyTargetUpdate builder.
func (ntuo *NotifyTargetUpdateOne) Where(ps ...predicate.NotifyTarget) *NotifyTargetUpdateOne {
	ntuo.mutation.Where(ps...)
	return ntuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ntuo *NotifyTargetUpdateOne) Select(field string, fields ...string) *NotifyTargetUpdateOne {
	ntuo.fields = append([]string{field}, fields...)
	return ntuo
}

// Save executes the query and returns the updated NotifyTarget entity.
func (ntuo *NotifyTargetUpdateOne) Save(ctx context.Context) (*NotifyTarget, error) {
	ntuo.defaults()
	return withHooks[*NotifyTarget, NotifyTargetMutation](ctx, ntuo.sqlSave, ntuo.mutation, ntuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ntuo *NotifyTargetUpdateOne) SaveX(ctx context.Context) *NotifyTarget {
	node, err := ntuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ntuo *NotifyTargetUpdateOne) Exec(ctx context.Context) error {
	_, err := ntuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntuo *NotifyTargetUpdateOne) ExecX(ctx context.Context) {
	if err := ntuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ntuo *NotifyTargetUpdateOne) defaults() {
	if _, ok := ntuo.mutation.UpdatedAt(); !ok {
		v := notifytarget.UpdateDefaultUpdatedAt()
		ntuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntuo *NotifyTargetUpdateOne) check() error {
	if v, ok := ntuo.mutation.GetType(); ok {
		if err := notifytarget.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "NotifyTarget.type": %w`, err)}
		}
	}
	if v, ok := ntuo.mutation.Status(); ok {
		if err := notifytarget.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "NotifyTarget.status": %w`, err)}
		}
	}
	if _, ok := ntuo.mutation.OwnerID(); ntuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NotifyTarget.owner"`)
	}
	return nil
}

func (ntuo *NotifyTargetUpdateOne) sqlSave(ctx context.Context) (_node *NotifyTarget, err error) {
	if err := ntuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notifytarget.Table, notifytarget.Columns, sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64))
	id, ok := ntuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotifyTarget.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ntuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notifytarget.FieldID)
		for _, f := range fields {
			if !notifytarget.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notifytarget.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ntuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ntuo.mutation.Token(); ok {
		_spec.SetField(notifytarget.FieldToken, field.TypeString, value)
	}
	if value, ok := ntuo.mutation.Name(); ok {
		_spec.SetField(notifytarget.FieldName, field.TypeString, value)
	}
	if value, ok := ntuo.mutation.Description(); ok {
		_spec.SetField(notifytarget.FieldDescription, field.TypeString, value)
	}
	if value, ok := ntuo.mutation.GetType(); ok {
		_spec.SetField(notifytarget.FieldType, field.TypeEnum, value)
	}
	if value, ok := ntuo.mutation.Status(); ok {
		_spec.SetField(notifytarget.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ntuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notifytarget.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ntuo.mutation.CreatedAt(); ok {
		_spec.SetField(notifytarget.FieldCreatedAt, field.TypeTime, value)
	}
	if ntuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifytarget.OwnerTable,
			Columns: []string{notifytarget.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifytarget.OwnerTable,
			Columns: []string{notifytarget.OwnerColumn},
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
	if ntuo.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTable,
			Columns: notifytarget.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		createE := &NotifyFlowTargetCreate{config: ntuo.config, mutation: newNotifyFlowTargetMutation(ntuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntuo.mutation.RemovedNotifyFlowIDs(); len(nodes) > 0 && !ntuo.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTable,
			Columns: notifytarget.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowTargetCreate{config: ntuo.config, mutation: newNotifyFlowTargetMutation(ntuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntuo.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTable,
			Columns: notifytarget.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowTargetCreate{config: ntuo.config, mutation: newNotifyFlowTargetMutation(ntuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ntuo.mutation.NotifyFlowTargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTargetTable,
			Columns: []string{notifytarget.NotifyFlowTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntuo.mutation.RemovedNotifyFlowTargetIDs(); len(nodes) > 0 && !ntuo.mutation.NotifyFlowTargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTargetTable,
			Columns: []string{notifytarget.NotifyFlowTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ntuo.mutation.NotifyFlowTargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTargetTable,
			Columns: []string{notifytarget.NotifyFlowTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NotifyTarget{config: ntuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ntuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notifytarget.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ntuo.mutation.done = true
	return _node, nil
}
