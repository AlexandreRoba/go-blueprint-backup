package backup

import (
	"os"
	"path/filepath"
	"archive/zip"
	"io"
)

type Archiver interface {
	Archive(src, dest string) error
}

type zipper struct{}

func (z *zipper) Archive(src, dest string) error {
	if err:= os.MkdirAll(filepath.Dir(dest),0777);err!=nil{
		return err
	}
	out,err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	return filepath.Walk(src,func(path string, info os.FileInfo,err error) error {
		if info.IsDir() {
			return nil
		}
		if err!=nil{
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		f,err := w.Create(path)
		if err != nil {
			return err
		}
		_,err=io.Copy(f,in)
		if err != nil {
			return err
		}
		return nil
	})
}

//Zip is an archiver that zips and unzips files.
var Zip Archiver = (*zipper)(nil)
