package github

type EventType string

const (
	PushEvent                     EventType = "PushEvent"
	PullRequestEvent              EventType = "PullRequestEvent"
	IssueEvent                    EventType = "IssueEvent"
	ReleaseEvent                  EventType = "ReleaseEvent"
	CommitCommentEvent            EventType = "CommitCommentEvent"
	CreateEvent                   EventType = "CreateEvent"
	DeleteEvent                   EventType = "DeleteEvent"
	ForkEvent                     EventType = "ForkEvent"
	GollumEvent                   EventType = "GollumEvent"
	IssueCommentEvent             EventType = "IssueCommentEvent"
	IssuesEvent                   EventType = "IssuesEvent"
	MemberEvent                   EventType = "MemberEvent"
	PublicEvent                   EventType = "PublicEvent"
	PullRequestReviewEvent        EventType = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent EventType = "PullRequestReviewCommentEvent"
	PullRequestReviewThreadEvent  EventType = "PullRequestReviewThreadEvent"
	SponsorshipEvent              EventType = "SponsorshipEvent"
	WatchEvent                    EventType = "WatchEvent"
)

type UserActivity struct {
	ID    string    `json:"id"`
	Type  EventType `json:"type"`
	Actor struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		AvatarURL    string `json:"avatar_url"`
		DisplayLogin string `json:"display_login"`
		URL          string `json:"url"`
		GravatarID   string `json:"gravatar_id"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		RepositoryID int    `json:"repository_id"`
		PushID       int    `json:"push_id"`
		Size         int    `json:"size"`
		DistinctSize int    `json:"distinct_size"`
		Ref          string `json:"ref"`
		Head         string `json:"head"`
		Before       string `json:"before"`
		Commits      []struct {
			SHA    string `json:"sha"`
			Author struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
			Message  string `json:"message"`
			Distinct bool   `json:"distinct"`
			URL      string `json:"url"`
		} `json:"commits"`
	} `json:"payload"`
	Public    bool   `json:"public"`
	CreatedAt string `json:"created_at"`
}

type errMessage struct {
	Message          string
	DocumentationURL string
	Status           string
}

func (e errMessage) Error() string {
	return e.Message
}
