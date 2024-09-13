package provider1

import (
	couriorproviders "delivery/pkg/couriorproviders"
)

type provider1 struct {
}

func NewProvider1() couriorproviders.CouriorSender {
	return provider1{}
}
