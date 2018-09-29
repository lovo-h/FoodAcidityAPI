package interfaces

import (
	"testing"
	"reflect"
)

type MockHandlerDB struct {
	sqlStr      string
	queryParams []interface{}
}

func (repo *MockHandlerDB) Query(query string, queryParams []interface{}, fnSignal func([][]byte)) error {
	repo.sqlStr = query
	repo.queryParams = queryParams
	// fdgrp_desc fdgrp_cd
	fnSignal([][]byte{[]byte("70705"), []byte("raw almonds"), []byte("1200"), []byte("Nut and Seed Products")})
	fnSignal([][]byte{[]byte("70706"), []byte("baked almonds"), []byte("1200"), []byte("Nut and Seed Products")})

	return nil
}

// =-=-=-=-=-=-=-=-=-= INIT =-=-=-=-=-=-=-=-=-=

// =-=-=-=-=-=-=-=-=-= TESTS =-=-=-=-=-=-=-=-=-=

func TestRepoFood_ManyLong_DescBySnippet(t *testing.T) {
	repoFood := new(RepoFood)
	mockDB := new(MockHandlerDB)
	repoFood.HandlerDB = mockDB

	longDesc, getErr := repoFood.ManyLong_DescBySnippet([]string{"raw"})

	if getErr != nil {
		t.Error(getErr)
		return
	}

	expectedQueryStr := "SELECT ndb_no, long_desc, fdgrp_cd, fdgrp_desc FROM food WHERE long_desc LIKE $1 LIMIT 10;"

	if mockDB.sqlStr != expectedQueryStr || len(mockDB.queryParams) != 1 {
		t.Error("Incorrect query string || len(queryParams)")
		return
	}

	t0 := reflect.TypeOf(longDesc).String() != "[]map[string]string"
	t1 := len(longDesc) != 2
	t2 := longDesc[0]["ndb_no"] != "70705"
	t3 := longDesc[1]["ndb_no"] != "70706"

	if t0 || t1 || t2 || t3 {
		t.Error("invalid results")
	}
}
