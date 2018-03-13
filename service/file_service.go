package service

import (
	"github.com/ndphu/fm/dao"
	"github.com/ndphu/fm/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FileService struct {
	db   *dao.DAO
	root *model.File
}

func NewFileService(dao *dao.DAO) *FileService {
	return &FileService{
		db: dao,
	}
}

func (s *FileService) Init() error {
	root := &model.File{
		Name:       "/",
		ServerPath: "",
		Type:       model.TYPE_FOLDER,
	}
	existingRoot, err := s.FindRootFolder()
	if err == mgo.ErrNotFound {
		root.Id = bson.NewObjectId()
		err = s.db.FileCollection().Insert(root)
	} else if err == nil {
		root.Id = existingRoot.Id
		err = s.db.FileCollection().UpdateId(root.Id, root)
	}
	if err != nil {
		panic(err)
	}
	s.root = root
	return nil
}

func (s *FileService) FileFileById(id string) (*model.File, error) {
	result := model.File{}
	err := s.db.FileCollection().FindId(bson.ObjectIdHex(id)).One(&result)
	return &result, err
}

func (s *FileService) CreateFile(f *model.File) error {
	f.Id = bson.NewObjectId()

	if f.ParentId.Hex() == "" {
		f.ParentId = s.root.Id
	}

	return s.db.FileCollection().Insert(f)
}

func (s *FileService) FindRootFolder() (*model.File, error) {
	var root model.File
	err := s.db.FileCollection().Find(bson.M{"parentId": ""}).One(&root)
	return &root, err
}

func (s *FileService) FindChildren(id string) ([]*model.File, error) {
	children := []*model.File{}
	err := s.db.FileCollection().Find(bson.M{"parentId": bson.ObjectIdHex(id)}).All(&children)
	return children, err
}
