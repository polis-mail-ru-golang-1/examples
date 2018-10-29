package index

type Index struct {
}

type Result struct {
	Name  string
	Score int
}

func New() Index {
	return Index{}
}

func (i Index) AddFile(name string) {}

func (i Index) Search(query string) []Result {
	return []Result{
		{
			Name:  "a",
			Score: 1,
		},
		{
			Name:  "b",
			Score: 2,
		},
	}
}

func (i Index) Count() int {
	return 2
}
