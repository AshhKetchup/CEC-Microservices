package model

import (
	"time"
)

type DeliveryStatus int

const (
	DeliveryStatusPending DeliveryStatus = iota
	DeliveryStatusInTransit
	DeliveryStatusDelivered
	DeliveryStatusDelayed
)

type Delivery struct {
	ID                string
	OrderID           string
	Address           string
	Status            DeliveryStatus
	EstimatedDelivery time.Time
	CreatedAt         time.Time
}

type DeliveryRequest struct {
	OrderID string
	Address string
}

type DeliveryStorage interface {
	SaveDelivery(*Delivery) (*Delivery, error)
	GetDelivery(string) (*Delivery, error)
	UpdateDeliveryStatus(string, DeliveryStatus) (*Delivery, error)
}

var ErrDeliveryNotFound = errors.New("delivery not found")
