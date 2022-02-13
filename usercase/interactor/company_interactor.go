package interactor

import (
	"errors"
	"github.com/umerm-work/xm-exercise/domain/models"
	inputdata "github.com/umerm-work/xm-exercise/usercase/input/data"
	ouputdata "github.com/umerm-work/xm-exercise/usercase/output/data"
	"github.com/umerm-work/xm-exercise/usercase/output/outputport"
	"github.com/umerm-work/xm-exercise/usercase/repository"
	"log"
)

// CompanyInteractor express the flow of authentication
type CompanyInteractor struct {
	outputport        outputport.CompanyOutputPort
	companyRepository repository.CompanyRepository
}

// NewCompanyInteractor ...
func NewCompanyInteractor(
	outputport outputport.CompanyOutputPort,
	companyRepository repository.CompanyRepository,
) *CompanyInteractor {
	return &CompanyInteractor{
		outputport:        outputport,
		companyRepository: companyRepository,
	}
}

// List companies
func (r *CompanyInteractor) List(filter string, value string) ([]*ouputdata.Company, error) {
	if len(filter) == 0 {
		list, err := r.companyRepository.List()
		if err != nil {
			return nil, err
		}
		return r.outputport.List(list), nil
	}
	list, err := r.companyRepository.Find(filter, value)
	if err != nil {
		return nil, err
	}

	return r.outputport.List(list), nil
}

// Create companies
func (r *CompanyInteractor) Create(in inputdata.Company) (*ouputdata.Company, error) {
	rest := models.NewCompany(in.Name, in.Code, in.Country, in.Website, in.Phone)
	err := r.companyRepository.Save(rest)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return r.outputport.Item(*rest), nil
}

// Update company
func (r *CompanyInteractor) Update(in inputdata.Company) error {
	if in.ID == 0 {
		return errors.New("id is required")
	}
	currentCompanyData, err := r.companyRepository.Get(models.Company{ID: in.ID})
	if err != nil {
		return err
	}
	if len(in.Name) > 0 {
		currentCompanyData.Name = in.Name
	}
	if len(in.Country) > 0 {
		currentCompanyData.Country = in.Country
	}
	if len(in.Phone) > 0 {
		currentCompanyData.Phone = in.Phone
	}
	if len(in.Code) > 0 {
		currentCompanyData.Code = in.Code
	}
	if len(in.Website) > 0 {
		currentCompanyData.Code = in.Website
	}

	err = r.companyRepository.Update(currentCompanyData)
	log.Println(err)
	if err != nil {
		return err
	}

	return nil
}

// Remove companies
func (r *CompanyInteractor) Remove(id int) error {
	err := r.companyRepository.Remove(id)
	if err != nil {
		return err
	}

	return nil
}
