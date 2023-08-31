package repository

import (
	"database/sql"
	"go-laundry/model"
)

type EmployeeRepository interface {
	Save(employee model.Employee) error
	FindAll() ([]model.Employee, error)
	FindById(id string) (model.Employee, error)
	Update(employee model.Employee) error
	Delete(id string) error
}

type employeeRepository struct {
	db *sql.DB
}

// Delete implements employeeRepository.
func (repository *employeeRepository) Delete(id string) error {
	_, err := repository.db.Exec("delete from employee where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements employeeRepository.
func (repository *employeeRepository) FindAll() ([]model.Employee, error) {
	rows, err := repository.db.Query("select id, name, phone_number, address from employee")
	if err !=nil {
		return nil, err
	}
	
	var employees []model.Employee
	for rows.Next() {
		employee := model.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

// FindById implements employeeRepository.
func (repository *employeeRepository) FindById(id string) (model.Employee, error) {
	row := repository.db.QueryRow("SELECT id, name, phone_number, address  FROM employee WHERE id = $1", id)
	var employee model.Employee
	err := row.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

// Save implements employeeRepository.
func (repository *employeeRepository) Save(employee model.Employee) error {
	_, err := repository.db.Exec("INSERT INTO employee VALUES ($1, $2, $3, $4)", employee.Id, employee.Name, employee.PhoneNumber, employee.Address)
	if err != nil {
		return err
	}
	return nil
}

// Update implements employeeRepository.
func (repository *employeeRepository) Update(employee model.Employee) error {
	SQL := "update employee set name = $1, phone_number = $2, address = $3 where id = $4"

	_, err := repository.db.Exec(SQL, employee.Name, employee.PhoneNumber, employee.Address, employee.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}

}
