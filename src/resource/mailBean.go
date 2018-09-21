package resource

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Mail struct {
	gorm.Model `gorm:"primary_key:id;AUTO_INCREMEN"`
	JobName    string    `gorm:"type:varchar(255);not null;"`
	SendTime   time.Time `gorm:"column:send_time;default null"`
}
