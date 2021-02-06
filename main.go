package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
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
	_ = json.Unmarshal(byteValue, &alarms)

	systemTimeTicker := time.NewTicker(time.Second * 1)
	go func() {
		for range systemTimeTicker.C {
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
		}
	}()

	wait := make(chan bool)

	<-wait
}
