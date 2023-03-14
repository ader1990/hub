package cmd

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Release struct {
	Version string
	Build   string
}

const mountAction = "/mountAction"

var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "This is an action for performing a reboot",
	Run: func(cmd1 *cobra.Command, args []string) {

		log.Info("Rebooting system")
		var defaultCommand = "shutdown"
		cmd := exec.Command(defaultCommand, "-r", "+1")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		err := cmd.Start()
		if err != nil {
			log.Fatalf("Error starting [%s] [%v]", defaultCommand, err)
		}
		err = cmd.Wait()
		if err != nil {
			log.Fatalf("Error running [%s] [%v]", defaultCommand, err)
		}

	},
}

// Execute - starts the command parsing process.
func Execute() {
	if err := rebootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
