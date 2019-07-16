package api

import (
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/model"
)

// GenImgID returns a randomly generated unique ID.
func genImgID(f string) (string, string) {
	id := uuid.Must(uuid.NewV4()).String() + path.Ext(f)
	p := "app/mac/" + id

	return id, p
}

// OpenFile opens and returns the FileHeader's associated File.
func openFile(c echo.Context) (*multipart.FileHeader, multipart.File, error) {
	f, err := c.FormFile("img")
	if err != nil {
		return nil, nil, err
	}
	s, err := f.Open()
	if err != nil {
		return nil, nil, err
	}
	defer s.Close()

	return f, s, err
}

// CreateFile creates the named file.
func createFile(p string, s multipart.File) {
	d, err := os.Create(p)
	if err != nil {
		return
	}
	defer d.Close()
	// Copy copies from src to dst
	if _, err := io.Copy(d, s); err != nil {
		return
	}
}

// RemoveFile removes the named file or (empty) directory.
func removeFile(m *model.Device) {
	if id := m.Img; id != "" {
		p := "app/mac/" + id

		// Remove image in Storage
		if err := os.Remove(p); err != nil {
			return
		}
	}
}
