package repository

import "github.com/umerm-work/xm-exercise/domain/models"

// CompanyRepository is an interface for DB operation
type CompanyRepository interface {
	Save(company *models.Company) error
	Update(company *models.Company) error
	Find(filter string, value string) ([]models.Company, error)
	List() ([]models.Company, error)
	Remove(id int) error
	Get(company models.Company) (*models.Company, error)
}
