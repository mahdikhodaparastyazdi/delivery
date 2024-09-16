package constants

type CouriorStatus string

const (
	COURIOR_STATUS_PENDING       CouriorStatus = "pending"
	COURIOR_STATUS_ASSIGNED      CouriorStatus = "assigned"
	COURIOR_STATUS_DELIVERED     CouriorStatus = "deliverd"
	COURIOR_STATUS_NOT_AVAILABLE CouriorStatus = "not_available"
)

func (n CouriorStatus) String() string {
	return string(n)
}
