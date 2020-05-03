package taskManager

import (
	"TaskInfoBot/loadConfig"
	"database/sql"
	"fmt"
)

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

func SetNotifyChannnlIDs(notifyChannels loadConfig.Notify) map[string]string {
	var notifys = map[string]string{}
	major := notifyChannels.Major
	notifys[major.General] = "MajorGeneral"
	notifys[major.M] = "MajorM"
	notifys[major.E] = "MajorE"
	notifys[major.I] = "MajorI"
	notifys[major.C] = "MajorC"

	minor := notifyChannels.Minor
	notifys[minor.M] = "MinorM"
	notifys[minor.E] = "MinorE"
	notifys[minor.I] = "MinorI"
	notifys[minor.C] = "MinorC"
	notifys[minor.G] = "MinorG"
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
