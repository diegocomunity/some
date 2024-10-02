package some

import "fmt"

type some struct {
}

func (s *some) Message(m string) {
	fmt.Printf(m)

}

func New() *some {
	return new(some)
}
