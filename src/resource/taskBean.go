package resource

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	JobName string `gorm:"type:varchar(255);not null;"`
	Name    string `gorm:"type:varchar(255);not null;"`
	Message string `gorm:"type:varchar(255);not null;"`
}
