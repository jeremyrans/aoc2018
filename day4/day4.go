package day4

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

type action string

const (
	START action = "start"
	SLEEP action = "sleep"
	WAKE  action = "wake"
)

type entry struct {
	t       time.Time
	guardID int
	action  action
}

func part1(input string) int {
	var entries []entry
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		entries = append(entries, parseEntry(l))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].t.Before(entries[j].t)
	})

	guardSleepMinutes := map[int][]int{}
	var guard int
	var sleepFrom time.Time
	var sleepTo time.Time
	for _, e := range entries {
		switch e.action {
		case START:
			guard = e.guardID
			sleepFrom = time.Time{}
			sleepTo = time.Time{}
			break
		case SLEEP:
			sleepFrom = e.t
			break
		case WAKE:
			sleepTo = e.t
			for i := sleepFrom.Minute(); i < sleepTo.Minute(); i++ {
				guardSleepMinutes[guard] = append(guardSleepMinutes[guard], i)
			}
			break
		}
	}

	// find sleepiest guard
	sleepiestGuard := 0
	sleepiestGuardMinutes := 0
	for guardID, sleepMinutes := range guardSleepMinutes {
		if len(sleepMinutes) > sleepiestGuardMinutes {
			sleepiestGuardMinutes = len(sleepMinutes)
			sleepiestGuard = guardID
		}
	}

	// find most common sleep minute
	minuteCounts := map[int]int{}
	for _, m := range guardSleepMinutes[sleepiestGuard] {
		minuteCounts[m]++
	}
	bestMinute := -1
	bestMinuteCount := -1
	for m, c := range minuteCounts {
		if c > bestMinuteCount {
			bestMinuteCount = c
			bestMinute = m
		}
	}

	return sleepiestGuard * bestMinute
}

// Example entry:
// [1518-11-01 00:00] Guard #10 begins shift
func parseEntry(s string) entry {
	e := entry{}
	parts := strings.Split(s, " ")
	e.t, _ = time.Parse("2006-01-0215:04", strings.Trim(parts[0]+parts[1], "[]"))
	if len(parts) > 4 {
		e.action = START
		e.guardID, _ = strconv.Atoi(strings.TrimPrefix(parts[3], "#"))

	} else {
		if parts[2] == "falls" {
			e.action = SLEEP
		} else {
			e.action = WAKE
		}
	}
	return e
}

func part2(input string) int {
	var entries []entry
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		entries = append(entries, parseEntry(l))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].t.Before(entries[j].t)
	})

	guardSleepMinutes := map[int][]int{}
	var guard int
	var sleepFrom time.Time
	var sleepTo time.Time
	for _, e := range entries {
		switch e.action {
		case START:
			guard = e.guardID
			sleepFrom = time.Time{}
			sleepTo = time.Time{}
			break
		case SLEEP:
			sleepFrom = e.t
			break
		case WAKE:
			sleepTo = e.t
			for i := sleepFrom.Minute(); i < sleepTo.Minute(); i++ {
				guardSleepMinutes[guard] = append(guardSleepMinutes[guard], i)
			}
			break
		}
	}

	sleepiestMinute := -1
	sleepiestMinuteCount := -1
	sleepiestGuard := -1
	minuteCounts := map[int]map[int]int{} // minute: { guard: count }
	for g, mins := range guardSleepMinutes {
		for _, m := range mins {
			_, ok := minuteCounts[m]
			if !ok {
				minuteCounts[m] = map[int]int{}
			}
			minuteCounts[m][g]++
			if minuteCounts[m][g] > sleepiestMinuteCount {
				sleepiestMinuteCount = minuteCounts[m][g]
				sleepiestMinute = m
				sleepiestGuard = g
			}
		}
	}

	return sleepiestGuard * sleepiestMinute
}
