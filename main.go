package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

func main() {
	var isOutputEnabled bool

	app := &cli.App{
		Name:  "clipy",
		Usage: "Copy file contents or pipe data to clipboard",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Continue the flow of data by outputting whatever is input",
				Value:       false,
				Destination: &isOutputEnabled,
			},
		},
		Action: func(c *cli.Context) error {
			if isInputFromPipe() {
				return toClipboard(os.Stdin, isOutputEnabled)
			}
			path := c.Args().Get(0)
			file, e := getFile(path)
			if e != nil {
				return e
			}
			defer file.Close()
			return toClipboard(file, isOutputEnabled)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func toClipboard(r io.Reader, isOutputEnabled bool) error {
	var content string
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		text := scanner.Text()
		content += text + "\n"
		if isOutputEnabled {
			if _, e := fmt.Println(text); e != nil {
				return e
			}
		}
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
