package model

import (
	"errors"
	"fmt"
	"strings"
)

type (
	ListQuery struct {
		Filters map[string]interface{} `json:"filters"`
		Search  map[string]interface{} `json:"search"`
		Orders  map[string]interface{} `json:"orders"`
		Offset  int64                  `json:"offset"`
		Limit   int64                  `json:"limit"`
	}
)

func (l *ListQuery) Query(colums []string) (string, []interface{}, error) {
	pos := 1
	where := "WHERE"
	var args []interface{}

	filtersQuery := ""
	for k, v := range l.Filters {

		if !l.isExist(k, colums) {
			return "", args, errors.New("column not found!")
		}

		filtersQuery = fmt.Sprintf("%s %s = $%d", filtersQuery, k, pos)
		filtersQuery = fmt.Sprintf("%s AND", filtersQuery)
		args = append(args, v)
		pos++
	}

	searchQuery := ""
	for k, v := range l.Search {

		if !l.isExist(k, colums) {
			return "", args, errors.New("column not found!")
		}

		searchQuery = fmt.Sprintf("%s %s LIKE $%d", searchQuery, k, pos)
		searchQuery = fmt.Sprintf("%s AND", searchQuery)
		args = append(args, "%"+fmt.Sprint(v)+"%")
		pos++
	}
	if searchQuery == "" && filtersQuery != "" {
		filtersQuery = strings.TrimRight(filtersQuery, " AND")
	}

	searchQuery = strings.TrimRight(searchQuery, " AND")

	ordersQuery := ""
	for k, v := range l.Orders {

		if !l.isExist(k, colums) {
			return "", args, errors.New("column not found!")
		}

		ordersQuery = fmt.Sprintf("ORDER BY %s %v", k, v)
		break
	}

	if filtersQuery == "" && searchQuery == "" {
		where = ""
	}

	query := fmt.Sprintf(" %s%s%s %s OFFSET $%d LIMIT $%d", where, filtersQuery, searchQuery, ordersQuery, pos, pos+1)
	args = append(args, l.Offset)
	args = append(args, l.Limit)

	return query, args, nil
}

func (l *ListQuery) isExist(n string, list []string) bool {
	for _, v := range list {
		if n == v {
			return true
		}
	}

	return false
}
