// GoCD CLI tool

package main

import (
	gocli "github.com/drewsonne/go-gocd/cli"
	"github.com/urfave/cli"
	"os"
)

// GoCDUtilityName is used in help text to identify the gocd cli util by name
const GoCDUtilityName = "gocd"

// GoCDUtilityUsageInstructions providers user facing support on operation of the gocd cli tool. @TODO Expand this content.
const GoCDUtilityUsageInstructions = "CLI Tool to interact with GoCD server"

func main() {

	app := cli.NewApp()
	app.Name = GoCDUtilityName
	app.Usage = GoCDUtilityUsageInstructions
	app.Version = gocli.Version()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		*gocli.ConfigureCommand(),
		*gocli.ListAgentsCommand(),
		*gocli.ListPipelineTemplatesCommand(),
		*gocli.GetAgentCommand(),
		*gocli.GetPipelineTemplateCommand(),
		*gocli.CreatePipelineTemplateCommand(),
		*gocli.UpdateAgentCommand(),
		*gocli.UpdateAgentsCommand(),
		*gocli.UpdatePipelineConfigCommand(),
		*gocli.UpdatePipelineTemplateCommand(),
		*gocli.DeleteAgentCommand(),
		*gocli.DeleteAgentsCommand(),
		*gocli.DeletePipelineTemplateCommand(),
		*gocli.DeletePipelineConfigCommand(),
		*gocli.ListPipelineGroupsCommand(),
		*gocli.GetPipelineHistoryCommand(),
		*gocli.CreatePipelineConfigCommand(),
		*gocli.GenerateJSONSchemaCommand(),
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "server", EnvVar: gocli.EnvVarServer},
		cli.StringFlag{Name: "username", EnvVar: gocli.EnvVarUsername},
		cli.StringFlag{Name: "password", EnvVar: gocli.EnvVarPassword},
	}

	app.Run(os.Args)
}