package main

import "time"

type syncData struct {
	via     string
	port    uint
	src     string
	dst     string
	verbose bool
}

func debounceChannel(interval time.Duration, output chan syncData) chan syncData {
	input := make(chan syncData)

	go func() {
		var buffer syncData
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
