package testassets

import (
	"encoding/json"

	goquery "github.com/macinnir/query"
)

const (
	JobSales_TableName goquery.TableName = "JobSales"

	JobSales_Column_JobSalesID        goquery.Column = "JobSalesID"
	JobSales_Column_JobID             goquery.Column = "JobID"
	JobSales_Column_UserID            goquery.Column = "UserID"
	JobSales_Column_CommissionPercent goquery.Column = "CommissionPercent"
	JobSales_Column_DateCreated       goquery.Column = "DateCreated"
	JobSales_Column_IsDeleted         goquery.Column = "IsDeleted"
	JobSales_Column_CommissionDollars goquery.Column = "CommissionDollars"
	JobSales_Column_IsHouse           goquery.Column = "IsHouse"
)

// JobSales is a `JobSales` data model
type JobSales struct {
	JobSalesID        int64   `db:"JobSalesID" json:"JobSalesID"`
	JobID             int64   `db:"JobID" json:"JobID"`
	UserID            int64   `db:"UserID" json:"UserID"`
	CommissionPercent float64 `db:"CommissionPercent" json:"CommissionPercent"`
	DateCreated       int64   `db:"DateCreated" json:"DateCreated"`
	IsDeleted         int     `db:"IsDeleted" json:"IsDeleted"`
	CommissionDollars float64 `db:"CommissionDollars" json:"CommissionDollars"`
	IsHouse           int     `db:"IsHouse" json:"IsHouse"`
}

// Comment_TableName is the name of the table
func (c *JobSales) Table_Name() goquery.TableName {
	return JobSales_TableName
}

func (c *JobSales) Table_Columns() []goquery.Column {
	return []goquery.Column{
		JobSales_Column_JobSalesID,
		JobSales_Column_JobID,
		JobSales_Column_UserID,
		JobSales_Column_CommissionPercent,
		JobSales_Column_DateCreated,
		JobSales_Column_IsDeleted,
		JobSales_Column_CommissionDollars,
		JobSales_Column_IsHouse,
	}
}

func (c *JobSales) Table_Column_Types() map[goquery.Column]string {
	return map[goquery.Column]string{
		"JobSalesID":        "%d",
		"JobID":             "%d",
		"UserID":            "%d",
		"CommissionPercent": "%d",
		"DateCreated":       "%d",
		"IsDeleted":         "%d",
		"CommissionDollars": "%d",
		"IsHouse":           "%d",
	}
}

func (c *JobSales) Table_Column_Values() map[goquery.Column]interface{} {
	return map[goquery.Column]interface{}{
		"JobSalesID":        c.JobSalesID,
		"JobID":             c.JobID,
		"UserID":            c.UserID,
		"CommissionPercent": c.CommissionPercent,
		"DateCreated":       c.DateCreated,
		"IsDeleted":         c.IsDeleted,
		"CommissionDollars": c.CommissionDollars,
		"IsHouse":           c.IsHouse,
	}
}

// Comment_PrimaryKey is the name of the table's primary key
func (c *JobSales) Table_PrimaryKey() goquery.Column {
	return "JobSalesID"
}

func (c *JobSales) Table_PrimaryKey_Value() int64 {
	return c.JobSalesID
}

// Comment_InsertColumns is a list of all insert columns for this model
func (c *JobSales) Table_InsertColumns() []goquery.Column {
	return []goquery.Column{"DateCreated", "Content", "ObjectType", "ObjectID"}
}

// Comment_UpdateColumns is a list of all update columns for this model
func (c *JobSales) Table_UpdateColumns() []goquery.Column {
	return []goquery.Column{"Content", "ObjectType", "ObjectID"}
}

func (c *JobSales) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c *JobSales) Create(db goquery.IDB) error {
	return nil
}

func (c *JobSales) Update(db goquery.IDB) error {
	return nil
}

func (c *JobSales) Delete(db goquery.IDB) error {
	return nil
}

func (c *JobSales) FromID(db goquery.IDB, id int64) (goquery.IModel, error) {
	return nil, nil
}
