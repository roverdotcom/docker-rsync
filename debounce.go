package main

import "time"

// SyncData holds data to be used in the method signature of Sync
type SyncData struct {
	via     string
	c       SSHCredentials
	src     string
	dst     string
	verbose bool
}

func debounceChannel(interval time.Duration, output chan SyncData) chan SyncData {
	input := make(chan SyncData)

	go func() {
		var buffer SyncData
		var ok bool

		// We do not start waiting for interval until called at least once
		buffer, ok = <-input
		// If channel closed exit, we could also close output
		if !ok {
			return
		}

		// We start waiting for an interval
		for {
			select {
			case buffer, ok = <-input:
				// If channel closed exit, we could also close output
				if !ok {
					return
				}

			case <-time.After(interval):
				// Interval has passed and we have data, so send it
				output <- buffer
				// Wait for data again before starting waiting for an interval
				buffer, ok = <-input
				if !ok {
					return
				}
				// If channel is not closed we have more data and start waiting for interval
			}
		}
	}()

	return input
}
