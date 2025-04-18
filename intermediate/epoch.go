package basic

import (
	"fmt"
	"time"
)

func main() {
	//00:00:00 UTC on Jan 1, 1970
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time : ", unixTime)
	//convert unix time ke human readable
	t := time.Unix(unixTime, 0)
	fmt.Println("human readable from epoch :", t)
	fmt.Println("time formmated :", t.Format("2006/01/02"))

}
