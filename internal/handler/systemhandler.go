package handler

import (
	"encoding/json"
	"github.com/SongOf/edge-storage-core/pkg/eslog"
	"github.com/SongOf/edge-storage-core/stat/system"
	"net/http"
	"strconv"
	"time"
)

func SystemHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	sec, err := strconv.ParseInt(req.FormValue("seconds"), 10, 64)
	if sec <= 0 || err != nil {
		sec = 30
	}

	sc := system.New(nil)
	time.Sleep(time.Duration(sec) * time.Second)

	sstats := sc.Once()

	//var buf strings.Builder
	//for k, v := range sstats.Values() {
	//	buf.WriteString(fmt.Sprintf("%s=%v\n", k, v))
	//}
	resp, err := json.Marshal(sstats)
	if err != nil {
		eslog.Err(err)
	}
	w.Write(resp)
}
