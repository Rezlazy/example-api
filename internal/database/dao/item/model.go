package item

type CreateInput struct {
	Title       string  `db:"title"`
	Description *string `db:"description"`
	ListID      int64   `db:"list_id"`
}

type Item struct {
	ID int64 `db:"id"`
	CreateInput
}

type ListFilter struct {
	ListIDs []int64
}
