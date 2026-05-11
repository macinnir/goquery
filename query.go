package goquery

import (
	"errors"
	"fmt"
	"strings"
)

type Column string
type TableName string

type Q struct {
	fields      []*Field
	noAlias     []int
	alias       string
	raw         string
	model       ModelInterface
	queryType   QueryType
	where       *whereClause
	limit       int64
	offset      int64
	orderBy     [][]string
	groupBy     []string
	setSorter   []Column
	sets        map[Column]interface{}
	columnTypes map[Column]string
	errors      []string
	inst        int64
}

func Query(model ModelInterface) *Q {
	return &Q{
		fields:      []*Field{},
		noAlias:     []int{},
		model:       model,
		orderBy:     [][]string{},
		groupBy:     []string{},
		setSorter:   []Column{},
		sets:        map[Column]interface{}{},
		columnTypes: model.Table_Column_Types(),
		alias:       "t",
		errors:      []string{},
		inst:        0,
	}
}

func (q *Q) Alias(alias string) *Q {
	q.alias = alias
	return q
}

func (q *Q) Limit(limit, offset int64) *Q {
	q.limit = limit
	q.offset = offset
	return q
}

func (q *Q) LimitPage(limit, page int64) *Q {
	q.limit = limit
	q.offset = limit * page
	return q
}

func (q *Q) OrderBy(col Column, dir OrderDir) *Q {
	q.orderBy = append(q.orderBy, []string{string(col), dir.String()})
	return q
}

func (q *Q) GroupBy(col ...Column) *Q {
	for _, c := range col {
		q.groupBy = append(q.groupBy, string(c))
	}
	return q
}

func (q *Q) GroupByColumns(cols ...Column) *Q {
	for _, c := range cols {
		q.groupBy = append(q.groupBy, string(c))
	}
	return q
}

// func Save(model ModelInterface) *Q {

// 	var q *Q
// 	colMap := model.Table_Column_Values()

// 	if model.Table_PrimaryKey_Value() > 0 {
// 		q = Update(model)
// 		updateColumns := model.Table_UpdateColumns()
// 		for _, col := range updateColumns {
// 			q.Set(col, colMap[col])
// 		}
// 		q.Where(
// 			EQ(model.Table_PrimaryKey(), model.Table_PrimaryKey_Value()),
// 		)
// 	} else {
// 		q = Insert(model)
// 		insertColumns := model.Table_InsertColumns()
// 		for _, col := range insertColumns {
// 			q.Set(col, colMap[col])
// 		}
// 	}

// 	return q
// }

// func Destroy(model ModelInterface) *Q {

// 	var q *Q
// 	colMap := model.Table_Column_Values()

// 	if colMap[model.Table_PrimaryKey()].(int64) > 0 {
// 		q = Delete(model)
// 		q.Where(
// 			EQ(model.Table_PrimaryKey(), model.Table_PrimaryKey_Value()),
// 		)
// 	}
// 	return q

// }

type QueryErrorType string

const (
	QUERY_ERROR_INVALID_VALUE      QueryErrorType = "Invalid value"
	QUERY_ERROR_INVALID_COLUMN     QueryErrorType = "Invalid Column Name"
	QUERY_ERROR_EMPTY_WHERE_CLAUSE QueryErrorType = "Empty where clause"
)

func (q *Q) String() (string, error) {

	var sb strings.Builder

	switch q.queryType {
	case QueryTypeRaw:
		sb.WriteString(q.raw)
		return sb.String(), nil
	case QueryTypeSelect:

		sb.WriteString("SELECT ")

		if len(q.fields) > 0 {
			for k := range q.fields {
				sb.WriteString(q.fields[k].String(q.alias))
				if k < len(q.fields)-1 {
					sb.WriteString(", ")
				}
			}
		} else {
			sb.WriteString("`" + q.alias + "`.*")
		}

		sb.WriteString(" FROM")
	case QueryTypeInsert:
		sb.WriteString("INSERT INTO")
		q.alias = ""
	case QueryTypeUpdate:
		sb.WriteString("UPDATE")
		q.alias = ""
	case QueryTypeDelete:
		sb.WriteString("DELETE FROM")
		q.alias = ""
	}

	sb.WriteString(" `" + string(q.model.Table_Name()) + "`")

	if len(q.alias) > 0 && q.queryType == QueryTypeSelect {
		sb.WriteString(" `" + q.alias + "`")
	}

	if q.queryType == QueryTypeUpdate && len(q.sets) > 0 {

		sb.WriteString(" SET ")
		// sql += " SET "

		setStmts := []string{}

		for k := range q.setSorter {

			colName := q.setSorter[k]

			val := fmt.Sprint(q.sets[colName])

			if _, ok := q.columnTypes[colName]; !ok {
				q.errorInvalidColumn(QUERY_ERROR_INVALID_COLUMN, "UPDATE...SET", string(colName))
				val = "'" + val + "'"
			} else {
				if q.columnTypes[colName] == "%s" {
					val = "'" + EscapeString(val) + "'"
				}
			}

			setStmts = append(setStmts, q.col(string(colName))+" = "+val)
		}

		sb.WriteString(strings.Join(setStmts, ", "))
	}

	if q.queryType == QueryTypeInsert && len(q.sets) > 0 {

		cols := []string{}
		vals := []string{}

		for k := range q.setSorter {

			colName := q.setSorter[k]

			if _, ok := q.columnTypes[colName]; !ok {
				q.errorInvalidColumn(QUERY_ERROR_INVALID_COLUMN, "INSERT...SET", string(colName))
			}

			val := fmt.Sprint(q.sets[colName])

			if q.columnTypes[colName] == "%s" {
				val = "'" + EscapeString(val) + "'"
			}

			cols = append(cols, q.col(string(colName)))
			vals = append(vals, val)
		}

		sb.WriteString(" ( " + strings.Join(cols, ", ") + " ) VALUES ( " + strings.Join(vals, ", ") + " )")
		// sql += " ( " + strings.Join(cols, ", ") + " ) VALUES ( " + strings.Join(vals, ", ") + " )"
	}

	if q.where != nil && q.queryType != QueryTypeInsert {
		whereClause := q.printWhereClause(q.columnTypes, q.where.WhereParts)
		if len(whereClause) > 0 {
			sb.WriteString(" WHERE ")
			sb.WriteString(whereClause)
		}
	}

	if q.queryType == QueryTypeSelect {

		if len(q.groupBy) > 0 {
			groupBys := make([]string, len(q.groupBy))
			for k := range q.groupBy {

				// Validate the group by column
				if _, ok := q.columnTypes[Column(q.groupBy[k])]; !ok {
					q.errorInvalidColumn(QUERY_ERROR_INVALID_COLUMN, "GROUP BY", q.groupBy[k])
				}

				groupBys[k] = q.col(q.groupBy[k])
			}
			sb.WriteString(" GROUP BY " + strings.Join(groupBys, ", "))
		}

		if len(q.orderBy) > 0 {
			orderBys := make([]string, len(q.orderBy))
			for k := range q.orderBy {

				// Validate the order by column
				if _, ok := q.columnTypes[Column(q.orderBy[k][0])]; !ok {
					q.errorInvalidColumn(QUERY_ERROR_INVALID_COLUMN, "ORDER BY", q.orderBy[k][0])
				}

				orderBys[k] = q.col(q.orderBy[k][0]) + " " + strings.ToUpper(q.orderBy[k][1])
			}
			sb.WriteString(" ORDER BY " + strings.Join(orderBys, ", "))
		}
	}

	if q.limit > 0 {
		sb.WriteString(" LIMIT " + fmt.Sprint(q.limit))
	}

	if q.offset > 0 {
		sb.WriteString(" OFFSET " + fmt.Sprint(q.offset))
	}

	var e error

	if len(q.errors) > 0 {
		e = errors.New(strings.Join(q.errors, "\n--"))
	}

	return sb.String(), e
}

func (q *Q) col(colName string) string {
	if len(q.alias) > 0 {
		return "`" + q.alias + "`.`" + colName + "`"
	}
	return "`" + string(colName) + "`"
}

func (q *Q) error(err string) {
	q.errors = append(q.errors, err)
}

func (q *Q) errorInvalidColumn(errType QueryErrorType, queryErrorLocation, comment string) {
	q.error(fmt.Sprintf("%s at %s in model `%s` -- %s", errType, queryErrorLocation, q.model.Table_Name(), comment))
}

// Set adds a column and value to be set in an update or insert query
func (q *Q) Set(fieldName Column, value interface{}) *Q {
	if _, ok := q.sets[fieldName]; !ok {
		q.sets[fieldName] = value
		q.setSorter = append(q.setSorter, fieldName)
	}
	return q
}
