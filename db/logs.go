package db

import (
	"github.com/Scalingo/cli/config"
	"github.com/Scalingo/cli/logs"
	errgo "gopkg.in/errgo.v1"
)

type LogsOpts struct {
	Follow bool
	Count  int
}

func Logs(app, addon string, opts LogsOpts) error {
	c := config.ScalingoClient()
	url, err := c.AddonLogsURL(app, addon)
	if err != nil {
		return errgo.Notef(err, "fail to get log URL")
	}

	err = logs.Dump(url, opts.Count, "")
	if err != nil {
		return errgo.Notef(err, "fail to dump logs")
	}

	if opts.Follow {
		err := logs.Stream(url, "")
		if err != nil {
			return errgo.Notef(err, "fail to stream logs")
		}
	}
	return nil
}
