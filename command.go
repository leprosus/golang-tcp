package tcp

import "fmt"

type Command struct {
	Object string
	Action string
}

func (com *Command) String() string {
	return fmt.Sprintf("%s/%s", com.Object, com.Action)
}
