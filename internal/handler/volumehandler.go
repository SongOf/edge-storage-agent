package handler

import (
	"encoding/json"
	"github.com/SongOf/edge-storage-core/pkg/eslog"
	"github.com/SongOf/edge-storage-core/stat/filesystem"
	"net/http"
	"strconv"
	"time"
)

func VolumeHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	sec, err := strconv.ParseInt(req.FormValue("seconds"), 10, 64)
	if sec <= 0 || err != nil {
		sec = 30
	}

	path := req.FormValue("path")
	fs := filesystem.New(nil, path)
	time.Sleep(time.Duration(sec) * time.Second)

	diskStat := fs.Once()

	//var buf strings.Builder
	//for k, v := range diskStat.Values() {
	//	buf.WriteString(fmt.Sprintf("%s=%v\n", k, v))
	//}
	resp, err := json.Marshal(diskStat)
	if err != nil {
		eslog.Err(err)
	}
	w.Write(resp)
}
