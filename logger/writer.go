// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

type (
	// LevelWriter defined level writer
	LevelWriter struct {
		W     io.Writer
		Level LEVEL
	}
	// MutiWriter write with level
	MutiWriter struct {
		writers []io.Writer
		Level   LEVEL
	}
	// RotateWriter rotate file
	RotateWriter struct {
		file     *os.File
		fileName string
		Dirname  string
		MaxSize  int64
	}
)

func (c *LevelWriter) Write(p []byte) (int, error) {
	level := toLevelString(c.Level)
	colorLevel := toColorLevel(c.Level, level)
	pbyte := []byte(fmt.Sprintf("%s %s", colorLevel, string(p)))
	return c.W.Write(pbyte)
}

func (t *MutiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		if r, ok := w.(*LevelWriter); ok {
			r.Level = t.Level
			n, err = w.Write(p)
			if err != nil {
				return
			}
		} else {
			pbyte := []byte(fmt.Sprintf("%s", string(p)))
			n, err = w.Write(pbyte)
			if err != nil {
				return
			}
		}
	}
	return len(p), nil
}

func multiLevelWriter(writers ...io.Writer) *MutiWriter {
	allWriters := make([]io.Writer, 0, len(writers))
	for _, w := range writers {
		if mw, ok := w.(*MutiWriter); ok {
			allWriters = append(allWriters, mw.writers...)
		} else {
			allWriters = append(allWriters, w)
		}
	}
	return &MutiWriter{writers: allWriters}
}

func (r *RotateWriter) Write(p []byte) (n int, err error) {
	if n, err = r.file.Write(p); err == nil {
		is := false
		if is, err = r.isCapped(); is {
			r.createFile()
		}
	}
	return n, err
}

func (r *RotateWriter) isCapped() (bool, error) {
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

func (r *RotateWriter) createFile() error {
	fileName := time.Now().Format("2006_01.02_03:04")
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

func (r *RotateWriter) initFilePath() (string, error) {
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
	fileName := time.Now().Format("2006_01.02_03:04")
	fileName = fmt.Sprintf("%s.log", fileName)
	filePath = path.Join(r.Dirname, fileName)
	r.fileName = fileName
	return filePath, nil
}

// NewRotateWriter defined return a rotate writer
func NewRotateWriter(dirName string, maxSize int64, flag int, perm os.FileMode) io.Writer {
	tf := &RotateWriter{
		Dirname: dirName,
		MaxSize: maxSize,
	}
	filePath, err := tf.initFilePath()
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filePath, flag, perm)
	if err != nil {
		panic(err)
	}
	tf.file = f
	return tf
}

// NewLevelWriter defined return a level writer
func NewLevelWriter(l LEVEL) io.Writer {
	return &LevelWriter{
		W:     os.Stdout,
		Level: l,
	}
}
