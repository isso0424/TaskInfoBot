package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New()
	discord.Token = loadTokenFromEnv()

	if err != nil {
		fmt.Println(err)
	}

	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	defer discord.Close()

	if err != nil {
		fmt.Println(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func onMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	TaskManager(session, event)
}

func loadTokenFromEnv() string {
	fp, err := os.Open(".env")
	if err != nil {
		fmt.Println(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var token string
	for scanner.Scan() {
		token = scanner.Text()
	}
	return fmt.Sprintf("Bot %s", token)
}
