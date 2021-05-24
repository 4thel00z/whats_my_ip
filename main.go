package main

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"github.com/gorilla/pat"

	j "encoding/json"
)

type Response struct {
	Error string `json:"error"`
	IP    string `json:"ip"`
}

var preSharedSecret string

// use JSON logging when run by Up (including `up start`).
func init() {

	if os.Getenv("UP_STAGE") == "" {
		log.SetHandler(text.Default)
	} else {
		log.SetHandler(json.Default)
	}
	preSharedSecret = os.Getenv("PRE_SHARED_SECRET")
	if preSharedSecret == "" {
		log.Fatal("PRE_SHARED_SECRET was not set!")
	}
	log.Info("Finished init()")

}

// setup.
func main() {
	addr := ":" + os.Getenv("PORT")
	app := pat.New()
	app.Get("/", get)
	log.Infof("Listening on %s", addr)
	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

// curl http://localhost:3000/
func get(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("Authorization")

	if subtle.ConstantTimeCompare([]byte(preSharedSecret), []byte(key)) == 0 {
		err := errors.New("invalid pre shared key")
		log.WithError(err).Errorf("\"%s\"", key)
		log.WithError(err).Errorf("expected \"%s\"", preSharedSecret)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		j.NewEncoder(w).Encode(Response{
			Error: err.Error(),
			IP:    "",
		})
		return
	}
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ips := strings.Split(xForwardedFor, ", ")

	if len(ips) <= 0 {
		err := errors.New(fmt.Sprintf("splitting of X-Forwarded-For header failed: %s", xForwardedFor))
		log.WithError(err).Error("failed to extract ip address")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		j.NewEncoder(w).Encode(Response{
			Error: err.Error(),
			IP:    "",
		})
		return
	}

	externalIp := ips[0]
	if externalIp == "" {
		log.Infof("X-Forwarded-For header is empty, using RemoteAddr instead %s", r.RemoteAddr)
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			externalIp = ip
		}
	}

	log.Infof("received request from %s", externalIp)
	w.Header().Set("Content-Type", "application/json")
	j.NewEncoder(w).Encode(Response{
		Error: "",
		IP:    externalIp,
	})
}
