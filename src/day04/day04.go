package day04

import (
	"adventOfCode"
	"strings"
	"sort"
	"time"
	"fmt"
)

type Event struct{
	timestamp time.Time
	guardID int
	action string
}

func parseEvents(rawEvent string) Event {
	var event Event
	var year, month, day, hr, min int
	var rawAction string
	fmt.Sscanf(rawEvent,
		"[%d-%d-%d %d:%d] %s",
		&year,
		&month,
		&day,
		&hr,
		&min,
		&rawAction,
	)
	desc := strings.Split(rawEvent, "]")[1]
	event.timestamp = time.Date(
		year,
		time.Month(month),
		day,
		hr,
		min,
		1,
		1,
		time.UTC,
	)
	fmt.Sscanf(desc, "  Guard #%d begins shift", &event.guardID)
	if(event.guardID != 0){
		event.action = "START_SHIFT"
	} else {
		if(rawAction == "wakes"){
			event.action = "WAKES"
		} else {
			event.action = "SLEEPS"
		}
	}


	return event
}

func PartOne()  {
	rawEvents := adventOfCode.ReadInputFile("04", "input.txt")
	events := []Event{}

	for _, rawEvent := range rawEvents {
		event := parseEvents(rawEvent)
		events = append(events, event)
	}

	sort.Slice(events, func(i, j int) bool { return events[i].timestamp.Before(events[j].timestamp)})

	var guardID int
	var TimeMap map[int] time.Time
	for _, event := range(events) {
		event := *&event
		if(event.action == "START_SHIFT"){
			// What if multiple days?
			TimeMap[event.guardID] = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
		} else {
			if(event.action == "WAKES"){

			} else {

			}
		}

	}

}

func PartTwo()  {
	return
	fmt.Printf("Answer: %d \n", 42)
}