package main

import (
	"TaskInfoBot/loadConfig"
	"TaskInfoBot/taskManager"
	"TaskInfoBot/taskNotify"
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

func main() {
	dbFileName := "./db.sqlite3"
	if fileNotExists(dbFileName) {
		file, err := os.OpenFile(dbFileName, os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}

	tmp, err := sql.Open("sqlite3", dbFileName)
	db = tmp
	if err != nil {
		panic(err)
	}
	createFirstTable()

	config := loadConfig.LoadConfig()
	taskManager.SetConfig(config)

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
	tc := time.NewTicker(time.Hour * 3)

	fmt.Println("Bot booted!!!")

	loopContinue := true

	for loopContinue {
		select {
		case <-sc:
			loopContinue = false
		case <-tc.C:
			taskNotify.TaskNotify(discord, db, config)
		}
	}
	<-sc
}

func createFirstTable() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "TASKS" ("ID" INTEGER PRIMARY KEY, "TASK" TEXT, "LIMIT" TEXT, "SUBJECT" TEXT, "COURSE" TEXT);`)
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

func fileNotExists(filename string) bool {
	_, err := os.Stat(filename)
	return err != nil
}
