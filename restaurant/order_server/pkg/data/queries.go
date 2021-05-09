package data

const (
	readAllOrdersQuery = `SELECT orders.id, orders.date, orders.table_number, orders.waiter_id, orders.price, orders.payment FROM "orders"`
	readOrdersQuery    = `SELECT orders.id, orders.date, orders.table_number, orders.waiter_id, orders.price, orders.payment FROM "orders" WHERE orders.id = $1`
	updateOrderQuery   = `UPDATE "orders" SET "payment"=$1 WHERE orders.id = $2`
	deleteOrderQuery   = `DELETE FROM "orders" WHERE orders.id = $1`
	allOrders          = "orders.id, orders.date, orders.table_number, orders.waiter_id, orders.price, orders.payment"
	ordersTable        = "orders"
	readWhere          = "orders.id = ?"
	readOrder          = "orders.id, orders.date, orders.table_number, orders.waiter_id, orders.price, orders.payment"
	updateColumn       = "payment"
)
