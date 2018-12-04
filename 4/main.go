package main

import (
	"fmt"
	"github.com/mathew/advent-of-code-2018"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const filePath = "4/data.txt"

type sleepyTime struct {
	start    time.Time
	end      time.Time
	duration time.Duration
}

type guardShift struct {
	guardId     string
	duration    int
	sleepyTimes []sleepyTime
}

type shifts []guardShift

func (s shifts) getSleepiestGuard() string {

	napTimes := map[string]int{}

	for _, gs := range s {
		napTimes[gs.guardId] += gs.duration
	}

	var naughtyGuard string
	for guardId, duration := range napTimes {
		if naughtyGuard == "" {
			naughtyGuard = guardId

			continue
		}

		if napTimes[naughtyGuard] < duration {
			naughtyGuard = guardId
		}
	}

	return naughtyGuard
}

func (s shifts) getSleepiestMinuteForGuard(guardId string) int {
	napMinutes := map[int]int{}

	for _, gs := range s {

		if gs.guardId != guardId {
			continue
		}

		for _, st := range gs.sleepyTimes {
			for i := st.start.Minute(); i < st.end.Minute(); i++ {
				napMinutes[i]++
			}
		}
	}

	sleepiestMinute := 0
	for minute, duration := range napMinutes {
		if sleepiestMinute == 0 {
			sleepiestMinute = minute
			continue
		}

		if duration > napMinutes[sleepiestMinute] {
			sleepiestMinute = minute
		}

	}

	return sleepiestMinute
}

func (s shifts) getSleepiestMinuteWithGuard() (string, int)  {
	napMinutes := map[string]map[int]int{}

	for _, gs := range s {
		for _, st := range gs.sleepyTimes {
			for i := st.start.Minute(); i < st.end.Minute(); i++ {

				if _, ok := napMinutes[gs.guardId]; !ok {
					napMinutes[gs.guardId] = map[int]int{}
				}

				napMinutes[gs.guardId][i]++
			}
		}
	}

	sleepiestMinute := 0
	sleepiestTimes := 0
	sleepiestGuardId := ""

	for guardId, minutesToCount := range napMinutes {

		for minute, times := range minutesToCount {
			if sleepiestMinute == 0 || times > sleepiestTimes {
				sleepiestMinute = minute
				sleepiestTimes = times
				sleepiestGuardId = guardId
			}
		}


	}

	return sleepiestGuardId, sleepiestMinute
}

func parseGuardShiftNotes(rawGuardShiftNotes string) shifts {
	unparsedGuardShiftNotes := strings.Split(rawGuardShiftNotes, "\n")

	guardNotesMap := map[time.Time]string{}
	var keys []time.Time

	for _, unparsedGuardShiftNote := range unparsedGuardShiftNotes {
		dt := parseGuardShiftNoteDateTime(unparsedGuardShiftNote)
		guardNotesMap[dt] = unparsedGuardShiftNote
		keys = append(keys, dt)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	return calculateShifts(keys, guardNotesMap)
}

func calculateShifts(orderedGuardShifts []time.Time, guardShiftNotes map[time.Time]string) shifts {
	shifts := shifts{}
	sleepyTimes := []sleepyTime{}
	shift := guardShift{}
	asleepTime := sleepyTime{}

	for _, k := range orderedGuardShifts {
		switch strings.TrimSpace(strings.SplitAfter(guardShiftNotes[k], "]")[1]) {

		case "wakes up":
			asleepTime.end = k
			asleepTime.duration = asleepTime.end.Sub(asleepTime.start)
			sleepyTimes = append(sleepyTimes, asleepTime)

		case "falls asleep":
			asleepTime = sleepyTime{
				start: k,
			}

		default:
			// deal with previous if they exist
			if shift.guardId != "" && sleepyTimes != nil {
				shift.sleepyTimes = sleepyTimes

				sleepy := 0
				for _, st := range sleepyTimes {
					sleepy += int(st.duration)
				}
				shift.duration = sleepy

				shifts = append(shifts, shift)
			}

			guardIdRe := regexp.MustCompile(`#[0-9]+`)

			shift = guardShift{
				guardId: guardIdRe.FindAllString(guardShiftNotes[k], 1)[0],
			}
			sleepyTimes = []sleepyTime{}
		}

		shift.sleepyTimes = sleepyTimes
	}

	return shifts
}

func parseGuardShiftNoteDateTime(unparsedGuardShift string) time.Time {
	squareBracketRe := regexp.MustCompile(`\[[0-9-:\s]+\]`)
	dt := squareBracketRe.FindAllString(unparsedGuardShift, 1)[0]
	dt = strings.Replace(dt, "[", "", -1)
	dt = strings.Replace(dt, "]", "", -1)

	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, dt)

	if err != nil {
		fmt.Println(err)
	}

	return t
}

func main() {
	rawGuardShifts := files.LoadFile(filePath)

	shifts := parseGuardShiftNotes(rawGuardShifts)
	sleepiestGuard := shifts.getSleepiestGuard()
	sleepiestMinute := shifts.getSleepiestMinuteForGuard(sleepiestGuard)
	intGuard, _ := strconv.Atoi(strings.TrimSpace(strings.Split(sleepiestGuard, "#")[1]))

	fmt.Println("Sleepiest Guard: ", sleepiestGuard, "Sleepiest Minute: ", sleepiestMinute, "Answer:", intGuard*sleepiestMinute)

	sleepiestGuard, sleepiestMinute = shifts.getSleepiestMinuteWithGuard()
	intGuard, _ = strconv.Atoi(strings.TrimSpace(strings.Split(sleepiestGuard, "#")[1]))

	fmt.Println("Sleepiest Guard: ", sleepiestGuard, "Sleepiest Minute: ", sleepiestMinute, "Answer:", intGuard*sleepiestMinute)
}
