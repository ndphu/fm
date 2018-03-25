package service

import (
	"encoding/json"
	"fmt"
	"github.com/ndphu/fm/dao"
	"github.com/ndphu/fm/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	APP_TOKEN      = os.Getenv("FACEBOOK_APP_TOKEN")
	APP_SECRET     = os.Getenv("FACEBOOK_APP_SECRET")
	APP_ID         = os.Getenv("FACEBOOK_APP_ID")
	LOGIN_CALLBACK = "http://localhost:8080/pfm/login/callback"
)

type LoginService struct {
	db          *dao.DAO
	userService *UserService
}

func NewLoginService(db *dao.DAO) *LoginService {
	return &LoginService{
		db: db,
	}
}

func (s *LoginService) SetUserService(us *UserService) {
	s.userService = us
}

func (s *LoginService) ProcessAccessCode(code string) {
	accessToken, err := s.GetAccessTokenFromAccessCode(code)
	if err != nil {
		panic(err)
	}
	fat, err := s.ValidateFacebookAcessToken(accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", fat.UserId)
	user, err := s.userService.FindUserByExternalId(fat.UserId)
	if err == mgo.ErrNotFound {
		user = &model.User{
			Id:         bson.NewObjectId(),
			ExternalId: fat.UserId,
		}
		s.userService.CreateUser(user)
	}

	getUserDetails := fmt.Sprintf("https://graph.facebook.com/v2.12/%s?fields=id,gender,email,name,first_name,last_name,address&access_token=%s", fat.UserId, accessToken)

	response, err := http.Get(getUserDetails)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fu := model.FacebookUser{}

	if err := json.Unmarshal(content, &fu); err != nil {
		panic(err)
	}
	user.FirstName = fu.FirstName
	user.LastName = fu.LastName

	s.userService.UpdateUser(user)
}

func (s *LoginService) GetAccessTokenFromAccessCode(code string) (string, error) {
	if APP_ID == "" || APP_SECRET == "" || APP_TOKEN == "" {
		panic("Invalid settings for facebook app")
	}
	getTokenUrl := fmt.Sprintf("https://graph.facebook.com/v2.12/oauth/access_token?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s",
		APP_ID, APP_SECRET, LOGIN_CALLBACK, code)

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

func (s *LoginService) ProcessAccessToken(accessToken string) *model.User {
	fat, err := s.ValidateFacebookAcessToken(accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", fat.UserId)
	user, err := s.userService.FindUserByExternalId(fat.UserId)
	if err == mgo.ErrNotFound {
		user = &model.User{
			Id:         bson.NewObjectId(),
			ExternalId: fat.UserId,
		}
		s.userService.CreateUser(user)
	}

	getUserDetails := fmt.Sprintf("https://graph.facebook.com/v2.12/%s?fields=id,gender,email,name,first_name,last_name,address&access_token=%s", fat.UserId, accessToken)

	response, err := http.Get(getUserDetails)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fu := model.FacebookUser{}
	fmt.Println(string(content))

	if err := json.Unmarshal(content, &fu); err != nil {
		panic(err)
	}
	user.FirstName = fu.FirstName
	user.LastName = fu.LastName
	s.userService.UpdateUser(user)

	return user
}
