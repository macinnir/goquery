package goquery

// Custom type
type Column string
type TableName string

type QueryType int

const (
	QueryTypeNotSet QueryType = iota
	QueryTypeSelect
	QueryTypeRaw
	QueryTypeUpdate
	QueryTypeDelete
	QueryTypeInsert
)

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

type Col struct {
	Name string
	Type string
}

type BaseModel struct {
	Columns       []Col
	PrimaryKey    Col
	InsertColumns []Col
	UpdateColumns []Col
	ColumnTypes   map[Col]string
}

func (b *BaseModel) Select() *BaseModel {
	return b
}

type TestModel struct {
	BaseModel
}

func NewTestModel() *TestModel {
	var m = &TestModel{}

	m.Select()
}
