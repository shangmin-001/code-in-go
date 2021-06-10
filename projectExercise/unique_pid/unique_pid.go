package unique_pid

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

func CheckPidFile(file string) error {
	contents, err := ioutil.ReadFile(file)
	if err == nil {
		pid, err := strconv.Atoi(strings.TrimSpace(string(contents)))
		if err != nil {
			return fmt.Errorf("reading proccess id from pidfile '%s' error: %v", file, err)
		}

		process, err := os.FindProcess(pid)
		// on Windows, err != nil if the process cannot be found
		if runtime.GOOS == "windows" {
			if err == nil {
				return fmt.Errorf("process %d is already running", pid)
			}
		} else if process != nil {
			// err is always nil on POSIX, so we have to send the process
			// a signal to check whether it exists
			if err = process.Signal(syscall.Signal(0)); err == nil {
				return fmt.Errorf("process %d is already running", pid)

			}
		}
	}
	if err = ioutil.WriteFile(file, []byte(strconv.Itoa(os.Getpid())), 0644); err != nil {
		return fmt.Errorf("unable to write pidfile '%s': %s", file, err)
	}

	return nil
}
