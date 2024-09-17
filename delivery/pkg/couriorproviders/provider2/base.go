package provider2

import (
	couriorproviders "delivery/pkg/couriorproviders"
)

type provider2 struct {
}

func NewProvider2() couriorproviders.CouriorSender {
	return provider2{}
}
