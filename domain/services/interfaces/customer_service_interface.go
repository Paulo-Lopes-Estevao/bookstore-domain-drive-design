package interfaces

import "bookstore/domain/aggregate"

type ICustumerService interface {
	IBaseService[aggregate.Customer]
}
