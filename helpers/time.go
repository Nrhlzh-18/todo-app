package helpers

import (
	"errors"
	"fmt"
	"time"
)

func Appends[T any](slices [][]T) []T {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int
	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

type JsonDate time.Time

func (jd *JsonDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// TODO(https://go.dev/issue/47353): Properly unescape a JSON string.
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Time.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	t, err := time.ParseInLocation("02/01/2006", string(data), time.Local)
	*jd = JsonDate(t)
	return err
}

func (jd JsonDate) MarshalJSON() ([]byte, error) {
	st := fmt.Sprintf("\"%s\"", time.Time(jd).Format("02/01/2006"))
	return []byte(st), nil
}

type JsonTime time.Time

func (jt *JsonTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// TODO: Properly unescape a JSON string.
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("JsonTime.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	t, err := time.Parse("15:04:05", string(data))
	*jt = JsonTime(t)
	return err
}

func (jt JsonTime) MarshalJSON() ([]byte, error) {
	st := fmt.Sprintf("\"%s\"", time.Time(jt).Format("15:04:05"))
	return []byte(st), nil
}

type JsonDateTime time.Time

func (jt *JsonDateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// TODO(https://go.dev/issue/47353): Properly unescape a JSON string.
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Time.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	t, err := time.ParseInLocation("02/01/2006 15:04:05", string(data), time.Local)
	*jt = JsonDateTime(t)
	return err
}

func (jt JsonDateTime) MarshalJSON() ([]byte, error) {
	st := fmt.Sprintf("\"%s\"", time.Time(jt).Format("02/01/2006 15:04:05"))
	return []byte(st), nil
}
