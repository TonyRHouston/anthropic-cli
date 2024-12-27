// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func createCompletionsCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("completions.create", flag.ExitOnError)

	flagSet.Func(
		"max-tokens-to-sample",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "max_tokens_to_sample", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"model",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "model", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"prompt",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "prompt", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"metadata.user_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "metadata.user_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"stop-sequences",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "stop_sequences.#", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"+stop_sequence",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "stop_sequences.-1", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"temperature",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "temperature", number)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"top-k",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "top_k", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"top-p",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "top_p", number)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			var err error
			json, err = jsonSet(json, "stream", true)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}
			stream := client.Completions.NewStreaming(
				context.TODO(),
				anthropic.CompletionNewParams{},
				option.WithRequestBody("application/json", json),
			)
			for stream.Next() {
				fmt.Printf("%s\n", stream.Current().JSON.RawJSON())
			}
		},
	}
}
