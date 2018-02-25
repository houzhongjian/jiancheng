package base

//Menu  栏目.
type Menu struct {
	Model
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	IsShow bool   `json:"is_show"`
}

//TableName Menu.
func (Menu) TableName() string {
	return "menu"
}
