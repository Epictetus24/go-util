package utc

import "time"

/*
A package for literally just putting UTC in the format I like, that's it. Honestly.
*/

const UTCFormat = "02-01-2006 15:04:05"

func Now() string {

	time := time.Now().UTC().Format(UTCFormat)

	return time
}

func Stamp2String(timeStamp time.Time) string {

	return timeStamp.UTC().Format(UTCFormat)

}
