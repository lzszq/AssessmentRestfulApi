package controllers

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lzszq/AssessmentRestfulApi/api/models"
	"github.com/lzszq/AssessmentRestfulApi/api/responses"
	"github.com/lzszq/AssessmentRestfulApi/api/utils"
)

type Image models.Image

func (server *Server) UploadHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	} else {
		req.ParseForm()
		uploadFile, handle, err := req.FormFile("image")
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}

		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".png" {
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("只支持jpg/png图片上传"))
			return
		}

		os.Mkdir("./uploaded/", 0777)
		filename := utils.RandStringRunes(20) + ext
		saveFile, err := os.OpenFile("./uploaded/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}
		io.Copy(saveFile, uploadFile)

		defer uploadFile.Close()
		defer saveFile.Close()

		image := Image{
			Url: "/uploaded/" + filename,
		}
		responses.JSON(w, http.StatusOK, image)
	}
}

func (server *Server) ShowPicHandle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	file, err := os.Open("./uploaded/" + vars["filename"])
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	responses.IMAGE(w, http.StatusOK, buff)
}

const html = `<html>
    <head></head>
    <body>
        <form method="post" enctype="multipart/form-data">
            <input type="file" name="image" />
            <input type="submit" />
        </form>
    </body>
</html>`
