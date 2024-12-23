package projects

import (
	"time"
)

// ----------------------------------------------------------------------------
// Project Structures
// ----------------------------------------------------------------------------

// Project represents a project in the database.
type Project struct {
	ProjectID    int              `json:"project_id"`
	OwnerID      int              `json:"owner_id"`
	Title        string           `json:"title"`
	Introduction string           `json:"introduction"`
	Sections     []ProjectSection `json:"sections"`
	Description  string           `json:"description"`
	Status       *string          `json:"status"`
	CreatedAt    time.Time        `json:"created_at"`
	Tags         []string         `json:"tags"`
}

type T struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Introduction string `json:"introduction"`
	OwnerId      string `json:"owner_id"`
	Tags         string `json:"tags"`
}

// Feedback represents feedback for a project.
type Feedback struct {
	FeedbackID int    `json:"feedback_id"`
	ProjectID  int    `json:"project_id"`
	UserID     int    `json:"user_id"`
	Feedback   string `json:"feedback"`
}

// ProjectSection represents an update for a project.
type ProjectSection struct {
	Title         string `json:"title"`
	SectionNumber int    `json:"section_number"`
	Body          string `json:"body"`
	ProjectID     int    `json:"project_id"`
}

type ProjectTag struct {
	ProjectID      int    `json:"project_id"`
	TagDescription string `json:"tag_description"`
}

// Event represents an event in the database.
type Event struct {
	EventID int    `json:"event_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

// UserEventParticipation represents the participation of a user in an event.
type UserEventParticipation struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}
