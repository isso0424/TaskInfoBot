package taskManager

import (
	"TaskInfoBot/loadConfig"
	"database/sql"
	"fmt"
)

// SetConfig is a function that sets taskManager Config
func SetConfig(assignConfig loadConfig.Config, givenDB *sql.DB) {
	config = assignConfig
	for _, course := range config.Courses {
		for _, majorSubjects := range course.Subjects.Major {
			availabilitySubjects = append(availabilitySubjects, majorSubjects)
		}

		for _, minorSubjects := range course.Subjects.Minor {
			availabilitySubjects = append(availabilitySubjects, minorSubjects)
		}
	}

	notifyChannelIDs = SetNotifyChannnlIDs(assignConfig.Channels.Notify)
	setCourseSubjects(assignConfig.Courses)
	db = givenDB
	isSetupped = true
}

// SetNotifyChannnlIDs is a function that creates notify channels map
// Key is channelID
// Value is course name
func SetNotifyChannnlIDs(notifyChannels loadConfig.Notify) (notifys map[string]string) {
	major := notifyChannels.Major
	notifys = map[string]string{}
	for course, channelID := range major {
		notifys[channelID] = fmt.Sprintf("Major%s", course)
	}

	minor := notifyChannels.Minor
	for course, channelID := range minor {
		notifys[channelID] = fmt.Sprintf("Minor%s", course)
	}

	return notifys
}

func setCourseSubjects(courses []loadConfig.Course) {
	for _, course := range courses {
		if len(course.Subjects.Major) != 0 {
			courseSubjects[fmt.Sprintf("Major%s", course.Alias)] = course.Subjects.Major
		}

		if len(course.Subjects.Minor) != 0 {
			courseSubjects[fmt.Sprintf("Minor%s", course.Alias)] = course.Subjects.Minor
		}
	}
}
