package telegram

import (
	"sync"
	"telegram-api-service/internal/entitiy"
)

var TestInstance *Test
var onceTest sync.Once

type Test struct {
	State1 *entitiy.State
	State2 *entitiy.State
	State3 *entitiy.State
}

func (t *Test) Init() {
	t.State2 = &entitiy.State{}
	t.State1 = &entitiy.State{}
	t.State3 = &entitiy.State{}
}

func TestGet() *Test {
	if TestInstance == nil {
		onceTest.Do(func() {
			TestInstance = new(Test)
			TestInstance.Init()
		})
	}
	return TestInstance
}

type FSMContext struct {
	state *entitiy.State
}

func NewFSMContext() *FSMContext {
	return &FSMContext{state: nil}
}

func (fsm *FSMContext) SetState(state *entitiy.State) {
	fsm.state = state
}

func (fsm *FSMContext) GetState() *entitiy.State {
	return fsm.state
}

func (fsm *FSMContext) ClearState() {
	fsm.state = nil
}
