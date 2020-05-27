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
	"bytes"
)

// BufferPool implements a pool of bytes.Buffers in the form of a bounded
// channel.
type BufferPool struct {
	buffers chan *bytes.Buffer
}

// NewBufferPool creates a new BufferPool bounded to the given size.
func NewBufferPool(size int) (bp *BufferPool) {
	return &BufferPool{
		buffers: make(chan *bytes.Buffer, size),
	}
}

// Get gets a Buffer from the BufferPool, or creates a new one if none are
// available in the pool.
func (this *BufferPool) Get() (buffer *bytes.Buffer) {
	select {
	case buffer = <-this.buffers:
	// reuse existing buffer
	default:
		// create new buffer
		buffer = bytes.NewBuffer([]byte{})
	}
	return
}

// Set returns the given Buffer to the BufferPool.
func (this *BufferPool) Set(buffer *bytes.Buffer) {
	buffer.Reset()
	select {
	case this.buffers <- buffer:
	default: // Discard the buffer if the pool is full.
	}
}
