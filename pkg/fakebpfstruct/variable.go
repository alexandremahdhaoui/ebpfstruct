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
package fakebpfstruct

import "github.com/alexandremahdhaoui/ebpfstruct"

var _ ebpfstruct.Variable[any] = &Variable[any]{}

func NewVariable[T any]() *Variable[T] {
	return &Variable[T]{
		V:        *new(T),
		doneCh:   make(chan struct{}),
		expector: expector{},
	}
}

type Variable[T any] struct {
	V      T
	doneCh chan struct{}
	expector
}

// Set implements ebpfstruct.Variable.
func (bv *Variable[T]) Set(v T) error {
	if err := bv.checkExpectation("Set"); err != nil {
		return err
	}
	bv.V = v
	return nil
}

func (bv *Variable[T]) Done() <-chan struct{} {
	return bv.doneCh
}

// It will close the channel returned by Done(), notifying when closed
// that the work done on behalf of this Variable[T] has been gracefully
// terminated.
func (bv *Variable[T]) CloseDoneChannel() {
	close(bv.doneCh)
}
