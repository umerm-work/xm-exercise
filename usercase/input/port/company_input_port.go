package inputport

import (
	inputdata "github.com/umerm-work/xm-exercise/usercase/input/data"
	ouputdata "github.com/umerm-work/xm-exercise/usercase/output/data"
)

// CompanyInputPort ...
type CompanyInputPort interface {
	List(filter string, value string) ([]*ouputdata.Company, error)
	Remove(id int) error
	Create(company inputdata.Company) (*ouputdata.Company, error)
	Update(company inputdata.Company) error
}
