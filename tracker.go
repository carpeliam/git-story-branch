package storybranch

import (
	"net/http"

	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

// PivotalTrackerReader comment
type PivotalTrackerReader interface {
	GetStoryDescription(storyID int) string
}

// PivotalTrackerStoryService comment
type PivotalTrackerStoryService interface {
	GetByID(storyID int) (*pivotal.Story, *http.Response, error)
}

// Tracker comment
type Tracker struct {
	storyService PivotalTrackerStoryService
}

// GetStoryDescription comment
func (tracker Tracker) GetStoryDescription(storyID int) string {
	storyService := tracker.storyService
	story, _, _ := storyService.GetByID(storyID)
	return story.Description
}

// NewTracker returns a new tracker
func NewTracker(storyService PivotalTrackerStoryService) *Tracker {
	return &Tracker{storyService}
}
