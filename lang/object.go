//Package lang contains all object definitions for hub
package lang

//Object is the basic interface for all objects in hub
type Object interface {
	GetHash() int64
	GetType() string
	Equals(obj Object) bool
}
