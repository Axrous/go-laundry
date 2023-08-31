package repository

import (
	"database/sql"
	"go-laundry/model"
)

type UomRepository interface {
	Save(uom model.Uom) error              // INSERT
	FindById(id string) (model.Uom, error) // SELECT by id
	FindAll() ([]model.Uom, error)         // SELECT *
	Update(uom model.Uom) error
	DeleteById(id string) error // DELETE FROM apa WHERE id = ?
}

type uomRepository struct {
	db *sql.DB
}

// DeleteById implements UomRepository.
func (u *uomRepository) DeleteById(id string) error {
	_, err := u.db.Exec("delete from uom where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements UomRepository.
func (u *uomRepository) FindAll() ([]model.Uom, error) {
	
	var uoms []model.Uom
	rows, err := u.db.Query("select id, name from uom")
	if err !=nil {
		return uoms, err
	}
	
	for rows.Next() {
		uom := model.Uom{}
		err := rows.Scan(&uom.Id, &uom.Name)
		if err != nil {
			return nil, err
		}
		uoms = append(uoms, uom)
	}

	return uoms, nil
}

// FindById implements UomRepository.
func (u *uomRepository) FindById(id string) (model.Uom, error) {
	row := u.db.QueryRow("SELECT id, name FROM uom WHERE id = $1", id)
	var uom model.Uom
	err := row.Scan(&uom.Id, &uom.Name)
	if err != nil {
		return model.Uom{}, err
	}
	return uom, nil
}

// Save implements UomRepository.
func (u *uomRepository) Save(uom model.Uom) error {
	_, err := u.db.Exec("INSERT INTO uom VALUES ($1, $2)", uom.Id, uom.Name)
	if err != nil {
		return err
	}
	return nil
}

// Update implements UomRepository.
func (u *uomRepository) Update(uom model.Uom) error {
	_, err := u.db.Exec("update uom set id = $1, name = $2 where id = $3", uom.Id, uom.Name, uom.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
