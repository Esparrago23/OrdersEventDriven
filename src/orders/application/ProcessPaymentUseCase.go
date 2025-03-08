package application

import (
	"order_Event_Driven/src/orders/domain"
	"order_Event_Driven/src/orders/domain/entities"
)

type ProcessPaymentUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewProcessPaymentUseCase(OrderRepository domain.OrderRepository) *ProcessPaymentUseCase {
	return &ProcessPaymentUseCase{OrderRepository: OrderRepository}
}

func (useCase *ProcessPaymentUseCase) Execute(order_id int) (*entities.Order, error){
	order, err := useCase.OrderRepository.ProccessPayment(order_id)
	if err != nil {
		return nil, err
	}
	return &order, nil
	
	
}
