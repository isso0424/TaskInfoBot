package taskNotify

import (
	"TaskInfoBot/loadConfig"
	"TaskInfoBot/taskManager"
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var db *sql.DB

// TaskNotify is a function that sends task regularly notications
func TaskNotify(session *discordgo.Session, config loadConfig.Config) {
	fmt.Println("start notify")
	deleteDeadlinePassedTask(time.Now().Add(time.Duration(-24) * time.Hour))
	for notifyChannel, course := range taskManager.SetNotifyChannnlIDs(config.Channels.Notify) {
		notifyMessages := createNotify(session, notifyChannel, course)
		if len(notifyMessages) == 3 {
			continue
		}

		for _, notify := range notifyMessages {
			session.ChannelMessageSend(notifyChannel, notify)
		}
	}
	fmt.Println("finish notify")
}

func createNotify(session *discordgo.Session, notifyChannel string, course string) []string {
	notifyDay := []string{"today", "tomorrow"}
	notifyMessages := []string{"***課題お知らせTIME***"}
	for _, day := range notifyDay {
		tasks := getTaskWithLimit(course, day)

		if len(tasks) == 0 {
			tasks = []string{fmt.Sprintf("%s提出期限の課題はありません", getDay(day))}
		} else {
			tasks = insertToHead(tasks, fmt.Sprintf("%s提出期限の課題は以下のとおりです", getDay(day)))
		}

		notifyMessages = append(
			notifyMessages,
			tasks...,
		)
	}
	return notifyMessages
}

func getDay(day string) string {
	switch day {
	case "today":
		return "今日"
	case "tomorrow":
		return "明日"
	default:
		return ""
	}
}

func insertToHead(slice []string, insertValue string) []string {
	slice = append(slice[:1], slice[0:]...)
	slice[0] = insertValue
	return slice
}

func getDate(date time.Time) string {
	return fmt.Sprintf("%d-%d-%d", date.Year(), int(date.Month()), date.Day())
}

func getTaskWithLimit(course string, targetDay string) []string {
	var date string

	switch targetDay {
	case "today":
		date = getDate(time.Now())
	case "tomorrow":
		date = getDate(time.Now().Add(time.Duration(24) * time.Hour))
	}

	rows, err := db.Query(`SELECT * FROM TASKS WHERE "LIMIT"=? AND "COURSE"=?`, date, course)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return []string{"Error with database"}
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

	return sendMessages
}

func deleteDeadlinePassedTask(date time.Time) {
	dateString := getDate(date)
	_, err := db.Exec(`DELETE FROM TASKS WHERE "LIMIT"=?`, dateString)
	if err != nil {
		fmt.Println(err)
	}
}

// SetDB is a function that set database to this package
func SetDB(givenDB *sql.DB) {
	db = givenDB
}
