package carbot

import (
	"errors"
	"fmt"
)

type AdvertiserInterface interface {
	ProcessQuery(q *Query, b *Browser) (ads []Ad, err error)
	GetName() string
	GetFormURL() string
}

type AdvertisersRegistry struct {
	Advertisers []AdvertiserInterface
}

// NewAdvertisersRegistry returns a new AdvertisersRegistry
func NewAdvertisersRegistry() *AdvertisersRegistry {
	return &AdvertisersRegistry{}
}

// ProcessQuery finds advertiser in the registry and processes the query
func (ar *AdvertisersRegistry) ProcessQuery(q *Query, b *Browser) (ads []Ad, err error) {
	// log.Printf("[AdvertisersRegistry] Processing query '%s' for advertiser: '%s'", q.Name, q.AdvertiserName)
	for _, a := range ar.Advertisers {
		if a.GetName() == q.AdvertiserName {
			return a.ProcessQuery(q, b)
		}
	}
	return nil, errors.New(fmt.Sprintf("[AdvertisersRegistry] Advertiser '%s' not found", q.AdvertiserName))
}

// AddAdvertiser adds a new advertiser to the registry
func (ar *AdvertisersRegistry) Add(a AdvertiserInterface) {
	ar.Advertisers = append(ar.Advertisers, a)
}
