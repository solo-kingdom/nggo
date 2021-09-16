package analyst

import "log"

type UserAnalyst struct {
	total map[string]bool
	error int
}

var KUid = "distinct_id"
var KCost = "cost"

func (impl *UserAnalyst) Process(records []map[string]interface{}) interface{} {
	for _, record := range records {
		if record[KUid] == nil {
			impl.error++
			log.Fatal("[ERROR] unexpect record: ", record)
		}
		uid := record[KUid].(string)
		cost := record[KCost].(float64)
		if len(uid) > 0 && cost > 0.6 {
			impl.total[uid] = true
		}
	}
	return nil
}

func (impl *UserAnalyst) Show() {
	log.Print("total: ", len(impl.total), ", error: ", impl.error)
}

func GetUserAnalyst() Analyst {
	r := UserAnalyst{total: map[string]bool{}}
	return &r
}
