package storybranch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type PivotalTrackerReader interface {
	GetStoryDescription(storyID int) string
}

type Tracker struct {
}

func (tracker Tracker) GetStoryDescriptionFromTracker(storyID int) string {
	return "HI"
}

var _ = Describe("Tracker", func() {

	It("should be able to look up the description of a story given the story ID", func() {
		// arrange
		// some kind of interface
		// maybe we need to mock out tracker to have a story id of 123456789
		gub := Tracker{}

		description := gub.GetStoryDescriptionFromTracker(123456789)

		// assert
		Expect(description).To(Equal("I dunno, uh, cool story... bro.. or something."))
	})
})
