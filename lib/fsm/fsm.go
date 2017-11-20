package fsm

type Fsm struct {
	states map[string]string
	currentState string
}

func (fsm *Fsm) Init() *Fsm {
	fsm.states = map[string]string{
		"waiting": "waiting",
		"tagParsing": "tagParsing",
	}
	fsm.currentState = "waiting"
	return fsm
}
func (fsm *Fsm) GetCurrentState() string {
	return fsm.currentState
}

func (fsm *Fsm) SetCurrentState(state string) {
	if newState, ok := fsm.states[state]; ok {
		fsm.currentState = newState
	}
}

func Init() *Fsm {
	return new(Fsm).Init()
}
