package clipy

import (
	"bufio"
	"io"
	"os"

	"github.com/pkg/errors"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

func main() {
	var path string

	app := &cli.App{
		Name:  "clipy",
		Usage: "Copy file contents or pipe data to clipboard",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Value:       "",
				Usage:       "Optional path to file",
				Destination: &path,
			},
		},
		Action: func(c *cli.Context) error {
			if isInputFromPipe() {
				return toClipboard(os.Stdin)
			} else {
				file, e := getFile(path)
				if e != nil {
					return e
				}
				defer file.Close()
				return toClipboard(file)
			}
		},
	}

	app.Run(os.Args)
}

func toClipboard(r io.Reader) error {
	var content string
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		content += scanner.Text()
	}
	return clipboard.WriteAll(content)
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func getFile(filepath string) (file *os.File, e error) {
	if filepath == "" {
		return nil, errors.New("please provide a path to a file")
	}
	if !fileExists(filepath) {
		return nil, errors.New("the provided file does not exist")
	}
	file, e = os.Open(filepath)
	if e != nil {
		return nil, errors.Wrapf(e, "unable to read the provided file")
	}
	return
}

func fileExists(filepath string) bool {
	info, e := os.Stat(filepath)
	return !os.IsNotExist(e) && !info.IsDir()
}
