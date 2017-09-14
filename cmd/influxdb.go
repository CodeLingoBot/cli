package cmd

import (
	"github.com/Scalingo/cli/appdetect"
	"github.com/Scalingo/cli/cmd/autocomplete"
	"github.com/Scalingo/cli/db"
	"github.com/urfave/cli"
)

var (
	InfluxDBConsoleCommand = cli.Command{
		Name:     "influxdb-console",
		Category: "Databases",
		Usage:    "Run an interactive console with your InfluxDB addon",
		Flags: []cli.Flag{appFlag,
			cli.StringFlag{Name: "size, s", Value: "", Usage: "Size of the container"},
		},
		Description: ` Run an interactive console with your InfluxDB addon.

   Examples
    scalingo --app myapp influxdb-console
    scalingo --app myapp influxdb-console --size L

   The --size flag makes it easy to specify the size of the container executing
   the InfluxDB console. Each container size has different price and performance.
   You can read more about container sizes here:
   http://doc.scalingo.com/internals/container-sizes.html

    # See also 'mongo-console' and 'mysql-console'
`,
		Before: AuthenticateHook,
		Action: func(c *cli.Context) {
			currentApp := appdetect.CurrentApp(c)
			opts := db.InfluxDBConsoleOpts{
				App:  currentApp,
				Size: c.String("s"),
			}
			if len(c.Args()) != 0 {
				cli.ShowCommandHelp(c, "influxdb-console")
			} else if err := db.InfluxDBConsole(opts); err != nil {
				errorQuit(err)
			}
		},
		BashComplete: func(c *cli.Context) {
			autocomplete.CmdFlagsAutoComplete(c, "influxdb-console")
		},
	}
)