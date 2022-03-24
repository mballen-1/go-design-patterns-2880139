package main

// Define the interface for the observable type
type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

// The DataSubject will have a list of listeners
// and a field that gets changed, triggering them
type DataSubject struct {
	observers []DataListener
	field     string
}

// The ChangeItem function will cause the Listeners to be called
func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

// This function adds an observer to the list
func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

// This function removes an observer from the list
func (ds *DataSubject) unregisterObserver(o DataListener) {
	obs := make([]DataListener, 0)
	for _, ob := range ds.observers {
		if ob != o {
			obs = append(obs, ob)
		}
	}
	ds.observers = obs
}

// The notifyAll function calls all the listeners with the changed data
func (ds *DataSubject) notifyAll() {
	for _, o := range ds.observers {
		o.onUpdate(ds.field)
	}
}
