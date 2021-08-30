package model

import (
	"time"
)

type Sensitive struct {
	Id           int       `json:"id" xorm:"not null pk INT(8)"`
	Words        string    `json:"words" xorm:"not null default '' unique VARCHAR(32)"`
	Lastmodified time.Time `json:"lastmodified" xorm:"DATE"`
}

func (t Sensitive) TableName() string {
	return "sensitive"
}
