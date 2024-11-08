package main

type OfferType string

const (
	OfferTypeNoOffer OfferType = "NoOffer"
	OfferType1       OfferType = "Offer1"
	OfferType2       OfferType = "Offer2"
)

type Offer interface {
	GetOfferType() OfferType
}

type NoOffer struct{}

func NewNoOffer() *NoOffer {
	return &NoOffer{}
}

func (no *NoOffer) GetOfferType() OfferType {
	return OfferTypeNoOffer
}

type Offer1 struct {
}

func (o1 *Offer1) GetOfferType() OfferType {
	return OfferType1
}

type Offer2 struct {
}

func (o1 *Offer2) GetOfferType() OfferType {
	return OfferType2
}
