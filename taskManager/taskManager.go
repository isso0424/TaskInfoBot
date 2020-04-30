package taskManager

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
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
		taskList(session, event, db)
	case "remove":
		deleteValue := messages[2]
		taskDelete(session, event, db, deleteValue)
	}
}

func taskAdd(session *discordgo.Session, event *discordgo.MessageCreate, messages []string, db *sql.DB) {
	var task string
	var limit time.Time
	var subject string
	switch len(messages) {
	case 2:
		// 引数不足により失敗
		return
	case 3:
		// taskだけ指定
		task = messages[2]
		t := time.Now()
		limit = time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, jst)
		subject = ""

	case 4:
		// limitまで指定
		var err error
		task = messages[2]
		limit, err = strToLimit(messages[3])
		if err != nil {
			session.ChannelMessageSend(event.ChannelID, "日付の指定は n/m でまともな日付の範囲で指定してください")
			return
		}

	case 5:
		// 全指定
		var err error
		task = messages[2]
		limit, err = strToLimit(messages[3])
		if err != nil {
			session.ChannelMessageSend(event.ChannelID, "日付の指定は n/m で指定してください")
			return
		}
		subject = messages[4]
	}
	err := createTask(task, limit, subject, db)

	if err != nil {
		session.ChannelMessageSend(event.ChannelID, "データの作成に失敗しました\n課題の名前の重複などが無いか確認してください")
	}

	message := fmt.Sprintf("```name: %s\nlimit: %d/%d\nsubject: %s```\nで新しい課題を作成しました。", task, int(limit.Month()), limit.Day(), subject)
	session.ChannelMessageSend(event.ChannelID, message)
}

func taskList(session *discordgo.Session, event *discordgo.MessageCreate, db *sql.DB) {
	rows, err := db.Query(`SELECT * FROM TASKS`)
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

func strToLimit(message string) (time.Time, error) {
	now := time.Now()
	nowYear := now.Year()
	dateStrings := strings.Split(message, "/")

	rawMonth := dateStrings[0]
	rawDay := dateStrings[1]

	month, err := strconv.Atoi(rawMonth)
	if err != nil || month < 1 || month > 12 {
		return time.Now(), errors.New("mouth cannot convert to int")
	}

	day, err := strconv.Atoi(rawDay)
	if err != nil || day < 1 || day > 31 {
		return time.Now(), errors.New("day cannot convert to int")
	}

	createdTime := time.Date(nowYear, time.Month(month), day, 0, 0, 0, 0, jst)

	if createdTime.Before(now) {
		createdTime = createdTime.Add(time.Duration(8760) * time.Hour)
	}
	return createdTime, nil
}

func createTask(task string, limit time.Time, subject string, db *sql.DB) error {
	date := fmt.Sprintf("%d-%d-%d", limit.Year(), int(limit.Month()), limit.Day())
	err := insertToDB(task, date, subject, db)

	if err != nil {
		return error(err)
	}
	return nil
}

func insertToDB(task string, limit string, subject string, db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO TASKS("TASK","LIMIT","SUBJECT") VALUES(?,?,?)`, task, limit, subject)
	if err != nil {
		return error(err)
	}
	return nil
}
