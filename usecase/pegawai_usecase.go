package usecase

import (
	"errors"
	"go-gin-sqlx/domain"
	"go-gin-sqlx/repository"
	"strconv"
)

type PegawaiUsecase interface {
	FindAllPegawai() ([]domain.Pegawai, error)
	FindPegawaiById(id int) (domain.Pegawai, error)
	CreatePegawai(request *domain.PegawaiRequest) (*domain.Pegawai, error)
	UpdatePegawai(id int, request *domain.PegawaiRequest) (*domain.Pegawai, error)
	DeletePegawai(id int) error
}

type pegawaiUsecase struct {
	pegawaiRepository repository.RepositoryPegawai
}

func NewPegawaiUsecase(pegawai repository.RepositoryPegawai) PegawaiUsecase {
	return &pegawaiUsecase{pegawaiRepository: pegawai}
}

func (uc *pegawaiUsecase) FindAllPegawai() ([]domain.Pegawai, error) {
	pegawai, err := uc.pegawaiRepository.FindAllPegawai()

	if err != nil {
		return nil, err
	}
	return pegawai, nil
}

func (uc *pegawaiUsecase) FindPegawaiById(id int) (domain.Pegawai, error) {
	var pegawai domain.Pegawai
	findPegawai, err := uc.pegawaiRepository.FindPegawaiById(id)

	if err != nil {
		return pegawai, err
	}
	return findPegawai, nil
}

func (uc pegawaiUsecase) ValidatePegawaiRequest(request *domain.PegawaiRequest) error {
	if request.Name == "" || request.Address == "" || request.Age == 0 {
		return errors.New("Data tidak boleh kosong")
	}
	return nil
}

func (uc pegawaiUsecase) CheckDuplicateName(name string) error {
	findPegawai, err := uc.pegawaiRepository.FindPegawaiByName(name)

	if err != nil {
		return err
	}
	if findPegawai != nil && findPegawai.Name == name {
		return errors.New("nama sudah ada")
	}
	return nil
}
func (uc pegawaiUsecase) CheckDuplicateNameById(name string, id int) error {
	findPegawai, err := uc.pegawaiRepository.FindPegawaiByName(name)

	if err != nil {
		return err
	}
	if findPegawai != nil && findPegawai.ID != strconv.Itoa(id) {
		return errors.New("nama sudah ada")
	}
	return nil
}
func (uc *pegawaiUsecase) CreatePegawai(request *domain.PegawaiRequest) (*domain.Pegawai, error) {

	if err := uc.ValidatePegawaiRequest(request); err != nil {
		return nil, err
	}

	if err := uc.CheckDuplicateName(request.Name); err != nil {
		return nil, err
	}
	pegawaiReq := &domain.Pegawai{
		Name:    request.Name,
		Address: request.Address,
		Age:     request.Age,
	}

	newPegawai, err := uc.pegawaiRepository.CreatePegawai(pegawaiReq)
	if err != nil {
		return nil, err
	}
	return newPegawai, nil
}

func (uc pegawaiUsecase) UpdatePegawai(id int, Request *domain.PegawaiRequest) (*domain.Pegawai, error) {

	tx, err := uc.pegawaiRepository.BeginTx()

	if err != nil {
		return nil, err
	}
	findPegawai, err := uc.pegawaiRepository.FindPegawaiById(id)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = uc.CheckDuplicateNameById(Request.Name, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	findPegawai.Name = Request.Name
	findPegawai.Address = Request.Address
	findPegawai.Age = Request.Age

	updatedPegawai, err := uc.pegawaiRepository.UpdatePegawai(&findPegawai)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return updatedPegawai, nil
}

func (uc pegawaiUsecase) DeletePegawai(id int) error {
	tx, err := uc.pegawaiRepository.BeginTx()
	if err != nil {
		return err
	}
	findUser, err := uc.pegawaiRepository.FindPegawaiById(id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = uc.pegawaiRepository.DeletePegawai(&findUser)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
