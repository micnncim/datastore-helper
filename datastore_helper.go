package datastorehelper

type Query struct {
	Cursor  string
	Limit   int
	Filters []Filter
	Orders  []string
}

type Filter struct {
	Property string
	Operator string
	Value    interface{}
}
