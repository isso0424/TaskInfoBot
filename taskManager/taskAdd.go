package taskManager

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func taskAdd(session *discordgo.Session, channelID string, messages []string) {
	var task string
	var limit time.Time
	var subject string
	if channelID != config.Channels.Regist {
		session.ChannelMessageSend(channelID, fmt.Sprintf("課題を登録する際は<#%s>で行ってください", config.Channels.Regist))
		return
	}
	switch len(messages) {
	case 2, 3:
		// 引数不足により失敗
		session.ChannelMessageSend(channelID, "引数が足りません\n最低でも2個は必要です")
		return

	case 4:
		// limitまで指定
		task = messages[2]
		subject = messages[3]
		t := time.Now()
		limit = time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, jst)

	case 5:
		// 全指定
		var err error
		task = messages[2]
		subject = messages[3]
		limit, err = strToLimit(messages[4])
		if err != nil {
			session.ChannelMessageSend(channelID, "日付の指定は n/m で指定してください")
			return
		}
	}
	if !checkSubjectIsDefine(subject) {
		session.ChannelMessageSend(channelID, "データの作成に失敗しました\n有効な教科の名前を指定してください")
		return
	}
	course := searchCourseWithSubject(subject)
	err := createTask(task, limit, subject, course)

	if err != nil {
		session.ChannelMessageSend(channelID, "データの作成に失敗しました\n課題の名前の重複などが無いか確認してください")
		return
	}

	message := fmt.Sprintf("```name: %s\nlimit: %d/%d\nsubject: %s```\nで新しい課題を作成しました。", task, int(limit.Month()), limit.Day(), subject)
	session.ChannelMessageSend(channelID, message)
}

func strToLimit(message string) (time.Time, error) {
	now := time.Now()
	nowYear := now.Year()
	dateStrings := strings.Split(message, "/")

	if len(dateStrings) < 1 {
		return time.Now(), errors.New("invalid patarn")
	}

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

	createdTime := time.Date(nowYear, time.Month(month), day, 23, 59, 59, 0, jst)

	if createdTime.Before(now) {
		createdTime = createdTime.Add(time.Duration(8760) * time.Hour)
	}
	return createdTime, nil
}

func createTask(task string, limitDate time.Time, subject string, course string) error {
	limit := fmt.Sprintf("%d-%d-%d", limitDate.Year(), int(limitDate.Month()), limitDate.Day())
	if checkTaskNameConflict(task) {
		return errors.New("NAME IS CONFLICTED")
	}
	_, err := db.Exec(`INSERT INTO TASKS("TASK","LIMIT","SUBJECT","COURSE") VALUES(?,?,?,?)`, task, limit, subject, course)
	if err != nil {
		return error(err)
	}
	return nil
}

func checkTaskNameConflict(task string) bool {
	rows, err := db.Query(`SELECT * FROM TASKS WHERE TASK=?`, task)
	defer rows.Close()
	if err != nil {
		return true
	}

	for rows.Next() {
		return true
	}
	return false
}
