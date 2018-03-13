package base

//Menu  栏目.
type Menu struct {
	Model
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	IsShow bool   `json:"is_show"`
	PId    int64  `json:"p_id"`
}

//TableName Menu.
func (Menu) TableName() string {
	return "menu"
}

//QueryMenu .
type QueryMenu struct {
	Menu
	MenuList []*Menu `json:"menu_list"`
}
