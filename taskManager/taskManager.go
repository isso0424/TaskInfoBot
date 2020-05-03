package taskManager

import (
	"TaskInfoBot/loadConfig"
	"database/sql"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var jst *time.Location = time.FixedZone("Asia/Tokyo", 9*60*60)
var config loadConfig.Config
var isSetupped = false
var availabilitySubjects []string
var notifyChannelIDs = map[string]string{}
var courseSubjects = map[string][]string{}
var db *sql.DB

func TaskManager(session *discordgo.Session, event *discordgo.MessageCreate) {
	messages := strings.Split(event.Content, " ")
	command := messages[1]

	if !isSetupped {
		session.ChannelMessageSend(event.ChannelID, "授業一覧が登録されていません")
		return
	}

	channelID := event.ChannelID

	switch command {
	case "add":
		taskAdd(session, channelID, messages)
	case "list":
		taskList(session, channelID, messages)
	case "remove":
		taskDelete(session, channelID, messages)
	case "help":
		help(session, channelID, messages)
	}
}
