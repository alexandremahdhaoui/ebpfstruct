/*
 * Copyright 2025 Alexandre Mahdhaoui
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package ebpfstruct

import (
	"errors"

	"github.com/alexandremahdhaoui/tooling/pkg/flaterrors"
	"github.com/cilium/ebpf"
)

var ErrCreatingNewVariable = errors.New("creating new variable")

// This is only used in order to write tests. No fancy feature.
// Maybe in the future add support for:
//   - Get()
//   - caching & no-cache for Get().
//   - Set with differable switchover in order to sync switchover
//     with many bpf data structures.
type Variable[T any] interface {
	// Done returns a channel that's closed when work done on behalf of this
	// interface has been gracefully terminated.
	Done() <-chan struct{}

	// Set the variable.
	Set(v T) error
}

// doneCh is a channel used to notify the bpf data structures or bpf
// program has been closed and they can no longer be used.
func NewVariable[T any](obj *ebpf.Variable, doneCh <-chan struct{}) (Variable[T], error) {
	if obj == nil {
		return nil, flaterrors.Join(ErrEBPFObjectsMustNotBeNil, ErrCreatingNewVariable)
	}

	return &bpfVariable[T]{obj: obj}, nil
}

type bpfVariable[T any] struct {
	obj *ebpf.Variable
	// doneCh is a channel used to notify the bpf data structures or bpf
	// program has been closed and they can no longer be used.
	doneCh <-chan struct{}
}

func (bv *bpfVariable[T]) Done() <-chan struct{} {
	return bv.doneCh
}

// Set implements BPFVariable.
func (bv *bpfVariable[T]) Set(v T) error {
	return bv.obj.Set(v)
}
