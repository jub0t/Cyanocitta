package prom

import (
	"disco/structs"
	"disco/utils"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func StartNodeInstance(i structs.NodeInstance) error {
	// Initialize
	cmd := exec.Command("node", i.IndexFile)

	if !utils.FileExists(i.IndexFile) {
		return nil
	}

	// Start the Node.js application
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting Node.js application:", err)
		return nil
	}

	// Create a channel to receive termination signals
	exitCh := make(chan struct{})

	// Monitor the application in a goroutine
	go func() {
		err := cmd.Wait()
		if err != nil {
			fmt.Println("Node.js application exited with error:", err)
		} else {
			fmt.Println("Node.js application exited gracefully.")
		}
		close(exitCh) // Notify that the application has exited
	}()

	// Periodically check if the application is still running
	for {
		proc := cmd.Process
		select {
		case <-time.After(time.Duration(i.CheckInterval)):
			if isRunning(proc) {
				fmt.Println("Node.js application is still running...")
			} else {
				fmt.Println("Node.js application has stopped.")
			}
		case <-exitCh:
			fmt.Printf("Application Has Exited\n")
			return nil // Application has exited
		}

		if ram_usage, err := GetRamUsage(proc); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ram_usage)
		}
	}
}

func isRunning(process *os.Process) bool {
	err := process.Signal(os.Signal(os.Interrupt))
	return err == nil
}
