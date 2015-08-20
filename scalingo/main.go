package main

import (
	"fmt"
	"os"

	"github.com/Scalingo/cli/Godeps/_workspace/src/github.com/Scalingo/codegangsta-cli"
	"github.com/Scalingo/cli/Godeps/_workspace/src/github.com/stvp/rollbar"
	"github.com/Scalingo/cli/cmd"
	"github.com/Scalingo/cli/config"
	"github.com/Scalingo/cli/signals"
	"github.com/Scalingo/cli/update"
)

func main() {
	app := cli.NewApp()
	app.Name = "Scalingo Client"
	app.Author = "Scalingo Team"
	app.Email = "hello@scalingo.com"
	app.Usage = "Manage your apps and containers"
	app.Version = config.Version
	app.CategorizedHelp = true
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "app, a", Value: "<name>", Usage: "Name of the app", EnvVar: "SCALINGO_APP"},
	}
	app.EnableBashCompletion = true
	app.Action = cmd.HelpCommand.Action
	app.Commands = []cli.Command{
		// Apps
		cmd.AppsCommand,
		cmd.CreateCommand,
		cmd.DestroyCommand,

		// Apps Actions
		cmd.LogsCommand,
		cmd.RunCommand,

		// Apps Process Actions
		cmd.PsCommand,
		cmd.ScaleCommand,
		cmd.RestartCommand,

		// Environment
		cmd.EnvCommand,
		cmd.EnvSetCommand,
		cmd.EnvUnsetCommand,

		// Domains
		cmd.DomainsListCommand,
		cmd.DomainsAddCommand,
		cmd.DomainsRemoveCommand,
		cmd.DomainsSSLCommand,

		// Collaborators
		cmd.CollaboratorsListCommand,
		cmd.CollaboratorsAddCommand,
		cmd.CollaboratorsRemoveCommand,

		// Addons
		cmd.AddonProvidersListCommand,
		cmd.AddonProvidersPlansCommand,
		cmd.AddonsListCommand,
		cmd.AddonsAddCommand,
		cmd.AddonsRemoveCommand,
		cmd.AddonsUpgradeCommand,

		// DB Access
		cmd.DbTunnelCommand,
		cmd.RedisConsoleCommand,
		cmd.MongoConsoleCommand,
		cmd.MySQLConsoleCommand,
		cmd.PgSQLConsoleCommand,

		// SSH keys
		cmd.ListSSHKeyCommand,
		cmd.AddSSHKeyCommand,
		cmd.RemoveSSHKeyCommand,

		// Sessions
		cmd.LoginCommand,
		cmd.LogoutCommand,
		cmd.SignUpCommand,

		// Version
		cmd.VersionCommand,
		cmd.UpdateCommand,

		// Help
		cmd.HelpCommand,
	}

	go signals.Handle()

	if len(os.Args) >= 2 && os.Args[1] == cmd.UpdateCommand.Name {
		err := update.Check()
		if err != nil {
			rollbar.Error(rollbar.ERR, err)
		}
		return
	} else {
		defer update.Check()
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Fail to run scalingo", err)
	}
}
