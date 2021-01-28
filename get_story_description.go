package storybranch

import (
	"regexp"
	"strconv"
)

type Story struct {
	ID string
	Description string
	State string
}

func getPivotalTrackerTaskID(branchName string) int {
	re := regexp.MustCompile(`\d+$`)
	taskIDString := re.FindString(branchName)
	taskID, _ := strconv.Atoi(taskIDString)
	return taskID
}

// GetStoryDescription comment
func GetStory(repo Repository, tracker Tracker) *Story {
	currentBranchName := repo.GetBranchName()
	storyID := getPivotalTrackerTaskID(currentBranchName)
	
	return &Story{
		ID: "1234",
		Description: tracker.GetStoryDescription(storyID),
		State: "delivered",
	}
}