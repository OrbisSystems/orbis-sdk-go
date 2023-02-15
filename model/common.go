package model

type Paging struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type OrderDirection string

const (
	OrderASC  OrderDirection = "asc"
	OrderDESC OrderDirection = "desc"
)

type Ordering struct {
	// Sort direction:
	// * asc - Ascending, from A to Z.
	// * desc - Descending, from Z to A.
	Field     string
	Direction OrderDirection `enums:"asc,desc" default:"asc"`
}
