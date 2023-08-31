package repository

import (
	"database/sql"
	"go-laundry/model"
)

type CustomerRepository interface {
	Save(customer model.Customer) error
	FindAll() ([]model.Customer, error)
	FindById(id string) (model.Customer, error)
	Update(customer model.Customer) error
	Delete(id string) error
}

type customerRepository struct {
	db *sql.DB
}

// Delete implements CustomerRepository.
func (repository *customerRepository) Delete(id string) error {
	_, err := repository.db.Exec("delete from customer where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements CustomerRepository.
func (repository *customerRepository) FindAll() ([]model.Customer, error) {
	rows, err := repository.db.Query("select id, name, phone_number, address from customer")
	if err !=nil {
		return nil, err
	}
	
	var customers []model.Customer
	for rows.Next() {
		customer := model.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// FindById implements CustomerRepository.
func (repository *customerRepository) FindById(id string) (model.Customer, error) {
	row := repository.db.QueryRow("SELECT id, name, phone_number, address  FROM customer WHERE id = $1", id)
	var customer model.Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

// Save implements CustomerRepository.
func (repository *customerRepository) Save(customer model.Customer) error {
	_, err := repository.db.Exec("INSERT INTO customer VALUES ($1, $2, $3, $4)", customer.Id, customer.Name, customer.PhoneNumber, customer.Address)
	if err != nil {
		return err
	}
	return nil
}

// Update implements CustomerRepository.
func (repository *customerRepository) Update(customer model.Customer) error {
	SQL := "update customer set name = $1, phone_number = $2, address = $3 where id = $4"

	_, err := repository.db.Exec(SQL, customer.Name, customer.PhoneNumber, customer.Address, customer.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}

}
