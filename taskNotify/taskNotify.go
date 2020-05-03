package taskNotify

import (
	"TaskInfoBot/loadConfig"
	"TaskInfoBot/taskManager"
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func TaskNotify(session *discordgo.Session, db *sql.DB, config loadConfig.Config) {
	for notifyChannel, course := range taskManager.SetNotifyChannnlIDs(config.Channels.Notify) {
		session.ChannelMessageSend(notifyChannel, "***課題お知らせTIME***")
		getTaskWithLimit(session, db, notifyChannel, course, "today")
		getTaskWithLimit(session, db, notifyChannel, course, "tomorrow")
	}
}

func getDate(date time.Time) string {
	return fmt.Sprintf("%d-%d-%d", date.Year(), int(date.Month()), date.Day())
}

func getTaskWithLimit(session *discordgo.Session, db *sql.DB, notifyChannel string, course string, targetDay string) {
	var date string
	var day string

	switch targetDay {
	case "today":
		date = getDate(time.Now())
		day = "今日"
	case "tomorrow":
		date = getDate(time.Now().Add(time.Duration(24) * time.Hour))
		day = "明日"
	}

	rows, err := db.Query(`SELECT * FROM TASKS WHERE "LIMIT"=? AND "COURSE"=?`, date, course)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

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

		sendMessages = append(sendMessages, fmt.Sprintf("```name: %s\nsubject: %s```", task, subject))
	}
	if len(sendMessages) == 0 {
		session.ChannelMessageSend(notifyChannel, fmt.Sprintf("%s提出期限の課題はありません", day))
		return
	}
	session.ChannelMessageSend(notifyChannel, fmt.Sprintf("%s提出期限の課題は以下のとおりです", day))

	for _, message := range sendMessages {
		session.ChannelMessageSend(notifyChannel, message)
	}

}
