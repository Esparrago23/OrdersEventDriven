package application

import (
	"order_Event_Driven/src/orders/domain"
	"order_Event_Driven/src/orders/domain/entities"
)

type FindOrderByIdUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewFindOrderByIdUseCase(OrderRepository domain.OrderRepository) *FindOrderByIdUseCase {
	return &FindOrderByIdUseCase{OrderRepository: OrderRepository}
}

func (useCase *FindOrderByIdUseCase) Execute(order_id int) (*entities.Order, error) {
	order, err := useCase.OrderRepository.FindById(order_id)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
