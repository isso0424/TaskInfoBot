package taskManager

import (
	"TaskInfoBot/loadConfig"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var jst *time.Location = time.FixedZone("Asia/Tokyo", 9*60*60)
var config loadConfig.Config
var configuration = false
var availabilitySubjects []string
var notifyChannelIDs = map[string]string{}
var courseSubjects = map[string][]string{}

func TaskManager(session *discordgo.Session, event *discordgo.MessageCreate, db *sql.DB) {
	messages := strings.Split(event.Content, " ")
	command := messages[1]

	if !configuration {
		session.ChannelMessageSend(event.ChannelID, "授業一覧が登録されていません")
		return
	}

	switch command {
	case "add":
		taskAdd(session, event, messages, db)
	case "list":
		taskList(session, event, messages, db)
	case "remove":
		deleteValue := messages[2]
		taskDelete(session, event, db, deleteValue)
	case "help":
		if len(messages) > 2 && messages[2] == "subject" {
			subjectHelp(session, event, messages)
		}
		help(session, event)
	}
}

func subjectHelp(session *discordgo.Session, event *discordgo.MessageCreate, messages []string) {
	sendMessageBase := ""
	for key, subjects := range courseSubjects {
		sendMessage := ""
		for index, subject := range subjects {
			if index == 0 {
				sendMessage = subject
				continue
			}
			sendMessage += fmt.Sprintf(", %s", subject)
		}
		sendMessageBase += fmt.Sprintf("%s\n```%s```\n", key, sendMessage)
	}
	session.ChannelMessageSend(event.ChannelID, sendMessageBase)
	return
}

func help(session *discordgo.Session, event *discordgo.MessageCreate) {
	helpMessage := "***課題管理BOT***\n```!task add <task> <limit> <subject>```\ntask: 課題名\nlimit: 締め切り(初期値=翌日)\nsubject: 教科(初期値='')\n教科は省略できる\n"
	helpMessage += "```!task list <subject>```\n課題一覧を表示します\n<subject>を指定すると教科ごとの絞り込みが可能です\n"
	helpMessage += "```!task remove <task>```\n課題を課題名から検索して削除します"
	helpMessage += "```!task help (subject)```\n使い方を表示します\nsubjectを付けると利用可能な教科を表示します"
	session.ChannelMessageSend(event.ChannelID, helpMessage)
}

func taskList(session *discordgo.Session, event *discordgo.MessageCreate, messages []string, db *sql.DB) {
	var rows *sql.Rows
	var err error
	var sendMessages []string

	if !checkChannelExistsInMap(event.ChannelID) {
		session.ChannelMessageSend(event.ChannelID, "このチャンネルは課題確認用に設定されていません")
		return
	}

	course := notifyChannelIDs[event.ChannelID]

	if len(messages) < 3 {
		rows, err = db.Query(`SELECT * FROM TASKS WHERE COURSE=?`, course)
	} else {
		if checkSubjectInCourse(course, messages[2]) {
			session.ChannelMessageSend(event.ChannelID, "指定された教科は別な系の教科です")
		}
		rows, err = db.Query(`SELECT * FROM TASKS WHERE COURSE=? OR SUBJECT=?`, course, messages[2])
	}

	if err != nil {
		session.ChannelMessageSend(event.ChannelID, "値の取り出しでエラーが発生しました")
		return
	}
	defer rows.Close()
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

	for _, message := range sendMessages {
		session.ChannelMessageSend(event.ChannelID, message)
	}

	if len(sendMessages) == 0 {
		session.ChannelMessageSend(event.ChannelID, "このチャンネル向けに作成された課題はありません")
	}
}

func taskDelete(session *discordgo.Session, event *discordgo.MessageCreate, db *sql.DB, deleteValue string) {
	_, err := db.Exec(`DELETE FROM TASKS WHERE TASK=?`, deleteValue)
	if err != nil {
		session.ChannelMessageSend(event.ChannelID, "指定された名前の課題は存在しません")
		return
	}
	session.ChannelMessageSend(event.ChannelID, fmt.Sprintf("%sを削除しました", deleteValue))
}

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
