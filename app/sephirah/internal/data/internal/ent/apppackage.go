// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// AppPackage is the model entity for the AppPackage schema.
type AppPackage struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// Source holds the value of the "source" field.
	Source apppackage.Source `json:"source,omitempty"`
	// SourceID holds the value of the "source_id" field.
	SourceID model.InternalID `json:"source_id,omitempty"`
	// SourcePackageID holds the value of the "source_package_id" field.
	SourcePackageID string `json:"source_package_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Public holds the value of the "public" field.
	Public bool `json:"public,omitempty"`
	// BinaryName holds the value of the "binary_name" field.
	BinaryName string `json:"binary_name,omitempty"`
	// BinarySizeByte holds the value of the "binary_size_byte" field.
	BinarySizeByte int64 `json:"binary_size_byte,omitempty"`
	// BinaryPublicURL holds the value of the "binary_public_url" field.
	BinaryPublicURL string `json:"binary_public_url,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppPackageQuery when eager-loading is set.
	Edges            AppPackageEdges `json:"edges"`
	app_app_package  *model.InternalID
	user_app_package *model.InternalID
}

// AppPackageEdges holds the relations/edges for other nodes in the graph.
type AppPackageEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// App holds the value of the app edge.
	App *App `json:"app,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppPackageEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// AppOrErr returns the App value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppPackageEdges) AppOrErr() (*App, error) {
	if e.loadedTypes[1] {
		if e.App == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: app.Label}
		}
		return e.App, nil
	}
	return nil, &NotLoadedError{edge: "app"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppPackage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apppackage.FieldPublic:
			values[i] = new(sql.NullBool)
		case apppackage.FieldID, apppackage.FieldSourceID, apppackage.FieldBinarySizeByte:
			values[i] = new(sql.NullInt64)
		case apppackage.FieldSource, apppackage.FieldSourcePackageID, apppackage.FieldName, apppackage.FieldDescription, apppackage.FieldBinaryName, apppackage.FieldBinaryPublicURL:
			values[i] = new(sql.NullString)
		case apppackage.FieldUpdatedAt, apppackage.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case apppackage.ForeignKeys[0]: // app_app_package
			values[i] = new(sql.NullInt64)
		case apppackage.ForeignKeys[1]: // user_app_package
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppPackage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppPackage fields.
func (ap *AppPackage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apppackage.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ap.ID = model.InternalID(value.Int64)
			}
		case apppackage.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				ap.Source = apppackage.Source(value.String)
			}
		case apppackage.FieldSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_id", values[i])
			} else if value.Valid {
				ap.SourceID = model.InternalID(value.Int64)
			}
		case apppackage.FieldSourcePackageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_package_id", values[i])
			} else if value.Valid {
				ap.SourcePackageID = value.String
			}
		case apppackage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ap.Name = value.String
			}
		case apppackage.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ap.Description = value.String
			}
		case apppackage.FieldPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public", values[i])
			} else if value.Valid {
				ap.Public = value.Bool
			}
		case apppackage.FieldBinaryName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field binary_name", values[i])
			} else if value.Valid {
				ap.BinaryName = value.String
			}
		case apppackage.FieldBinarySizeByte:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field binary_size_byte", values[i])
			} else if value.Valid {
				ap.BinarySizeByte = value.Int64
			}
		case apppackage.FieldBinaryPublicURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field binary_public_url", values[i])
			} else if value.Valid {
				ap.BinaryPublicURL = value.String
			}
		case apppackage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ap.UpdatedAt = value.Time
			}
		case apppackage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ap.CreatedAt = value.Time
			}
		case apppackage.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_app_package", values[i])
			} else if value.Valid {
				ap.app_app_package = new(model.InternalID)
				*ap.app_app_package = model.InternalID(value.Int64)
			}
		case apppackage.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_app_package", values[i])
			} else if value.Valid {
				ap.user_app_package = new(model.InternalID)
				*ap.user_app_package = model.InternalID(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the AppPackage entity.
func (ap *AppPackage) QueryOwner() *UserQuery {
	return NewAppPackageClient(ap.config).QueryOwner(ap)
}

// QueryApp queries the "app" edge of the AppPackage entity.
func (ap *AppPackage) QueryApp() *AppQuery {
	return NewAppPackageClient(ap.config).QueryApp(ap)
}

// Update returns a builder for updating this AppPackage.
// Note that you need to call AppPackage.Unwrap() before calling this method if this AppPackage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ap *AppPackage) Update() *AppPackageUpdateOne {
	return NewAppPackageClient(ap.config).UpdateOne(ap)
}

// Unwrap unwraps the AppPackage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ap *AppPackage) Unwrap() *AppPackage {
	_tx, ok := ap.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppPackage is not a transactional entity")
	}
	ap.config.driver = _tx.drv
	return ap
}

// String implements the fmt.Stringer.
func (ap *AppPackage) String() string {
	var builder strings.Builder
	builder.WriteString("AppPackage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ap.ID))
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", ap.Source))
	builder.WriteString(", ")
	builder.WriteString("source_id=")
	builder.WriteString(fmt.Sprintf("%v", ap.SourceID))
	builder.WriteString(", ")
	builder.WriteString("source_package_id=")
	builder.WriteString(ap.SourcePackageID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ap.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ap.Description)
	builder.WriteString(", ")
	builder.WriteString("public=")
	builder.WriteString(fmt.Sprintf("%v", ap.Public))
	builder.WriteString(", ")
	builder.WriteString("binary_name=")
	builder.WriteString(ap.BinaryName)
	builder.WriteString(", ")
	builder.WriteString("binary_size_byte=")
	builder.WriteString(fmt.Sprintf("%v", ap.BinarySizeByte))
	builder.WriteString(", ")
	builder.WriteString("binary_public_url=")
	builder.WriteString(ap.BinaryPublicURL)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ap.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ap.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AppPackages is a parsable slice of AppPackage.
type AppPackages []*AppPackage
