package utc

import "time"

/*
A package for literally just putting UTC in the format I like, that's it. Honestly.
*/

const utcformat = "02-01-2006 15:04:05"

func Now() string {

	time := time.Now().UTC().Format(utcformat)

	return time
}

func Stamp2String(timeStamp time.Time) string {

	return timeStamp.UTC().Format(utcformat)

}
