package service

import (
	"encoding/json"
	"fmt"
	"github.com/ndphu/fm/dao"
	"github.com/ndphu/fm/model"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	APP_TOKEN      = os.Getenv("FACEBOOK_APP_TOKEN")
	APP_SECRET     = os.Getenv("FACEBOOK_APP_SECRET")
	APP_ID         = "568365100192445"
	LOGIN_CALLBACK = "http://localhost:8080/pfm/login/callback"
)

type LoginService struct {
	db *dao.DAO
}

func NewLoginService(db *dao.DAO) *LoginService {
	return &LoginService{
		db: db,
	}
}

func (s *LoginService) ProcessAccessCode(code string) {
	accessToken, err := s.GetAccessTokenFromAccessCode(code)
	if err != nil {
		panic(err)
	}

	fmt.Println("Access Token = " + accessToken)

	fat, err := s.ValidateFacebookAcessToken(accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", fat.UserId)
}

func (s *LoginService) GetAccessTokenFromAccessCode(code string) (string, error) {
	fmt.Println("App Token = ", APP_TOKEN)
	fmt.Println("App Secret = ", APP_SECRET)
	getTokenUrl := fmt.Sprintf("https://graph.facebook.com/v2.12/oauth/access_token?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s",
		APP_ID, APP_SECRET, LOGIN_CALLBACK, code)
	fmt.Println("Get Token URL " + getTokenUrl)
	response, err := http.Get(getTokenUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(content))

	cvr := model.CodeValidationResponse{}

	err = json.Unmarshal(content, &cvr)
	//TODO: should validate expires in field here
	return cvr.AccessToken, err
}

func (s *LoginService) ValidateFacebookAcessToken(fbAccessToken string) (*model.FacebookAccessToken, error) {
	fbUrl := fmt.Sprintf("https://graph.facebook.com/debug_token?input_token=%s&access_token=%s", fbAccessToken, APP_TOKEN)
	fmt.Println("FB URL = " + fbUrl)
	response, err := http.Get(fbUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fbResponse := model.FacebookValidateResponse{}

	err = json.Unmarshal(content, &fbResponse)
	return &fbResponse.FacebookAccessToken, err
}
