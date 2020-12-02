package storybranch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	storybranch "github.com/carpeliam/git-story-branch"
)

type MockGitRepository struct {
}

func (mockGitRepo MockGitRepository) GetBranchName() string {
	return "Insert Branch Name Here"
}

var _ = Describe("Git Tracker name translator", func() {
	It("should derive the Pivotal Tracker task ID from the given branch", func() {
		pivotalTrackerTaskID := storybranch.GetPivotalTrackerTaskID("current-branch-123456789")
		Expect(pivotalTrackerTaskID).To(Equal(123456789))
	})

	It("should derive the Pivotal Tracker task ID from the given branch even with the octothorpe", func() {
		pivotalTrackerTaskID := storybranch.GetPivotalTrackerTaskID("current-branch-#123456789")
		Expect(pivotalTrackerTaskID).To(Equal(123456789))
	})

	It("should retrieve a Pivotal Tracker Story based on the current git branch name", func() {
		mockGitRepo := MockGitRepository{}
		// Some execution
		storyDescription := storybranch.GetStoryDescription(mockGitRepo)
		// Some expectation
		Expect(storyDescription).To(Equal("Description"))
	})
})
