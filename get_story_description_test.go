package storybranch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	storybranch "github.com/carpeliam/git-story-branch"
)

type MockGitRepository struct {
}

func (mockGitRepo MockGitRepository) GetBranchName() string {
	return "Insert Branch Name Here-#1234567890"
}

type MockPivotalTrackerReader struct {
}

func (mockTrackerReader MockPivotalTrackerReader) GetStory(storyID int) *storybranch.Story {
	if storyID == 1234567890 {
		return &storybranch.Story{
			ID: "1234567890",
			Description: "Description",
			State: "delivered",
		}
	}
	return nil
}

func (mockTrackerReader MockPivotalTrackerReader) GetStoryDescription(storyID int) string {
	if storyID == 1234567890 {
		return "Description"
	}
	return "Wrong Description"
}

func (mockTrackerReader MockPivotalTrackerReader) GetStoryState(storyID int) string {
	if storyID == 1234567890 {
		return "delivered"
	}
	return "unstarted"
}

var _ = Describe("Git Tracker name translator", func() {
	It("should retrieve a Pivotal Tracker Story based on the current git branch name", func() {
		mockGitRepo := MockGitRepository{}
		mockTrackerReader := MockPivotalTrackerReader{}

		story := storybranch.GetStory(mockGitRepo, mockTrackerReader)

		Expect(story.Description).To(Equal("Description"))
		Expect(story.State).To(Equal("delivered"))
	})
})
