package analyst

import (
	"log"
	"time"
)

type UserAnalyst struct {
	total map[string]bool
	error int
}

// KUid var KUid = "uid"
var KUid = "distinct_id"
var KCost = "cost"
var KTime = "time"
var StartTS = time.Date(2021, time.Month(9), 7, 0, 0, 0, 0, time.UTC)
var EndTS = time.Date(2021, time.Month(9), 7, 1, 50, 0, 0, time.UTC)

func (impl *UserAnalyst) Process(records []map[string]interface{}) interface{} {
	for _, record := range records {
		if record[KUid] == nil {
			impl.error++
			log.Fatal("[ERROR] unexpect record: ", record)
		}
		uid := record[KUid].(string)
		cost := record[KCost].(float64)
		tm := record[KTime].(time.Time)
		if len(uid) > 0 && cost > 0.5 &&
			tm.Unix() >= StartTS.Unix() && tm.Unix() <= EndTS.Unix() {
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
