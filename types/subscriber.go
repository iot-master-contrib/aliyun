package types

import (
	"time"
)

type Subscriber struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Level     uint      `json:"level,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
}
