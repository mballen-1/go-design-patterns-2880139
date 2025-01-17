package main

import "fmt"

// The NotificationBuilder has fields exported
type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.SubTitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(pri int) {
	nb.Priority = pri
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.NotType = notType
}

// The Build method returns a fully finished Notification object
func (nb *NotificationBuilder) Build() (*Notification, error) {
	// Error checking can be done at the Build stage
	if nb.Title == "" {
		return nil, fmt.Errorf("title is nil")
	} else {
		// Return a newly created Notification object using the current settings
		return &Notification{
			title:    nb.Title,
			subtitle: nb.SubTitle,
			message:  nb.Message,
			image:    nb.Image,
			icon:     nb.Icon,
			priority: nb.Priority,
			notType:  nb.NotType,
		}, nil
	}
	return nil, nil
}
