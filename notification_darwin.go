package main

import notifier "github.com/deckarep/gosx-notifier"

func Notification(title, subTitle, timeTitle string) {

	note := notifier.NewNotification(timeTitle)

	note.Title = title
	note.Subtitle = subTitle

	note.Sound = notifier.Submarine

	note.Group = "com.alarm.clock.identifier"
	note.Sender = "com.apple.Stickies"

	err := note.Push()
	if err != nil {
		log.Fatal("Notification " + err.Error())
	}
}
