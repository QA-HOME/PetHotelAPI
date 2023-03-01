package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"v1/src/auth"
	"v1/src/db/postgre"
	"v1/src/errors"
	"v1/src/models/common"
)

func Register(w http.ResponseWriter, r *http.Request) {

	NotPost(w, r)
	SetContentJson(w)

	user := common.User{}

	body, err := io.ReadAll(r.Body)
	errors.CheckErr(err)

	err = json.Unmarshal([]byte(body), &user)
	errors.CheckErr(err)

	if !user.CheckRegisterRequest() {
		EmptyParameters(w, r)
	} else {

		user.UserPassword, err = auth.HashPassword(user.UserPassword)
		errors.CheckErr(err)

		err = postgre.InsertUser(user)
		errors.CheckErr(err)

		w.WriteHeader(http.StatusAccepted)

		m := make(map[string]string)
		m["status"] = "accepted"
		m["message"] = "register precess is successfully"
		jsonM, _ := json.Marshal(m)

		_, err := w.Write(jsonM)
		errors.CheckErr(err)
	}

}
