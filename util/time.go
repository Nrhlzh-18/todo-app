package util

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	strData := string(data)
	if strData == "null" {
		*jt = JsonTime(time.Time{})
		return nil
	}

	// Remove surrounding quotes if present
	if len(strData) >= 2 && strData[0] == '"' && strData[len(strData)-1] == '"' {
		strData = strData[1 : len(strData)-1]
	}

	t, err := time.Parse("15:04:05", strData)
	if err != nil {
		return err
	}

	*jt = JsonTime(t)
	return nil
}

func (jt JsonTime) MarshalJSON() ([]byte, error) {
	if time.Time(jt).IsZero() {
		return []byte("null"), nil
	}
	st := fmt.Sprintf("\"%s\"", time.Time(jt).Format("15:04:05"))
	return []byte(st), nil
}

func (jt JsonTime) Value() (driver.Value, error) {
	if time.Time(jt).IsZero() {
		return nil, nil
	}
	// Return the formatted time as a string
	return time.Time(jt).Format("15:04:05"), nil
}

func (jt *JsonTime) Scan(value interface{}) error {
	if value == nil {
		*jt = JsonTime(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*jt = JsonTime(v)
		return nil
	case string:
		t, err := time.Parse("15:04:05", v)
		if err != nil {
			return err
		}
		*jt = JsonTime(t)
		return nil
	default:
		return fmt.Errorf("JsonTime.Scan: unable to convert %v to time.Time", value)
	}
}

func (jt *JsonTime) GormDataType() string {
	return "time"
}

func (jt *JsonTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "time"
}

type JsonDateTime time.Time

func (jt *JsonDateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
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
