package modelfeed

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Feed struct {
	ID          model.InternalID `json:"internal_id,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Link        string           `json:"link,omitempty"`
	Authors     []*Person        `json:"authors,omitempty"`
	Language    string           `json:"language,omitempty"`
	Image       *Image           `json:"image,omitempty"`
	Items       []*Item          `json:"items"`
	FeedType    string           `json:"feedType"`
	FeedVersion string           `json:"feedVersion"`
}

type Item struct {
	ID                model.InternalID `json:"internal_id,omitempty"`
	Title             string           `json:"title,omitempty"`
	Description       string           `json:"description,omitempty"`
	Content           string           `json:"content,omitempty"`
	Link              string           `json:"link,omitempty"`
	Updated           string           `json:"updated,omitempty"`
	UpdatedParsed     *time.Time       `json:"updatedParsed,omitempty"`
	Published         string           `json:"published,omitempty"`
	PublishedParsed   *time.Time       `json:"publishedParsed,omitempty"`
	Authors           []*Person        `json:"authors,omitempty"`
	GUID              string           `json:"guid,omitempty"`
	Image             *Image           `json:"image,omitempty"`
	Enclosures        []*Enclosure     `json:"enclosures,omitempty"`
	PublishPlatform   string           `json:"publish_platform,omitempty"`
	DigestDescription string           `json:"digest_description,omitempty"`
	DigestImages      []*Image         `json:"digest_images,omitempty"`
}

// Person is an individual specified in a feed
// (e.g. an author).
type Person struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// Image is an image that is the artwork for a given
// feed or item.
type Image struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

// Enclosure is a file associated with a given Item.
type Enclosure struct {
	URL    string `json:"url,omitempty"`
	Length string `json:"length,omitempty"`
	Type   string `json:"type,omitempty"`
}

// Len returns the length of Items.
func (f Feed) Len() int {
	return len(f.Items)
}

// Less compares PublishedParsed of Items[i], Items[k]
// and returns true if Items[i] is less than Items[k].
func (f Feed) Less(i, k int) bool {
	return f.Items[i].PublishedParsed.Before(
		*f.Items[k].PublishedParsed,
	)
}

// Swap swaps Items[i] and Items[k].
func (f Feed) Swap(i, k int) {
	f.Items[i], f.Items[k] = f.Items[k], f.Items[i]
}
