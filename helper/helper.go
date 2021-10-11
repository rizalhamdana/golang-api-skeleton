package helper

type Meta struct {
	Limit              int
	Offset             int
	CountAllData       int
	CountRetrievedData int
}

type Param struct {
	Limit  int
	Offset int
	Filter interface{}
}

type Filter struct {
	Search   string `query:"search"`
	Category string `query:"category"`
}
