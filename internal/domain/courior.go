package domain

import "time"

type COURIOR struct {
	ID         uint
	ProviderId uint
	TemplateID uint
	Mobile     string
	Content    string
	Status     CouriorStatus
	ExpiresAt  time.Time
}
