package observer

// Observer - interface or abstract class defining the operations
// to be used to notify this object.
type Observer interface {
	Update(value ...interface{})
}
