package main

import(
	"fmt"
	"time"
)

var now = time.Now()
var election = time.Date(2016, time.November, 8, 0, 0, 0, 0, time.UTC)
func main() {
	//get duration between election date and now
	tillElection := election.Sub(now)
	//get duration in nanoseconds
	toNanoseconds := tillElection.Nanoseconds()
	//calculate hours from toNanoseconds
	hours := toNanoseconds/3600000000000
	remainder := toNanoseconds%3600000000000
	//derive minutes from remainder of hours
	minutes := remainder/60000000000
	remainder = remainder%60000000000
	//derive seconds from remainder of minutes
	seconds := remainder/1000000000 
	//calculate days and get hours left from remainder
	days := hours/24
	hoursLeft := hours%24

	fmt.Printf("\nHow long until the 2016 U.S. Presidential election?\n\n%v Days %v Hours %v Minutes %v Seconds\n\n", days, hoursLeft, minutes, seconds)

}
