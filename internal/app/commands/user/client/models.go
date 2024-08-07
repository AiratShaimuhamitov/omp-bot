package client

const (
	PrevButtonText = "Prev page"
	NextButtonText = "Next page"
)

const (
	DefaultListLimit = 3
)

type CallbackListData struct {
	Cursor uint64
	Limit  uint64
}

type ClientInput struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
}
