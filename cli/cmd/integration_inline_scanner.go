//
// Author:: Salim Afiune Maya (<afiune@lacework.net>)
// Copyright:: Copyright 2022, Lacework Inc.
// License:: Apache License, Version 2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cmd

import (
	"strconv"

	"github.com/AlecAivazis/survey/v2"

	"github.com/lacework/go-sdk/api"
)

func createInlineScannerIntegration() error {
	questions := []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Name: "},
			Validate: survey.Required,
		},
		{
			Name: "identifier_tag",
			Prompt: &survey.Multiline{
				Message: "List of 'key:value' tags:"},
		},
		{
			Name: "limit_num_scan",
			Prompt: &survey.Input{
				Message: "Limit number of scans: ",
				Default: "60",
			},
			Validate: promptRequiredInt(
				"The limit must be a number.",
			),
		},
	}

	answers := struct {
		Name          string
		IdentifierTag string `survey:"identifier_tag"`
		LimitNumScan  string `survey:"limit_num_scan"`
	}{}

	if err := survey.Ask(questions, &answers,
		survey.WithIcons(promptIconsFunc),
	); err != nil {
		return err
	}

	limitNumScan, err := strconv.Atoi(answers.LimitNumScan)
	if err != nil {
		cli.Log.Warnw("unable to convert limit_num_scan, using default",
			"error", err,
			"input", answers.LimitNumScan,
			"default", "5",
		)
		limitNumScan = 60
	}

	inline := api.NewContainerRegistry(answers.Name,
		api.InlineScannerContainerRegistry,
		api.InlineScannerData{
			IdentifierTag: castStringToLimitByLabel(answers.IdentifierTag),
			LimitNumScan:  limitNumScan,
		},
	)

	cli.StartProgress("Creating integration...")
	_, err = cli.LwApi.V2.ContainerRegistries.Create(inline)
	cli.StopProgress()
	return err
}
