package storybranch

import (
	"regexp"
	"strconv"
)

// GetPivotalTrackerTaskID returns a string containing the task ID of the current branch name of the repository.
func GetPivotalTrackerTaskID(branchName string) int {
	re := regexp.MustCompile(`\d+$`)
	taskIDString := re.FindString(branchName)
	taskID, _ := strconv.Atoi(taskIDString)
	return taskID
}

// GetStoryDescription comment
func GetStoryDescription(gitRepo GitRepository, trackerReader PivotalTrackerReader) string {
	currentBranchName := gitRepo.GetBranchName()
	storyID := GetPivotalTrackerTaskID(currentBranchName)
	return trackerReader.GetStoryDescription(storyID)
}
