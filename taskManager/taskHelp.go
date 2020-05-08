package taskManager

import (
	"TaskInfoBot/messageController"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var helpMessage = messageController.CreateTaskHelpMessage()

func subjectHelp(session *discordgo.Session, channelID string) {
	sendMessageBase := ""
	for key, subjects := range courseSubjects {
		sendMessage := ""
		for index, subject := range subjects {
			if index == 0 {
				sendMessage = subject
				continue
			}
			sendMessage += fmt.Sprintf(helpMessage.Subjects.EachSubject, subject)
		}
		sendMessageBase += fmt.Sprintf(helpMessage.Subjects.EachCourse, key, sendMessage)
	}
	session.ChannelMessageSend(channelID, sendMessageBase)
	return
}

func help(session *discordgo.Session, channelID string, messages []string) {
	if len(messages) > 2 && messages[2] == "subject" {
		subjectHelp(session, channelID)
		return
	}
	session.ChannelMessageSend(channelID, helpMessage.General)
}
