package resource

import (
	"github.com/jinzhu/gorm"
)

type Job struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);not null;"`
	Context string `gorm:"type:varchar(255);not null;"`
}
