package logic

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type AppLogic struct {
}

func (l *AppLogic) GetIP(w http.ResponseWriter, r *http.Request) {
	ip := net.ParseIP(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]).String()

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
