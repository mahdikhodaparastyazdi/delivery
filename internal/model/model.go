package model

import "delivery/internal/domain"

type DbModel interface {
	TableName() string
	ToDomain() domain.Domain
}
