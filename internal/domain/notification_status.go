package domain

type CouriorStatus string

const (
	CouriorStatusPending CouriorStatus = "PENDING"
	CouriorStatusSent    CouriorStatus = "SENT"
	CouriorStatusRetry   CouriorStatus = "RETRY"
	CouriorStatusFailed  CouriorStatus = "FAILED"
	CouriorStatusExpired CouriorStatus = "EXPIRED"
)

func (n CouriorStatus) String() string {
	return string(n)
}
