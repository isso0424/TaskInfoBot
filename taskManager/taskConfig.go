package taskManager

import (
	"TaskInfoBot/loadConfig"
	"fmt"
)

func SetConfig(assignConfig loadConfig.Config) {
	config = assignConfig
	for _, course := range config.Courses {
		for _, majorSubjects := range course.Subjects.Major {
			availabilitySubjects = append(availabilitySubjects, majorSubjects)
		}

		for _, minorSubjects := range course.Subjects.Minor {
			availabilitySubjects = append(availabilitySubjects, minorSubjects)
		}
	}

	setNotifyChannnlIDs(assignConfig.Channels.Notify)
	setCourseSubjects(assignConfig.Courses)
	configuration = true
}

func setNotifyChannnlIDs(notifyChannels loadConfig.Notify) {
	major := notifyChannels.Major
	notifyChannelIDs[major.General] = "MajorGeneral"
	notifyChannelIDs[major.M] = "MajorM"
	notifyChannelIDs[major.E] = "MajorE"
	notifyChannelIDs[major.I] = "MajorI"
	notifyChannelIDs[major.C] = "MajorC"

	minor := notifyChannels.Minor
	notifyChannelIDs[minor.M] = "MinorM"
	notifyChannelIDs[minor.E] = "MinorE"
	notifyChannelIDs[minor.I] = "MinorI"
	notifyChannelIDs[minor.C] = "MinorC"
	notifyChannelIDs[minor.G] = "MinorG"
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
