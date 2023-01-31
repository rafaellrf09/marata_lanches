package models

type Address struct {
	Street       string `json:"street,omitempty"`
	Number       string `json:"number,omitempty"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	ZipCode      string `json:"zipcode,omitempty"`
}
