package main

import (
	"gopkg.in/toast.v1"
	"log"
)

func Notification(title, subTitle, timeTitle string) {
	notification := toast.Notification{
		AppID: title,
		Title: subTitle,
		Message: timeTitle,
	}
	err := notification.Push()
	if err != nil {
		log.Fatal("Notification " + err.Error())
	}
}