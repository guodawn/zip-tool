package zip

import (
	"archive/zip"
	_ "bytes"
	_"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	_ "strconv"
	_ "sync"
)
// 压缩文件
func Zip(srcFile string, destFile string) error{
	srcFile = filepath.Join(srcFile, "/")
	//create destFile, destFile will be truncate if it is already exists
	zipFile, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		fmt.Println(srcFile, path)
		if(srcFile == path){
			return nil
		}
		header.Name = path[len(srcFile)+1:]
		if info.IsDir() {
			header.Name += "/"
		}else {
			header.Method = zip.Deflate
		}
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir(){
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
//	wait.Done();
	return  err
}
// 解压
func Unzip(srcfile string, destfile string) error{
	zipReader, err := zip.OpenReader(srcfile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destfile, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		}else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return  nil
}

//func ZipCourses()  {
//	srcFileDir :=  "/Users/guochenguang/Downloads/v2/"
//	destFileDir := "/Users/guochenguang/Downloads/v2/"
//	wait.Add(14)
//	for i:=1;i<=14;i++ {
//		tmpName  := "HRNL0"
//		if(i < 10){
//			tmpName += strconv.Itoa(0) + strconv.Itoa(i)
//		}else {
//			tmpName += strconv.Itoa(i)
//		}
//		go Zip(
//			srcFileDir + tmpName  + "/",
//			destFileDir + tmpName + ".zip")
//	}
//	wait.Wait()
//}