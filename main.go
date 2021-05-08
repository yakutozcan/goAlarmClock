package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"
)

type Alarms struct {
	Alarms []Alarm `json:"alarms"`
}

type Alarm struct {
	AlarmDateTime string `json:"alarm_date_time"`
	AlarmTitle    string `json:"alarm_title"`
	AlarmSubTitle string `json:"alarm_sub_title"`
}

func main() {
	alarmFile := flag.String("c", "alarms.json", "Input alarm filename: alarm.json")
	jsonFile, err := os.Open(*alarmFile)
	Notification("Running", "Hello SubTitle", "Hi! Title")
	if err != nil {
		log.Fatal("Can not open file \n" +
			"Input alarm filename: alarms.json " + err.Error())
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var alarms Alarms
	err = json.Unmarshal(byteValue, &alarms)
	if err != nil {
		log.Fatalf("Unmarshaling alarms failed: %s", err.Error())
	}

	//create a  buffered channel for interupt signal
	is := make(chan os.Signal, 1)
	//notify this channel on interupt
	signal.Notify(is, os.Interrupt)

	//create a channel for done signal
	ds := make(chan struct{})

	//create ticker
	systemTimeTicker := time.NewTicker(time.Second * 1)

	//start main loop
	go func() {
		defer close(ds)
		for {
			select {
			case <-systemTimeTicker.C:
				localTime := time.Now().Format("15:04 02.01.2006")
				for i := 0; i < len(alarms.Alarms); i++ {
					if localTime == alarms.Alarms[i].AlarmDateTime {
						Notification(alarms.Alarms[i].AlarmTitle, alarms.Alarms[i].AlarmSubTitle, "⏰ "+alarms.Alarms[i].AlarmDateTime+" ⏰")
						alarms.Alarms[i].AlarmDateTime = time.Now().AddDate(0, 0, 1).Format("15:04 02.01.2006")
						file, _ := json.MarshalIndent(alarms, "", " ")
						writeErr := ioutil.WriteFile("alarm.json", file, 0644)
						if writeErr != nil {
							log.Fatal("Can not write file " + err.Error())
						}
					}
				}
			case <-ds:
				return
			}
		}
	}()

	//wait for interupt signal
	<-is

	//close interupt signal channel
	close(is)

	//stop ticker
	systemTimeTicker.Stop()

	//send done signal
	ds <- struct{}{}

	//wait until done signal channel is closed
	<-ds
}
