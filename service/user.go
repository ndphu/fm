package service

import (
	"github.com/ndphu/fm/dao"
	"github.com/ndphu/fm/model"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	DB *dao.DAO
}

func NewUserService(d *dao.DAO) *UserService {
	return &UserService{
		DB: d,
	}
}

func (s *UserService) FindUserByExternalId(extId string) (*model.User, error) {
	user := model.User{}
	err := s.DB.UserCollection().Find(bson.M{"externalId": extId}).One(&user)
	return &user, err
}

func (s *UserService) CreateUser(newUser *model.User) (*model.User, error) {
	err := s.DB.UserCollection().Insert(newUser)
	return newUser, err
}
func (s *UserService) UpdateUser(update *model.User) (*model.User, error) {
	err := s.DB.UserCollection().UpdateId(update.Id, update)
	return update, err
}

func (s *UserService) FindAll() ([]model.User, error) {
	var users []model.User
	err := s.DB.UserCollection().Find(nil).All(&users)
	return users, err
}
