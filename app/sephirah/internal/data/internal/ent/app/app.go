// Code generated by ent, DO NOT EDIT.

package app

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the app type in the database.
	Label = "app"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldSourceAppID holds the string denoting the source_app_id field in the database.
	FieldSourceAppID = "source_app_id"
	// FieldSourceURL holds the string denoting the source_url field in the database.
	FieldSourceURL = "source_url"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldShortDescription holds the string denoting the short_description field in the database.
	FieldShortDescription = "short_description"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIconImageURL holds the string denoting the icon_image_url field in the database.
	FieldIconImageURL = "icon_image_url"
	// FieldHeroImageURL holds the string denoting the hero_image_url field in the database.
	FieldHeroImageURL = "hero_image_url"
	// FieldLogoImageURL holds the string denoting the logo_image_url field in the database.
	FieldLogoImageURL = "logo_image_url"
	// FieldReleaseDate holds the string denoting the release_date field in the database.
	FieldReleaseDate = "release_date"
	// FieldDeveloper holds the string denoting the developer field in the database.
	FieldDeveloper = "developer"
	// FieldPublisher holds the string denoting the publisher field in the database.
	FieldPublisher = "publisher"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgePurchasedByAccount holds the string denoting the purchased_by_account edge name in mutations.
	EdgePurchasedByAccount = "purchased_by_account"
	// EdgePurchasedByUser holds the string denoting the purchased_by_user edge name in mutations.
	EdgePurchasedByUser = "purchased_by_user"
	// EdgeAppPackage holds the string denoting the app_package edge name in mutations.
	EdgeAppPackage = "app_package"
	// EdgeBindInternal holds the string denoting the bind_internal edge name in mutations.
	EdgeBindInternal = "bind_internal"
	// EdgeBindExternal holds the string denoting the bind_external edge name in mutations.
	EdgeBindExternal = "bind_external"
	// Table holds the table name of the app in the database.
	Table = "apps"
	// PurchasedByAccountTable is the table that holds the purchased_by_account relation/edge. The primary key declared below.
	PurchasedByAccountTable = "account_purchased_app"
	// PurchasedByAccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	PurchasedByAccountInverseTable = "accounts"
	// PurchasedByUserTable is the table that holds the purchased_by_user relation/edge. The primary key declared below.
	PurchasedByUserTable = "user_purchased_app"
	// PurchasedByUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	PurchasedByUserInverseTable = "users"
	// AppPackageTable is the table that holds the app_package relation/edge.
	AppPackageTable = "app_packages"
	// AppPackageInverseTable is the table name for the AppPackage entity.
	// It exists in this package in order to avoid circular dependency with the "apppackage" package.
	AppPackageInverseTable = "app_packages"
	// AppPackageColumn is the table column denoting the app_package relation/edge.
	AppPackageColumn = "app_app_package"
	// BindInternalTable is the table that holds the bind_internal relation/edge.
	BindInternalTable = "apps"
	// BindInternalColumn is the table column denoting the bind_internal relation/edge.
	BindInternalColumn = "app_bind_external"
	// BindExternalTable is the table that holds the bind_external relation/edge.
	BindExternalTable = "apps"
	// BindExternalColumn is the table column denoting the bind_external relation/edge.
	BindExternalColumn = "app_bind_external"
)

// Columns holds all SQL columns for app fields.
var Columns = []string{
	FieldID,
	FieldSource,
	FieldSourceAppID,
	FieldSourceURL,
	FieldName,
	FieldType,
	FieldShortDescription,
	FieldDescription,
	FieldIconImageURL,
	FieldHeroImageURL,
	FieldLogoImageURL,
	FieldReleaseDate,
	FieldDeveloper,
	FieldPublisher,
	FieldVersion,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "apps"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"app_bind_external",
}

var (
	// PurchasedByAccountPrimaryKey and PurchasedByAccountColumn2 are the table columns denoting the
	// primary key for the purchased_by_account relation (M2M).
	PurchasedByAccountPrimaryKey = []string{"account_id", "app_id"}
	// PurchasedByUserPrimaryKey and PurchasedByUserColumn2 are the table columns denoting the
	// primary key for the purchased_by_user relation (M2M).
	PurchasedByUserPrimaryKey = []string{"user_id", "app_id"}
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

// Source defines the type for the "source" enum field.
type Source string

// Source values.
const (
	SourceInternal Source = "internal"
	SourceSteam    Source = "steam"
)

func (s Source) String() string {
	return string(s)
}

// SourceValidator is a validator for the "source" field enum values. It is called by the builders before save.
func SourceValidator(s Source) error {
	switch s {
	case SourceInternal, SourceSteam:
		return nil
	default:
		return fmt.Errorf("app: invalid enum value for source field: %q", s)
	}
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeUnknown Type = "unknown"
	TypeGame    Type = "game"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeUnknown, TypeGame:
		return nil
	default:
		return fmt.Errorf("app: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the App queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// BySourceAppID orders the results by the source_app_id field.
func BySourceAppID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceAppID, opts...).ToFunc()
}

// BySourceURL orders the results by the source_url field.
func BySourceURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceURL, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByShortDescription orders the results by the short_description field.
func ByShortDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShortDescription, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIconImageURL orders the results by the icon_image_url field.
func ByIconImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIconImageURL, opts...).ToFunc()
}

// ByHeroImageURL orders the results by the hero_image_url field.
func ByHeroImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeroImageURL, opts...).ToFunc()
}

// ByLogoImageURL orders the results by the logo_image_url field.
func ByLogoImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogoImageURL, opts...).ToFunc()
}

// ByReleaseDate orders the results by the release_date field.
func ByReleaseDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleaseDate, opts...).ToFunc()
}

// ByDeveloper orders the results by the developer field.
func ByDeveloper(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeveloper, opts...).ToFunc()
}

// ByPublisher orders the results by the publisher field.
func ByPublisher(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublisher, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByPurchasedByAccountCount orders the results by purchased_by_account count.
func ByPurchasedByAccountCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPurchasedByAccountStep(), opts...)
	}
}

// ByPurchasedByAccount orders the results by purchased_by_account terms.
func ByPurchasedByAccount(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPurchasedByAccountStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPurchasedByUserCount orders the results by purchased_by_user count.
func ByPurchasedByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPurchasedByUserStep(), opts...)
	}
}

// ByPurchasedByUser orders the results by purchased_by_user terms.
func ByPurchasedByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPurchasedByUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAppPackageCount orders the results by app_package count.
func ByAppPackageCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppPackageStep(), opts...)
	}
}

// ByAppPackage orders the results by app_package terms.
func ByAppPackage(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppPackageStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBindInternalField orders the results by bind_internal field.
func ByBindInternalField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBindInternalStep(), sql.OrderByField(field, opts...))
	}
}

// ByBindExternalCount orders the results by bind_external count.
func ByBindExternalCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBindExternalStep(), opts...)
	}
}

// ByBindExternal orders the results by bind_external terms.
func ByBindExternal(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBindExternalStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPurchasedByAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PurchasedByAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PurchasedByAccountTable, PurchasedByAccountPrimaryKey...),
	)
}
func newPurchasedByUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PurchasedByUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PurchasedByUserTable, PurchasedByUserPrimaryKey...),
	)
}
func newAppPackageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppPackageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AppPackageTable, AppPackageColumn),
	)
}
func newBindInternalStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BindInternalTable, BindInternalColumn),
	)
}
func newBindExternalStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BindExternalTable, BindExternalColumn),
	)
}
