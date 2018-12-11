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

package util

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type Path struct{}

func NewPath() *Path {
	return &Path{}
}

func (this *Path) ExecPath() string {
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
func (this *Path) WorkDir() string {
	execPath := this.ExecPath()
	return path.Dir(strings.Replace(execPath, "\\", "/", -1))
}

func (this *Path) WorkName() string {
	execPath := this.ExecPath()
	return path.Base(strings.Replace(execPath, "\\", "/", -1))
}

// 判断文件夹是否存在
func (this *Path) IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (this *Path) Create(dir string, perm os.FileMode) error {
	isExist, err := this.IsExist(dir)
	if err == nil {
		if !isExist {
			// 创建文件夹
			err = os.Mkdir(dir, perm)
		}
	}

	return err
}

func (this *Path) Abs(filePath string) (string, error) {
	return filepath.Abs(filePath)
}

func (this *Path) Dir(filePath string) (string, error) {
	p, err := this.Abs(filePath)
	if err == nil {
		p = path.Dir(p)
	}

	return p, err
}

func (this *Path) FileName(filePath string) (string, error) {
	filePath, err := this.Abs(filePath)
	if err != nil {
		return "", err
	}

	basePath, err := this.Dir(filePath)
	if err == nil {
		filePath = strings.Replace(filePath, "\\", "/", -1)
		basePath = strings.Replace(basePath, "\\", "/", -1)

		basePath = strings.Replace(filePath, basePath+"/", "", 1)
	}

	return basePath, err
}

func (this *Path) Split(filePath string) []string {
	filePath = strings.Replace(filePath, "\\", "/", -1)

	return strings.Split(filePath, "/")
}
