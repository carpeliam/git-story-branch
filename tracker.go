package storybranch

import "net/http"

// GetStoryDescriptionFromTracker returns the story description
func GetStoryDescriptionFromTracker(storyID int) string {
	return ""
}

type PivotalTrackerClient interface {
	Do(req *http.Request, v interface{}) (*http.Response, error)
	NewRequest(method, urlPath string, body interface{}) (*http.Request, error)
	SetBaseURL(baseURL string) error
	SetUserAgent(agent string)
}

type PivotalTrackerReader interface {
	GetStoryDescription(storyID int) string
}

type RealPivotalTrackerReader struct {
}

func (realPivotalTrackerReader RealPivotalTrackerReader) GetStoryDescription(storyID int) string {
	rptrDeleteMe := RealPivotalTrackerReader{&pivotal}
	return "Do this later"
}
