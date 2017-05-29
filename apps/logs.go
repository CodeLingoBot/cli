package apps

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	stdio "io"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/Scalingo/cli/config"
	"github.com/Scalingo/cli/debug"
	"github.com/Scalingo/cli/io"
	"github.com/Scalingo/go-scalingo"
	"github.com/fatih/color"
	"golang.org/x/net/websocket"
	"gopkg.in/errgo.v1"
)

type WSEvent struct {
	Type      string    `json:"event"`
	Log       string    `json:"log"`
	Timestamp time.Time `json:"timestamp"`
}

type LogsRes struct {
	LogsURL string        `json:"logs_url"`
	App     *scalingo.App `json:"app"`
}

func Logs(appName string, stream bool, n int, filter string) error {
	err := checkFilter(appName, filter)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	c := config.ScalingoClient()
	res, err := c.LogsURL(appName)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errgo.Newf("fail to query logs: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	debug.Println("[API-Response] ", string(body))

	logsRes := &LogsRes{}
	if err = json.Unmarshal(body, &logsRes); err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	if err = dumpLogs(logsRes.LogsURL, n, filter); err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	if stream {
		if err = streamLogs(logsRes.LogsURL, filter); err != nil {
			return errgo.Mask(err, errgo.Any)
		}
	}
	return nil
}

func dumpLogs(logsURL string, n int, filter string) error {
	c := config.ScalingoClient()
	res, err := c.Logs(logsURL, n, filter)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 || res.StatusCode == 204 {
		io.Error("There is not log for this application")
		io.Info("Ensure your application is writing to the standard output")
		return nil
	}

	sr := bufio.NewReader(res.Body)
	for {
		bline, err := sr.ReadBytes('\n')
		if err != nil {
			break
		}

		colorizeLogs(string(bline))
	}

	return nil
}

func streamLogs(logsRawURL string, filter string) error {
	var (
		err   error
		event WSEvent
	)

	logsURL, err := url.Parse(logsRawURL)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}
	if logsURL.Scheme == "https" {
		logsURL.Scheme = "wss"
	} else {
		logsURL.Scheme = "ws"
	}

	logsURLString := fmt.Sprintf("%s&stream=true", logsURL.String())
	if filter != "" {
		logsURLString = fmt.Sprintf("%s&filter=%s", logsURLString, filter)
	}

	conn, err := websocket.Dial(logsURLString, "", "http://scalingo-cli.local/"+config.Version)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	for {
		err := websocket.JSON.Receive(conn, &event)
		if err != nil {
			conn.Close()
			if err == stdio.EOF {
				debug.Println("Remote server broke the connection, reconnecting")
				for err != nil {
					conn, err = websocket.Dial(logsURLString, "", "http://scalingo-cli.local/"+config.Version)
					time.Sleep(time.Second * 1)
				}
				continue
			} else {
				return errgo.Mask(err, errgo.Any)
			}
		} else {
			switch event.Type {
			case "ping":
			case "log":
				colorizeLogs(strings.TrimSpace(event.Log))
			}
		}
	}
}

func checkFilter(appName string, filter string) error {
	if filter != "" {
		c := config.ScalingoClient()
		processes, err := c.AppsPs(appName)
		if err != nil {
			return errgo.Mask(err)
		}

		filters := strings.Split(filter, "|")
		for _, f := range filters {

			ctName := ""
			for _, ct := range processes {

				ctName = ct.Name
				if strings.HasPrefix(f, ctName+"-") || f == ctName {
					break
				}
			}
			if !strings.HasPrefix(f, ctName+"-") && f != ctName {
				return errgo.Newf(
					"%s is not a valid container filter\n\nEXAMPLES:\n"+
						"\"scalingo logs -F web\": logs of every web containers\n"+
						"\"scalingo logs -F web-1\": logs of web container 1\n"+
						"\"scalingo logs -F web|worker\": logs of every web and worker containers\n",
					f)
			}
		}
	}

	return nil
}

type colorFunc func(...interface{}) string

var (
	containerColors = []colorFunc{
		color.New(color.FgBlue).SprintFunc(),
		color.New(color.FgCyan).SprintFunc(),
		color.New(color.FgGreen).SprintFunc(),
		color.New(color.FgMagenta).SprintFunc(),
		color.New(color.FgRed).SprintFunc(),
		color.New(color.FgHiYellow).SprintFunc(),
		color.New(color.FgHiBlue).SprintFunc(),
		color.New(color.FgHiCyan).SprintFunc(),
		color.New(color.FgHiGreen).SprintFunc(),
		color.New(color.FgHiMagenta).SprintFunc(),
		color.New(color.FgHiRed).SprintFunc(),
	}
)

func colorizeLogs(logs string) {
	lines := strings.Split(logs, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		lineSplit := strings.Split(line, " ")
		if len(lineSplit) < 5 {
			fmt.Println(line)
			continue
		}
		content := strings.Join(lineSplit[5:], " ")

		headerSplit := lineSplit[:5]
		date := strings.Join(headerSplit[:4], " ")
		containerWithSurround := headerSplit[4]
		container := containerWithSurround[1 : len(containerWithSurround)-1]

		hash := md5.Sum([]byte(container))
		colorId := int(hash[0]+hash[1]+hash[2]+hash[3]) % len(containerColors)

		fmt.Printf(
			"%s [%s] %s\n",
			color.New(color.FgYellow).SprintFunc()(date),
			containerColors[colorId](container),
			content,
		)
	}
}
