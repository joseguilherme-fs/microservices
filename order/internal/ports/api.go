package ports

 import "github.com/joseguilherme-fs/microservices/order/internal/application/core/domain"

 type APIPort interface {
	PlaceOrder ( order domain.Order ) ( domain.Order,error )
 }