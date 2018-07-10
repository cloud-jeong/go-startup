package utils

import (
	"os"
	"io"
	"math/rand"
	"log"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)

		if err == nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		} else {
			return err
		}
	} else {
		return err
	}

	return nil
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				log.Printf(err.Error())
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				log.Printf(err.Error())
			}
		}

	}
	return
}

func ReadFile(filePath string, buf *[]byte) {

	file, err := os.Open(filePath)
	CheckError(err)

	defer file.Close()

	fi, err := file.Stat()
	CheckError(err)

	*buf = make([]byte, fi.Size())

	_, err = file.Read(*buf)
	CheckError(err)
}

func WriteFile(filePath string, buf *[]byte) error {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Printf("error while open file: %s\n", err.Error())
		return err
	}
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		log.Printf("fail to write file[1]: %s\n", err.Error())
		return err
	}

	_, err = f.WriteString(string(*buf))
	if err != nil {
		log.Printf("fail to write file[2]: %s\n", err.Error())
		return err
	}

	return nil
}

func WriteFileString(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Printf("error while open file: %s\n", err.Error())
		return err
	}
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		log.Printf("fail to write file[1]: %s\n", err.Error())
		return err
	}

	_, err = f.WriteString(content)
	if err != nil {
		log.Printf("fail to write file[2]: %s\n", err.Error())
		return err
	}

	return nil
}

func CheckError(err error) {
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}

func CheckErrorWithMsg(err error, msg string) {
	if err != nil {
		log.Printf(err.Error())
		log.Printf(msg)
		//os.Exit(1)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
