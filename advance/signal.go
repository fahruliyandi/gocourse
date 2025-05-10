package advance

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Working...")

	go func() {
		sig := <-sigs
		fmt.Println("Signal Receive", sig)
		fmt.Println("Gracefully Shutdown")
		os.Exit(0)
	}()

	for {
		time.Sleep(time.Second)
	}

}
