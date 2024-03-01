package prom

import (
	"disco/structs"
	"os"
)

// prom - Custom PROcess Manager
func GetRamUsage(proc *os.Process) (int64, error) {
	return 10, nil
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
