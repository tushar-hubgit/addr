package logic

import (
	"encoding/json"
	"fmt"
	"addr/functions"
	"net/http"
	"strings"
)

type AppLogic struct {
	AppFunctions functions.AppFunctions
}

func (l *AppLogic) GetIP(w http.ResponseWriter, r *http.Request) {
	ip := l.AppFunctions.ExtractIP(r.RemoteAddr)

	format := r.URL.Query().Get("format")

	if strings.Compare(format, "json") == 0 {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			map[string]string{
				"ip": ip,
			},
		)
		return
	}

	fmt.Fprint(w, ip)
}
