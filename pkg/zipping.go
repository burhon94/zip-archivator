package pkg

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func SetLogOut() (file *os.File, err error) {
	_, err = os.Stat("log.txt")
	if !os.IsNotExist(err) {
		file, err = os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Printf("can't create logger: %v", err)
			return file, err
		}
		defer func() {
			err := file.Close()
			if err != nil {
				log.Printf("can't close logger: %v", err)
			}
		}()
	}
	file, err = os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("can't open logger: %v", err)
	}
	log.Println("logger setup success")
	return file, nil
}

func DirChecker(dirZipFiles string) (err error) {
	logOut, err := SetLogOut()
	if err != nil {
		log.Printf("can't set log out: %v", err)
		return err
	}
	log.SetOutput(logOut)
	log.Printf("check dir status: %s", dirZipFiles)
	_, err = os.Stat(dirZipFiles)
	if err != nil {
		log.Println("dir is not exist, try create it")
		err = os.Mkdir(dirZipFiles, 0755)
		if err != nil {
			log.Printf("can't create dir: %v", err)
			return
		}
		log.Printf("dir %s, create success", dirZipFiles)
	}
	log.Printf("dir %s, is exist", dirZipFiles)
	return nil
}

func CreateZipFile(scrFile, zipDir string) (err error) {
	logOut, err := SetLogOut()
	if err != nil {
		log.Printf("can't set log out: %v", err)
		return err
	}
	log.SetOutput(logOut)
	zipName := zipDir + scrFile + ".zip"
	log.Println("check zip directory files status")
	err = DirChecker(zipDir)
	if err != nil {
		log.Printf("Error with dir operations: %v", err)
		return err
	}
	log.Printf("try create zip file: %s", zipName)
	newZipFile, err := os.Create(zipName)
	if err != nil {
		log.Printf("can't create zip file: %s, error: %v", zipName, err)
		return err
	}
	defer func() {
		err := newZipFile.Close()
		if err != nil {
			log.Printf("can't close zip file")
			return
		}
	}()
	err = writeToZip(scrFile, newZipFile)
	if err != nil {
		log.Printf("can't write data to zip file: %s", newZipFile.Name())
		return err
	}
	return nil
}

func writeToZip(file string, file2 *os.File) (err error) {
	logOut, err := SetLogOut()
	if err != nil {
		log.Printf("can't set log out: %v", err)
		return err
	}
	log.SetOutput(logOut)
	zipWriter := zip.NewWriter(file2)
	defer func() {
		err := zipWriter.Close()
		if err != nil {
			log.Printf("can't close writer zip: %v", err)
		}
	}()
	log.Printf("try open source file: %s", file)
	zipfile, err := os.Open(file)
	if err != nil {
		log.Printf("can't find source file: %v", err)
		return err
	}
	defer func() {
		err = zipfile.Close()
		if err != nil {
			log.Printf("can't close source file: %v", err)
		}
	}()
	log.Printf("try get status from file: %s", zipfile.Name())
	info, err := zipfile.Stat()
	if err != nil {
		log.Printf("can't get status: %v", err)
		return err
	}
	log.Printf("try create zip header %s and set size: %d by bytes", info.Name(), info.Size())
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		log.Printf("can't create header: %v", err)
		return err
	}
	header.Name = file
	header.Method = zip.Deflate
	log.Printf("try write header zip to zip file: %s", header.Name)
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		log.Printf("can't write header: %v", err)
		return err
	}
	log.Printf("try write data to zip file: %s", header.Name)
	if _, err = io.Copy(writer, zipfile); err != nil {
		log.Printf("can't write data to zip file: %v", err)
		return err
	}
	return nil
}