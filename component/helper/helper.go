package helper

import (
	"os"
	"runtime"
	"time"
)

// WatchMaster watches child procs
// Inspired by https://github.com/gofiber/fiber/blob/master/prefork.go
func WatchMaster(interval time.Duration) {
	if runtime.GOOS == "windows" {
		// finds parent process,
		// and waits for it to exit
		p, err := os.FindProcess(os.Getppid())
		if err == nil {
			_, _ = p.Wait() //nolint:errcheck // It is fine to ignore the error here
		}
		os.Exit(1) //nolint:revive // Calling os.Exit is fine here in the prefork
	}
	// if it is equal to 1 (init process ID),
	// it indicates that the master process has exited
	for range time.NewTicker(interval).C {
		if os.Getppid() == 1 {
			os.Exit(1) //nolint:revive // Calling os.Exit is fine here in the prefork
		}
	}
}
