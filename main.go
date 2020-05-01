package main

import (
	"TaskInfoBot/taskManager"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var notifyChannel = "574884574778359844"

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
	tc := time.NewTicker(time.Second * 10)

	loopContinue := true

	for loopContinue {
		select {
		case <-sc:
			loopContinue = false
		case <-tc.C:
			taskNotify(discord)
		}
	}
	<-sc
}

func taskNotify(session *discordgo.Session) {
	session.ChannelMessageSend(notifyChannel, "***課題お知らせTIME***")
	now := getDate(time.Now())
	session.ChannelMessageSend(notifyChannel, "今日提出期限の課題は以下のとおりです")
	getTaskWithLimit(now, session)
	tomorrow := getDate(time.Now().Add(time.Duration(24) * time.Hour))
	session.ChannelMessageSend(notifyChannel, "明日提出期限の課題は以下のとおりです")
	getTaskWithLimit(tomorrow, session)
}

func getDate(date time.Time) string {
	return fmt.Sprintf("%d-%d-%d", date.Year(), int(date.Month()), date.Day())
}

func getTaskWithLimit(limit string, session *discordgo.Session) {
	rows, err := db.Query(`SELECT * FROM TASKS WHERE "LIMIT"=?`, limit)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var id int
		var task string
		var limit string
		var subject string

		if err := rows.Scan(&id, &task, &limit, &subject); err != nil {
			fmt.Println(err)
			continue
		}

		session.ChannelMessageSend(notifyChannel, fmt.Sprintf("```name: %s\nsubject: %s```", task, subject))
	}
}

func createFirstTable() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "TASKS" ("ID" INTEGER PRIMARY KEY, "TASK" TEXT, "LIMIT" TEXT, "SUBJECT" TEXT);`)
	if err != nil {
		panic(err)
	}
}

func onMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(event.Content, "!task") && len(event.Content) >= 8 {
		taskManager.TaskManager(session, event, db)
		return
	}
}

func loadTokenFromEnv() string {
	fp, err := os.Open(".env")
	defer fp.Close()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(fp)
	var token string
	for scanner.Scan() {
		token = scanner.Text()
	}
	return fmt.Sprintf("Bot %s", token)
}
