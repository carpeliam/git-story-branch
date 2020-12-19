package storybranch_test

import (
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

type PivotalTrackerStoryService interface {
	GetByID(storyID int) (*pivotal.Story, *http.Response, error)
}

type MockPivotalTrackerStoryService struct {
}

func (mockPivotalTrackerStoryService MockPivotalTrackerStoryService) GetByID(storyID int) (*pivotal.Story, *http.Response, error) {
	if storyID == 123456789 {
		newStory := pivotal.Story{
			ID:            123456789,
			ProjectID:     0,
			Name:          "My cool story",
			Description:   "I dunno, uh, cool story... bro.. or something.",
			Type:          "Type",
			State:         "State",
			Estimate:      nil,
			AcceptedAt:    nil,
			Deadline:      nil,
			RequestedByID: 0,
			OwnerIDs:      []int{},
			LabelIDs:      []int{},
			Labels:        []*pivotal.Label{},
			TaskIDs:       []int{},
			Tasks:         []int{},
			FollowerIDs:   []int{},
			CommentIDs:    []int{},
			CreatedAt:     nil,
			UpdatedAt:     nil,
			BeforeID:      0,
			AfterID:       0,
			IntegrationID: 0,
			ExternalID:    "ExternalID",
			URL:           "URL",
		}
		return &newStory, nil, nil
	}

	return nil, nil, errors.New("Story does not exist.")
}

type Tracker struct {
	storyService *PivotalTrackerStoryService
}

func (tracker Tracker) GetStoryDescriptionFromTracker(storyID int) string {
	// Tracker is going to ask the Story Service to get the description
	//GetByID(storyID int) (*pivotal.Story, *http.Response, error)
	storyServicePointer := tracker.storyService
	pivotalTrackerStory, _, _ := (*storyServicePointer).GetByID(storyID)
	return pivotalTrackerStory.Description
}

var _ = Describe("Tracker", func() {

	It("should be able to look up the description of a story given the story ID", func() {
		storyService := &MockPivotalTrackerStoryService{}
		// arrange
		// TODO Why can't I pass in storyService directly?
		var i PivotalTrackerStoryService
		i = storyService
		gub := Tracker{&i}
		description := gub.GetStoryDescriptionFromTracker(123456789)

		// assert
		Expect(description).To(Equal("I dunno, uh, cool story... bro.. or something."))
	})
})
