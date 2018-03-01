package base

//Access 访问记录.
type Access struct {
	Model
	Ip string `json:"ip"`
}

//TableName Access.
func (Access) TableName() string {
	return "access"
}
