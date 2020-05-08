package yi

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"

	_ "github.com/festum/yi/statik"
	"github.com/rakyll/statik/fs"
)

const filePath = "gua.data"

func libCompress(b []byte) error {
	buff := bytes.NewBuffer(b)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	w := zlib.NewWriter(file)
	if _, err := io.Copy(w, buff); err != nil {
		return err
	}
	w.Flush()
	defer w.Close()
	return nil
}

func libDecompress() ([]byte, error) {
	buff := bytes.Buffer{}
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r, err := zlib.NewReader(file)
	if err == nil {
		if _, err := io.Copy(&buff, r); err == nil {
			return buff.Bytes(), nil
		}
	}
	defer r.Close()
	return nil, err
}

func libDecompressStatik() ([]byte, error) {
	buff := bytes.Buffer{}
	sfs, err := fs.New()
	if err != nil {
		return nil, err
	}

	file, err := sfs.Open("/" + filePath)
	if err != nil {
		return nil, err
	}
	r, err := zlib.NewReader(file)
	if err == nil {
		if _, err := io.Copy(&buff, r); err == nil {
			return buff.Bytes(), nil
		}
	}
	defer r.Close()
	return nil, err
}
