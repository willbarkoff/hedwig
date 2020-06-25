package main

import (
	"flag"
	"fmt"
	"runtime"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

const version = "0.1.0"

func main() {
	versionFlag := flag.Bool("version", false, "Shows the version of the Hedwig CLI")
	recipientFlag := flag.String("recipient", "", "The recipent of the message. If this isn't specified, it defaults to every recipeint in the config file")
	addRecipientGlobalFlag := flag.String("add-recipient-global", "", "Adds a recipient to the config file")
	title := flag.String("title", "Message from Hedwig!", "The title of the notification")
	ttl := flag.Int("ttl", 0, "Time to live, in seconds, for the message")
	badge := flag.Int("badge", 0, "Number to put on the badge on in iOS app")
	sound := flag.String("sound", "default", "The sound to play on iOS")

	flag.Parse()
	recipients := readConfigFile()

	if *versionFlag {
		printVersion()
		return
	} else if *addRecipientGlobalFlag != "" {
		addRecipientToConfigFile(*addRecipientGlobalFlag)
		fmt.Println("Added recipient to global config")
		return
	} else if *recipientFlag != "" {
		recipients = []string{*recipientFlag}
	} else if len(recipients) == 0 {
		panic("No recipients were specified. You can add recipients using the --add-recipient-global flag to add a default recipient, or the --recipient flag to add a single recipient")
	}

	client := expo.NewPushClient(nil)

	if flag.Arg(0) == "" {
		panic("No message was specified!")
	}

	for _, recipient := range recipients {
		expoToken, err := expo.NewExponentPushToken(recipient)
		if err != nil {
			panic(err)
		}
		response, err := client.Publish(
			&expo.PushMessage{
				To:         expoToken,
				Body:       flag.Arg(0),
				Sound:      *sound,
				Title:      *title,
				TTLSeconds: *ttl,
				Badge:      *badge,
				Priority:   expo.DefaultPriority,
			},
		)
		if err != nil {
			panic(err)
		}
		if response.ValidateResponse() != nil {
			fmt.Println(response.PushMessage.To, "failed")
		}
	}

}

func printVersion() {
	fmt.Printf("🦉 Hedwig v%s (%s/%s)\n", version, runtime.GOOS, runtime.GOARCH)
}
