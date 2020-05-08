package taskManager

import (
	"TaskInfoBot/messageController"
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var listTemplates = messageController.CreateTaskListMessageTemplate()
var listError = messageController.CreateListErrorMessage()

func taskList(session *discordgo.Session, channelID string, messages []string) {
	var rows *sql.Rows
	var err error

	if !checkChannelExistsInMap(channelID) {
		session.ChannelMessageSend(channelID, listError.NotCheckTaskChannel)
		return
	}

	course := notifyChannelIDs[channelID]

	if len(messages) < 3 {
		rows, err = db.Query(`SELECT * FROM TASKS WHERE COURSE=?`, course)
	} else {
		if checkSubjectInCourse(course, messages[2]) {
			session.ChannelMessageSend(channelID, listError.SubjectIsNotInCourse)
		}
		rows, err = db.Query(`SELECT * FROM TASKS WHERE COURSE=? AND SUBJECT=?`, course, messages[2])
	}

	if err != nil {
		session.ChannelMessageSend(channelID, listError.GetValueError)
		return
	}
	defer rows.Close()

	sendMessages := createList(rows)

	for _, message := range sendMessages {
		session.ChannelMessageSend(channelID, message)
	}

	if len(sendMessages) == 0 {
		session.ChannelMessageSend(channelID, listError.TaskNotFound)
	}
}

func createList(rows *sql.Rows) []string {
	var sendMessages = []string{}
	for rows.Next() {
		var id int
		var task string
		var limit string
		var subject string
		var course string

		if err := rows.Scan(&id, &task, &limit, &subject, &course); err != nil {
			fmt.Println(err)
			continue
		}

		sendMessages = append(sendMessages, fmt.Sprintf(listTemplates, task, limit, subject))
	}
	return sendMessages
}
