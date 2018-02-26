package base

//Banner 广告图.
type Banner struct {
	Model
	Path   string `json:"path"`
	IsShow bool   `json:"is_show"`
	Sort   int64  `json:"sort"`
}
