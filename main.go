package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	// Choose the pin to control, for example, GPIO 17
	pin := rpio.Pin(17)

	// Set the pin to output mode
	pin.Output()

	// Blink the LED
	for i := 0; i < 10; i++ {
		pin.High() // Set pin high
		time.Sleep(time.Second)

		pin.Low() // Set pin low
		time.Sleep(time.Second)
	}
}
