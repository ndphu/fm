package service

import (
	"github.com/ndphu/fm/dao"
	"github.com/ndphu/fm/model"
	"gopkg.in/mgo.v2/bson"
)

type FileService struct {
	d    *dao.DAO
	root *model.File
}

func NewFileService(_dao *dao.DAO) *FileService {
	return &FileService{
		d: _dao,
	}
}

func (s *FileService) Init() error {
	root, err := s.d.FindRootFolder()

	if err != nil {
		root, err = s.d.SaveOrUpdateFile(&model.File{
			Name:       "root",
			IsRoot:     true,
			ServerPath: "",
		})
		if err != nil {
			return err
		}
	}

	s.root = root

	return nil
}

func (s *FileService) FileFileById(id string) (*model.File, error) {
	result := model.File{}
	err := s.d.FileCollection().FindId(bson.ObjectIdHex(id)).One(&result)
	return &result, err
}

func (s *FileService) CreateFile(f *model.File) error {
	f.Id = bson.NewObjectId()

	if f.Parent.Id == nil {
		f.Parent.Collection = s.d.FileCollection().Name
		f.Parent.Id = s.root.Id
	}

	return s.d.FileCollection().Insert(f)
}

func (s *FileService) FindRootFolder() (*model.File, error) {
	var root model.File
	err := s.d.FileCollection().Find(bson.M{"isRoot": true}).One(&root)
	return &root, err
}
