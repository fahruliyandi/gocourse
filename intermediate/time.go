package basic

import (
	"fmt"
	"time"
)

func main() {
	//current local time
	fmt.Println(time.Now())

	//spesific time
	specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println("spesific time :", specificTime)

	//parse time
	//layout value : Mon Jan 2 15:04:05 MST 2006
	parsedTime, _ := time.Parse("2006-01-02", "2006-05-01")
	parsedTime1, _ := time.Parse("06-01-02", "06-05-01")
	fmt.Println("parsed time : ", parsedTime)
	fmt.Println("parsed time : ", parsedTime1)

	//formatin timme
	t := time.Now()
	fmt.Println("Formatted time :", t.Format("06-01-02 15-04"))

	//
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("One Day Later : ", oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	//rounding time
	fmt.Println("Rounded Time : ", t.Round(time.Hour))

	//truncate time
	fmt.Println("Truncated time :", t.Truncate(time.Hour))

	//substraction
	t1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("Duration :", duration)

	//compares time
	fmt.Println("t2 after t1?", t2.After(t1))

}
