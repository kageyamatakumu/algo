package main

import "fmt"

const (
    HoursPerDay int  = 24
    MinutesPerHour int = 60
	SecondsPerMinute int = 60
)

func main() {
	fmt.Println(HoursPerDay * MinutesPerHour * SecondsPerMinute)
}