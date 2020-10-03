package page

type Page struct {
	TotalElements uint64 `json:"totalElements"`
	TotalPages    uint64 `json:"totalPages"`
	PageSize      uint64 `json:"pageSize"`
	PageNumber    uint64 `json:"pageNumber"`
}

type PageInput struct {
	PageNumber uint64 `json:"pageNumber"`
	PageSize   uint64 `json:"pageSize"`
}

func (p PageInput) Page(itemCount uint64) Page {
	if p.PageSize == 0 || p.PageSize > 1000 {
		p.PageSize = 20
	}

	if p.PageNumber == 0 {
		p.PageNumber = 1
	}

	totalPageCount := itemCount / p.PageSize
	if itemCount%p.PageSize > 0 {
		totalPageCount++
	}

	return Page{
		TotalElements: itemCount,
		TotalPages:    totalPageCount,
		PageNumber:    p.PageNumber,
		PageSize:      p.PageSize,
	}
}

func (p Page) SQLOffset() uint64 {
	return (p.PageSize * (p.PageNumber - 1))
}
