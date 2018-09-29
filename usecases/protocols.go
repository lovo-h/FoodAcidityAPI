package usecases


type M map[string]interface{}

type RepoFood interface {
	OneFoodByNdb_No(string) ([]map[string]string, error)
	ManyLong_DescBySnippet([]string) ([]map[string]string, error)
}

type HandlerLogger interface {
	Log(string) error
}
