package goquery

type FieldType int

const (
	FieldTypeBasic FieldType = iota
	FieldTypeRaw
	FieldTypeCount
	FieldTypeSum
	FieldTypeAvg
	FieldTypeMin
	FieldTypeMax
)

type Field struct {
	FieldType FieldType
	Name      Column
	As        string
	Raw       string
}

// NewField creates a new field.
//
//	NewField(FieldTypeBasic, "Foo")
//	NewField(FieldTypeBasic, "Foo", "Bar") <-- `Foo` AS `Bar`
func NewField(fieldType FieldType, column Column, opts ...string) *Field {

	as := ""

	if len(opts) > 0 {
		as = opts[0]
	}

	return &Field{
		FieldType: fieldType,
		Name:      column,
		As:        as,
		Raw:       "",
	}
}

// NewRawField creates a new field.
//
//	NewRawField("`t`.`Foo` AS `Bar`)
func NewRawField(raw string) *Field {

	return &Field{
		FieldType: FieldTypeRaw,
		Name:      "",
		As:        "",
		Raw:       raw,
	}
}

func (f *Field) String(alias string) string {

	as := ""

	if len(f.As) > 0 {
		as = " AS `" + f.As + "`"
	}

	switch f.FieldType {
	case FieldTypeCount:
		return "COUNT(`" + alias + "`.`" + string(f.Name) + "`)" + as
	case FieldTypeSum:
		return "COALESCE(SUM(`" + alias + "`.`" + string(f.Name) + "`), 0)" + as
	case FieldTypeAvg:
		return "COALESCE(AVG(`" + alias + "`.`" + string(f.Name) + "`), 0)" + as
	case FieldTypeMin:
		return "COALESCE(MIN(`" + alias + "`.`" + string(f.Name) + "`), 0)" + as
	case FieldTypeMax:
		return "COALESCE(MAX(`" + alias + "`.`" + string(f.Name) + "`), 0)" + as
	case FieldTypeRaw:
		return f.Raw
	// FieldTypeBasic
	default:
		return "`" + alias + "`.`" + string(f.Name) + "`" + as
	}
}
