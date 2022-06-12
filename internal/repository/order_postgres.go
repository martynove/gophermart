package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/martynove/gophermart/internal/models"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type OrderPostgres struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewOrderPostgres(logger *logrus.Logger, db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{
		db:     db,
		logger: logger}
}
func (o *OrderPostgres) UploadOrder(userID, orderNumber int) error {
	uploadedAt := time.Now()
	q := fmt.Sprintf("INSERT INTO %s (number, user_id, status, uploaded_at) VALUES ($1, $2, $3, $4)", "orders")
	_, err := o.db.Query(q, orderNumber, userID, "PROCESSED", uploadedAt)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderPostgres) GetAllOrders(userID int) ([]models.Order, error) {
	var orders []models.Order
	q := fmt.Sprintf("SELECT number, status, accrual, uploaded_at from %s WHERE user_id=$1", "orders")
	rows, err := o.db.Query(q, userID)
	if err != nil {
		log.Println(err)
		return orders, err
	}
	for rows.Next() {
		var r models.Order
		err = rows.Scan(&r.Number, &r.Status, &r.Accrual, &r.UploadedAt)
		if err != nil {
			log.Println(err)
			return orders, nil
		}
		orders = append(orders, r)
	}
	return orders, nil
}
