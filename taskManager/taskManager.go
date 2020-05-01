package taskManager

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var jst *time.Location = time.FixedZone("Asia/Tokyo", 9*60*60)

func TaskManager(session *discordgo.Session, event *discordgo.MessageCreate, db *sql.DB) {
	messages := strings.Split(event.Content, " ")
	command := messages[1]

	switch command {
	case "add":
		taskAdd(session, event, messages, db)
	case "list":
		taskList(session, event, messages, db)
	case "remove":
		deleteValue := messages[2]
		taskDelete(session, event, db, deleteValue)
	case "help":
		help(session, event)
	}
}

func help(session *discordgo.Session, event *discordgo.MessageCreate) {
	helpMessage := "***課題管理BOT***\n```!task add <task> <limit> <subject>```\ntask: 課題名\nlimit: 締め切り(初期値=翌日)\nsubject: 教科(初期値='')\nこれは後方を削って使用することが可能です\n"
	helpMessage += "```!task list <subject>```\n課題一覧を表示します\n<subject>を指定すると教科ごとの絞り込みが可能です\n"
	helpMessage += "```!task remove <task>```\n課題を課題名から検索して削除します"
	session.ChannelMessageSend(event.ChannelID, helpMessage)
}

func taskList(session *discordgo.Session, event *discordgo.MessageCreate, messages []string, db *sql.DB) {
	var rows *sql.Rows
	var err error

	if len(messages) < 3 {
		rows, err = db.Query(`SELECT * FROM TASKS`)
	} else {
		rows, err = db.Query(`SELECT * FROM TASKS WHERE SUBJECT=?`, messages[2])
	}
	defer rows.Close()
	if err != nil {
		session.ChannelMessageSend(event.ChannelID, "値の取り出しでエラーが発生しました")
		return
	}
	for rows.Next() {
		var id int
		var task string
		var limit string
		var subject string

		if err := rows.Scan(&id, &task, &limit, &subject); err != nil {
			fmt.Println(err)
			continue
		}

		session.ChannelMessageSend(event.ChannelID, fmt.Sprintf("```task: %s\nlimit: %s\nsubject: %s```", task, limit, subject))
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
