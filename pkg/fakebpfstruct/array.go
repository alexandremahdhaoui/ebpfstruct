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

var _ ebpfstruct.Array[any] = &Array[any]{}

func NewArray[T any]() *Array[T] {
	return &Array[T]{
		a:         make([]T, 0),
		b:         make([]T, 0),
		activePtr: false,
		expector:  expector{},
	}
}

type Array[T any] struct {
	a, b      []T
	activePtr bool
	expector
}

// Set implements Array.
func (a *Array[T]) Set(values []T) error {
	a.setPassive(values)
	return a.checkExpectation("Set")
}

// SetAndDeferSwitchover implements Array.
func (a *Array[T]) SetAndDeferSwitchover(values []T) (func(), error) {
	a.setPassive(values)
	return a.switchover, a.checkExpectation("SetAndDeferSwitchover")
}

// -- GET ACTIVE

// It returns the actual state of the array in active state.
func (a *Array[T]) GetActiveArray() []T {
	if a.activePtr {
		return a.b
	}
	return a.a
}

// -- HELPERS

func (a *Array[T]) setPassive(values []T) {
	if a.activePtr {
		a.a = values
	} else {
		a.b = values
	}
}

func (a *Array[T]) switchover() {
	a.activePtr = !a.activePtr
}
