package taskManager

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var jst *time.Location = time.FixedZone("Asia/Tokyo", 9*60*60)

func TaskManager(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	if !strings.HasPrefix(event.Content, "!task") {
		return
	}

	if len(event.Content) < 7 {
		return
	}

	messages := strings.Split(event.Content, " ")
	command := messages[1]

	switch command {
	case "add":
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
				fmt.Println(err)
				return
			}

		case 5:
			// 全指定
			var err error
			task = messages[2]
			limit, err = strToLimit(messages[3])
			if err != nil {
				fmt.Println(err)
				return
			}
			subject = messages[4]
		}
		createTask(task, limit, subject)
	}
}

func strToLimit(message string) (time.Time, error) {
	now := time.Now()
	nowYear := now.Year()
	dateStrings := strings.Split(message, "/")

	rawMonth := dateStrings[0]
	rawDay := dateStrings[1]

	month, err := strconv.Atoi(rawMonth)
	if err != nil {
		return time.Now(), errors.New("mouth cannot convert to int")
	}

	day, err := strconv.Atoi(rawDay)
	if err != nil {
		return time.Now(), errors.New("day cannot convert to int")
	}

	createdTime := time.Date(nowYear, time.Month(month), day, 0, 0, 0, 0, jst)

	if createdTime.Before(now) {
		createdTime = createdTime.Add(time.Duration(8760) * time.Hour)
	}
	return createdTime, nil
}

func createTask(task string, limit time.Time, subject string) {

}
