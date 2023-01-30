package common

// Pagination 分页请求（新）
type Pagination struct {
	Page   int `json:"page" form:"page" db:"-"`
	Offset int `json:"-" form:"-" db:"offset"`
	Size   int `json:"size" form:"size" db:"size"`
}

// HandleOffset 初始化页码，计算偏移量
func (p *Pagination) HandleOffset() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Size <= 0 {
		p.Size = 0
	}
	p.Offset = (p.Page - 1) * p.Size
}

// PaginationResponse 翻页响应
type PaginationResponse struct {
	List  interface{} `json:"list"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int64       `json:"total"`
}

// PaginationList 响应列表结构
func PaginationList(pg Pagination, totalCount int64, items interface{}) *PaginationResponse {
	return &PaginationResponse{
		List:  items,
		Total: totalCount,
		Page:  pg.Page,
		Size:  pg.Size,
	}
}
