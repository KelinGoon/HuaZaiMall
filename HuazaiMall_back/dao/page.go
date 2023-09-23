package dao

type Pagination struct {
	PageNum  int
	PageSize int
}

// SitePaginationSelectModel 分页的模型，用于接收前端的数据
type SitePaginationSelectModel struct {
	UserId int `json:"UserID"`
	Pagination
}
