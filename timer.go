package main

import (
	"fmt"
	"time"
)

// ======multiple timer=======
func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	for range 2 {
		select {
		case <-timer1.C:
			fmt.Println("Timer1 expired")
		case <-timer2.C:
			fmt.Println("Timer2 expired")
		}
		time.Sleep(1 * time.Second)
	}
}

// =========DELAY OPERATION=========
// func main() {
// 	timer := time.NewTimer(3 * time.Second)
// 	go func() {
// 		<-timer.C
// 		fmt.Println("Operation Delayed")
// 	}()
// 	defer timer.Stop()

// 	fmt.Println("Waiting")
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("End of Program")

// }

//========TIME OUT==========
// func longRunnigOps() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(3 * time.Second)
// 	done := make(chan bool)
// 	go func() {
// 		longRunnigOps()
// 		done <- true
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Time out")
// 	case <-done:
// 		fmt.Println("Proses finish")
// 	}
// }

//========BASIC TIMER USER===========
// func main() {
// 	fmt.Println("Start App")
// 	timer := time.NewTimer(5 * time.Second)
// 	fmt.Println("Waiting for timer expired")
// 	stoped := timer.Stop()
// 	if stoped {
// 		fmt.Println("timer stop")
// 	}
// 	timer.Reset(1 * time.Second)
// 	fmt.Println("timer reset")
// 	<-timer.C
// 	fmt.Println("Timer expired")

// }
