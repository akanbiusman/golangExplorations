package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Pomodoro Timer Started!")
	startTimer(25*time.Minute, 5*time.Minute)
}

func startTimer(workDuration, breakDuration time.Duration) {
	for {
		fmt.Println("\nFocus time! Work for", workDuration.Minutes(), "minutes.")
		countdown(workDuration)

		fmt.Println("\nBreak time! Relax for", breakDuration.Minutes(), "minutes.")
		countdown(breakDuration)
	}
}

func countdown(duration time.Duration) {
	timer := time.NewTimer(duration)
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			duration -= time.Second
			fmt.Printf("\rTime left: %v ", duration.Round(time.Second))
		case <-timer.C:
			fmt.Println("\nTime’s up!")
			ticker.Stop()
			return
		}
	}
}
