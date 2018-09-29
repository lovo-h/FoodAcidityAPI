package usecases

import (
	"testing"
	"strings"
)

type MockRepoFood struct {
	long_descs []map[string]string
	foods       map[string]map[string]map[string]string
}

func (repo *MockRepoFood) OneFoodByNdb_No(getNdbNo string) ([]map[string]string, error) {
	retFood := []map[string]string{}

	if food, ok := repo.foods[getNdbNo]; ok {
		proteinData := food["203"]
		calciumData := food["301"]

		proteinNutrData := map[string]string{}
		calciumNutrData := map[string]string{}

		proteinNutrData["ndb_no"] = getNdbNo
		proteinNutrData["nutr_no"] = "203"
		proteinNutrData["nutr_val"] = proteinData["nutr_val"]
		proteinNutrData["units"] = proteinData["units"]
		proteinNutrData["nutrdesc"] = proteinData["nutr_desc"]

		calciumNutrData["ndb_no"] = getNdbNo
		calciumNutrData["nutr_no"] = "301"
		calciumNutrData["nutr_val"] = calciumData["nutr_val"]
		calciumNutrData["units"] = calciumData["units"]
		calciumNutrData["nutrdesc"] = calciumData["nutr_desc"]

		retFood = append(retFood, proteinNutrData, calciumNutrData)
	}

	return retFood, nil
}

func (repo *MockRepoFood) ManyLong_DescBySnippet(snippet []string) ([]map[string]string, error) {
	retlong_descsArr := []map[string]string{}
	for _, long_descDict := range repo.long_descs {
		long_desc := long_descDict["long_desc"]

		for _, snip := range snippet {
			if strings.Contains(strings.ToLower(long_desc), strings.ToLower(snip)) {
				retlong_descsArr = append(retlong_descsArr, long_descDict)
			}
		}


	}

	return retlong_descsArr, nil
}

// =-=-=-=-=-=-=-=-=-= INIT =-=-=-=-=-=-=-=-=-=

func addFood2Map(foodsMap map[string]map[string]map[string]string, ndb_no, nutr_val_a, nutr_val_b string) {
	foodsMap[ndb_no] = map[string]map[string]string{
		"203": {
			"nutr_val":  nutr_val_a,
			"units":     "g",
			"nutr_desc": "Protein",
		},
		"301": {
			"nutr_val":  nutr_val_b,
			"units":     "mg",
			"nutr_desc": "Calcium, Ca",
		},
	}
}

func GetFoods() map[string]map[string]map[string]string {
	foods := map[string]map[string]map[string]string{}

	addFood2Map(foods, "12061", "21.15", "269")
	addFood2Map(foods, "11529", "0.88", "10")
	addFood2Map(foods, "05006", "18.6", "11")
	addFood2Map(foods, "11739", "2.98", "48")

	return foods
}

func newLong_descMap(ndb_no, long_desc, fdgrp_cd, fdgrp_desc string) map[string]string {
	return map[string]string{
		"ndb_no":     ndb_no,
		"long_desc":  long_desc,
		"fdgrp_cd":   fdgrp_cd,
		"fdgrp_desc": fdgrp_desc,
	}
}

func GetLongDescs() []map[string]string {
	long_descs := []map[string]string{}

	long_descs = append(long_descs, newLong_descMap("12061", "Nuts, almonds", "1200", "Nut and Seed Products"))
	long_descs = append(long_descs, newLong_descMap("11529", "Tomatoes, red, ripe, raw, year round average", "1100", "Vegetables and Vegetable Products"))
	long_descs = append(long_descs, newLong_descMap("05006", "	Chicken, broilers or fryers, meat and skin, raw", "500", "Poultry Products"))
	long_descs = append(long_descs, newLong_descMap("11739", "Broccoli, leaves, raw", "1100", "Vegetables and Vegetable Products"))

	return long_descs
}

func newMockRepoFood() *MockRepoFood {
	repo := new(MockRepoFood)
	repo.long_descs = GetLongDescs()
	repo.foods = GetFoods()

	return repo
}

// =-=-=-=-=-=-=-=-=-= TESTS =-=-=-=-=-=-=-=-=-=

func TestInteractorFood_ManyLong_DescBySnippet(t *testing.T) {
	interactorFood := new(InteractorFood)
	interactorFood.RepoFood = newMockRepoFood()

	long_descs, getErr := interactorFood.ManyLong_DescBySnippet([]string{"raw"})

	if getErr != nil || len(long_descs) != 3 {
		t.Error("getErr OR len != 3")
	}
}

func TestInteractorFood_OneFoodByNDB_No(t *testing.T) {
	interactorFood := new(InteractorFood)
	interactorFood.RepoFood = newMockRepoFood()

	food, getErr := interactorFood.OneFoodByNDB_No("12061")

	// Fail if there's an error OR no food is returned OR if it's the wrong food
	if getErr != nil || len(food) == 0 || food[0]["nutr_val"] != "21.15" {
		t.Error("getErr OR nutr_val != '21.15'")
		return
	}
}
