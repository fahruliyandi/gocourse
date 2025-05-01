package advance

import (
	"fmt"
	"time"
)

//=======jika buat ticker, maka harus buat stopnya==========//

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Prosesing something ", tick)
		case <-stop:
			fmt.Println("stoping ticker")
			return
		}
	}
}

//==============================
// func periodicTask() {
// 	fmt.Println("Performinc periodic task")
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }
//==============================

//==============================
// func main() {
// 	//scheduler pada go
// 	// ticker := time.NewTicker(time.Second)
// 	// for tick := range ticker.C {
// 	// 	fmt.Println("tick at : ", tick)
// 	// }

// 	// i := 0
// 	// for range ticker.C {
// 	// 	i++
// 	// 	fmt.Println(i)
// 	// }

// 	//stop ticker
// 	// i := 0
// 	// defer ticker.Stop()
// 	// for range 5 {
// 	// 	i++
// 	// 	fmt.Println(i)
// 	// }
// }
//==============================
