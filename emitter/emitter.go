package emitter

import (
	"log"
	"math"
	"reflect"
)

const NewListener Event = math.MaxUint32
const RemoveListener Event = math.MaxUint32 - 1

type Event uint

type listener struct {
	callback func(args ...interface{})
	once     bool
}

type Emitter struct {
	listeners map[Event][]listener
	maxCount  uint
}

func New() *Emitter {
	return &Emitter{
		listeners: map[Event][]listener{},
		maxCount:  10,
	}
}

func (e *Emitter) SetMaxListeners(n uint) {
	e.maxCount = n
}

// 添加 listener 函数到名为 eventName 的事件的监听器数组的末尾。 不会检查 listener 是否已被添加。 多次调用并传入相同的 eventName 与 listener 会导致 listener 会被添加多次。
func (e *Emitter) On(event Event, callback func(args ...interface{})) *Emitter {
	e.listeners[event] = append(e.listeners[event], listener{callback: callback})
	e.Emit(NewListener, callback)
	if e.maxCount > 0 && len(e.listeners[event]) > int(e.maxCount) {
		log.Println("listener reach maxCount")
	}
	return e
}

func (e *Emitter) Once(event Event, callback func(args ...interface{})) *Emitter {
	e.listeners[event] = append(e.listeners[event], listener{callback: callback, once: true})
	e.Emit(NewListener, callback)
	return e
}

// Off() 最多只会从监听器数组中移除一个监听器。 如果监听器被多次添加到指定 eventName 的监听器数组中，则必须多次调用 Off() 才能移除所有实例。
// 一旦事件被触发，所有绑定到该事件的监听器都会按顺序依次调用。 这意味着，在事件触发之后、且最后一个监听器执行完成之前， Off() 或 RemoveAllListeners() 不会从 emit() 中移除它们。
func (e *Emitter) Off(event Event, callback func(...interface{})) *Emitter {
	return e.removeListenerInternal(event, callback, true)
}

func (e *Emitter) removeListenerInternal(event Event, callback func(...interface{}), needEmitRemoveEvent bool) *Emitter {
	if listeners, ok := e.listeners[event]; !ok {
		return e
	} else {
		for k, v := range listeners {
			if reflect.ValueOf(v.callback).Pointer() == reflect.ValueOf(callback).Pointer() {
				e.listeners[event] = append(e.listeners[event][:k], e.listeners[event][k+1:]...)
				if needEmitRemoveEvent {
					e.Emit(RemoveListener, callback)
				}
				return e
			}
		}
		return e
	}
}

func (e *Emitter) RemoveAllListeners(events ...Event) *Emitter {
	if len(events) == 0 {
		e.listeners = map[Event][]listener{}
		return e
	}
	for _, value := range events {
		if _, ok := e.listeners[value]; ok {
			delete(e.listeners, value)
		}
	}
	return e
}

func (e *Emitter) Emit(event Event, args ...interface{}) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Emitter error:", err)
		}
	}()
	var ret bool
	for _, l := range e.listeners[event] {
		if l.once {
			e.removeListenerInternal(event, l.callback, false)
		}
		ret = true
		l.callback(args...)
	}
	return ret
}
