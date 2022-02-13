package database

import (
	"fmt"
	"github.com/umerm-work/xm-exercise/domain/models"
	"github.com/umerm-work/xm-exercise/infrastructure/db"
	"github.com/umerm-work/xm-exercise/usercase/repository"
	"gorm.io/gorm"
)

// CompanyDatabase ...
type CompanyDatabase struct {
	db *gorm.DB
}

// NewCompanyDatabase ...
func NewCompanyDatabase() repository.CompanyRepository {
	return &CompanyDatabase{db.CONN}
}

func (r *CompanyDatabase) Save(table *models.Company) error {
	if err := r.db.Create(&table).Error; err != nil {
		return err
	}
	return nil
}

func (r *CompanyDatabase) Find(filter string, value string) ([]models.Company, error) {
	var list []models.Company
	format := `%q='%s'`
	if filter == "id" {
		format = `%q=%s`
	}
	if err := r.db.Where(fmt.Sprintf(format, filter, value)).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *CompanyDatabase) List() ([]models.Company, error) {
	var list []models.Company
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *CompanyDatabase) Update(table *models.Company) error {
	if err := r.db.Save(&table).Error; err != nil {
		return err
	}
	return nil
}
func (r *CompanyDatabase) Get(condition models.Company) (*models.Company, error) {
	var out models.Company
	if err := r.db.Where(&condition).Find(&out).Error; err != nil {
		return nil, err
	}
	return &out, nil
}
func (r *CompanyDatabase) Remove(id int) error {
	var out models.Company
	if err := r.db.Where(models.Company{ID: id}).Delete(&out).Error; err != nil {
		return err
	}
	return nil
}
