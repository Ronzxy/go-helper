/* Copyright 2018 Ron Zhang <ronzxy@mx.aketi.cn>. All rights reserved.
 *
 * Licensed under the Apache License, version 2.0 (the "License").
 * You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package helper

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

var File = NewFileHelper()

type FileHelper struct{}

func NewFileHelper() *FileHelper {
	return &FileHelper{}
}

// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the number of bytes
// copied and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
//
// If src implements the WriterTo interface,
// the copy is implemented by calling src.WriteTo(dst).
// Otherwise, if dst implements the ReaderFrom interface,
// the copy is implemented by calling dst.ReadFrom(src).
func (this *FileHelper) CopyFile(src, dst string) (int64, error) {
	stat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !stat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	reader, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer reader.Close()

	writer, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer writer.Close()

	return io.Copy(writer, reader)
}

// Use gzip to compress src to dst
// until either EOF is reached or an error occurs.
func (this *FileHelper) GZipFile(src, dst string) error {
	reader, err := os.Open(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer writer.Close()

	gw, err := gzip.NewWriterLevel(writer, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer gw.Close()

	var bytes = make([]byte, 4096)
	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err.Error() != "EOF" {
				return err
			}

			break
		}

		gw.Write(bytes[:n])
		gw.Flush()
	}

	return nil
}

func (this *FileHelper) Save(fileName string, body []byte) error {
	var (
		file *os.File
		n    int
		err  error
	)

	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	n, err = file.Write(body)
	if err != nil {
		return err
	}

	if n == len(body) {
		return file.Sync()
	}

	return nil
}
