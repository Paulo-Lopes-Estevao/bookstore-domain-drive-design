package services

import (
	"bookstore/domain/entities/order"
)

func TotalOrderPrice(order *order.Order) int {
	total := order.Total()
	return total
}
