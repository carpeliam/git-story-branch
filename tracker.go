package storybranch

// GetStoryDescriptionFromTracker returns the story description
func GetStoryDescriptionFromTracker(storyID int) string {
	return ""
}

type PivotalTrackerReader interface {
	GetStoryDescription(storyID int) string
}
