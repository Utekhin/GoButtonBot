package main

import (
	"bufio"
	"math/rand"
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Replace <token> with your bot's token
	bot, err := tgbotapi.NewBotAPI("<token>")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Create a new updates channel
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	// Open the text file
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	// Read the lines of the file into a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Start a loop to listen for updates
	for update := range updates {
		// If the update has a message
		if update.Message == nil {
			continue
		}

		// Check if the message is "Please"
		if update.Message.Text == "Please" {
			// Choose a random line from the file
			randomLine := lines[rand.Intn(len(lines))]

			// Send the random line back to the user
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, randomLine)
			reply.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(reply)
			if err != nil {
				log.Println("Error:", err)
			}
		} else {
			// Echo back the message to the user
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			reply.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(reply)
			if err != nil {
				log.Println("Error:", err)
			}
		}
	}
}
