package analyst

import "log"

type RequestAnalyst struct {
	total map[string]bool
}

var KLogId = "log_id"

func (impl *RequestAnalyst) Process(records []map[string]interface{}) interface{} {
	for _, record := range records {
		uid := record[KLogId].(string)
		if len(uid) > 0 {
			impl.total[uid] = true
		}
	}
	return nil
}

func (impl *RequestAnalyst) Show() {
	log.Print("total: ", len(impl.total))
}

func GetRequestAnalyst() Analyst {
	r := UserAnalyst{total: map[string]bool{}}
	return &r
}
