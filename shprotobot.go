package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"gopkg.in/telegram-bot-api.v4"
)

type Configuration struct {
	Token string
}

var err error
var bot *tgbotapi.BotAPI

func main() {
	reg, _ := regexp.Compile("^пр(([уеыаэию]т)|о)")
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	fatal(err)
	bot, err = tgbotapi.NewBotAPI(config.Token)
	fatal(err)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	fatal(err)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		rand.Seed(time.Now().UTC().UnixNano())
		words := strings.Split(update.Message.Text, " ")
		shprots := []string{}
		for _, word := range words {
			word = strings.ToLower(word)
			matched := reg.MatchString(word)
			fatal(err)
			if matched {
				shprots = append(shprots, fmt.Sprintf("ш%s", word))
			}
		}
		if len(shprots) > 0 {
			text := shprots[rand.Intn(len(shprots))]
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}

func fatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
