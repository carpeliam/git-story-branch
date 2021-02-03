package storybranch

import (
	"net/http"
	"strconv"

	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

// Tracker comment
type Tracker interface {
	GetStoryDescription(storyID int) string
	GetStoryState(storyID int) string
	GetStory(storyID int) *Story
}

// StoryService comment
type StoryService interface {
	GetByID(storyID int) (*pivotal.Story, *http.Response, error)
}

// PivotalTracker comment
type PivotalTracker struct {
	storyService StoryService
}

// GetStoryDescription comment
func (tracker PivotalTracker) GetStoryDescription(storyID int) string {
	storyService := tracker.storyService
	story, _, _ := storyService.GetByID(storyID)
	return story.Description
}

func (tracker PivotalTracker) GetStoryState(storyID int) string {
	storyService := tracker.storyService
	story, _, _ := storyService.GetByID(storyID)
	return story.State
}

func (tracker PivotalTracker) GetStory(storyID int) *Story {
	storyService := tracker.storyService
	story, _, _ := storyService.GetByID(storyID)
	
	return &Story{
		ID: strconv.Itoa(story.ID),
		Description: story.Description,
		State: story.State,
	}
}

// NewPivotalTracker returns a new tracker
func NewPivotalTracker(storyService StoryService) *PivotalTracker {
	return &PivotalTracker{storyService}
}
