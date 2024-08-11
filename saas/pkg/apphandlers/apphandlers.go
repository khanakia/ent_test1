package apphandlers

import (
	"configmgr/gen/appconfig"
	"io"
	"lace/ginserver"
	"lace/medialibrary"
	"lace/nlog"
	"lace/util"
	"log"
	"net/http"
	"saas/gen/ent"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Nlog      nlog.Logger
	Server    ginserver.Server
	EntClient *ent.Client
	AppConfig *appconfig.AppConfig
}

type AppHandlers struct {
	Config
}

func New(config Config) {
	router := config.Server.Router

	handler := &AppHandlers{Config: config}

	router.POST("/upload", handler.UploadHandler)
}

func (pkg AppHandlers) UploadHandler(c *gin.Context) {
	// _, err := auth.GetUserFromContext(c)
	// if err != nil {
	// 	error := util.ResponseError{
	// 		Message: "access denied",
	// 		// Errors:  errorList,
	// 	}
	// 	error.Send(c, 401)
	// 	return
	// }

	multipart, err := c.Request.MultipartReader()
	if err != nil {
		log.Println("Failed to create MultipartReader", err)
		error := util.ResponseError{
			Message: err.Error(),
		}
		error.Send(c, http.StatusUnauthorized)
		return
	}

	// var files []media_type.File
	var medias []*medialibrary.Media

	for {
		mimePart, err := multipart.NextPart()
		if err == io.EOF {
			log.Printf("multipart.NextPart: %v", err)
			break
		}
		if err != nil {
			log.Printf("Error reading multipart section: %v", err)
			break
		}

		// fileAdder := FileAdder{}
		fileAdder := medialibrary.NewFileAdder("/Volumes/D/www/projects/khanakia/saasfly/saasfly_api/uploads")
		media, err := fileAdder.CreateFromMultiPart(mimePart)
		medias = append(medias, media)
		if err != nil {
			log.Println("Failed to create MultipartReader", err)
			error := util.ResponseError{
				Message: err.Error(),
			}
			error.Send(c, http.StatusUnauthorized)
			return
		}

	}

	// fmt.Println(files)
	c.JSON(http.StatusOK, gin.H{
		"data": medias,
	})
}
