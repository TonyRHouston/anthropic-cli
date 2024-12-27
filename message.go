// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func createMessagesCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.create", flag.ExitOnError)

	flagSet.Func(
		"max-tokens",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "max_tokens", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"messages.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.tool_use_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.tool_use_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.is_error",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "messages.#.content.#.is_error", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_use_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_result_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+document_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.role",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.role", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+message",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.-1", map[string]interface{}{})
			if err != nil {
				return err
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
		"system.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+system",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "system.-1", map[string]interface{}{})
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
		"tool-choice.tool_choice_auto.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_auto.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"tools.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tools.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tools.description",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.description", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+tool",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "tools.-1", map[string]interface{}{})
			if err != nil {
				return err
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
			stream := client.Messages.NewStreaming(
				context.TODO(),
				anthropic.MessageNewParams{},
				option.WithRequestBody("application/json", json),
			)
			for stream.Next() {
				fmt.Printf("%s\n", stream.Current().JSON.RawJSON())
			}
		},
	}
}

func createMessagesCountTokensSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.count_tokens", flag.ExitOnError)

	flagSet.Func(
		"messages.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.tool_use_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.tool_use_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.is_error",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "messages.#.content.#.is_error", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_use_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_result_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+document_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.role",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.role", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+message",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.-1", map[string]interface{}{})
			if err != nil {
				return err
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
		"system.union_member_0",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.union_member_1.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.union_member_1.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.union_member_1.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"system.+union_member_1",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "system.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"tool-choice.tool_choice_auto.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_auto.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"tools.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tools.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tools.description",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.description", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+tool",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "tools.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.CountTokens(
				context.TODO(),
				anthropic.MessageCountTokensParams{},
				option.WithRequestBody("application/json", json),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}
