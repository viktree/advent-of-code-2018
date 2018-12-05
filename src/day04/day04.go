package day04

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"adventOfCode"
)

type Event struct {
	date    time.Time
	action  action
	guardID int
}

type action int

const (
	beginShift action = iota
	startSleeping
	stopSleeping
)

func parseDate(rawEvent string) time.Time {
	var y, m, d, hr, mn int
	n, err := fmt.Sscanf(rawEvent, "[%d-%d-%d %d:%d]", &y, &m, &d, &hr, &mn)
	if n < 5 || err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(2)
	}
	return time.Date(y, time.Month(m), d, hr, mn, 0, 0, time.UTC)
}

func setAction(message string) (int, action) {
	var guardID int
	var action action
	n, _ := fmt.Sscanf(message, "Guard #%d begins shift", &guardID)
	switch {
	case n == 1:
		action = beginShift
	case message == "falls asleep":
		action = startSleeping
	case message == "wakes up":
		action = stopSleeping
	default:
		fmt.Fprintf(os.Stderr, "ERROR: unknown action\n")
		os.Exit(2)
	}
	return guardID, action
}

func parseEvents(rawEvent []string) []Event {
	events := []Event{}
	rawEvents := adventOfCode.ReadInputFile("04", "input.txt")

	for _, rawEvent := range rawEvents {
		event := Event{guardID: -1}
		event.date = parseDate(rawEvent)

		i := strings.Index(rawEvent, "] ")
		message := rawEvent[i+2:]
		event.guardID, event.action = setAction(message)
		events = append(events, event)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].date.Before(events[j].date)
	})

	return events
}

func PartOne() {
	rawEvents := adventOfCode.ReadInputFile("04", "input.txt")
	events := parseEvents(rawEvents)

	var sleepyguard int
	MinsSleeping := make(map[int]int)
	var guard, minsStartSleeping int

	for _, event := range events {
		switch event.action {
		case beginShift:
			guard = event.guardID
		case startSleeping:
			minsStartSleeping = event.date.Minute()
		case stopSleeping:
			MinsSleeping[guard] += event.date.Minute() - minsStartSleeping
			if MinsSleeping[guard] > MinsSleeping[sleepyguard] {
				sleepyguard = guard
			}
		}
	}

	minutes := [60]int{}
	guard = -1
	var sleepyminute int

	for _, event := range events {
		if event.action == beginShift {
			guard = event.guardID
		} else if guard == sleepyguard {
			switch event.action {
			case startSleeping:
				minsStartSleeping = event.date.Minute()
			case stopSleeping:
				for i := minsStartSleeping; i < event.date.Minute(); i++ {
					minutes[i]++
					if minutes[i] > minutes[sleepyminute] {
						sleepyminute = i
					}
				}
			}
		}

	}

	fmt.Printf("Answer: %d\n", sleepyguard*sleepyminute)
}

type GuardSleepAtMinutePair struct {
	guardID int
	minute  int
}

func PartTwo() {

	rawEvents := adventOfCode.ReadInputFile("04", "input.txt")
	events := parseEvents(rawEvents)

	GuardSleepAtMinFreq := make(map[GuardSleepAtMinutePair]int)
	var guard, minsStartSleeping, maxFreq int
	var bestPair GuardSleepAtMinutePair

	for _, event := range events {
		switch event.action {
		case beginShift:
			guard = event.guardID
		case startSleeping:
			minsStartSleeping = event.date.Minute()
		case stopSleeping:
			for i := minsStartSleeping; i < event.date.Minute(); i++ {
				pair := GuardSleepAtMinutePair{guardID: guard, minute: i}
				GuardSleepAtMinFreq[pair]++
				if GuardSleepAtMinFreq[pair] > maxFreq {
					maxFreq = GuardSleepAtMinFreq[pair]
					bestPair = pair
				}
			}
		}
	}

	answer := bestPair.guardID * bestPair.minute

	fmt.Printf("Answer: %d\n", answer)
}
