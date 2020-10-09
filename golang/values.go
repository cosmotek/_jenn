package golang

import (
	"fmt"
	"time"
)

var currentTimestamp = time.Now().UTC()

type Value struct {
	Literal string
}

var (
	EmptyTime = Value{
		Literal: "time.Time{}",
	}

	TimeNow = Value{
		Literal: fmt.Sprintf(
			"time.Date(%d, %d, %d, %d, %d, %d, %d, time.UTC)",
			currentTimestamp.Year(),
			currentTimestamp.Month(),
			currentTimestamp.Day(),
			currentTimestamp.Hour(),
			currentTimestamp.Minute(),
			currentTimestamp.Second(),
			currentTimestamp.Nanosecond(),
		),
	}

	Zero = Value{
		Literal: "0",
	}

	EmptyString = Value{
		Literal: "\"\"",
	}

	Nil = Value{
		Literal: "nil",
	}
)

func PointInTimeUTC(point time.Time) Value {
	return Value{
		Literal: fmt.Sprintf(
			"time.Date(%d, %d, %d, %d, %d, %d, %d, time.UTC)",
			point.Year(),
			point.Month(),
			point.Day(),
			point.Hour(),
			point.Minute(),
			point.Second(),
			point.Nanosecond(),
		),
	}
}
