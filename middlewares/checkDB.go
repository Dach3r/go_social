package middlewares

import (
	"net/http"

	"github.com/dach3r/vecker/db"
)

/*VerifyDB if exist connection with DB*/
func VerifyDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.VerifyConnection() == false {
			http.Error(w, "Lost connection with database", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
