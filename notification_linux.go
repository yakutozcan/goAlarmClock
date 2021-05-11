package main

import(
	"github.com/godbus/dbus/v5"
	"log"
)

// https://github.com/godbus/dbus/blob/master/_examples/notification.go
func Notification(title, subTitle, timeTitle string) {

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Fatal("Notification "+ err.Error())
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		title, subTitle, timeTitle, []string{},
		map[string]dbus.Variant{}, int32(5000))
	if call.Err != nil {
		log.Fatal("Notification "+ call.Err.Error())
	}
}
