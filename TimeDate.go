package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	//Get the abbreviated name of the time zone (such as "CET") and its offset in seconds east of UTC.
	fmt.Println(start.Zone()) //returns PST -28800

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	//loc, _ := time.LoadLocation("Europe/Berlin") 		//e.g., "America/New_York"
	loc, _ := time.LoadLocation(time.Now().Location().String())

	t, _ = time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	// Note: without explicit zone, returns time in given location.
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

	//easy time zone conversions
	timeHere := time.Now()
	fmt.Printf("Time here=%v\n", timeHere.String())

	loc, _ = time.LoadLocation("UTC")
	utcTime := timeHere.In(loc)
	fmt.Printf("Time UTC =%v\n", utcTime)

	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("Runtime=%v\n", diff.String())

} //end main()
