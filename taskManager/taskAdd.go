package taskManager

import (
	"errors"
	"fmt"
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
	datesMap, err := checkDatePatarn(message)

	if err != nil {
		return time.Now(), nil
	}

	month := datesMap["month"]
	day := datesMap["day"]

	createdTime := time.Date(nowYear, time.Month(month), day, 23, 59, 59, 0, jst)

	if createdTime.Before(now) {
		createdTime = createdTime.Add(time.Duration(8760) * time.Hour)
	}
	return createdTime, nil
}

func createTask(task string, limitDate time.Time, subject string, course string) error {
	limit := fmt.Sprintf("%d-%d-%d", limitDate.Year(), int(limitDate.Month()), limitDate.Day())
	if checkTaskNameConflict(task) {
		return errors.New("NAME IS DUPLICATED")
	}
	_, err := db.Exec(`INSERT INTO TASKS("TASK","LIMIT","SUBJECT","COURSE") VALUES(?,?,?,?)`, task, limit, subject, course)
	if err != nil {
		return error(err)
	}
	return nil
}
