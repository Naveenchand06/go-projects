package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}


func getProducts(db *sql.DB) ([]product, error) {
	query := "SELECT id, name, quantity, price FROM products"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var products []product
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil { 
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}



func (p *product) getProduct(db *sql.DB) error {
	fmt.Println("get P ---> 6")

	query := fmt.Sprintf("SELECT id, name, quantity, price FROM products WHERE id=%v", p.ID)
	fmt.Println("get P ---> 7")

	row := db.QueryRow(query)
	err := row.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
	fmt.Println("get P ---> 8")

	if err != nil {
	fmt.Println("get P ---> 9", err)
		return err
	}
	return nil
}


func (p *product) createProduct(db *sql.DB) error {
	query := fmt.Sprintf("insert into products(name,quantity,price) values('%v', %v, %v)", p.Name, p.Quantity, p.Price)
	fmt.Println("Post product ", query, p)
	result, err := db.Exec(query)
	if err != nil {
	fmt.Println("Post product 1 ", query, p)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil
}


func (p *product) updateProduct(db *sql.DB) error {
	query := fmt.Sprintf("update products set name='%v', quantity=%v, price=%v where id=%v", p.Name, p.Quantity, p.Price, p.ID)
	result, err :=  db.Exec(query)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no such row exists")
	}
	return nil
}


func (p *product) deleteProduct(db *sql.DB) error {
	query := fmt.Sprintf("delete from products where id=%v", p.ID)
	_, err := db.Exec(query)
	return err
}