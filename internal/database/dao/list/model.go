package list

type CreateInput struct {
	Name string `db:"name"`
}

type List struct {
	ID int64 `db:"id"`
	CreateInput
}

type ListFilter struct {
	IDs []int64
}
