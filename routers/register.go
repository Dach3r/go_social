package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dach3r/vecker/db"
	"github.com/dach3r/vecker/models"
)

/*Register can register user on db*/
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 || len(user.Password) < 8 {
		http.Error(w, "Email or Password not can't be blank", 401)
		return
	}

	_, found, _ := db.VerifyUserExist(user.Email)

	if found == true {
		http.Error(w, "This user exist on DB", 401)
		return
	}

	_, status, err := db.InsertRegister(user)

	if err != nil {
		http.Error(w, "Error: "+err.Error(), 401)
		return
	}

	if status == false {
		http.Error(w, "Register is invalid", 401)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
