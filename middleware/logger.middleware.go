package middleware

import (
	errors2 "ecommerce/services/utils/errors"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/rs/xid"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
)

func accessLogger(h http.Handler) http.Handler {
	year, week := time.Now().ISOWeek()
	logFile := fmt.Sprintf("%s/access-w%v-%d.log", conf.Paths.Logs, week, year)
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return handlers.CombinedLoggingHandler(io.Writer(f), h)
	// Append request to access log and move to next middleware or complete the HTTP request
}

// errorLogger
func errorLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			var code int
			rec := recover()
			if rec != nil {
				if conf.Environment != "production" {
					fmt.Printf("Error Recovered\n%v\n", rec)
				}

				year, week := time.Now().ISOWeek()
				logFile := fmt.Sprintf("%s/access-w%v-%d.log", conf.Paths.Logs, week, year)
				f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer f.Close()

				if reflect.TypeOf(rec) == reflect.TypeOf(errors2.Error{}) {
					code = rec.(errors2.Error).Code
				} else {
					code = http.StatusInternalServerError
				}
				conf.Session.Values["error"] = code
				// Create a Unique ID for the error
				uuid := xid.NewWithTime(time.Now())
				logErr := log.New(f, "", log.LstdFlags)
				logErr.Println(fmt.Sprintf("| error_id:%v | remote_ip:%v | user_agent:%v | status_code:%v | path:%v query:%v | error:%v", uuid, r.RemoteAddr, code, r.UserAgent(), r.URL.Path, r.URL.RawQuery, rec))

				http.Redirect(w, r, "/error", http.StatusSeeOther)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
