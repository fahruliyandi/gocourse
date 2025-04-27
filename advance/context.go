package advance

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation Cancelled :", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func logWithContext(ctx context.Context, msg string) {
	reqId := ctx.Value("requestId")
	log.Printf("LOG: %v-%v", reqId, msg)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "requestId", "12345")

	go doWork(ctx)

	time.Sleep(5 * time.Second)
	reqId := ctx.Value("requestId")
	if reqId != nil {
		fmt.Println("Request Id", reqId)
	} else {
		fmt.Println("No ReqId Found")
	}

	logWithContext(ctx, "Tes loging with")
}

// func CheckEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation Cancelled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()
// 	result := CheckEvenOdd(ctx, 5)
// 	fmt.Println("result with context.TODO()", result)

// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancel()

// 	result = CheckEvenOdd(ctx, 10)
// 	fmt.Println("Result from timeout context ", result)

// 	time.Sleep(2 * time.Second)
// 	result = CheckEvenOdd(ctx, 15)
// 	fmt.Println("Result after timeout ", result)

// }

// func main() {

// 	todoCtx := context.TODO()
// 	ctx1 := context.WithValue(todoCtx, "name", "jhon")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("name"))

// 	bgCtx := context.Background()
// 	ctx2 := context.WithValue(bgCtx, "city", "new york")
// 	fmt.Println(ctx2)
// 	fmt.Println(ctx2.Value("city"))

// }
