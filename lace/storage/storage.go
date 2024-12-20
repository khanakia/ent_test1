package storage

import (
	"lace/filesystem"
)

const (
	UploadsDir = "/Volumes/D/www/projects/khanakia/saasfly/saasfly_api/uploads"
)

var PublicURLs []string = []string{"http://localhost:2303/assets"}

var disks map[string]*filesystem.FileSystem = make(map[string]*filesystem.FileSystem)

var _storage *Storage

func init() {
	_storage = NewStorage()
}

func GetStorage() *Storage {
	return _storage
}

type Storage struct {
	disks map[string]*filesystem.FileSystem
}

func NewStorage() *Storage {

	adpater := filesystem.NewLocalFileSystemAdapter(UploadsDir)
	fs := filesystem.NewFileSystem(adpater, filesystem.WithPublicURL(PublicURLs))
	disks["local"] = fs

	s := &Storage{disks: disks}
	return s
}

func (s *Storage) GetDefault() *filesystem.FileSystem {
	return disks["local"]
}

func (s *Storage) GetDisk(diskName string) (*filesystem.FileSystem, bool) {
	disk, ok := s.disks[diskName]
	return disk, ok
}

func (s *Storage) AddDisk(diskName string, fs *filesystem.FileSystem) {
	s.disks[diskName] = fs
}

func (s *Storage) RemoveDisk(diskName string) {
	delete(s.disks, diskName)
}

func (s *Storage) GetDisks() map[string]*filesystem.FileSystem {
	return s.disks
}
