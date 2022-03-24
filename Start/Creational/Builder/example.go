package main

import "fmt"

func main() {
	// Create a NotificationBuilder and use it to set properties
	nb := newNotificationBuilder()

	// Use the builder to set some properties
	nb.SetImage("people running around a block")
	nb.SetIcon("people_running")
	nb.SetTitle("imageTitle.png")

	// Use the Build function to create a finished object
	obj, err := nb.Build()
	if err != nil {
		fmt.Println(fmt.Errorf("there was an error => %v", err))
	} else {
		fmt.Println("Final Object is \n", obj)
	}
}
