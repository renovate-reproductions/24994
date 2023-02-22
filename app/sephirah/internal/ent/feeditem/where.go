// Code generated by ent, DO NOT EDIT.

package feeditem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldID, id))
}

// InternalID applies equality check predicate on the "internal_id" field. It's identical to InternalIDEQ.
func InternalID(v int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldInternalID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldDescription, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldContent, v))
}

// GUID applies equality check predicate on the "guid" field. It's identical to GUIDEQ.
func GUID(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldGUID, v))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldLink, v))
}

// Published applies equality check predicate on the "published" field. It's identical to PublishedEQ.
func Published(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldPublished, v))
}

// PublishedParsed applies equality check predicate on the "published_parsed" field. It's identical to PublishedParsedEQ.
func PublishedParsed(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldPublishedParsed, v))
}

// Updated applies equality check predicate on the "updated" field. It's identical to UpdatedEQ.
func Updated(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldUpdated, v))
}

// UpdatedParsed applies equality check predicate on the "updated_parsed" field. It's identical to UpdatedParsedEQ.
func UpdatedParsed(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldUpdatedParsed, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldCreatedAt, v))
}

// InternalIDEQ applies the EQ predicate on the "internal_id" field.
func InternalIDEQ(v int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldInternalID, v))
}

// InternalIDNEQ applies the NEQ predicate on the "internal_id" field.
func InternalIDNEQ(v int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldInternalID, v))
}

// InternalIDIn applies the In predicate on the "internal_id" field.
func InternalIDIn(vs ...int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldInternalID, vs...))
}

// InternalIDNotIn applies the NotIn predicate on the "internal_id" field.
func InternalIDNotIn(vs ...int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldInternalID, vs...))
}

// InternalIDGT applies the GT predicate on the "internal_id" field.
func InternalIDGT(v int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldInternalID, v))
}

// InternalIDGTE applies the GTE predicate on the "internal_id" field.
func InternalIDGTE(v int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldInternalID, v))
}

// InternalIDLT applies the LT predicate on the "internal_id" field.
func InternalIDLT(v int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldInternalID, v))
}

// InternalIDLTE applies the LTE predicate on the "internal_id" field.
func InternalIDLTE(v int64) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldInternalID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContainsFold(FieldDescription, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContainsFold(FieldContent, v))
}

// GUIDEQ applies the EQ predicate on the "guid" field.
func GUIDEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldGUID, v))
}

// GUIDNEQ applies the NEQ predicate on the "guid" field.
func GUIDNEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldGUID, v))
}

// GUIDIn applies the In predicate on the "guid" field.
func GUIDIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldGUID, vs...))
}

// GUIDNotIn applies the NotIn predicate on the "guid" field.
func GUIDNotIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldGUID, vs...))
}

// GUIDGT applies the GT predicate on the "guid" field.
func GUIDGT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldGUID, v))
}

// GUIDGTE applies the GTE predicate on the "guid" field.
func GUIDGTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldGUID, v))
}

// GUIDLT applies the LT predicate on the "guid" field.
func GUIDLT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldGUID, v))
}

// GUIDLTE applies the LTE predicate on the "guid" field.
func GUIDLTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldGUID, v))
}

// GUIDContains applies the Contains predicate on the "guid" field.
func GUIDContains(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContains(FieldGUID, v))
}

// GUIDHasPrefix applies the HasPrefix predicate on the "guid" field.
func GUIDHasPrefix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasPrefix(FieldGUID, v))
}

// GUIDHasSuffix applies the HasSuffix predicate on the "guid" field.
func GUIDHasSuffix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasSuffix(FieldGUID, v))
}

// GUIDEqualFold applies the EqualFold predicate on the "guid" field.
func GUIDEqualFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEqualFold(FieldGUID, v))
}

// GUIDContainsFold applies the ContainsFold predicate on the "guid" field.
func GUIDContainsFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContainsFold(FieldGUID, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContainsFold(FieldLink, v))
}

// PublishedEQ applies the EQ predicate on the "published" field.
func PublishedEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldPublished, v))
}

// PublishedNEQ applies the NEQ predicate on the "published" field.
func PublishedNEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldPublished, v))
}

// PublishedIn applies the In predicate on the "published" field.
func PublishedIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldPublished, vs...))
}

// PublishedNotIn applies the NotIn predicate on the "published" field.
func PublishedNotIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldPublished, vs...))
}

// PublishedGT applies the GT predicate on the "published" field.
func PublishedGT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldPublished, v))
}

// PublishedGTE applies the GTE predicate on the "published" field.
func PublishedGTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldPublished, v))
}

// PublishedLT applies the LT predicate on the "published" field.
func PublishedLT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldPublished, v))
}

// PublishedLTE applies the LTE predicate on the "published" field.
func PublishedLTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldPublished, v))
}

// PublishedContains applies the Contains predicate on the "published" field.
func PublishedContains(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContains(FieldPublished, v))
}

// PublishedHasPrefix applies the HasPrefix predicate on the "published" field.
func PublishedHasPrefix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasPrefix(FieldPublished, v))
}

// PublishedHasSuffix applies the HasSuffix predicate on the "published" field.
func PublishedHasSuffix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasSuffix(FieldPublished, v))
}

// PublishedEqualFold applies the EqualFold predicate on the "published" field.
func PublishedEqualFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEqualFold(FieldPublished, v))
}

// PublishedContainsFold applies the ContainsFold predicate on the "published" field.
func PublishedContainsFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContainsFold(FieldPublished, v))
}

// PublishedParsedEQ applies the EQ predicate on the "published_parsed" field.
func PublishedParsedEQ(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldPublishedParsed, v))
}

// PublishedParsedNEQ applies the NEQ predicate on the "published_parsed" field.
func PublishedParsedNEQ(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldPublishedParsed, v))
}

// PublishedParsedIn applies the In predicate on the "published_parsed" field.
func PublishedParsedIn(vs ...time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldPublishedParsed, vs...))
}

// PublishedParsedNotIn applies the NotIn predicate on the "published_parsed" field.
func PublishedParsedNotIn(vs ...time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldPublishedParsed, vs...))
}

// PublishedParsedGT applies the GT predicate on the "published_parsed" field.
func PublishedParsedGT(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldPublishedParsed, v))
}

// PublishedParsedGTE applies the GTE predicate on the "published_parsed" field.
func PublishedParsedGTE(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldPublishedParsed, v))
}

// PublishedParsedLT applies the LT predicate on the "published_parsed" field.
func PublishedParsedLT(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldPublishedParsed, v))
}

// PublishedParsedLTE applies the LTE predicate on the "published_parsed" field.
func PublishedParsedLTE(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldPublishedParsed, v))
}

// UpdatedEQ applies the EQ predicate on the "updated" field.
func UpdatedEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldUpdated, v))
}

// UpdatedNEQ applies the NEQ predicate on the "updated" field.
func UpdatedNEQ(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldUpdated, v))
}

// UpdatedIn applies the In predicate on the "updated" field.
func UpdatedIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldUpdated, vs...))
}

// UpdatedNotIn applies the NotIn predicate on the "updated" field.
func UpdatedNotIn(vs ...string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldUpdated, vs...))
}

// UpdatedGT applies the GT predicate on the "updated" field.
func UpdatedGT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldUpdated, v))
}

// UpdatedGTE applies the GTE predicate on the "updated" field.
func UpdatedGTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldUpdated, v))
}

// UpdatedLT applies the LT predicate on the "updated" field.
func UpdatedLT(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldUpdated, v))
}

// UpdatedLTE applies the LTE predicate on the "updated" field.
func UpdatedLTE(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldUpdated, v))
}

// UpdatedContains applies the Contains predicate on the "updated" field.
func UpdatedContains(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContains(FieldUpdated, v))
}

// UpdatedHasPrefix applies the HasPrefix predicate on the "updated" field.
func UpdatedHasPrefix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasPrefix(FieldUpdated, v))
}

// UpdatedHasSuffix applies the HasSuffix predicate on the "updated" field.
func UpdatedHasSuffix(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldHasSuffix(FieldUpdated, v))
}

// UpdatedEqualFold applies the EqualFold predicate on the "updated" field.
func UpdatedEqualFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEqualFold(FieldUpdated, v))
}

// UpdatedContainsFold applies the ContainsFold predicate on the "updated" field.
func UpdatedContainsFold(v string) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldContainsFold(FieldUpdated, v))
}

// UpdatedParsedEQ applies the EQ predicate on the "updated_parsed" field.
func UpdatedParsedEQ(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldUpdatedParsed, v))
}

// UpdatedParsedNEQ applies the NEQ predicate on the "updated_parsed" field.
func UpdatedParsedNEQ(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldUpdatedParsed, v))
}

// UpdatedParsedIn applies the In predicate on the "updated_parsed" field.
func UpdatedParsedIn(vs ...time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldUpdatedParsed, vs...))
}

// UpdatedParsedNotIn applies the NotIn predicate on the "updated_parsed" field.
func UpdatedParsedNotIn(vs ...time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldUpdatedParsed, vs...))
}

// UpdatedParsedGT applies the GT predicate on the "updated_parsed" field.
func UpdatedParsedGT(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldUpdatedParsed, v))
}

// UpdatedParsedGTE applies the GTE predicate on the "updated_parsed" field.
func UpdatedParsedGTE(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldUpdatedParsed, v))
}

// UpdatedParsedLT applies the LT predicate on the "updated_parsed" field.
func UpdatedParsedLT(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldUpdatedParsed, v))
}

// UpdatedParsedLTE applies the LTE predicate on the "updated_parsed" field.
func UpdatedParsedLTE(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldUpdatedParsed, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FeedItem {
	return predicate.FeedItem(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FeedItem) predicate.FeedItem {
	return predicate.FeedItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FeedItem) predicate.FeedItem {
	return predicate.FeedItem(func(s *sql.Selector) {
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
func Not(p predicate.FeedItem) predicate.FeedItem {
	return predicate.FeedItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
