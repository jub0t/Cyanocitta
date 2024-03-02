package prom

import (
	"disco/structs"
)

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

type ProcessStatus struct {
    AnonMem int64
    THPEnabled bool
    VmHWM int64
    VmRss int64

    RealMemory int64
}

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
        
        if (len(fields) < 1) {
          continue;
        }

        switch fields[0] {
 
        case "VmHWM:":
            val, err := strconv.ParseInt(fields[1], 10, 64)
            if err != nil {
                return ProcessStatus{}, err
            }
            ps.VmHWM = val
            real_mem += val
        case "VmRSS:":
            val, err := strconv.ParseInt(fields[1], 10, 64)
            if err != nil {
                return ProcessStatus{}, err
            }
            ps.VmRss = val
            real_mem += val
        case "RssAnon:":
            val, err := strconv.ParseInt(fields[1], 10, 64)
            if err != nil {
                return ProcessStatus{}, err
            }
            ps.AnonMem = val
            real_mem += val
        case "THP_enabled:":
            val, err := strconv.ParseBool(fields[1])
            if err != nil {
                return ProcessStatus{}, err
            }
            ps.THPEnabled = val
        default:
            continue
        }
    }

    ps.RealMemory = real_mem

    if ps.AnonMem == 0 {
        return ProcessStatus{}, fmt.Errorf("could not find 'RssAnon' field in '%s'", path)
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
