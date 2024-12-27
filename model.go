// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

func createModelsRetrieveSubcommand() Subcommand {
	var modelID *string = nil
	var flagSet = flag.NewFlagSet("models.retrieve", flag.ExitOnError)

	flagSet.Func(
		"model-id",
		"",
		func(string string) error {
			modelID = &string
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Models.Get(context.TODO(), *modelID)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createModelsListSubcommand() Subcommand {
	var flagSet = flag.NewFlagSet("models.list", flag.ExitOnError)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Models.List(context.TODO(), anthropic.ModelListParams{})
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}
