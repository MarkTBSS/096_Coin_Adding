package models

type Item struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Price       uint   `json:"price"`
}

type ItemFilter struct {
	Name        string `query:"name" validate:"omitempty,max=64"`
	Description string `query:"description" validate:"omitempty,max=128"`
	Paginate
}

type Paginate struct {
	Page int64 `query:"page" validate:"required,min=1"`
	Size int64 `query:"size" validate:"required,min=1,max=20"`
}

type ItemResult struct {
	Items    []*Item        `json:"items"`
	Paginate PaginateResult `json:"paginate"`
}

type PaginateResult struct {
	Page      int64 `json:"page"`
	TotalPage int64 `json:"totalPage"`
}
