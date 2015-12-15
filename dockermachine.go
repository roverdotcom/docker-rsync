package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func RunSSHCommand(machineName, command string, verbose bool) (out []byte, err error) {
	if verbose {
		fmt.Println(`docker-machine ssh ` + machineName + ` '` + command + `'`)
	}
	return exec.Command("/bin/sh", "-c", `docker-machine ssh `+machineName+` '`+command+`'`).CombinedOutput()
}

func GetSSHPort(machineName string) (port uint, err error) {
	out, err := exec.Command("/bin/sh", "-c", `docker-machine inspect `+machineName).CombinedOutput()
	if err != nil {
		return 0, err
	}

	return PortFromMachineJSON(out)
}

func PortFromMachineJSON(jsonData []byte) (port uint, err error) {
	var v struct {
		Driver struct {
			Driver struct {
				SSHPort uint
			}
			SSHPort uint
		}
	}

	if err := json.Unmarshal(jsonData, &v); err != nil {
		return 0, err
	}

	if v.Driver.SSHPort == 0 {
		return v.Driver.Driver.SSHPort, nil
	} else {
		return v.Driver.SSHPort, nil
	}
}
