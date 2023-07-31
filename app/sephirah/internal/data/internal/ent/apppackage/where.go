// Code generated by ent, DO NOT EDIT.

package apppackage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldID, id))
}

// SourceID applies equality check predicate on the "source_id" field. It's identical to SourceIDEQ.
func SourceID(v model.InternalID) predicate.AppPackage {
	vc := int64(v)
	return predicate.AppPackage(sql.FieldEQ(FieldSourceID, vc))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDescription, v))
}

// Public applies equality check predicate on the "public" field. It's identical to PublicEQ.
func Public(v bool) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPublic, v))
}

// BinaryName applies equality check predicate on the "binary_name" field. It's identical to BinaryNameEQ.
func BinaryName(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldBinaryName, v))
}

// BinarySizeBytes applies equality check predicate on the "binary_size_bytes" field. It's identical to BinarySizeBytesEQ.
func BinarySizeBytes(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldBinarySizeBytes, v))
}

// BinaryPublicURL applies equality check predicate on the "binary_public_url" field. It's identical to BinaryPublicURLEQ.
func BinaryPublicURL(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldBinaryPublicURL, v))
}

// BinarySha256 applies equality check predicate on the "binary_sha256" field. It's identical to BinarySha256EQ.
func BinarySha256(v []byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldBinarySha256, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldCreatedAt, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v Source) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v Source) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...Source) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...Source) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldSource, vs...))
}

// SourceIDEQ applies the EQ predicate on the "source_id" field.
func SourceIDEQ(v model.InternalID) predicate.AppPackage {
	vc := int64(v)
	return predicate.AppPackage(sql.FieldEQ(FieldSourceID, vc))
}

// SourceIDNEQ applies the NEQ predicate on the "source_id" field.
func SourceIDNEQ(v model.InternalID) predicate.AppPackage {
	vc := int64(v)
	return predicate.AppPackage(sql.FieldNEQ(FieldSourceID, vc))
}

// SourceIDIn applies the In predicate on the "source_id" field.
func SourceIDIn(vs ...model.InternalID) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppPackage(sql.FieldIn(FieldSourceID, v...))
}

// SourceIDNotIn applies the NotIn predicate on the "source_id" field.
func SourceIDNotIn(vs ...model.InternalID) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppPackage(sql.FieldNotIn(FieldSourceID, v...))
}

// SourceIDGT applies the GT predicate on the "source_id" field.
func SourceIDGT(v model.InternalID) predicate.AppPackage {
	vc := int64(v)
	return predicate.AppPackage(sql.FieldGT(FieldSourceID, vc))
}

// SourceIDGTE applies the GTE predicate on the "source_id" field.
func SourceIDGTE(v model.InternalID) predicate.AppPackage {
	vc := int64(v)
	return predicate.AppPackage(sql.FieldGTE(FieldSourceID, vc))
}

// SourceIDLT applies the LT predicate on the "source_id" field.
func SourceIDLT(v model.InternalID) predicate.AppPackage {
	vc := int64(v)
	return predicate.AppPackage(sql.FieldLT(FieldSourceID, vc))
}

// SourceIDLTE applies the LTE predicate on the "source_id" field.
func SourceIDLTE(v model.InternalID) predicate.AppPackage {
	vc := int64(v)
	return predicate.AppPackage(sql.FieldLTE(FieldSourceID, vc))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldDescription, v))
}

// PublicEQ applies the EQ predicate on the "public" field.
func PublicEQ(v bool) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPublic, v))
}

// PublicNEQ applies the NEQ predicate on the "public" field.
func PublicNEQ(v bool) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldPublic, v))
}

// BinaryNameEQ applies the EQ predicate on the "binary_name" field.
func BinaryNameEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldBinaryName, v))
}

// BinaryNameNEQ applies the NEQ predicate on the "binary_name" field.
func BinaryNameNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldBinaryName, v))
}

// BinaryNameIn applies the In predicate on the "binary_name" field.
func BinaryNameIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldBinaryName, vs...))
}

// BinaryNameNotIn applies the NotIn predicate on the "binary_name" field.
func BinaryNameNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldBinaryName, vs...))
}

// BinaryNameGT applies the GT predicate on the "binary_name" field.
func BinaryNameGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldBinaryName, v))
}

// BinaryNameGTE applies the GTE predicate on the "binary_name" field.
func BinaryNameGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldBinaryName, v))
}

// BinaryNameLT applies the LT predicate on the "binary_name" field.
func BinaryNameLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldBinaryName, v))
}

// BinaryNameLTE applies the LTE predicate on the "binary_name" field.
func BinaryNameLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldBinaryName, v))
}

// BinaryNameContains applies the Contains predicate on the "binary_name" field.
func BinaryNameContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldBinaryName, v))
}

// BinaryNameHasPrefix applies the HasPrefix predicate on the "binary_name" field.
func BinaryNameHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldBinaryName, v))
}

// BinaryNameHasSuffix applies the HasSuffix predicate on the "binary_name" field.
func BinaryNameHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldBinaryName, v))
}

// BinaryNameIsNil applies the IsNil predicate on the "binary_name" field.
func BinaryNameIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldBinaryName))
}

// BinaryNameNotNil applies the NotNil predicate on the "binary_name" field.
func BinaryNameNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldBinaryName))
}

// BinaryNameEqualFold applies the EqualFold predicate on the "binary_name" field.
func BinaryNameEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldBinaryName, v))
}

// BinaryNameContainsFold applies the ContainsFold predicate on the "binary_name" field.
func BinaryNameContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldBinaryName, v))
}

// BinarySizeBytesEQ applies the EQ predicate on the "binary_size_bytes" field.
func BinarySizeBytesEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldBinarySizeBytes, v))
}

// BinarySizeBytesNEQ applies the NEQ predicate on the "binary_size_bytes" field.
func BinarySizeBytesNEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldBinarySizeBytes, v))
}

// BinarySizeBytesIn applies the In predicate on the "binary_size_bytes" field.
func BinarySizeBytesIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldBinarySizeBytes, vs...))
}

// BinarySizeBytesNotIn applies the NotIn predicate on the "binary_size_bytes" field.
func BinarySizeBytesNotIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldBinarySizeBytes, vs...))
}

// BinarySizeBytesGT applies the GT predicate on the "binary_size_bytes" field.
func BinarySizeBytesGT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldBinarySizeBytes, v))
}

// BinarySizeBytesGTE applies the GTE predicate on the "binary_size_bytes" field.
func BinarySizeBytesGTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldBinarySizeBytes, v))
}

// BinarySizeBytesLT applies the LT predicate on the "binary_size_bytes" field.
func BinarySizeBytesLT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldBinarySizeBytes, v))
}

// BinarySizeBytesLTE applies the LTE predicate on the "binary_size_bytes" field.
func BinarySizeBytesLTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldBinarySizeBytes, v))
}

// BinarySizeBytesIsNil applies the IsNil predicate on the "binary_size_bytes" field.
func BinarySizeBytesIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldBinarySizeBytes))
}

// BinarySizeBytesNotNil applies the NotNil predicate on the "binary_size_bytes" field.
func BinarySizeBytesNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldBinarySizeBytes))
}

// BinaryPublicURLEQ applies the EQ predicate on the "binary_public_url" field.
func BinaryPublicURLEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldBinaryPublicURL, v))
}

// BinaryPublicURLNEQ applies the NEQ predicate on the "binary_public_url" field.
func BinaryPublicURLNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldBinaryPublicURL, v))
}

// BinaryPublicURLIn applies the In predicate on the "binary_public_url" field.
func BinaryPublicURLIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldBinaryPublicURL, vs...))
}

// BinaryPublicURLNotIn applies the NotIn predicate on the "binary_public_url" field.
func BinaryPublicURLNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldBinaryPublicURL, vs...))
}

// BinaryPublicURLGT applies the GT predicate on the "binary_public_url" field.
func BinaryPublicURLGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldBinaryPublicURL, v))
}

// BinaryPublicURLGTE applies the GTE predicate on the "binary_public_url" field.
func BinaryPublicURLGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldBinaryPublicURL, v))
}

// BinaryPublicURLLT applies the LT predicate on the "binary_public_url" field.
func BinaryPublicURLLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldBinaryPublicURL, v))
}

// BinaryPublicURLLTE applies the LTE predicate on the "binary_public_url" field.
func BinaryPublicURLLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldBinaryPublicURL, v))
}

// BinaryPublicURLContains applies the Contains predicate on the "binary_public_url" field.
func BinaryPublicURLContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldBinaryPublicURL, v))
}

// BinaryPublicURLHasPrefix applies the HasPrefix predicate on the "binary_public_url" field.
func BinaryPublicURLHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldBinaryPublicURL, v))
}

// BinaryPublicURLHasSuffix applies the HasSuffix predicate on the "binary_public_url" field.
func BinaryPublicURLHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldBinaryPublicURL, v))
}

// BinaryPublicURLIsNil applies the IsNil predicate on the "binary_public_url" field.
func BinaryPublicURLIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldBinaryPublicURL))
}

// BinaryPublicURLNotNil applies the NotNil predicate on the "binary_public_url" field.
func BinaryPublicURLNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldBinaryPublicURL))
}

// BinaryPublicURLEqualFold applies the EqualFold predicate on the "binary_public_url" field.
func BinaryPublicURLEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldBinaryPublicURL, v))
}

// BinaryPublicURLContainsFold applies the ContainsFold predicate on the "binary_public_url" field.
func BinaryPublicURLContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldBinaryPublicURL, v))
}

// BinarySha256EQ applies the EQ predicate on the "binary_sha256" field.
func BinarySha256EQ(v []byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldBinarySha256, v))
}

// BinarySha256NEQ applies the NEQ predicate on the "binary_sha256" field.
func BinarySha256NEQ(v []byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldBinarySha256, v))
}

// BinarySha256In applies the In predicate on the "binary_sha256" field.
func BinarySha256In(vs ...[]byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldBinarySha256, vs...))
}

// BinarySha256NotIn applies the NotIn predicate on the "binary_sha256" field.
func BinarySha256NotIn(vs ...[]byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldBinarySha256, vs...))
}

// BinarySha256GT applies the GT predicate on the "binary_sha256" field.
func BinarySha256GT(v []byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldBinarySha256, v))
}

// BinarySha256GTE applies the GTE predicate on the "binary_sha256" field.
func BinarySha256GTE(v []byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldBinarySha256, v))
}

// BinarySha256LT applies the LT predicate on the "binary_sha256" field.
func BinarySha256LT(v []byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldBinarySha256, v))
}

// BinarySha256LTE applies the LTE predicate on the "binary_sha256" field.
func BinarySha256LTE(v []byte) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldBinarySha256, v))
}

// BinarySha256IsNil applies the IsNil predicate on the "binary_sha256" field.
func BinarySha256IsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldBinarySha256))
}

// BinarySha256NotNil applies the NotNil predicate on the "binary_sha256" field.
func BinarySha256NotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldBinarySha256))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasApp applies the HasEdge predicate on the "app" edge.
func HasApp() predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppWith applies the HasEdge predicate on the "app" edge with a given conditions (other predicates).
func HasAppWith(preds ...predicate.App) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		step := newAppStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		p(s.Not())
	})
}
