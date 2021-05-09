package data

import (
	"database/sql"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func NewGorm(db *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	})
	gormDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gormDb
}

var testOrder = Order{
	Id:          1,
	Date:        time.Now(),
	TableNumber: 1,
	WaiterId:    1,
	Price:       124,
	Payment:     true,
}

func TestReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	rows := sqlmock.NewRows([]string{"id", "Date", "TableNumber", "WaiterId", "Price", "Payment"}).
		AddRow(testOrder.Id, testOrder.Date, testOrder.TableNumber, testOrder.WaiterId, testOrder.Price, testOrder.Payment)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnRows(rows)
	orders, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(orders)
	assert.Equal(orders[0], testOrder)
	assert.Len(orders, 1)
}

func TestReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnError(errors.New("something went wrong..."))
	users, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(users)
}

func TestRead(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	rows := sqlmock.NewRows([]string{"id", "Date", "TableNumber", "WaiterId", "Price", "Payment"}).
		AddRow(testOrder.Id, testOrder.Date, testOrder.TableNumber, testOrder.WaiterId, testOrder.Price, testOrder.Payment)
	mock.ExpectQuery(readOrdersQuery).WithArgs(testOrder.Id).WillReturnRows(rows)
	orders, err := data.Read(testOrder.Id)
	assert.NoError(err)
	assert.Equal(orders, testOrder)
}

func TestReadErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectQuery(readOrdersQuery).WithArgs(testOrder.Id).
		WillReturnError(errors.New("something went wrong..."))
	users, err := data.Read(testOrder.Id)
	assert.Error(err)
	assert.Empty(users)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(deleteOrderQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := data.Delete(1)
	assert.NoError(err)
}

func TestDeleteErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(deleteOrderQuery).
		WithArgs(1).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	err := data.Delete(1)
	assert.Error(err)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(updateOrderQuery).
		WithArgs(testOrder.Payment, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := data.Update(1, testOrder.Payment)
	assert.NoError(err)
}

func TestUpdateErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(updateOrderQuery).
		WithArgs(testOrder.Payment, 1).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	err := data.Update(1, testOrder.Payment)
	assert.Error(err)
}
