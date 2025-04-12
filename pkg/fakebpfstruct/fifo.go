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

var _ ebpfstruct.FIFO[any] = &FIFO[any]{}

func NewFIFO[T any]() *FIFO[T] {
	return &FIFO[T]{
		Chan:     make(chan T),
		expector: expector{},
	}
}

type FIFO[T any] struct {
	Chan chan T
	expector
}

// Subscribe implements FIFO.
func (f *FIFO[T]) Subscribe() (<-chan T, error) {
	return f.Chan, f.checkExpectation("Subscribe")
}
