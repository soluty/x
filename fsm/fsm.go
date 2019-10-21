package fsm

import (
	"errors"
	"strings"
)

/*
现态：是指当前所处的状态。

条件：又称为事件。当一个条件被满足，可能将会触发一个动作，或者执行一次状态的迁移。

动作：条件满足后执行的动作行为。动作执行完毕后，可以迁移到新的状态，也可以仍旧保持原状态。动作不是必需的，当条件满足后，也可以不执行任何动作，直接迁移到新状态。

*次态：条件满足后要迁往的新状态。“次态”是相对于“现态”而言的，“次态”一旦被激活，就转变成新的“现态”了。

摩尔型有限状态机（Moore机），输出只依赖于当前状态。即：
次态 = f(现态,输入)，输出 = f(现态)

米利型有限状态机（Mealy机），输出依赖于当前状态和输入。即：
次态 = f(现态,输入)，输出 = f(现态,输入)
*/

type (
	CallbackKey string
	State       string
	Event       string
	States      []State
)

func (s States) GetState() []State {
	return s
}

const BeforeAllEvent = CallbackKey("before_")
const LeaveAllState = CallbackKey("leave_")
const EnterAllState = CallbackKey("enter_")
const AfterAllEvent = CallbackKey("after_")
const invalid = State("")

func (s State) Enter() CallbackKey {
	return CallbackKey("enter_" + string(s))
}

func (s State) Leave() CallbackKey {
	return CallbackKey("leave_" + string(s))
}

func (s State) isValid() bool {
	return strings.TrimSpace(string(s)) != "" && !strings.Contains(string(s), "*")
}

func (s State) String() string {
	return string(s)
}

func (s Event) Before() CallbackKey {
	return CallbackKey("before_" + string(s))
}

func (s Event) After() CallbackKey {
	return CallbackKey("after_" + string(s))
}

func (s Event) String() string {
	return string(s)
}

type Stater interface {
	GetState() []State
	// Prefix() string todo 分层状态机
}

func (s State) GetState() []State {
	return []State{s}
}

type (
	Transition struct {
		Name Event
		From Stater
		To   State
	}
	LifeCycle struct {
		Transition Event
		From       State
		To         State
		Fsm        *StateMachine
	}
	// type ActionMoore func(l LifeCycle) (interface{}, error)  Moore机 输出和args无关
	Action  func(l LifeCycle, args ...interface{}) (interface{}, error)
	Options struct {
		Init        State
		Transitions []Transition
		Callbacks   map[CallbackKey]Action
	}
)

type (
	// StateMachine
	StateMachine struct {
		current State
		states  []State

		transitions map[eKey]State
		callbacks   map[CallbackKey][]Action

		history *History
	}
	eKey struct {
		event Event
		src   State
	}
)

func New(options *Options, historyMax ...int) *StateMachine {
	if !options.Init.isValid() {
		panic("初始状态不合法，必须不为空并且不包含*号")
	}
	fsm := &StateMachine{
		current:     options.Init,
		transitions: map[eKey]State{},
		callbacks:   map[CallbackKey][]Action{},
	}

	allEvents := make(map[string]bool)
	allStates := make(map[string]bool)
	for _, e := range options.Transitions {
		for _, value := range e.From.GetState() {
			if _, ok := fsm.transitions[eKey{e.Name, value}]; ok {
				panic("event和src重复")
			}
			fsm.transitions[eKey{e.Name, value}] = e.To
			allStates[value.String()] = true
			allStates[e.To.String()] = true
			allEvents[e.Name.String()] = true
			fsm.addState(value)
		}
		fsm.addState(e.To)
	}
	for key, value := range options.Callbacks {
		fsm.callbacks[key] = append(fsm.callbacks[key], value)
	}
	if len(historyMax) > 0 {
		fsm.history = &History{
			fsm:     fsm,
			states:  []State{options.Init},
			current: 0,
			maxSize: historyMax[0],
		}
	}
	return fsm
}

// todo 分层有限状态机
func (fsm *StateMachine) GetState() []State {
	//return fsm.current
	return nil
}

func (fsm *StateMachine) Fire(event Event, args ...interface{}) (output interface{}, err error) {
	to, err := fsm.seek(event)
	if err != nil {
		return nil, err
	}
	return fsm.transit(event, fsm.current, to, args)
}

func (fsm *StateMachine) State() State {
	return fsm.current
}

func (fsm *StateMachine) Is(s State) bool {
	return fsm.current == s
}

func (fsm *StateMachine) Can(e Event) bool {
	_, err := fsm.seek(e)
	return err != nil
}

func (fsm *StateMachine) Cannot(e Event) bool {
	return !fsm.Can(e)
}

func (fsm *StateMachine) Transitions() []Event {
	var ret []Event
	for key := range fsm.transitions {
		if key.src == fsm.current && !containsEvent(ret, key.event) {
			ret = append(ret, key.event)
		}
	}
	return ret
}

func (fsm *StateMachine) AllTransitions() []Event {
	var ret []Event
	for key := range fsm.transitions {
		if !containsEvent(ret, key.event) {
			ret = append(ret, key.event)
		}
	}
	return ret
}

func (fsm *StateMachine) AllStates() []State {
	return fsm.states
}

func (fsm *StateMachine) Observe(m map[CallbackKey]Action) {
	for key, value := range m {
		fsm.callbacks[key] = append(fsm.callbacks[key], value)
	}
}

func (fsm *StateMachine) History() *History {
	return fsm.history
}

func (fsm *StateMachine) addState(s State) {
	if !containsState(fsm.states, s) {
		fsm.states = append(fsm.states, s)
	}
}

func (fsm *StateMachine) seek(e Event) (State, error) {
	dst, ok := fsm.transitions[eKey{e, fsm.current}]
	if !ok {
		for key := range fsm.transitions {
			if key.event == e {
				return invalid, errors.New("invalid")
			}
		}
		return invalid, errors.New("invalid")
	}
	return dst, nil
}

func (fsm *StateMachine) transit(transition Event, from, to State, args ...interface{}) (output interface{}, err error) {
	changed := from != to
	l := &LifeCycle{
		Transition: transition,
		From:       from,
		To:         to,
		Fsm:        fsm,
	}
	if fn, ok := fsm.callbacks[BeforeAllEvent]; ok {
		for _, fn := range fn {
			if output, err = fn(*l, args...); err != nil {
				return
			}
		}
	}
	if fn, ok := fsm.callbacks[transition.Before()]; ok {
		for _, fn := range fn {
			if output, err = fn(*l, args...); err != nil {
				return
			}
		}
	}
	if changed {
		if fn, ok := fsm.callbacks[LeaveAllState]; ok {
			for _, fn := range fn {
				if output, err = fn(*l, args...); err != nil {
					return
				}
			}
		}
		if fn, ok := fsm.callbacks[from.Leave()]; ok {
			for _, fn := range fn {
				if output, err = fn(*l, args...); err != nil {
					return
				}
			}
		}
	}
	fsm.current = to
	if fsm.history != nil {
		fsm.history.states = append(fsm.history.states, to)
		fsm.history.current++
		if fsm.history.maxSize > 0 && len(fsm.history.states) > fsm.history.maxSize {
			fsm.history.current--
			fsm.history.states = fsm.history.states[1:]
		}
	}
	if changed {
		if fn, ok := fsm.callbacks[EnterAllState]; ok {
			for _, fn := range fn {
				if output, err = fn(*l, args...); err != nil {
					return
				}
			}
		}
		if fn, ok := fsm.callbacks[to.Enter()]; ok {
			for _, fn := range fn {
				if output, err = fn(*l, args...); err != nil {
					return
				}
			}
		}
	}
	if fn, ok := fsm.callbacks[AfterAllEvent]; ok {
		for _, fn := range fn {
			if output, err = fn(*l, args...); err != nil {
				return
			}
		}
	}
	if fn, ok := fsm.callbacks[transition.After()]; ok {
		for _, fn := range fn {
			if output, err = fn(*l, args...); err != nil {
				return
			}
		}
	}
	return
}

func containsEvent(transitions []Event, e Event) bool {
	for _, value := range transitions {
		if value == e {
			return true
		}
	}
	return false
}

func containsState(transitions []State, e State) bool {
	for _, value := range transitions {
		if value == e {
			return true
		}
	}
	return false
}
