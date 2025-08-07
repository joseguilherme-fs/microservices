package ports

import "github.com/joseguilherme-fs/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
