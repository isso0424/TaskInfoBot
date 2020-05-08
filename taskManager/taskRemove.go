package taskManager

import (
	"TaskInfoBot/messageController"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var removeMessage = messageController.CreateTaskRemoveMessage()
var removeError = messageController.CreateRemoveErrorMessage()

func taskDelete(session *discordgo.Session, channelID string, messages []string) {
	if len(messages) < 3 {
		session.ChannelMessageSend(channelID, removeError.NotEnoughArgs)
	}

	deleteValue := messages[2]

	_, err := db.Exec(`DELETE FROM TASKS WHERE TASK=?`, deleteValue)
	if err != nil {
		session.ChannelMessageSend(channelID, removeError.TaskNotFound)
		return
	}

	session.ChannelMessageSend(channelID, fmt.Sprintf(removeMessage, deleteValue))
}
