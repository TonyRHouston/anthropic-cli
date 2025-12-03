// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/stainless-sdks/anthropic-cli/internal/apiquery"
	"github.com/stainless-sdks/anthropic-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var betaSkillsVersionsCreate = cli.Command{
	Name:  "create",
	Usage: "Create Skill Version",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&requestflag.StringSliceFlag{
			Name:  "file",
			Usage: "Files to upload for the skill.\n\nAll files must be in the same top-level directory and must include a SKILL.md file at the root of that directory.",
			Config: requestflag.RequestConfig{
				BodyPath: "files",
			},
		},
		&requestflag.YAMLSliceFlag{
			Name:  "beta",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: requestflag.RequestConfig{
				HeaderPath: "anthropic-beta",
			},
		},
	},
	Action:          handleBetaSkillsVersionsCreate,
	HideHelpCommand: true,
}

var betaSkillsVersionsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get Skill Version",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&requestflag.StringFlag{
			Name:  "version",
			Usage: "Version identifier for the skill.\n\nEach version is identified by a Unix epoch timestamp (e.g., \"1759178010641129\").",
		},
		&requestflag.YAMLSliceFlag{
			Name:  "beta",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: requestflag.RequestConfig{
				HeaderPath: "anthropic-beta",
			},
		},
	},
	Action:          handleBetaSkillsVersionsRetrieve,
	HideHelpCommand: true,
}

var betaSkillsVersionsList = cli.Command{
	Name:  "list",
	Usage: "List Skill Versions",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&requestflag.IntFlag{
			Name:  "limit",
			Usage: "Number of items to return per page.\n\nDefaults to `20`. Ranges from `1` to `1000`.",
			Config: requestflag.RequestConfig{
				QueryPath: "limit",
			},
		},
		&requestflag.StringFlag{
			Name:  "page",
			Usage: "Optionally set to the `next_page` token from the previous response.",
			Config: requestflag.RequestConfig{
				QueryPath: "page",
			},
		},
		&requestflag.YAMLSliceFlag{
			Name:  "beta",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: requestflag.RequestConfig{
				HeaderPath: "anthropic-beta",
			},
		},
	},
	Action:          handleBetaSkillsVersionsList,
	HideHelpCommand: true,
}

var betaSkillsVersionsDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete Skill Version",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&requestflag.StringFlag{
			Name:  "version",
			Usage: "Version identifier for the skill.\n\nEach version is identified by a Unix epoch timestamp (e.g., \"1759178010641129\").",
		},
		&requestflag.YAMLSliceFlag{
			Name:  "beta",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: requestflag.RequestConfig{
				HeaderPath: "anthropic-beta",
			},
		},
	},
	Action:          handleBetaSkillsVersionsDelete,
	HideHelpCommand: true,
}

func handleBetaSkillsVersionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("skill-id") && len(unusedArgs) > 0 {
		cmd.Set("skill-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillVersionNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Skills.Versions.New(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "skill-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("beta:skills:versions create", json, format, transform)
}

func handleBetaSkillsVersionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version") && len(unusedArgs) > 0 {
		cmd.Set("version", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillVersionGetParams{
		SkillID: requestflag.CommandRequestValue[string](cmd, "skill-id"),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Skills.Versions.Get(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "version"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("beta:skills:versions retrieve", json, format, transform)
}

func handleBetaSkillsVersionsList(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("skill-id") && len(unusedArgs) > 0 {
		cmd.Set("skill-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillVersionListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Skills.Versions.List(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "skill-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("beta:skills:versions list", json, format, transform)
}

func handleBetaSkillsVersionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version") && len(unusedArgs) > 0 {
		cmd.Set("version", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillVersionDeleteParams{
		SkillID: requestflag.CommandRequestValue[string](cmd, "skill-id"),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Skills.Versions.Delete(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "version"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("beta:skills:versions delete", json, format, transform)
}
