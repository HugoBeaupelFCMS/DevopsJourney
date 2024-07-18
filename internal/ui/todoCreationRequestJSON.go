package ui

type todoCreationJSONRequest struct {
	TitleJSON       string `json:"title"`
	DescriptionJSON string `json:"description"`
	DueDateJSON     int64  `json:"duedate"`
}

func (t todoCreationJSONRequest) Title() string {
	return t.TitleJSON
}

func (t todoCreationJSONRequest) Description() string {
	return t.DescriptionJSON
}

func (t todoCreationJSONRequest) DueDate() int64 {
	return t.DueDateJSON
}
