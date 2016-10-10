package vm

import "github.com/PMoneda/hub/lang"

//Context is execution context instance
type Context struct {
	Outer *Context
	Store map[string]lang.Object
}
