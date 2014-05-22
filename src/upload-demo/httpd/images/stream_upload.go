package images

import (
	"io"
	"mime/multipart"
	"net/http"
	"upload-demo/config"
	"upload-demo/log"
	"upload-demo/models"

	"github.com/levicook/head"
)

func streamUpload(mr *multipart.Reader) models.File {

	// create our file
	file, err := config.FileRepo.Create()
	log.PanicIf(err)

	headWriter := head.NewWriter(512)

	limitedReader := new(io.LimitedReader)
	limitedReader.N = int64(MaxSize)

	for {
		part, err := mr.NextPart()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if file.Name() == "" && part.FileName() != "" {
			file.SetName(part.FileName())
		}

		limitedReader.R = part

		teeReader := io.TeeReader(limitedReader, headWriter)

		io.Copy(file, teeReader)

		part.Close()
	}

	// detect and set file's content type
	ct := http.DetectContentType(headWriter.Head())
	file.SetContentType(ct)
	file.Close()

	return file
}
