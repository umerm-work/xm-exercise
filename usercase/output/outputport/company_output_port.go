package outputport

import (
	"github.com/umerm-work/xm-exercise/domain/models"
	ouputdata "github.com/umerm-work/xm-exercise/usercase/output/data"
)

type CompanyOutputPort interface {
	List(company []models.Company) []*ouputdata.Company
	Item(company models.Company) *ouputdata.Company
}
