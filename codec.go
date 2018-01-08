package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func decodeRequest(r *http.Request, v interface{}) error {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(req, v)
	if err != nil {
		return err
	}
	return nil
}

func encodeResponse(w http.ResponseWriter, v interface{}) error {
	resp, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Apache/2.2.34'<!--")
	w.Header().Set("X-AspNet-Version", "2.0.50727'<!--")
	w.Header().Set("X-Powered-By", "PHP/5.5.38'<!--")
	_, _ = w.Write(resp)
	return nil
}

func decodeList(str *string) *[]string {
	if str == nil {
		return nil
	}
	list := []string{}
	for _, i := range strings.Split(*str, ", ") {
		list = append(list, strings.Replace(i, ",-", ",", -1))
	}
	return &list
}

func encodeList(list *[]string) *string {
	if list == nil {
		return nil
	}
	str := ""
	for _, i := range *list {
		str += ", "
		str += strings.Replace(i, ",", ",-", -1)
	}
	str = str[2:]
	return &str
}

func getUserID(r *http.Request) (*int64, error) {
	session, err := store.Get(r, "JSESSIONID")
	if err != nil {
		return nil, err
	}
	IUserID, ok := session.Values["userid"]
	if !ok {
		return nil, nil
	}
	userID, ok := IUserID.(int64)
	if !ok {
		return nil, nil
	}
	return &userID, nil
}

func setUserID(w http.ResponseWriter, r *http.Request, username string, userID int64) error {
	session, err := store.Get(r, "JSESSIONID")
	if err != nil {
		cookie, err := r.Cookie("JSESSIONID")
		if err != nil {
			return err
		}
		cookie.Value = ""
		session, err = store.Get(r, "JSESSIONID")
		if err != nil {
			return err
		}
	}
	session.Values["username"] = username
	session.Values["userid"] = userID
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
