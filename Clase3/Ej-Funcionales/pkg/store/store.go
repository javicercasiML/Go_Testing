package store
import (
    "encoding/json"
    "io/ioutil"
    "os"
)
type Store interface {
    Read(data interface{}) error
    Write(data interface{}) error
}
type Type string
const (
    FileType Type = "file"
)
func New(store Type, fileName string) Store {
    switch store {
    case FileType:
        return &FileStore{fileName, nil}
    }
    return nil
}
type FileStore struct {
    FileName string
    Mock     *Mock
}
type Mock struct {
    Data []byte
    Err  error
}
func (fs *FileStore) AddMock(mock *Mock) {
    fs.Mock = mock
}
func (fs *FileStore) ClearMock() {
    fs.Mock = nil
}
func (fs *FileStore) Read(data interface{}) error {
    if fs.Mock != nil {
        if fs.Mock.Err != nil {
            return fs.Mock.Err
        }
        return json.Unmarshal(fs.Mock.Data, data)
    }
    file, err := os.ReadFile(fs.FileName)
    if err != nil {
        return err
    }
    return json.Unmarshal(file, data)
}
func (fs *FileStore) Write(data interface{}) error {
    jsonData, err := json.MarshalIndent(data, "", " ")
    if err != nil {
        return err
    }
    if fs.Mock != nil {
        if fs.Mock.Err != nil {
            return fs.Mock.Err
        }
        fs.Mock.Data = jsonData
        return nil
    }
    return ioutil.WriteFile(fs.FileName, jsonData, 0644)
}