package testassets

import (
	"encoding/json"

	goquery "github.com/macinnir/query"
	"gopkg.in/guregu/null.v3"
)

const (

	// ListItem_SchemaName is the name of the schema group this model is in
	ListItem_SchemaName = "rfq"

	// Comment_TableName
	Comment_TableName goquery.TableName = "Comment"

	// Columns
	Comment_Column_CommentID   goquery.Column = "CommentID"
	Comment_Column_DateCreated goquery.Column = "DateCreated"
	Comment_Column_IsDeleted   goquery.Column = "IsDeleted"
	Comment_Column_Content     goquery.Column = "Content"
	Comment_Column_ObjectType  goquery.Column = "ObjectType"
	Comment_Column_ObjectID    goquery.Column = "ObjectID"
	Comment_Column_Name        goquery.Column = "Name"
)

// Comment is a `Comment` data model
type Comment struct {
	CommentID   int64       `db:"CommentID" json:"CommentID"`
	DateCreated int64       `db:"DateCreated" json:"DateCreated"`
	IsDeleted   int         `db:"IsDeleted" json:"IsDeleted"`
	Content     null.String `db:"Content" json:"Content"`
	ObjectType  int64       `db:"ObjectType" json:"ObjectType"`
	ObjectID    int64       `db:"ObjectID" json:"ObjectID"`
	Name        string      `db:"Name" json:"Name"`
}

// Comment_TableName is the name of the table
func (c *Comment) Table_Name() goquery.TableName {
	return Comment_TableName
}

func (c *Comment) Table_Columns() []goquery.Column {
	return []goquery.Column{
		Comment_Column_CommentID,
		Comment_Column_DateCreated,
		Comment_Column_IsDeleted,
		Comment_Column_Content,
		Comment_Column_ObjectType,
		Comment_Column_ObjectID,
		Comment_Column_Name,
	}
}

func (c *Comment) Table_Column_Types() map[goquery.Column]string {
	return map[goquery.Column]string{
		"CommentID":   "%d",
		"DateCreated": "%d",
		"IsDeleted":   "%d",
		"Content":     "%s",
		"ObjectType":  "%d",
		"ObjectID":    "%d",
		"Name":        "%s",
	}
}

func (c *Comment) Table_Column_Values() map[goquery.Column]interface{} {
	return map[goquery.Column]interface{}{
		"CommentID":   c.CommentID,
		"DateCreated": c.DateCreated,
		"IsDeleted":   c.IsDeleted,
		"Content":     c.Content.String,
		"ObjectType":  c.ObjectType,
		"ObjectID":    c.ObjectID,
		"Name":        c.Name,
	}
}

// Comment_PrimaryKey is the name of the table's primary key
func (c *Comment) Table_PrimaryKey() goquery.Column {
	return Comment_Column_CommentID
}

func (c *Comment) Table_PrimaryKey_Value() int64 {
	return c.CommentID
}

// Comment_InsertColumns is a list of all insert columns for this model
func (c *Comment) Table_InsertColumns() []goquery.Column {
	return []goquery.Column{
		Comment_Column_DateCreated,
		Comment_Column_Content,
		Comment_Column_ObjectType,
		Comment_Column_ObjectID,
		Comment_Column_Name,
	}
}

// Comment_UpdateColumns is a list of all update columns for this model
func (c *Comment) Table_UpdateColumns() []goquery.Column {
	return []goquery.Column{
		Comment_Column_Content,
		Comment_Column_ObjectType,
		Comment_Column_ObjectID,
		Comment_Column_Name,
	}
}

func (c *Comment) String() string {
	str, _ := json.Marshal(c)
	return string(str)
}

func (c *Comment) Create(db goquery.IDB) error {
	return nil
}

func (c *Comment) Update(db goquery.IDB) error {
	return nil
}

func (c *Comment) Delete(db goquery.IDB) error {
	return nil
}

func (c *Comment) FromID(db goquery.IDB, id int64) (goquery.IModel, error) {
	return nil, nil
}
