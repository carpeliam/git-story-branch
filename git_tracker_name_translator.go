package storybranch

import (
	"regexp"
	"strconv"
)

// GetPivotalTrackerTaskID returns a string containing the task ID of the current branch name of the repository.
func GetPivotalTrackerTaskID(branchName string) int {
	re := regexp.MustCompile(`[1-9]+$`)
	taskIdString := re.FindString(branchName)
	taskId, _ := strconv.Atoi(taskIdString)
	return taskId
}