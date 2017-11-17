package ctrl

type fsm struct {
	states map[string]string
	currentState string
}

func (c *fsm) Init() *fsm {
	c.states = map[string]string{
		"init": "init",
	}
	c.currentState = "init"
	return c
}
func (c *fsm) GetCurrentState() string {
	return c.currentState
}

func Init() *fsm {
	return new(fsm).Init()
}
