package main

import (
	"os"

	"github.com/google/logger"
	"github.com/urfave/cli"
)

const logPath = "tmux-pureillusion.log"

func main() {
	app := cli.NewApp()
	app.Name = "tmux-pureillusion"
	app.Version = "0.0.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Print info level logs to stdout",
		},
		cli.BoolFlag{
			Name:  "status-line-only",
			Usage: "Update status line only",
		},
		cli.StringFlag{
			Name:  "background",
			Value: "dark",
			Usage: "Color scheme to use",
		},
	}

	app.Usage = "apply Tmux theme to current tmux session"

	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print the version",
	}

	// const settings map[string]string = settings()

	app.Action = func(c *cli.Context) error {

		lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}
		defer lf.Close()

		defer logger.Init("TmuxPureillusion", c.Bool("verbose"), true, lf).Close()

		if c.Bool("status-line-only") != true {
			update_settings("pureillusion-tmux.json")
			update_colors(c.String("background"))
			tmux_setw("window-status-format", inactive_window())
			tmux_setw("window-status-current-format", active_window())
			tmux_setw("window-status-last-style", last_window())

			tmux_set("status-left", status_left())
			tmux_set("status-right", status_right())
		} else {
			tmux_set("status-left", status_left())
			tmux_set("status-right", status_right())
		}
		return nil
	}

	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Liss McCabe",
			Email: "liss@cuteasheck.com",
		},
	}

	app.Run(os.Args)
}
