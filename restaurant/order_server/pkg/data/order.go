package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id          int64
	Date        time.Time
	TableNumber int64
	WaiterId    int64
	Price       int64
	Payment     bool
}

func (o Order) String() string {
	return fmt.Sprintf("(%d %s %d %d %d %t)", o.Id, o.Date, o.TableNumber, o.WaiterId, o.Price, o.Payment)
}

type OrderData struct {
	db *gorm.DB
}

func NewOrderData(db *gorm.DB) *OrderData {
	return &OrderData{db: db}
}

func (o OrderData) ReadAll() ([]Order, error) {
	var orders []Order

	err := o.db.Table(ordersTable).
		Select(allOrders).
		Find(&orders)
	if err.Error != nil {
		return nil, err.Error
	}
	return orders, nil
}

func (o OrderData) Read(id int64) (Order, error) {
	var order Order
	err := o.db.Table(ordersTable).
		Where(readWhere, id).
		Select(readOrder).
		Find(&order)

	if err.Error != nil {
		return Order{}, err.Error
	}
	return order, nil
}

func (o OrderData) Create(order Order) error {

	err := o.db.Create(&order)
	if err.Error != nil {
		return fmt.Errorf("error: %s", err.Error)
	}
	return nil
}

func (o OrderData) Update(id int64, payment bool) error {
	err := o.db.Table(ordersTable).Where(readWhere, id).Update(updateColumn, payment)
	if err.Error != nil {
		return fmt.Errorf("error: %s", err.Error)
	}
	return nil
}

func (o OrderData) Delete(id int64) error {
	err := o.db.Where(readWhere, id).Delete(&Order{})
	if err.Error != nil {
		return fmt.Errorf("error: %s", err.Error)
	}
	return nil
}
