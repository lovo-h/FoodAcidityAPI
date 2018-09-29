package usecases

type InteractorFood struct {
	RepoFood RepoFood
	//	TODO: add logging
}

func (interactor *InteractorFood) OneFoodByNDB_No(ndb_no string) ([]map[string]string, error) {
	food, getErr := interactor.RepoFood.OneFoodByNdb_No(ndb_no)

	if getErr != nil {
		return nil, getErr
	}

	return food, getErr
}

func (interactor *InteractorFood) ManyLong_DescBySnippet(desc_snippet []string) ([]map[string]string, error) {
	long_descs, getErr := interactor.RepoFood.ManyLong_DescBySnippet(desc_snippet)

	if getErr != nil {
		return nil, getErr
	}

	return long_descs, getErr
}
