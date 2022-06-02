package oauth2

import (
	"encoding/json"
	"fjpc/go-oauth2/config"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func LoginUser(w http.ResponseWriter, r *http.Request, config config.Config) {
	protocol := r.Header.Get("x-forwarded-proto")
	host := r.Header.Get("x-forwarded-host")
	path := r.URL.Path
	cookie := &http.Cookie{
		Name:  config.Cookies.ForwardedTo,
		Value: fmt.Sprintf("%v://%v%v", protocol, host, path),
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, config.OAuth2.Login.Url, http.StatusFound)
}

func GetUserInfo(token string, config config.OAuth2) (*UserInfo, error) {
	var response *http.Response
	client := &http.Client{}
	request, err := http.NewRequest("GET", config.UserInfo.Url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 401 {
		return nil, &UnauthorizedTokenError{}
	}
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	user := &UserInfo{}
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

