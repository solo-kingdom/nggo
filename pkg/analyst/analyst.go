package analyst

type Analyst interface {
	Process(records []map[string]interface{}) interface{}
	Show()
}
