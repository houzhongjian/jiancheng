package base

import "time"

type Model struct {
	Id        int64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Create    string    `gorm:"-" json:"create"`
	Update    string    `gorm:"-" json:"update"`
	Page      int       `gorm:"-" form:"page" json:"page"`
	PageSize  int       `gorm:"-" form:"pagesize" json:"pagesize"`
}

func (x Model) Offset() int {
	if x.Page > 1 {
		return (x.Page - 1) * x.PageSize
	}
	return 0
}
