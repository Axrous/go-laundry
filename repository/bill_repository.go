package repository

import (
	"database/sql"
	"go-laundry/model"
)

type BillRepository interface {
	Save(bill model.Bill) error
	FindAll() ([]model.Bill, error)
	FindById(id string) ([]model.BillDetailShow, error)
}

type billRepository struct {
	db *sql.DB
}

// FindAll implements BillRepository.
func (repository *billRepository) FindAll() ([]model.Bill, error) {
	SQL := "select b.id, b.bill_date, b.entry_date, b.finish_date, e.name, c.name from bill as b join employee as e on e.id = b.employee_id join customer as c on c.id = b.customer_id"
	rows, err := repository.db.Query(SQL)
	if err != nil {
		return nil, err
	}

	var bills []model.Bill
	for rows.Next(){
		var bill model.Bill
		err := rows.Scan(&bill.Id, &bill.BillDate, &bill.EntryDate, &bill.FinishDate, &bill.Employee.Name, &bill.Customer.Name)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}

	return bills, err
}

// FindById implements BillRepository.
func (repository *billRepository) FindById(id string) ([]model.BillDetailShow, error) {
	SQL := `select b.id, b.bill_date, b.entry_date, b.finish_date, e.id, e.name, c.id, c.name, bd.qty, bd.product_price, p.name
	from bill as b
	left join employee as e on e.id = b.employee_id
	left join customer as c on c.id = b.customer_id
	left join bill_detail as bd on bd.bill_id = b.id
	left join product as p on p.id = bd.product_id
	where b.id = $1
	`
	result, err := repository.db.Query(SQL, id)
	if err != nil {
		return nil, err
	}

	var bills []model.BillDetailShow

	for result.Next(){
		var bill model.BillDetailShow
		err := result.Scan(
			&bill.Id,
			&bill.BillDate,
			&bill.EntryDate,
			&bill.FinishDate,
			&bill.Employee.Id,
			&bill.Employee.Name,
			&bill.Customer.Id,
			&bill.Customer.Name,
			&bill.Qty,
			&bill.ProductPrice,
			&bill.Product.Name,
		)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}

	return bills, err
}

// Save implements BillRepository.
func (repository *billRepository) Save(bill model.Bill) error {
	tx, err := repository.db.Begin()

	if err != nil {
		return err
	}

	SQL := "insert into bill values($1, $2, $3, $4, $5, $6)"
	_, err = tx.Exec(SQL, bill.Id, bill.BillDate, bill.EntryDate, bill.FinishDate, bill.Employee.Id, bill.Customer.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, billDetail := range bill.BillDetail {
		// err = repository.saveDetail(tx, billDetail)
		SQL := "insert into bill_detail values($1, $2, $3, $4, $5)"
		_, err := tx.Exec(SQL, billDetail.Id, billDetail.BillId, billDetail.Product.Id, billDetail.ProductPrice, billDetail.Qty)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{
		db: db,
	}
}
