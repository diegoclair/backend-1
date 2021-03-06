package entity

import (
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// Companies entity
type Companies struct {
	CompanyID      int64   `json:"company_id"`
	CompanyName    string  `json:"company_name"`
	TotalRating    int64   `json:"total_rating"`
	RatingQuantity float64 `json:"rating_quantity"`
	AverageRating  float64 `json:"average_rating"`
	Street         string  `json:"street"`
	Number         string  `json:"number"`
	City           string  `json:"city"`
	FederativeUnit string  `json:"federative_unit"`
}

// CompanyDetail entity
type CompanyDetail struct {
	ID             int64         `json:"id"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Description    string        `json:"description"`
	CountryCode    string        `json:"country_code"`
	AreaCode       string        `json:"area_code"`
	PhoneNumber    string        `json:"phone_number"`
	DocumentNumber string        `json:"document_number"`
	Website        string        `json:"website"`
	BusinessID     int64         `json:"business_id"`
	Address        Address       `json:"address"`
	SocialNetwork  SocialNetwork `json:"social_network"`
}

// Address entity
type Address struct {
	Country        string `json:"country"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	Complement     string `json:"complement"`
	ZipCode        int64  `json:"zip_code"`
	Neighborhood   string `json:"neighborhood"`
	City           string `json:"city"`
	FederativeUnit string `json:"federative_unit"`
}

// SocialNetwork entity
type SocialNetwork struct {
	InstagramURL string `json:"instagram_url"`
	FacebookURL  string `json:"facebook_url"`
	LinkedinURL  string `json:"linkedin_url"`
	TwitterURL   string `json:"twitter_url"`
}

// Validate to validate a user data
func (company *CompanyDetail) Validate() *resterrors.RestErr {

	company.Name = strings.TrimSpace(company.Name)
	if company.Name == "" {
		return resterrors.NewBadRequestError("Name is invalid")
	}

	company.Email = strings.TrimSpace(company.Email)
	if company.Email == "" {
		return resterrors.NewBadRequestError("Email is invalid")
	}

	company.AreaCode = strings.TrimSpace(company.AreaCode)
	if company.AreaCode == "" {
		return resterrors.NewBadRequestError("AreaCode is invalid")
	}

	company.PhoneNumber = strings.TrimSpace(company.PhoneNumber)
	if company.PhoneNumber == "" {
		return resterrors.NewBadRequestError("PhoneNumber is invalid")
	}

	if company.BusinessID == 0 {
		return resterrors.NewBadRequestError("BusinessID is invalid")
	}

	return nil
}
