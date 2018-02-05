package base

//Conf 结构体.
type Conf struct {
	DbName      string `json:"db_name"`
	DbHost      string `json:"db_host"`
	DbPort      string `json:"db_port"`
	DbUser      string `json:"db_user"`
	DbPassword  string `json:"db_password"`
	WebsitePost string `json:"website_port"`
}
