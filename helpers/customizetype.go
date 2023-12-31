package helpers

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type CustomTime struct {
	time.Time
}

func (CustomTime) GormDataType() string {
	return "time"
}

func (CustomTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "time"
}
