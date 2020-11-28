package storybranch

import (
	"regexp"
	"strconv"
)

// GetPivotalTrackerTaskID returns a string containing the task ID of the current branch name of the repository.
func GetPivotalTrackerTaskID(branchName string) int {
	re := regexp.MustCompile(`[1-9]+$`)
	taskIDString := re.FindString(branchName)
	taskID, _ := strconv.Atoi(taskIDString)
	return taskID
}
