package main

import (
	"fmt"

	"github.com/ader1990/hub/actions/reboot/v1/cmd"
)

func main() {
	fmt.Printf(" - Kernel Reboot\n------------------------\n")
	cmd.Execute()
}
