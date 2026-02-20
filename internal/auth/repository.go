package auth

import (
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) CreateTest(pekerjaan string, umur int, alamat string) (*Contoh, error){
	var contoh Contoh
	query := "INSERT INTO contoh (pekerjaan, umur, alamat) VALUES ($1, $2, $3) RETURNING pekerjaan, umur, alamat"
	err := r.db.QueryRow(query,pekerjaan,umur,alamat).Scan(&contoh.Pekerjaan,contoh.Umur,&contoh.Alamat)
	if err != nil {
		return nil, err
	}
	return &contoh ,nil
}

func (r *Repository) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) Create(name, email string) (*User, error) {
	var user User
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email"
	err := r.db.QueryRow(query, name, email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetByID(id int) (*User, error) {
	var user User
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *Repository) Update(user User) error {
	res, err := r.db.Exec("UPDATE users SET name=$1,email=$2 WHERE id = $3", user.Name , user.Email,user.ID)
	
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("users not found")
	}
	return  nil
}

