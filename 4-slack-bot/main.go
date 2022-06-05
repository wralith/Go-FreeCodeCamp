package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	// Gets env variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connects bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	// If bot see that phrase
	bot.Command("year of birth is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator",
		Example:     "my year of birth is 2022",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			// Turn it into int
			if err != nil {
				println(err)
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age) // Response with "age is ..."
			response.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters) // >> &{map[year:****]}
		fmt.Println("---")
	}
}
