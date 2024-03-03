package prom

import (
	"disco/structs"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type ProcessStatusCPUs struct {
	Min int64
	Max int64
}

type ProcessStatus struct {
	THPEnabled bool
	AnonMem    int64
	VmHWM      int64
	VmRss      int64

	RealMemory int64
	CPUs       ProcessStatusCPUs
}

// Simple implementation of reading the
func GetProcessStatus(pid string) (ProcessStatus, error) {
	path := fmt.Sprintf("/proc/%s/status", pid)
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return ProcessStatus{}, err
	}

	var ps ProcessStatus
	var real_mem int64

	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)

		if len(fields) > 1 {

			value := fields[1]
			switch fields[0] {

			case "VmHWM:":
				val, err := strconv.ParseInt(value, 0, 64)
				if err != nil {
					return ProcessStatus{}, err
				}
				ps.VmHWM = val
				real_mem += val
			case "VmRSS:":
				val, err := strconv.ParseInt(value, 0, 64)
				if err != nil {
					return ProcessStatus{}, err
				}
				ps.VmRss = val
				real_mem += val
			case "RssAnon:":
				val, err := strconv.ParseInt(value, 0, 64)
				if err != nil {
					return ProcessStatus{}, err
				}
				ps.AnonMem = val
				real_mem += val
			case "Cpus_allowed_list:":
				cut := strings.Split(value, "-")
				if len(cut) < 2 {
					continue
				}

				min_str := cut[0]
				max_str := cut[1]

				mini, err := strconv.ParseInt(min_str, 0, 8)
				if err != nil {
					return ProcessStatus{}, err
				}
				ps.CPUs.Min = mini

				maxi, err := strconv.ParseInt(max_str, 0, 8)
				if err != nil {
					return ProcessStatus{}, err
				}
				ps.CPUs.Max = maxi
			case "THP_enabled:":
				val, err := strconv.ParseBool(value)
				if err != nil {
					return ProcessStatus{}, err
				}
				ps.THPEnabled = val

				ps.RealMemory = real_mem

				if ps.AnonMem == 0 {
					return ProcessStatus{}, fmt.Errorf("could not find 'RssAnon' field in '%s'", path)
				}
			default:
				continue
			}
		}
	}

	return ps, nil
}

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
