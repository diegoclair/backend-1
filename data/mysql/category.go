package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Team-73/backend/domain/entity"
	"github.com/Team-73/backend/utils/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type categoryRepo struct {
	db *sql.DB
}

// newCategoryRepo returns a instance of dbrepo
func newCategoryRepo(db *sql.DB) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

const queryCategorySelectBase = `
					SELECT 	c.id,
									c.name
								
					FROM 		tab_product_category 		c 
					`

func (s *categoryRepo) parseCategorySet(rows *sql.Rows) (categories []entity.Category, err error) {
	for rows.Next() {
		category := entity.Category{}
		category, err = s.parseCategory(rows)
		if err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *categoryRepo) parseCategory(row scanner) (category entity.Category, err error) {

	err = row.Scan(
		&category.ID,
		&category.Name,
	)

	if err != nil {
		return category, err
	}

	return category, nil
}

//GetCategories - return a list os categories
func (s *categoryRepo) GetCategories() (*[]entity.Category, *resterrors.RestErr) {

	query := queryCategorySelectBase

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0001: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetCategoryByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var categories []entity.Category

	rows, err := stmt.Query()
	if err != nil {
		errorCode := "Error 0002: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in GetCategorys", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer rows.Close()

	categories, err = s.parseCategorySet(rows)
	if err != nil {
		errorCode := "Error 0003: "
		log.Println(fmt.Sprintf("%sError when trying to parse result in parseCategorySet", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &categories, nil
}

//GetCategoryByID - get a category by ID
func (s *categoryRepo) GetCategoryByID(id int64) (*entity.Category, *resterrors.RestErr) {

	query := queryCategorySelectBase + `
		WHERE 	c.id 		= ?;`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0004: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in GetCategoryByID", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var category entity.Category

	result := stmt.QueryRow(id)
	category, err = s.parseCategory(result)
	if err != nil {
		errorCode := "Error 0005: "
		log.Println(fmt.Sprintf("%sError when trying to execute QueryRow in GetCategoryByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &category, nil
}

// Create - to create a category on database
func (s *categoryRepo) Create(category entity.Category) (int64, *resterrors.RestErr) {

	query := `
		INSERT INTO tab_product_category (
			name) VALUES
		(?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0006: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Create a category", errorCode), err)
		return 0, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(
		category.Name)
	if err != nil {
		errorCode := "Error 0007: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Create category", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	categoryID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0008: "
		log.Println(fmt.Sprintf("%sError when trying to get LastInsertId in the Create category", errorCode), err)
		return 0, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return categoryID, nil
}

// Update - to update a category on database
func (s *categoryRepo) Update(category entity.Category) (*entity.Category, *resterrors.RestErr) {

	query := `
		UPDATE tab_product_category
			SET	name 												= ?
			
		WHERE id	= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0009: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Update a category", errorCode), err)
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		category.Name,
		category.ID)
	if err != nil {
		errorCode := "Error 0010: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Update category", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &category, nil
}

// Delete - to delete a category on database
func (s *categoryRepo) Delete(id int64) *resterrors.RestErr {

	query := `
		DELETE FROM tab_product_category
		WHERE 	id			= ?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorCode := "Error 0011: "
		log.Println(fmt.Sprintf("%sError when trying to prepare the query statement in the Delete category", errorCode), err)
		return resterrors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		errorCode := "Error 0012: "
		log.Println(fmt.Sprintf("%sError when trying to execute Query in the Delete category", errorCode), err)
		return mysqlutils.HandleMySQLError(errorCode, err)
	}

	return nil
}
