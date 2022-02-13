package presenter

import (
	"github.com/umerm-work/xm-exercise/domain/models"
	ouputdata "github.com/umerm-work/xm-exercise/usercase/output/data"
	"reflect"
	"testing"
)

func TestCompanyPresenter_Item(t *testing.T) {

	tests := []struct {
		name string
		in   models.Company
		want *ouputdata.Company
	}{
		{
			name: "Correct Result",
			in: models.Company{
				Name:    "abc",
				Code:    "abc",
				Country: "abc",
				Website: "abc",
				Phone:   "abc",
			},
			want: &ouputdata.Company{
				models.Company{
					Name:    "abc",
					Code:    "abc",
					Country: "abc",
					Website: "abc",
					Phone:   "abc",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presenter := &CompanyPresenter{}
			if got := presenter.Item(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Item() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyPresenter_List(t *testing.T) {

	tests := []struct {
		name   string
		tables []models.Company
		want   []*ouputdata.Company
	}{
		{
			name: "Correct Result",
			tables: []models.Company{models.Company{
				Name:    "abc",
				Code:    "abc",
				Country: "abc",
				Website: "abc",
				Phone:   "abc",
			}},
			want: []*ouputdata.Company{
				{models.Company{
					Name:    "abc",
					Code:    "abc",
					Country: "abc",
					Website: "abc",
					Phone:   "abc",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presenter := &CompanyPresenter{}
			if got := presenter.List(tt.tables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}
