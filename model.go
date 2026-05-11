package goquery

type ModelInterface interface {
	Table_Name() TableName
	Table_Columns() []Column
	Table_PrimaryKey() Column
	Table_PrimaryKey_Value() int64
	Table_InsertColumns() []Column
	Table_UpdateColumns() []Column
	Table_Column_Types() map[Column]string
	String() string
	Update(db DBInterface) error
	Create(db DBInterface) error
	Delete(db DBInterface) error
	// FromID(db DBInterface, id int64) (ModelInterface, error)

	// Table_Column_Values() map[string]interface{}
}
