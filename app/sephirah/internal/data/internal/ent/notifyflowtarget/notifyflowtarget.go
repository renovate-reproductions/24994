// Code generated by ent, DO NOT EDIT.

package notifyflowtarget

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the notifyflowtarget type in the database.
	Label = "notify_flow_target"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNotifyFlowID holds the string denoting the notify_flow_id field in the database.
	FieldNotifyFlowID = "notify_flow_id"
	// FieldNotifyTargetID holds the string denoting the notify_target_id field in the database.
	FieldNotifyTargetID = "notify_target_id"
	// FieldChannelID holds the string denoting the channel_id field in the database.
	FieldChannelID = "channel_id"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeNotifyFlow holds the string denoting the notify_flow edge name in mutations.
	EdgeNotifyFlow = "notify_flow"
	// EdgeNotifyTarget holds the string denoting the notify_target edge name in mutations.
	EdgeNotifyTarget = "notify_target"
	// Table holds the table name of the notifyflowtarget in the database.
	Table = "notify_flow_targets"
	// NotifyFlowTable is the table that holds the notify_flow relation/edge.
	NotifyFlowTable = "notify_flow_targets"
	// NotifyFlowInverseTable is the table name for the NotifyFlow entity.
	// It exists in this package in order to avoid circular dependency with the "notifyflow" package.
	NotifyFlowInverseTable = "notify_flows"
	// NotifyFlowColumn is the table column denoting the notify_flow relation/edge.
	NotifyFlowColumn = "notify_flow_id"
	// NotifyTargetTable is the table that holds the notify_target relation/edge.
	NotifyTargetTable = "notify_flow_targets"
	// NotifyTargetInverseTable is the table name for the NotifyTarget entity.
	// It exists in this package in order to avoid circular dependency with the "notifytarget" package.
	NotifyTargetInverseTable = "notify_targets"
	// NotifyTargetColumn is the table column denoting the notify_target relation/edge.
	NotifyTargetColumn = "notify_target_id"
)

// Columns holds all SQL columns for notifyflowtarget fields.
var Columns = []string{
	FieldID,
	FieldNotifyFlowID,
	FieldNotifyTargetID,
	FieldChannelID,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the NotifyFlowTarget queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNotifyFlowID orders the results by the notify_flow_id field.
func ByNotifyFlowID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotifyFlowID, opts...).ToFunc()
}

// ByNotifyTargetID orders the results by the notify_target_id field.
func ByNotifyTargetID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotifyTargetID, opts...).ToFunc()
}

// ByChannelID orders the results by the channel_id field.
func ByChannelID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelID, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByNotifyFlowField orders the results by notify_flow field.
func ByNotifyFlowField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotifyFlowStep(), sql.OrderByField(field, opts...))
	}
}

// ByNotifyTargetField orders the results by notify_target field.
func ByNotifyTargetField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotifyTargetStep(), sql.OrderByField(field, opts...))
	}
}
func newNotifyFlowStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotifyFlowInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, NotifyFlowTable, NotifyFlowColumn),
	)
}
func newNotifyTargetStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotifyTargetInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, NotifyTargetTable, NotifyTargetColumn),
	)
}
