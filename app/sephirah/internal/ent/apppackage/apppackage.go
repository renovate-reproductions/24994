// Code generated by ent, DO NOT EDIT.

package apppackage

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the apppackage type in the database.
	Label = "app_package"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldSourceID holds the string denoting the source_id field in the database.
	FieldSourceID = "source_id"
	// FieldSourcePackageID holds the string denoting the source_package_id field in the database.
	FieldSourcePackageID = "source_package_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// FieldBinaryName holds the string denoting the binary_name field in the database.
	FieldBinaryName = "binary_name"
	// FieldBinarySizeByte holds the string denoting the binary_size_byte field in the database.
	FieldBinarySizeByte = "binary_size_byte"
	// FieldBinaryPublicURL holds the string denoting the binary_public_url field in the database.
	FieldBinaryPublicURL = "binary_public_url"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// Table holds the table name of the apppackage in the database.
	Table = "app_packages"
	// AppTable is the table that holds the app relation/edge.
	AppTable = "app_packages"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "apps"
	// AppColumn is the table column denoting the app relation/edge.
	AppColumn = "app_app_package"
)

// Columns holds all SQL columns for apppackage fields.
var Columns = []string{
	FieldID,
	FieldSource,
	FieldSourceID,
	FieldSourcePackageID,
	FieldName,
	FieldDescription,
	FieldPublic,
	FieldBinaryName,
	FieldBinarySizeByte,
	FieldBinaryPublicURL,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "app_packages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"app_app_package",
}

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

// Source defines the type for the "source" enum field.
type Source string

// Source values.
const (
	SourceManual   Source = "manual"
	SourceSentinel Source = "sentinel"
)

func (s Source) String() string {
	return string(s)
}

// SourceValidator is a validator for the "source" field enum values. It is called by the builders before save.
func SourceValidator(s Source) error {
	switch s {
	case SourceManual, SourceSentinel:
		return nil
	default:
		return fmt.Errorf("apppackage: invalid enum value for source field: %q", s)
	}
}
