package query

import "strings"

type OrderDir int

const (
	OrderDirASC OrderDir = iota
	OrderDirDESC
)

func OrderDirFromString(s string) OrderDir {
	s = strings.ToLower(s)
	if s == "desc" {
		return OrderDirDESC
	}

	return OrderDirASC
}

func (q OrderDir) String() string {
	switch q {
	case OrderDirASC:
		return "ASC"
	default:
		// OrderDirDESC
		return "DESC"
	}
}
