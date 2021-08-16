package model

import (
	"time"
)

type Portion struct {
	Id           int       `json:"id" xorm:"not null pk INT(8)"`
	Word         string    `json:"word" xorm:"unique VARCHAR(8)"`
	Idx          string    `json:"idx" xorm:"VARCHAR(8)"`
	Category     string    `json:"category" xorm:"VARCHAR(8)"`
	Lastmodified time.Time `json:"lastmodified" xorm:"DATE"`
}

func (t Portion) TableName() string {
	return "portion"
}
