package main

func main() {
	// Construct two DataListener observers and
	// give each one a name
	listener1 := DataListener{
		Name: "Listener 1",
	}
	listener2 := DataListener{
		Name: "Listener 2",
	}

	// Create the DataSubject that the listeners will observe
	subj := &DataSubject{}
	// Register each listener with the DataSubject

	subj.registerObserver(listener1)
	subj.registerObserver(listener2)

	// Change the data in the DataSubject - this will cause the
	// onUpdate method of each listener to be called
	subj.ChangeItem("Colombia")
	subj.ChangeItem("Bolivia")

	// Unregister one of the observers
	subj.unregisterObserver(listener2)

	// change the data, again
	subj.ChangeItem("Venezuela")
}
