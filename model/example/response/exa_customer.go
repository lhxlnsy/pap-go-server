package response

import "github.com/lhxlnsy/pap-go-server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
