package main

import (
	"TaskInfoBot/taskManager"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	tmp, err := sql.Open("sqlite3", "./db.sqlite3")
	db = tmp

	if err != nil {
		panic(err)
	}
	createFirstTable()

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

func createFirstTable() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "TASKS" ("ID" INTEGER PRIMARY KEY, "TASK" TEXT, "LIMIT" TEXT, "SUBJECT" TEXT);`)
	if err != nil {
		panic(err)
	}
}

func onMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	taskManager.TaskManager(session, event)
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
