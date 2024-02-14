package main

import (
	"log"
	"os"
	"path"
	"pkt2xml/crypt"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	decryptCommand := &cli.Command{
		Name:    "decrypt",
		Aliases: []string{"d", "dec"},
		Usage:   "decrypt PKT/PKA and output XML",
		Action: func(ctx *cli.Context) error {
			input_file_path := ctx.Args().First()
			content, err := os.ReadFile(input_file_path)
			if err != nil {
				return err
			}

			out := strings.Split(path.Base(input_file_path), ".")
			out[len(out)-1] = "xml"
			outfile := strings.Join(out, ".")

			result, err := crypt.Decrypt(content)
			if err != nil {
				return err
			}

			file, err := os.Create(outfile)
			if err != nil {
				return err
			}

			file.Write(result)

			return nil
		},
	}
	encryptCommand := &cli.Command{
		Name:    "encrypt",
		Aliases: []string{"e", "enc"},
		Usage:   "encrypt XML and output PKT",
		Action: func(ctx *cli.Context) error {
			input_file_path := ctx.Args().First()
			content, err := os.ReadFile(input_file_path)
			if err != nil {
				return err
			}

			out := strings.Split(path.Base(input_file_path), ".")
			out[len(out)-1] = "pkt"
			outfile := strings.Join(out, ".")

			result, err := crypt.Encrypt(content)
			if err != nil {
				return err
			}

			file, err := os.Create(outfile)
			if err != nil {
				return err
			}

			file.Write(result)

			return nil
		},
	}

	app := &cli.App{
		Name:  "pka2xml",
		Usage: "Converts pka files to xml and vice versa",
		Commands: []*cli.Command{
			decryptCommand,
			encryptCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
