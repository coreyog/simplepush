package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/simplepush/simplepush-go"
)

// Arguments x
type Arguments struct {
	Title       string `short:"t" long:"title" description:"title of notification"`
	Event       string `short:"e" long:"event" description:"event name of the notification"`
	UnnamedArgs struct {
		Message []string
	} `positional-args:"yes" positional-arg-name:"MESSAGE"`
}

func main() {
	var args Arguments
	_, err := flags.Parse(&args)
	if flags.WroteHelp(err) {
		return
	} else if err != nil {
		panic(err)
	}
	msgText := strings.Join(args.UnnamedArgs.Message, " ")
	if len(msgText) > 0 {
		// Send notification
		simpleKey, hasKey := os.LookupEnv("SIMPLEPUSH_KEY")
		simplePassword, hasPassword := os.LookupEnv("SIMPLEPUSH_PASSWORD")
		simpleSalt, hasSalt := os.LookupEnv("SIMPLEPUSH_SALT")

		if hasKey && hasPassword && hasSalt {
			msg := simplepush.Message{
				SimplePushKey: simpleKey,
				Title:         args.Title,
				Message:       msgText,
				Event:         args.Event,
				Encrypt:       true,
				Salt:          simpleSalt,
				Password:      simplePassword,
			}
			simplepush.Send(msg)
		} else {
			fmt.Println("Missing environment variables: SIMPLEPUSH_KEY, SIMPLEPUSH_PASSWORD, SIMPLEPUSH_SALT")
		}
	}
}
