package taskManager

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func taskDelete(session *discordgo.Session, channelID string, messages []string) {
	if len(messages) < 3 {
		session.ChannelMessageSend(channelID, "引数が足りません\n削除する課題の名前を指定してください")
	}

	deleteValue := messages[2]

	_, err := db.Exec(`DELETE FROM TASKS WHERE TASK=?`, deleteValue)
	if err != nil {
		session.ChannelMessageSend(channelID, "指定された名前の課題は存在しません")
		return
	}

	session.ChannelMessageSend(channelID, fmt.Sprintf("%sを削除しました", deleteValue))
}
