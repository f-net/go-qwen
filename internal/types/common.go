package types

type GetId struct {
	Id int64 `json:"id" form:"id"`
}

type ListQuery struct {
	Page   int `json:"page" form:"page"`
	Size   int `json:"size" form:"size"`
	Offset int `json:"offset" form:"offset"`
}

func (l *ListQuery) InitPageSize() {
	if l.Page == 0 {
		l.Page = 1
	}
	if l.Size == 0 {
		l.Size = 10
	}
	l.Offset = (l.Page - 1) * l.Size
}

type ListCommonResp struct {
	List      interface{} `json:"list"`
	Total     int64       `json:"total"`
	Page      int         `json:"page"`
	Size      int         `json:"size"`
	TotalPage int         `json:"totalPage"`
}
