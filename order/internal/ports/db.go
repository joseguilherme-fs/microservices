package ports
2
3 import "github.com/joseguilherme-fs/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get ( id string ) ( domain . Order , error)
	Save (* domain . Order) error
}