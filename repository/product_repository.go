package repository

import (
	"database/sql"
	"go-laundry/model"
)

type ProductRepository interface {
	Save(product model.Product) error
	FindAll() ([]model.Product, error)
	FindById(id string) (model.Product, error)
	Update(product model.Product) error
	Delete(id string) error
}

type productRepository struct {
	db *sql.DB
}

// Delete implements ProductRepository.
func (p *productRepository) Delete(id string) error {
	SQL := "delete from product where id = $1"
	_, err := p.db.Exec(SQL, id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductRepository.
func (p *productRepository) FindAll() ([]model.Product, error) {
	SQL := "select p.id, p.name, p.price, u.name from product as p join uom as u on u.id = p.uom_id"

	rows, err := p.db.Query(SQL)
	if err != nil {
		return nil, err
	}

	var products []model.Product
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Uom.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// FindById implements ProductRepository.
func (p *productRepository) FindById(id string) (model.Product, error) {
	SQL := "select p.id, p.name, p.price, u.name from product as p join uom as u on u.id = p.uom_id where p.id = $1"

	row := p.db.QueryRow(SQL, id)
	var product model.Product
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Uom.Name)
	if err != nil {
		return model.Product{} , err
	}

	return product, nil
}

// Save implements ProductRepository.
func (p *productRepository) Save(product model.Product) error {
	SQL := "insert into product values($1, $2, $3, $4)"
	_, err := p.db.Exec(SQL, product.Id, product.Name, product.Price, product.Uom.Id)
	if err != nil {
		return err
	}

	return nil
}

// Update implements ProductRepository.
func (p *productRepository) Update(product model.Product) error {
	SQL := "update product set name = $1 where id = $2"

	_, err := p.db.Exec(SQL, product.Name, product.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
