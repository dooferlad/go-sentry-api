package sentry

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dooferlad/go-sentry-api/datatype"
)

// Tag is used for a event
type Tag struct {
	Value *string `json:"value,omitempty" bson:"value,omitempty"`
	Key   *string `json:"key,omitempty" bson:"key,omitempty"`
}

//User is the user that was affected
type User struct {
	Username *string `json:"username,omitempty" bson:"username,omitempty"`
	Email    *string `json:"email,omitempty" bson:"email,omitempty"`
	ID       *string `json:"id,omitempty" bson:"id,omitempty"`
}

// Entry is the entry for the message/stacktrace/etc...
type Entry struct {
	Type string          `json:"type,omitempty" bson:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty" bson:"data,omitempty"`
	S string             `json:"s" bson:"s"`
}

//GetInterface will convert the entry into a go interface
func (e *Entry) GetInterface() (string, interface{}, error) {
	var destination interface{}

	switch e.Type {
	case "message":
		destination = new(datatype.Message)
	case "stacktrace":
		destination = new(datatype.Stacktrace)
	case "exception":
		destination = new(datatype.Exception)
	case "request":
		destination = new(datatype.Request)
	case "template":
		destination = new(datatype.Template)
	case "user":
		destination = new(datatype.User)
	case "query":
		destination = new(datatype.Query)
	case "breadcrumbs":
		destination = new(datatype.Breadcrumb)
	}

	err := json.Unmarshal(e.Data, &destination)
	return e.Type, destination, err
}

// Event is the event that was created on the app and sentry reported on
type Event struct {
	EventID         string                 `json:"eventID,omitempty" bson:"eventID,omitempty"`
	UserReport      string                 `json:"userReport,omitempty" bson:"userReport,omitempty"`
	NextEventID     string                 `json:"nextEventID,omitempty" bson:"nextEventID,omitempty"`
	PreviousEventID string                 `json:"previousEventID,omitempty" bson:"previousEventID,omitempty"`
	Message         string                 `json:"message,omitempty" bson:"message,omitempty"`
	ID              string                 `json:"id,omitempty" bson:"_id,omitempty"`
	Size            int                    `json:"size,omitempty" bson:"size,omitempty"`
	Platform        string                 `json:"platform,omitempty" bson:"platform,omitempty"`
	Type            string                 `json:"type,omitempty" bson:"type,omitempty"`
	Metadata        map[string]string      `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Tags            []Tag                  `json:"tags,omitempty" bson:"tags,omitempty"`
	DateCreated     time.Time              `json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
	DateReceived    time.Time              `json:"dateReceived,omitempty" bson:"dateReceived,omitempty"`
	User            User                   `json:"user,omitempty" bson:"user,omitempty"`
	Entries         []Entry                `json:"entries,omitempty" bson:"entries,omitempty"`
	Packages        map[string]string      `json:"packages,omitempty" bson:"packages,omitempty"`
	SDK             map[string]interface{} `json:"sdk,omitempty" bson:"sdk,omitempty"`
	Contexts        map[string]interface{} `json:"contexts,omitempty" bson:"contexts,omitempty"`
	Context         map[string]interface{} `json:"context,omitempty" bson:"context,omitempty"`
	Release         Release                `json:"release,omitempty" bson:"release,omitempty"`
	GroupID         string                 `json:"groupID,omitempty" bson:"groupID,omitempty"`
}

// GetProjectEvent will fetch a event on a project
func (c *Client) GetProjectEvent(o Organization, p Project, eventID string) (Event, error) {
	var event Event
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/events/%s", *o.Slug, *p.Slug, eventID), &event, nil)
	return event, err
}

//GetLatestEvent will fetch the latest event for a issue
func (c *Client) GetLatestEvent(i Issue) (Event, error) {
	var event Event
	err := c.do("GET", fmt.Sprintf("issues/%s/events/latest", i.ID), &event, nil)
	return event, err
}

//GetOldestEvent will fetch the latest event for a issue
func (c *Client) GetOldestEvent(i Issue) (Event, error) {
	var event Event
	err := c.do("GET", fmt.Sprintf("issues/%s/events/oldest", i.ID), &event, nil)
	return event, err
}
