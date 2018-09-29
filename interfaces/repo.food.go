package interfaces

import (
	"fmt"
	"strconv"
)

type RepoFood struct {
	HandlerDB HandlerDB
}

func (repo *RepoFood) OneFoodByNdb_No(ndb_no string) ([]map[string]string, error) {
	query := "SELECT ndb_no, nutr_no, nutr_val, units, nutrdesc FROM nutrition WHERE ndb_no = $1"
	queryParams := []interface{}{ndb_no}

	food := []map[string]string{}

	getValidString := func(data []byte) string {
		if data == nil {
			return ""
		}
		return string(data)
	}

	fn := func(rawResults [][]byte) {
		food = append(food, map[string]string{
			"ndb_no":   getValidString(rawResults[0]),
			"nutr_no":  getValidString(rawResults[1]),
			"nutr_val": getValidString(rawResults[2]),
			"units":    getValidString(rawResults[3]),
			"nutrdesc": getValidString(rawResults[4]),
		})
	}

	queryErr := repo.HandlerDB.Query(query, queryParams, fn)

	if queryErr != nil {
		fmt.Println("Error occurred in query:", queryErr)
		return nil, queryErr
	}

	return food, nil
}

func (repo *RepoFood) ManyLong_DescBySnippet(snippets []string) ([]map[string]string, error) {
	query := "SELECT ndb_no, long_desc, fdgrp_cd, fdgrp_desc FROM food WHERE long_desc ILIKE $1 "
	var queryParams []interface{}

	for idx, snip := range snippets {
		queryParams = append(queryParams, fmt.Sprintf("%s%s%s", "%%", snip, "%%"))

		if idx > 0 {
			query += "AND long_desc ILIKE $" + strconv.Itoa(idx+1) + " "
		}
	}

	query += "LIMIT 10;"

	long_descs := []map[string]string{}

	getValidString := func(data []byte) string {
		if data == nil {
			return ""
		}
		return string(data)
	}

	fn := func(rawResults [][]byte) {
		long_descs = append(long_descs, map[string]string{
			"ndb_no":     getValidString(rawResults[0]),
			"long_desc":  getValidString(rawResults[1]),
			"fdgrp_cd":   getValidString(rawResults[2]),
			"fdgrp_desc": getValidString(rawResults[3]),
		})
	}

	queryErr := repo.HandlerDB.Query(query, queryParams, fn)

	if queryErr != nil {
		return nil, queryErr
	}

	return long_descs, nil
}
