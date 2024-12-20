package medialibrary

import (
	"bytes"
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"lace/filesystem"
	"lace/publicid"
	"lace/storage"
	"log"
	"mime"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/ubgo/goutil"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

const (
	VARIANT_NAME_ORIGINAL = "original"
)

type Media struct {
	ID            string    `bun:",pk"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Disk          string    `json:"disk"`
	Directory     string    `json:"directory"`
	Name          string    `json:"name"`
	OriginalName  string    `json:"originalName"`
	Extension     string    `json:"extension"`
	MimeType      string    `json:"mimeType"`
	AggregateType string    `json:"aggregateType"`
	Uid           string    `json:"uid"`
	Checksum      string    `json:"checksum"`
	IsVariant     bool      `json:"isVariant"`
	WorkspaceID   string    `json:"workspaceId"`
	Size          uint      `json:"size"`
	URL           string    `json:"url" bun:"-"`
	AppID         string    `json:"appId"`
}

func (m *Media) FilePath() string {
	if len(m.Directory) == 0 {
		return m.Name
	}
	return m.Directory + m.Name
}

// ytd add unique index media_id, mediable_type, mediable_id
type Mediable struct {
	ID           string    `bun:",pk"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	MediaID      string    `json:"mediaId"`
	MediableType string    `json:"medailedType"`
	MediableID   string    `json:"mediableId"`
	Tag          string    `json:"tag"`
	Order        int       `json:"order"`
	AppID        string    `json:"appId"`
	Media        Media     `bun:"rel:belongs-to,join:media_id=id"`
}

// call this once in app.go inside func New() Plugin { ... }
func Construct(sqldb *sql.DB) {
	InitDb(sqldb)
}

func InitDb(sqldb *sql.DB) {
	db = bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

}

// Responsible only for single Media
type FileAdder struct {
	rootPath    string
	filename    string
	contentType string
	filesystem  *filesystem.FileSystem
}

func NewFileAdder(diskName string) *FileAdder {
	if db == nil {
		panic("medialibrary not constructed yet")
	}

	// adpater := filesystem.NewLocalFileSystemAdapter(rootPath)
	// fs := filesystem.NewFileSystem(adpater, filesystem.WithPublicURL(publicURLs))
	strg := storage.GetStorage()
	disk, ok := strg.GetDisk(diskName)

	if !ok {
		panic("storage not found: " + diskName)
	}

	return &FileAdder{
		filesystem: disk,
		rootPath:   storage.UploadsDir,
	}
}

// GetChecksum computes and returns the MD5 checksum of the provided byte slice.
func (cls *FileAdder) GetChecksum(b []byte) string {
	checkSum := md5.Sum(b)
	return hex.EncodeToString(checkSum[:])
}

func (cls *FileAdder) CreateFromMultiPart(filePart *multipart.Part) (*Media, error) {
	fmt.Println(filePart.Header)
	_, params, err := mime.ParseMediaType(filePart.Header.Get("Content-Disposition"))
	if err != nil {
		log.Printf("Invalid Content-Disposition: %v", err)
		return nil, err
	}

	cls.filename = params["filename"]
	cls.contentType = filePart.Header.Get("Content-Type")

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(filePart)
	if err != nil {
		return nil, err
	}

	b := buf.Bytes()

	err = cls.filesystem.Write(cls.filename, b, nil)

	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(cls.filename))

	media := Media{
		ID:           publicid.Must12(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Checksum:     cls.GetChecksum(b),
		Name:         cls.filename, // actual filename of the file saved
		OriginalName: cls.filename, // is the name of file uploaded by user but we can save it filename-xxx.jpeg (name)
		MimeType:     cls.contentType,
		Disk:         "local",
		// relative path (test) = rootPath - path (./uploads - ./uploads/test/image.jpg )
		// for not it can be empty as we do not not allow users to upload in subdirectory
		Directory:   "",
		IsVariant:   false,
		Uid:         publicid.Must14(),
		Extension:   ext[1:],
		Size:        uint(len(b)), // in bytes
		WorkspaceID: "default",
		AppID:       "a1", // YTD - make dynamic
	}

	// goutil.PrintToJSON(media)

	_, err = db.NewInsert().Model(&media).Exec(context.Background())

	if err != nil {
		cls.filesystem.Delete(cls.filename)
		return nil, err
	}

	media.URL, _ = cls.filesystem.PublicURL(media.FilePath())

	// fmt.Println("res", res)
	return &media, nil
}

func SyncMedia(appID, tag, mediableType, mediableID string, mediaIds []string) error {
	ctx := context.Background()
	_, err := db.NewDelete().Model(&Mediable{}).
		Where("app_id = ?", appID).
		Where("tag = ?", tag).
		Where("mediable_type = ?", mediableType).
		Where("mediable_id = ?", mediableID).
		Exec(ctx)

	if err != nil {
		return err
	}

	// YTD - validate mediaIds exists in appId

	mediables := []Mediable{}

	for i, mediaID := range mediaIds {
		mediables = append(mediables, Mediable{
			ID:           publicid.Must(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			MediableType: mediableType,
			MediableID:   mediableID,
			MediaID:      mediaID,
			Tag:          tag,
			Order:        i,
			AppID:        appID,
		})
	}

	res, err := db.NewInsert().Model(&mediables).Exec(ctx)
	if err != nil {
		panic(err)
	}

	goutil.PrintToJSON(res)

	return nil
}

func FetchMedias(appID, tag, mediableType, mediableID string) ([]*Media, error) {

	mediables := []Mediable{}

	ctx := context.Background()
	err := db.NewSelect().Model(&Mediable{}).
		Relation("Media").
		Where("mediable.app_id = ?", appID).
		Where("mediable.tag = ?", tag).
		Where("mediable.mediable_type = ?", mediableType).
		Where("mediable.mediable_id = ?", mediableID).
		Limit(100).
		Order("mediable.order ASC").
		Scan(ctx, &mediables)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bunMedias := []*Media{}

	for _, mediable := range mediables {
		media := &mediable.Media
		media.URL, err = storage.GetStorage().GetDefault().PublicURL(media.Name)

		fmt.Println(err)
		// if err != nil {
		// 	return nil, err
		// }
		bunMedias = append(bunMedias, media)
	}

	goutil.PrintToJSON(bunMedias)

	return bunMedias, nil
}
