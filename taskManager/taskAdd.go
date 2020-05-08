package taskManager

import (
	"TaskInfoBot/messageController"
	"errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var addSuccess = messageController.CreateTaskAddMessage()
var addError = messageController.CreateAddErrorMessage()

func taskAdd(session *discordgo.Session, channelID string, messages []string) {
	var task string
	var limit time.Time
	var subject string
	if channelID != config.Channels.Regist {
		session.ChannelMessageSend(channelID, fmt.Sprintf(addError.InvalidChannel, config.Channels.Regist))
		return
	}
	switch len(messages) {
	case 2, 3:
		// 引数不足により失敗
		session.ChannelMessageSend(channelID, addError.NotEnoughArgs)
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
			session.ChannelMessageSend(channelID, addError.InvalidDatePatarn)
			return
		}
	}
	if !checkSubjectIsDefine(subject) {
		session.ChannelMessageSend(channelID, addError.InvalidSubjectName)
		return
	}
	course := searchCourseWithSubject(subject)
	err := createTask(task, limit, subject, course)

	if err != nil {
		session.ChannelMessageSend(channelID, addError.DuplicateName)
		return
	}

	message := fmt.Sprintf(addSuccess, task, int(limit.Month()), limit.Day(), subject)
	session.ChannelMessageSend(channelID, message)
}

func strToLimit(message string) (createdTime time.Time, err error) {
	now := time.Now()
	nowYear := now.Year()
	datesMap, err := checkDatePatarn(message)

	if err != nil {
		return time.Now(), nil
	}

	month := datesMap["month"]
	day := datesMap["day"]

	createdTime = time.Date(nowYear, time.Month(month), day, 23, 59, 59, 0, jst)

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
