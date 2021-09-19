package cli

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain/authing"
	"coryptex.com/bot/vip-signal/internal/domain/publishing"
	"coryptex.com/bot/vip-signal/pkg/excel"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

// CI Cli interface
type CI interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}

// CLI parameters and services
type CLI struct {
	Pubsvc  publishing.Service
	Authsvc authing.Service
}

// NewCLI return an instance of cli with suitable services
func NewCLI(pubsvc publishing.Service, authsvc authing.Service) CI {
	return &CLI{
		Pubsvc:  pubsvc,
		Authsvc: authsvc,
	}
}

// Start cli and get ready to answer commands
func (cli *CLI) Start(ctx context.Context) {
	///**** Commands Part
	commandLen := len(os.Args) - 1 // Exclude First that is File Path
	if commandLen == 0 {
		fmt.Println("You didn't provide any commandservice.")
	} else {
		if commandLen == 2 && strings.ToLower(os.Args[1]) == "setadmin" {
			fmt.Println("In set admin")
			err := cli.Authsvc.NewAdmin(ctx, os.Args[2])
			if err != nil {
				fmt.Println("Error Adding new Admin", errors.Wrap(err, "internal.presentation.cli"))
			}
		} else if commandLen == 2 && strings.ToLower(os.Args[1]) == "getfile" {
			fmt.Println("In get file")
			signals, err := excel.ReadSignals(os.Args[2], "")
			if err != nil {
				fmt.Println("Error Reading Signals from Excel", errors.Wrap(err, "internal.presentation.cli"))
			} else {
				fmt.Printf("Publishing...")
				err = cli.Pubsvc.Publish(ctx, signals)
				if err != nil {
					fmt.Println("Error Publishing Signals", errors.Wrap(err, "internal.presentation.cli"))
				}
				fmt.Printf("\r%s", "All messages Published.")
			}
		} else if commandLen == 1 && strings.ToLower(os.Args[1]) == "help" {
			fmt.Println("Usage: SetCommand commandservice description\nUsage: SetAdmin AdminChatId")
		} else {
			fmt.Println("Not allowed command, send help for usage")
		}
	}
}

// Stop cli
func (cli *CLI) Stop(ctx context.Context) {
	// TODO: Stop CLI?
}
