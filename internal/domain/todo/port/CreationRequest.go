package port

type ICreationRequest interface {
	Title() string
	Description() string
	DueDate() int64
}
