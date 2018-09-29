package domain

type Food struct {
	Ndb_no      string
	Long_desc   string
	Fdgrp_cd    string
	Fdgrp_desc  string

	Nutrients []Nutrient
}
