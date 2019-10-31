package sentry

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	// Resolved helps mark a issue or others as resolved
	Resolved Status = "resolved"
	// Unresolved helps mark a issue or others as unresolved
	Unresolved Status = "unresolved"
	// Ignored helps mark a issue or others as ignored
	Ignored Status = "ignored"
)

// Hash is returned via the issue_id/hashes
type Hash struct {
	ID string `json:"id,omitempty" bson:"id,omitempty"`
}

// Status is used to make consts for statuses
type Status string

// IssueStats is the stats of a issue
type IssueStats struct {
	TwentyFourHour *[]Stat `json:"24h,omitempty" bson:"24h,omitempty"`
	ThirtyDays     *[]Stat `json:"30d,omitempty" bson:"30d,omitempty"`
}

//IssueTagValue represents a tags value
type IssueTagValue struct {
	Count     *int64     `json:"count,omitempty" bson:"count,omitempty"`
	FirstSeen *time.Time `json:"firstSeen,omitempty" bson:"firstSeen,omitempty"`
	ID        *string    `json:"iD,omitempty" bson:"iD,omitempty"`
	Key       *string    `json:"key,omitempty" bson:"key,omitempty"`
	LastSeen  *time.Time `json:"lastSeen,omitempty" bson:"lastSeen,omitempty"`
	Name      *string    `json:"name,omitempty" bson:"name,omitempty"`
	Value     *string    `json:"value,omitempty" bson:"value,omitempty"`
}

// IssueTag is a tag used in a sentry issue
type IssueTag struct {
	UniqueValues int             `json:"uniqueValues,omitempty" bson:"uniqueValues,omitempty"`
	ID           string          `json:"id,omitempty" bson:"id,omitempty"`
	Key          string          `json:"key,omitempty" bson:"key,omitempty"`
	Name         string          `json:"name,omitempty" bson:"name,omitempty"`
	TopValues    []IssueTagValue `json:"topValues,omitempty" bson:"topValues,omitempty"`
}

//Avatar is used for a users avatar
type Avatar struct {
	AvatarType *string `json:"avatarType,omitempty" bson:"avatarType,omitempty"`
	AvatarUUID *string `json:"avatarUuid,omitempty" bson:"avatarUuid,omitempty"`
}

//InternalUser is a user on sentry and not a external customer
type InternalUser struct {
	AvatarURL  *string    `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	DateJoined *time.Time `json:"dateJoined,omitempty" bson:"dateJoined,omitempty"`
	Email      *string    `json:"email,omitempty" bson:"email,omitempty"`
	Has2FA     *bool      `json:"has2fa,omitempty" bson:"has2fa,omitempty"`
	ID         *string    `json:"iD,omitempty" bson:"id,omitempty"`
	IsActive   *bool      `json:"isActive,omitempty" bson:"isActive,omitempty"`
	IsManaged  *bool      `json:"isManaged,omitempty" bson:"isManaged,omitempty"`
	LastLogin  *time.Time `json:"lastLogin,omitempty" bson:"lastLogin,omitempty"`
	Name       *string    `json:"name,omitempty" bson:"name,omitempty"`
	Username   *string    `json:"username,omitempty" bson:"username,omitempty"`
}

//Activity is what current activity has happend on a issue
type Activity struct {
	Data        *map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"`
	DateCreated *time.Time              `json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
	ID          *string                 `json:"id,omitempty" bson:"id,omitempty"`
	Type        *string                 `json:"type,omitempty" bson:"type,omitempty"`
	User        *InternalUser           `json:"user,omitempty" bson:"user,omitempty"`
}

// Issue returns a issue found in sentry
type Issue struct {
	Annotations         []string               `json:"annotations,omitempty" bson:"annotations,omitempty"`
	AssignedTo          InternalUser           `json:"assignedTo,omitempty" bson:"assignedTo,omitempty"`
	Activity            []Activity             `json:"activity,omitempty" bson:"activity,omitempty"`
	Count               string                 `json:"count,omitempty" bson:"count,omitempty"`
	Culprit             string                 `json:"culprit,omitempty" bson:"culprit,omitempty"`
	FirstSeen           time.Time              `json:"firstSeen,omitempty" bson:"firstSeen,omitempty"`
	HasSeen             bool                   `json:"hasSeen,omitempty" bson:"hasSeen,omitempty"`
	ID                  string                 `json:"id,omitempty" bson:"_id,omitempty"`
	IsBookmarked        bool                   `json:"isBookmarked,omitempty" bson:"isBookmarked,omitempty"`
	IsPublic            bool                   `json:"isPublic,omitempty" bson:"isPublic,omitempty"`
	IsSubscribed        bool                   `json:"isSubscribed,omitempty" bson:"isSubscribed,omitempty"`
	LastSeen            time.Time              `json:"lastSeen,omitempty" bson:"lastSeen,omitempty"`
	Level               string                 `json:"level,omitempty" bson:"level,omitempty"`
	Logger              string                 `json:"logger,omitempty" bson:"logger,omitempty"`
	Metadata            map[string]string      `json:"metadata,omitempty" bson:"metadata,omitempty"`
	NumComments         int                    `json:"numComments,omitempty" bson:"numComments,omitempty"`
	Permalink           string                 `json:"permalink,omitempty" bson:"permalink,omitempty"`
	Project             Project                `json:"project,omitempty" bson:"project,omitempty"`
	ShareID             string                 `json:"shareId,omitempty" bson:"shareId,omitempty"`
	ShortID             string                 `json:"shortId,omitempty" bson:"shortId,omitempty"`
	Stats               IssueStats             `json:"stats,omitempty" bson:"stats,omitempty"`
	Status              Status                 `json:"status,omitempty" bson:"status,omitempty"`
	StatusDetails       map[string]interface{} `json:"statusDetails,omitempty" bson:"statusDetails,omitempty"`
	SubscriptionDetails map[string]string      `json:"subscriptionDetails,omitempty" bson:"subscriptionDetails,omitempty"`
	Tags                []IssueTag             `json:"tags,omitempty" bson:"tags,omitempty"`
	Title               string                 `json:"title,omitempty" bson:"title,omitempty"`
	Type                string                 `json:"type,omitempty" bson:"type,omitempty"`
	UserCount           int                    `json:"userCount,omitempty" bson:"userCount,omitempty"`
	UserReportCount     int                    `json:"userReportCount,omitempty" bson:"userReportCount,omitempty"`
	Events              []Event                `json:"_events,omitempty" bson:"events,omitempty"`
}

type issueQuery struct {
	StatsPeriod   *string
	ShortIDLookup *bool
	Query         *string
}

func (i *issueQuery) ToQueryString() string {
	query := url.Values{}
	if i.StatsPeriod != nil {
		query.Add("statsPeriod", *i.StatsPeriod)
	}
	if i.ShortIDLookup != nil {
		query.Add("shortIdLookup", strconv.FormatBool(*i.ShortIDLookup))
	}
	if i.Query != nil {
		query.Add("query", *i.Query)
	}

	return query.Encode()
}

//GetIssues will fetch all issues for organization and project
func (c *Client) GetIssues(o Organization, p Project, StatsPeriod *string, ShortIDLookup *bool, query *string) ([]Issue, *Link, error) {
	var issues []Issue

	issueFilter := &issueQuery{
		StatsPeriod:   StatsPeriod,
		ShortIDLookup: ShortIDLookup,
		Query:         query,
	}

	link, err := c.doWithPaginationQuery(
		"GET", fmt.Sprintf("projects/%s/%s/issues", *o.Slug, *p.Slug), &issues, nil, issueFilter)
	return issues, link, err
}

//GetIssue will fetch a issue by its ID as a string
func (c *Client) GetIssue(id string) (Issue, error) {
	var issue Issue
	err := c.do("GET", fmt.Sprintf("issues/%s", id), &issue, nil)
	return issue, err
}

//GetIssueHashes will fetch all hashes for a issue
func (c *Client) GetIssueHashes(i Issue) ([]Hash, *Link, error) {
	var hashes []Hash
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/hashes", i.ID), &hashes, nil)
	return hashes, link, err
}

//GetIssueTags will fetch all tags for a issue
func (c *Client) GetIssueTags(i Issue) ([]IssueTag, *Link, error) {
	var tags []IssueTag
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/tags", i.ID), &tags, nil)
	return tags, link, err
}

//GetIssueTag will fetch a tag used in a issue. Eg; environment, release, server
func (c *Client) GetIssueTag(i Issue, tagname string) (IssueTag, error) {
	var tag IssueTag
	err := c.do("GET", fmt.Sprintf("issues/%s/tags/%s", i.ID, tagname), &tag, nil)
	return tag, err
}

//GetIssueTagValues will fetch all values for a issues tag
func (c *Client) GetIssueTagValues(i Issue, tag IssueTag) ([]IssueTagValue, *Link, error) {
	var values []IssueTagValue
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/tags/%s/values", i.ID, tag.Key), &values, nil)
	return values, link, err
}

//GetIssueEvents will fetch all events for a issue
func (c *Client) GetIssueEvents(i Issue) ([]Event, *Link, error) {
	var events []Event
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/events", i.ID), &events, nil)
	return events, link, err
}

//GetIssueEvents will fetch all events for a issue
func (c *Client) GetIssueEventsFull(i Issue) ([]Event, *Link, error) {
	var events []Event
	link, err := c.doWithPagination("GET", fmt.Sprintf("issues/%s/events/?full=true", i.ID), &events, nil)
	return events, link, err
}

//UpdateIssue will update status, assign to, hasseen, isbookmarked and issubscribed
func (c *Client) UpdateIssue(i Issue) error {
	return c.do("PUT", fmt.Sprintf("issues/%s", i.ID), &i, &i)
}

//DeleteIssue will delete an issue
func (c *Client) DeleteIssue(i Issue) error {
	return c.do("DELETE", fmt.Sprintf("issues/%s", i.ID), nil, nil)
}
