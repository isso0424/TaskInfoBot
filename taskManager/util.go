package taskManager

import "fmt"

func checkChannelExistsInMap(channelID string) bool {
	for notifyChannel, _ := range notifyChannelIDs {
		if notifyChannel == channelID {
			return true
		}
	}
	return false
}

func checkSubjectInCourse(course string, subject string) bool {
	for _, subject := range courseSubjects[course] {
		if course == subject {
			return true
		}
	}
	return false
}

func checkSubjectIsDefine(subject string) bool {
	for _, tmpSubject := range availabilitySubjects {
		if subject == tmpSubject {
			return true
		}
	}
	return false
}

func searchCourseWithSubject(subject string) string {
	for key, courseSubject := range courseSubjects {
		fmt.Printf(key)
		for _, s := range courseSubject {
			if subject == s {
				return key
			}
		}
	}

	return ""
}
