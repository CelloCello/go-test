package pagination

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Paginator struct {
	Limit   int
	Page    int
	Count   int64
	Sort    string
	OrderBy string
}

type SortType string

const (
	Asc  SortType = "asc"
	Desc SortType = "desc"
)

const (
	DefaultLimit      = "10"
	DefaultPage       = "1"
	DefaultSort       = Desc
	DefaultOrderField = "created_at"
	MaxPageNum        = 100
)

type PageData struct {
	Count       int64 `json:"count"`
	TotalPage   int   `json:"totalPage"`
	CurrentPage int   `json:"currentPage"`

	// need redefine
	List interface{} `json:"list"`
}

func NewPaginatorFromRequest(c *gin.Context) *Paginator {
	// Initializing default
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", DefaultLimit))
	if limit > MaxPageNum {
		limit = MaxPageNum
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", DefaultPage))
	sort := c.DefaultQuery("sort", string(DefaultSort))
	orderBy := c.DefaultQuery("orderby", DefaultOrderField)

	return &Paginator{
		Limit:   limit,
		Page:    page,
		Sort:    sort,
		OrderBy: orderBy,
	}
}

func (p *Paginator) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

func (p *Paginator) GetOrderCmd(orderFields []string) string {
	foundField := false
	for _, field := range orderFields {
		if field == p.OrderBy {
			foundField = true
			break
		}
	}

	if foundField {
		return fmt.Sprintf("%s %s", p.OrderBy, p.Sort)
	}
	return fmt.Sprintf("%s %s", DefaultOrderField, p.Sort)
}

func (p *Paginator) GetTotalPage() int {
	if p.Count < 0 || p.Limit <= 0 {
		return 0
	}

	return int(math.Ceil(float64(p.Count) / float64(p.Limit)))
}
