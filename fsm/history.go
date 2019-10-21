package fsm

import "fmt"

// todo history 的处理是否需要lifecycle?
type History struct {
	fsm     *StateMachine
	states  []State
	maxSize int
	current int
}

func (h *History) String() string {
	return fmt.Sprint(h.states[:h.current+1])
}

func (h *History) Back() bool {
	if h.current > 0 {
		h.current -= 1
		h.fsm.current = h.states[h.current]
		return true
	}
	return false
}

func (h *History) Forward() bool {
	if h.current < len(h.states)-1 {
		h.current += 1
		h.fsm.current = h.states[h.current]
		return true
	}
	return false
}
