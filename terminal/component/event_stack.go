package component

import "sync"

type EventStack struct {
	stack  []Component
	locker sync.RWMutex
}

func NewEventStack() *EventStack {
	return &EventStack{}
}

func (e *EventStack) Push(comp Component) {
	e.locker.Lock()
	defer e.locker.Unlock()

	e.stack = append(e.stack, comp)
}

func (e *EventStack) Get() Component {
	e.locker.Lock()
	defer e.locker.Unlock()

	if len(e.stack) == 0 {
		return nil
	}

	comp := e.stack[0]
	e.stack = e.stack[1:]

	return comp
}

func (e *EventStack) Has() bool {
	e.locker.Lock()
	defer e.locker.Unlock()

	if len(e.stack) == 0 {
		return false
	}

	return true
}
