package formdata

import (
	"bytes"
	"mime/multipart"

	log "github.com/sirupsen/logrus"
	"go.k6.io/k6/js/modules"
)

type File struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

type FileFormData struct {
	Body     string `json:"body"`
	Boundary string `json:"boundary"`
}

func init() {
	modules.Register("k6/x/formdata", new(FormData))
}

type FormData struct{}

func (f *FormData) GetFileFormData(file File) FileFormData {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	w.FormDataContentType()
	fw, err := w.CreateFormFile("filename", file.Name)
	if err != nil {
		log.Fatal(err)
	}
	fw.Write(file.Data)
	w.Close()
	return FileFormData{
		Body:     buf.String(),
		Boundary: w.Boundary(),
	}
}
