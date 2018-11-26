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
