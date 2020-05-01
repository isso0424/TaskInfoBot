package taskNotify

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var notifyChannel = "574884574778359844"

func TaskNotify(session *discordgo.Session, db *sql.DB) {
	session.ChannelMessageSend(notifyChannel, "***課題お知らせTIME***")
	now := getDate(time.Now())
	session.ChannelMessageSend(notifyChannel, "今日提出期限の課題は以下のとおりです")
	getTaskWithLimit(now, session, db)
	tomorrow := getDate(time.Now().Add(time.Duration(24) * time.Hour))
	session.ChannelMessageSend(notifyChannel, "明日提出期限の課題は以下のとおりです")
	getTaskWithLimit(tomorrow, session, db)
}

func getDate(date time.Time) string {
	return fmt.Sprintf("%d-%d-%d", date.Year(), int(date.Month()), date.Day())
}

func getTaskWithLimit(limit string, session *discordgo.Session, db *sql.DB) {
	rows, err := db.Query(`SELECT * FROM TASKS WHERE "LIMIT"=?`, limit)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
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

		session.ChannelMessageSend(notifyChannel, fmt.Sprintf("```name: %s\nsubject: %s```", task, subject))
	}
}
