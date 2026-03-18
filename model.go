package goquery

type IModel interface {
	Table_Name() TableName
	Table_Columns() []Column
	Table_PrimaryKey() Column
	Table_PrimaryKey_Value() int64
	Table_InsertColumns() []Column
	Table_UpdateColumns() []Column
	Table_Column_Types() map[Column]string
	String() string
	// Update(db IDB) error
	// Create(db IDB) error
	// Delete(db IDB) error
	// FromID(db IDB, id int64) (IModel, error)
}
