package advance

import (
	"fmt"
)

type StateFull struct {
	ch    chan int
	count int
}

func (s *StateFull) start() {
	go func() {
		// for {
		// 	select {
		// 	case value := <-s.ch:
		// 		s.count += value
		// 		fmt.Println(s.count)
		// 	}
		// }

		for val := range s.ch {
			s.count += val
			fmt.Println(s.count)
		}

	}()
}

func (s *StateFull) send(value int) {
	s.ch <- value
}

func main() {

	sf := &StateFull{
		ch: make(chan int),
	}

	sf.start()

	for i := range 5 {
		sf.send(i)
	}

	close(sf.ch)
}
