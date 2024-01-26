package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"hash/crc32"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:            "flywaycs",
		Usage:           "Allow you to provide a filepath and receive its checksum",
		Version:         "0.0.1",
		ArgsUsage:       "file_name",
		HideHelpCommand: true,
		Action: func(context *cli.Context) error {
			fmt.Println(generateCheckSum(context.Args().First()))
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func generateCheckSum(filePath string) uint32 {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	crc32q := crc32.MakeTable(crc32.IEEE)
	scanner := bufio.NewScanner(file)
	var checksum uint32
	for scanner.Scan() {
		checksum = crc32.Update(checksum, crc32q, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return checksum
}
