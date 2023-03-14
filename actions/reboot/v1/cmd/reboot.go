package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

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
	Run: func(cmd *cobra.Command, args []string) {

		log.Infof("Running reboot")
		// Load the kernel configuration into memory
		log.Info("Rebooting system")
	},
}

// Execute - starts the command parsing process.
func Execute() {
	if err := reboot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
