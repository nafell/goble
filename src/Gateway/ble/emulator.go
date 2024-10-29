package ble

import (
	"fmt"
	"log/slog"
	"time"
)

func StartEmulator(interval time.Duration, startValue int) {
	slog.Info(fmt.Sprintf("Starting sensor device emulator with duration of %d ms", interval.Milliseconds()))
	ticker := time.NewTicker(interval)
	done := make(chan bool)
	go func() {
		slog.Info("Emulator started. Use --debug to show logs.")
		x := startValue
		for {
			select {
			case <-done:
				return

			case t := <-ticker.C:
				slog.Debug(fmt.Sprintf("=== Emulate value: %d %s", x, t.Format("")))
				EmulateCO2Callback(x)
				x = (x + 100) % 3000 // loop value between 0-2999
			}

		}
	}()
}
