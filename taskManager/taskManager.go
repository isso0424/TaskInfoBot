package taskManager

import (
	"TaskInfoBot/loadConfig"
	"TaskInfoBot/messageController"
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
var generalError = messageController.CreateGeneralErrorMessage()
var db *sql.DB

// TaskManager is a function that root of this package
func TaskManager(session *discordgo.Session, event *discordgo.MessageCreate) {
	messages := strings.Split(event.Content, " ")
	command := messages[1]

	if !isSetupped {
		session.ChannelMessageSend(event.ChannelID, generalError.NotSetupped)
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
