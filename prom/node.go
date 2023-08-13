package prom

import (
	"disco/structs"
	"fmt"
	"os/exec"
)

func StartNodeInstance(i structs.NodeInstance) {
	cmd := exec.Command("node", i.IndexFile)

	// Start the Node.js application
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting Node.js application:", err)
		return
	}
}
