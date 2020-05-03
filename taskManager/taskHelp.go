package taskManager

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func subjectHelp(session *discordgo.Session, channelID string) {
	sendMessageBase := ""
	for key, subjects := range courseSubjects {
		sendMessage := ""
		for index, subject := range subjects {
			if index == 0 {
				sendMessage = subject
				continue
			}
			sendMessage += fmt.Sprintf(", %s", subject)
		}
		sendMessageBase += fmt.Sprintf("%s\n```%s```\n", key, sendMessage)
	}
	session.ChannelMessageSend(channelID, sendMessageBase)
	return
}

func help(session *discordgo.Session, channelID string, messages []string) {
	if len(messages) > 2 && messages[2] == "subject" {
		subjectHelp(session, channelID)
		return
	}
	helpMessage := "***課題管理BOT***\n```!task add <task> <limit> <subject>```\ntask: 課題名\nlimit: 締め切り(初期値=翌日)\nsubject: 教科(初期値='')\n教科は省略できる\n"
	helpMessage += "```!task list <subject>```\n課題一覧を表示します\n<subject>を指定すると教科ごとの絞り込みが可能です\n"
	helpMessage += "```!task remove <task>```\n課題を課題名から検索して削除します"
	helpMessage += "```!task help (subject)```\n使い方を表示します\nsubjectを付けると利用可能な教科を表示します"
	session.ChannelMessageSend(channelID, helpMessage)
}
