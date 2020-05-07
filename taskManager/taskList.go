package taskManager

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func taskList(session *discordgo.Session, channelID string, messages []string) {
	var rows *sql.Rows
	var err error

	if !checkChannelExistsInMap(channelID) {
		session.ChannelMessageSend(channelID, "このチャンネルは課題確認用に設定されていません")
		return
	}

	course := notifyChannelIDs[channelID]

	if len(messages) < 3 {
		rows, err = db.Query(`SELECT * FROM TASKS WHERE COURSE=?`, course)
	} else {
		if checkSubjectInCourse(course, messages[2]) {
			session.ChannelMessageSend(channelID, "指定された教科は現在のチャンネルの系に存在しません")
		}
		rows, err = db.Query(`SELECT * FROM TASKS WHERE COURSE=? AND SUBJECT=?`, course, messages[2])
	}

	if err != nil {
		session.ChannelMessageSend(channelID, "値の取り出しでエラーが発生しました")
		return
	}
	defer rows.Close()

	sendMessages := createList(rows)

	for _, message := range sendMessages {
		session.ChannelMessageSend(channelID, message)
	}

	if len(sendMessages) == 0 {
		session.ChannelMessageSend(channelID, "このチャンネル向けに作成された課題はありません")
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

		sendMessages = append(sendMessages, fmt.Sprintf("```task: %s\nlimit: %s\nsubject: %s```", task, limit, subject))
	}
	return sendMessages
}
