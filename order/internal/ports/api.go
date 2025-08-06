package ports
2
3 import "github.com/joseguilherme-fs/microservices/order/internal/application/core/domain"
4
5 type APIPort interface {
6 PlaceOrder ( order domain . Order ) ( domain . Order , error )
7 }