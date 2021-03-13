package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Version: "v1.0.0",
		Commands: []*cli.Command{
			{
				Name:  "doctor",
				Usage: "System check",
				Action: func(c *cli.Context) error {
					dir, err := os.Getwd()
					if err != nil {
						return cli.Exit("Cannot get current working directory.", 1)
					}
					fmt.Println(dir)
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "Create new project",
				Subcommands: []*cli.Command{
					{
						Name:  "laravel",
						Usage: "Create new backend Laravel project",
						Action: func(c *cli.Context) error {
							// Check if Composer already installed
							path, err := exec.LookPath("composer")
							if err != nil {
								return cli.Exit("Composer not installed.", 1)
							}

							fmt.Printf("Composer found at %v\n", path)

							// Display Composer version
							cmd := exec.Command("bash", "-c", "composer -V")
							stdout, err := cmd.Output()

							if err != nil {
								return cli.Exit("Cannot get Composer version.", 1)
							}

							fmt.Print(string(stdout))

							// Composer Create Project
							projectDir := "laravel-app"

							if c.NArg() > 0 {
								projectDir = c.Args().Get(0)
							}

							cmd = exec.Command(
								"bash",
								"-c",
								"composer create-project laravel/laravel "+projectDir,
							)

							stdout, err = cmd.Output()

							if err != nil {
								return cli.Exit("Cannot install Laravel.", 1)
							}

							fmt.Print(string(stdout))
							return nil
						},
					},
					{
						Name:  "microsite",
						Usage: "Create new frontend microsite project  ",
						Action: func(c *cli.Context) error {
							fmt.Println("Success")
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
