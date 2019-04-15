/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [transport for rotateFile]
 */

package logger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

// RotateFile rotate file
type RotateFile struct {
	file     *os.File
	fileName string
	Dirname  string
	MaxSize  int64
}

func (r *RotateFile) Write(p []byte) (int, error) {
	// Write to file
	n, err := r.file.Write(p)
	// Check to see if we need to end the stream and create a new one.
	need, err := r.needsNewFile()
	if need {
		r.newFile()
	}
	return n, err
}

func (r *RotateFile) needsNewFile() (bool, error) {
	info, error := r.file.Stat()
	if error != nil {
		return false, error
	}
	size := info.Size()
	if size >= r.MaxSize {
		return true, nil
	}
	return false, nil
}

func (r *RotateFile) newFile() error {
	fileName := time.Now().Format("2019-01.02-03:04")
	fileName = fmt.Sprintf("%s.log", fileName)
	filePath := path.Join(r.Dirname, fileName)
	file, error := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if error != nil {
		return error
	}
	error = r.file.Close()
	if error != nil {
		return error
	}
	r.file = file
	r.fileName = fileName
	return nil
}

func (r *RotateFile) initFilePath() (string, error) {
	var filePath string
	if _, err := os.Stat(r.Dirname); os.IsNotExist(err) {
		os.Mkdir(r.Dirname, os.ModePerm)
	}
	filepath.Walk(r.Dirname, func(path string, info os.FileInfo, err error) error {
		if filePath == "" && info.IsDir() != true {
			fileSize := info.Size()
			sizeMatch := fileSize < r.MaxSize
			if sizeMatch {
				filePath = path
				r.fileName = info.Name()
			}
		}
		return nil
	})
	if filePath != "" {
		return filePath, nil
	}
	// Create level log file
	fileName := time.Now().Format("2006.01.02")
	fileName = fmt.Sprintf("%s.log", fileName)
	filePath = path.Join(r.Dirname, fileName)
	r.fileName = fileName
	return filePath, nil
}

// OpenFile open file
func OpenFile(dirName string, maxSize int64, flag int, perm os.FileMode) (*RotateFile, error) {
	tf := &RotateFile{
		Dirname: dirName,
		MaxSize: maxSize,
	}
	filePath, error := tf.initFilePath()
	if error != nil {
		return nil, error
	}
	f, error := os.OpenFile(filePath, flag, perm)
	tf.file = f
	return tf, error
}
