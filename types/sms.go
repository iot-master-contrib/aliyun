package types

import (
	"time"
)

type Sms struct {
	Id        int64     `json:"id"`
	Type      string    `json:"type,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Content   string    `json:"content,omitempty"`
	Result    string    `json:"result,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
}
