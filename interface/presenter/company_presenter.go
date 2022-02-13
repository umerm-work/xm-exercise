package presenter

import (
	"github.com/umerm-work/xm-exercise/domain/models"
	ouputdata "github.com/umerm-work/xm-exercise/usercase/output/data"
	"gorm.io/gorm"
)

// CompanyPresenter ...
type CompanyPresenter struct {
}

// NewCompanyPresenter ...
func NewCompanyPresenter() *CompanyPresenter {
	return &CompanyPresenter{}
}

// List ...
func (presenter *CompanyPresenter) List(tables []models.Company) []*ouputdata.Company {
	list := make([]*ouputdata.Company, 0, len(tables))
	for _, v := range tables {
		list = append(list, &ouputdata.Company{
			Company: models.Company{
				ID:      v.ID,
				Name:    v.Name,
				Code:    v.Code,
				Country: v.Country,
				Website: v.Website,
				Phone:   v.Phone,
				Model:   gorm.Model{},
			},
		})
	}
	return list
}

func (presenter *CompanyPresenter) Item(in models.Company) *ouputdata.Company {
	return &ouputdata.Company{
		Company: models.Company{
			ID:      in.ID,
			Name:    in.Name,
			Code:    in.Code,
			Country: in.Country,
			Website: in.Website,
			Phone:   in.Phone,
			Model:   gorm.Model{},
		},
	}
}
