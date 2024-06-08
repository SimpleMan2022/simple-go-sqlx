package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go-gin-sqlx/domain"
)

type RepositoryPegawai interface {
	FindAllPegawai() ([]domain.Pegawai, error)
	FindPegawaiById(id int) (domain.Pegawai, error)
	CreatePegawai(request *domain.Pegawai) (*domain.Pegawai, error)
	FindPegawaiByName(name string) (*domain.Pegawai, error)
	UpdatePegawai(pegawai *domain.Pegawai) (*domain.Pegawai, error)
	DeletePegawai(p *domain.Pegawai) error
	BeginTx() (*sqlx.Tx, error)
}

type pegawaiRepository struct {
	db *sqlx.DB
}

func NewPegawaiRepository(db *sqlx.DB) RepositoryPegawai {
	return &pegawaiRepository{db}
}

func (r *pegawaiRepository) BeginTx() (*sqlx.Tx, error) {
	return r.db.Beginx()
}

func (r *pegawaiRepository) FindAllPegawai() ([]domain.Pegawai, error) {
	var pegawaiList []domain.Pegawai

	err := r.db.Select(&pegawaiList, "SELECT * FROM pegawai")

	if err != nil {
		return nil, err
	}
	return pegawaiList, nil
}
func (r *pegawaiRepository) FindPegawaiById(id int) (domain.Pegawai, error) {
	var pegawai domain.Pegawai
	err := r.db.Get(&pegawai, "SELECT * FROM pegawai WHERE id = ?", id)
	if err != nil {
		return pegawai, err
	}
	return pegawai, nil
}

func (r *pegawaiRepository) CreatePegawai(request *domain.Pegawai) (*domain.Pegawai, error) {
	_, err := r.db.Exec("INSERT INTO pegawai (Name, Address, Age) VALUES (?, ?, ?)",
		request.Name, request.Address, request.Age)
	if err != nil {
		return nil, err
	}

	return request, err
}

func (r *pegawaiRepository) FindPegawaiByName(name string) (*domain.Pegawai, error) {
	query := "SELECT Id, Name, Address, Age FROM pegawai WHERE Name = ?"
	row := r.db.QueryRow(query, name)
	var pegawai domain.Pegawai
	err := row.Scan(&pegawai.ID, &pegawai.Name, &pegawai.Address, &pegawai.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &pegawai, nil
}
func (r *pegawaiRepository) UpdatePegawai(pegawai *domain.Pegawai) (*domain.Pegawai, error) {
	// mulai transaksi
	tx, err := r.db.Beginx()

	if err != nil {
		return nil, err
	}
	_, err = tx.Exec("UPDATE pegawai set Name = ?, Address = ?, Age = ? WHERE id = ?",
		pegawai.Name, pegawai.Address, pegawai.Age, pegawai.ID)

	// rollback jika ada kesalahan
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// commit jika sukses
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return pegawai, nil
}

func (r *pegawaiRepository) DeletePegawai(pegawai *domain.Pegawai) error {
	tx, err := r.db.Beginx()

	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM pegawai WHERE id = ?", pegawai.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
