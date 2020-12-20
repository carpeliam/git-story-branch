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
	StoryService PivotalTrackerStoryService
}

// GetStoryDescription comment
func (tracker Tracker) GetStoryDescription(storyID int) string {
	storyService := tracker.StoryService
	pivotalTrackerStory, _, _ := storyService.GetByID(storyID)
	return pivotalTrackerStory.Description
}
