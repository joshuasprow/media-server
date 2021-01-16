package main

import (
	"log"

	"github.com/alexflint/go-arg"
	"github.com/joshuasprow/media-server/configurator/cmd"
	"github.com/joshuasprow/media-server/configurator/sonarr"
)

func main() {
	args := cmd.Arguments{}

	arg.MustParse(&args)

	c := sonarr.NewConfigurator(args)

	err := c.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
