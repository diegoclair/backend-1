package contract

import (
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

// PingService holds a ping service operations
type PingService interface {
}

// UserService holds a user service operations
type UserService interface {
	GetUsers() (*[]entity.User, *resterrors.RestErr)
	GetUserByID(userID int64) (*entity.User, *resterrors.RestErr)
	CreateUser(entity.User) (int64, *resterrors.RestErr)
	UpdateUser(entity.User) (*entity.User, *resterrors.RestErr)
	DeleteUser(userID int64) *resterrors.RestErr
	LoginUser(request entity.LoginRequest) (*entity.User, *resterrors.RestErr)
}

// ProductService holds a product service operations
type ProductService interface {
	GetProducts(categoryID int64) (*[]entity.Product, *resterrors.RestErr)
	GetProductByID(productID int64) (*entity.Product, *resterrors.RestErr)
	CreateProduct(entity.Product) (int64, *resterrors.RestErr)
	UpdateProduct(entity.Product) (*entity.Product, *resterrors.RestErr)
	DeleteProduct(productID int64) *resterrors.RestErr
}

// CategoryService holds a category service operations
type CategoryService interface {
	GetCategories() (*[]entity.Category, *resterrors.RestErr)
	GetCategoryByID(categoryID int64) (*entity.Category, *resterrors.RestErr)
	CreateCategory(entity.Category) (int64, *resterrors.RestErr)
	UpdateCategory(entity.Category) (*entity.Category, *resterrors.RestErr)
	DeleteCategory(categoryID int64) *resterrors.RestErr
}

// BusinessService holds a business service operations
type BusinessService interface {
	GetBusinesses() (*[]entity.Business, *resterrors.RestErr)
	GetBusinessByID(businessID int64) (*entity.Business, *resterrors.RestErr)
	CreateBusiness(entity.Business) (int64, *resterrors.RestErr)
	UpdateBusiness(entity.Business) (*entity.Business, *resterrors.RestErr)
	DeleteBusiness(businessID int64) *resterrors.RestErr
}

// CompanyService holds a user service operations
type CompanyService interface {
	GetCompanies() (*[]entity.Companies, *resterrors.RestErr)
	GetCompanyByID(userID int64) (*entity.CompanyDetail, *resterrors.RestErr)
	CreateCompany(entity.CompanyDetail) (int64, *resterrors.RestErr)
	UpdateCompany(entity.CompanyDetail) (*entity.CompanyDetail, *resterrors.RestErr)
	DeleteCompany(userID int64) *resterrors.RestErr
}

// OrderService holds a user service operations
type OrderService interface {
	CreateOrder(entity.Order) (int64, *resterrors.RestErr)
	GetOrdersByUserID(userID int64) (*[]entity.OrdersByUserID, *resterrors.RestErr)
	GetOrderDetail(orderID int64) (*entity.OrderDetail, *resterrors.RestErr)
}

// RatingService holds a user service operations
type RatingService interface {
	GetCompanyUserRating(companyID, userID int64) (*entity.Rating, *resterrors.RestErr)
	UpdateRating(entity.Rating) (*entity.Rating, *resterrors.RestErr)
}
