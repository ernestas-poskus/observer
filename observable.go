package observer

// ObservableInterface - interface or abstract class defining the operations
// for attaching and de-attaching observers to the client.
// In the GOF book this class/interface is known as Subject.
type ObservableInterface interface {
	Attach(observer *Observer)
	Detach(observer *Observer)
	Notify(value ...interface{})
}

// Observable - concrete Observable class implements ObservableInterface.
// It maintains the state of the object and when a change in the
// state occurs it notifies the attached Observers.
type Observable []Observer

// Attach - adds new observers to observed object
func (observers *Observable) Attach(o Observer) {
	*observers = append(*observers, o)
}

// Detach - removes given observer from observed object
func (observers *Observable) Detach(o Observer) {
	newObservable := []Observer{}
	for _, observer := range *observers {
		if observer != o {
			newObservable = append(newObservable, observer)
		}
	}
	*observers = newObservable
}

// Notify - notifies all observers listening on object
func (observers *Observable) Notify(values ...interface{}) {
	for _, observer := range *observers {
		observer.Update(values)
	}
}
