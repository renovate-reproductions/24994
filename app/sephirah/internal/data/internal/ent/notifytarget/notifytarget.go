// Code generated by ent, DO NOT EDIT.

package notifytarget

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the notifytarget type in the database.
	Label = "notify_target"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeNotifyFlow holds the string denoting the notify_flow edge name in mutations.
	EdgeNotifyFlow = "notify_flow"
	// EdgeNotifyFlowTarget holds the string denoting the notify_flow_target edge name in mutations.
	EdgeNotifyFlowTarget = "notify_flow_target"
	// Table holds the table name of the notifytarget in the database.
	Table = "notify_targets"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "notify_targets"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_notify_target"
	// NotifyFlowTable is the table that holds the notify_flow relation/edge. The primary key declared below.
	NotifyFlowTable = "notify_flow_targets"
	// NotifyFlowInverseTable is the table name for the NotifyFlow entity.
	// It exists in this package in order to avoid circular dependency with the "notifyflow" package.
	NotifyFlowInverseTable = "notify_flows"
	// NotifyFlowTargetTable is the table that holds the notify_flow_target relation/edge.
	NotifyFlowTargetTable = "notify_flow_targets"
	// NotifyFlowTargetInverseTable is the table name for the NotifyFlowTarget entity.
	// It exists in this package in order to avoid circular dependency with the "notifyflowtarget" package.
	NotifyFlowTargetInverseTable = "notify_flow_targets"
	// NotifyFlowTargetColumn is the table column denoting the notify_flow_target relation/edge.
	NotifyFlowTargetColumn = "notify_target_id"
)

// Columns holds all SQL columns for notifytarget fields.
var Columns = []string{
	FieldID,
	FieldToken,
	FieldName,
	FieldDescription,
	FieldType,
	FieldStatus,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "notify_targets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_notify_target",
}

var (
	// NotifyFlowPrimaryKey and NotifyFlowColumn2 are the table columns denoting the
	// primary key for the notify_flow relation (M2M).
	NotifyFlowPrimaryKey = []string{"notify_flow_id", "notify_target_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeTelegram Type = "telegram"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeTelegram:
		return nil
	default:
		return fmt.Errorf("notifytarget: invalid enum value for type field: %q", _type)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusActive  Status = "active"
	StatusSuspend Status = "suspend"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusSuspend:
		return nil
	default:
		return fmt.Errorf("notifytarget: invalid enum value for status field: %q", s)
	}
}
