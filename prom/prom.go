package prom

import (
	"disco/structs"
	"os"
	"syscall"
)

// prom - Custom PROcess Manager

func GetRamUsage(proc *os.Process) (int64, error) {
	// Wait for the process to collect its state
	ps, err := proc.Wait()
	if err != nil {
		return 0, err
	}

	// Extract usage information from the process state
	rss := ps.SysUsage().(*syscall.Rusage).Maxrss

	// Convert RSS to bytes (on most systems, Maxrss is in kilobytes)
	rssBytes := rss * 1024

	return rssBytes, nil
}

// StartInstance() => Start{Language}Instance()
func StartInstance(i any) error {
	switch i.(type) {
	case structs.NodeInstance:
		{
			return StartNodeInstance(i.(structs.NodeInstance))
		}
	default:
		{
			return nil
		}
	}
}
