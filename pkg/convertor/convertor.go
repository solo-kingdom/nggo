package convertor

type Convertor struct {
	Run     func(string) interface{}
	IsQuote bool
	Name    string
}
