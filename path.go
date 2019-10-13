/* Copyright 2015 sky<skygangsta@hotmail.com>. All rights reserved.
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
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var Path = NewPathHelper()

type PathHelper struct{}

func NewPathHelper() *PathHelper {
	return &PathHelper{}
}

func (this *PathHelper) ExecPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return p
}

// WorkDir returns absolute path of work directory.
func (this *PathHelper) WorkDir() string {
	execPath := this.ExecPath()
	return path.Dir(strings.Replace(execPath, "\\", "/", -1))
}

func (this *PathHelper) WorkName() string {
	execPath := this.ExecPath()
	return path.Base(strings.Replace(execPath, "\\", "/", -1))
}

// 判断文件夹是否存在
func (this *PathHelper) IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func (this *PathHelper) Create(dir string, perm os.FileMode) error {
	isExist, err := this.IsExist(dir)
	if err != nil {
		return err
	}

	if !isExist {
		// 创建文件夹
		return os.Mkdir(dir, perm)
	}

	return nil
}

func (this *PathHelper) Abs(filePath string) (string, error) {
	return filepath.Abs(filePath)
}

func (this *PathHelper) Dir(filePath string) (string, error) {
	p, err := this.Abs(filePath)
	if err == nil {
		p = path.Dir(p)
	}

	return p, err
}

func (this *PathHelper) FileName(filePath string) (string, error) {
	filePath, err := this.Abs(filePath)
	if err != nil {
		return "", err
	}

	basePath, err := this.Dir(filePath)
	if err == nil {
		filePath = strings.Replace(filePath, "\\", "/", -1)
		basePath = strings.Replace(basePath, "\\", "/", -1)

		basePath = filePath[len(basePath)+1:]
	}

	return basePath, err
}

func (this *PathHelper) Split(filePath string) []string {
	filePath = strings.Replace(filePath, "\\", "/", -1)

	return strings.Split(filePath, "/")
}

func (this *PathHelper) CreateDir(filePath string, perm os.FileMode) error {
	var (
		dirs []string
		err  error
	)
	dirs = this.Split(filePath)

	filePath = ""
	for _, v := range dirs {
		if v == "" {
			filePath = "/"
		} else {
			filePath = path.Join(filePath, v)

			err = this.Create(filePath, perm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
