package domain

type Client struct {
	ID       uint
	Username string
	Token    string
	IsActive bool
}

func (c Client) IsDomain() bool {
	return true
}
