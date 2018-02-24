package base

type Model struct {
	Id        int64 `gorm:"primary_key" json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	Page      int   `gorm:"-" form:"page" json:"page"`
	PageSize  int   `gorm:"-" form:"pagesize" json:"pagesize"`
}

func (x Model) Offset() int {
	if x.Page > 1 {
		return (x.Page - 1) * x.PageSize
	}
	return 0
}
