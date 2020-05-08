package taskManager

func checkChannelExistsInMap(channelID string) (exist bool) {
	exist = true
	for notifyChannel := range notifyChannelIDs {
		if notifyChannel == channelID {
			return
		}
	}

	exist = false
	return
}

func checkSubjectInCourse(course string, subject string) (isSubjectInCourse bool) {
	isSubjectInCourse = true
	for _, subject := range courseSubjects[course] {
		if course == subject {
			return
		}
	}

	isSubjectInCourse = false
	return
}

func checkSubjectIsDefine(subject string) (subjectIsDefine bool) {
	subjectIsDefine = true
	for _, tmpSubject := range availabilitySubjects {
		if subject == tmpSubject {
			return
		}
	}

	subjectIsDefine = false
	return
}

func searchCourseWithSubject(subject string) (findSubject string) {
	findSubject = ""
	for key, courseSubject := range courseSubjects {
		for _, s := range courseSubject {
			if subject == s {
				findSubject = key
				return
			}
		}
	}
	return
}
