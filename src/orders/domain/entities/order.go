package entities

import (
	"time"
)

type Order struct {
	OrderID     int       `json:"order_id" db:"order_id"`
	ServiceName string    `json:"service_name" db:"service_name"`
	Description string    `json:"description" db:"description"`
	TotalAmount float64   `json:"total_amount" db:"total_amount"`
	Status      string    `json:"status" db:"status"` // 'pending', 'contratado', 'cancelled'
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewOrder(orderID int, serviceName string, description string, totalAmount float64, status string, createdAt time.Time, updatedAt time.Time) *Order {
	return &Order{
		OrderID:     orderID,
		ServiceName: serviceName,
		Description: description,
		TotalAmount: totalAmount,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func (o *Order) GetOrderID() int {
	return o.OrderID
}

func (o *Order) GetServiceName() string {
	return o.ServiceName
}

func (o *Order) GetDescription() string {
	return o.Description
}

func (o *Order) GetTotalAmount() float64 {
	return o.TotalAmount
}

func (o *Order) GetStatus() string {
	return o.Status
}

func (o *Order) GetCreatedAt() time.Time {
	return o.CreatedAt
}

func (o *Order) GetUpdatedAt() time.Time {
	return o.UpdatedAt
}
