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

import (
	"github.com/alexandremahdhaoui/ebpfstruct"
)

var _ ebpfstruct.Map[uint32, any] = &Map[uint32, any]{}

// -------------------------------------------------------------------
// -- FAKE MAP
// -------------------------------------------------------------------

type Map[K comparable, V any] struct {
	Map map[K]V
	expector
}

// BatchDelete implements Map.
func (f *Map[K, V]) BatchDelete(keys []K) error {
	// ----------------------
	// MOVE BELOW TO EXPECTOR
	//
	if len(f.expectationList) == 0 {
		panic("did not expect call to Map.BatchDelete")
	}
	if f.expectationList[0].Method != "BatchDelete" {
		panic("did not expect call to Map.BatchDelete")
	}
	if len(f.expectationList) > 1 {
		f.expectationList = f.expectationList[1:]
	}
	if err := f.expectationList[0].Err; err != nil {
		return err
	}
	//
	// MOVE ABOVE TO EXPECTOR
	// ----------------------

	for _, k := range keys {
		delete(f.Map, k)
	}
	return nil
}

// BatchUpdate implements Map.
func (f *Map[K, V]) BatchUpdate(kv map[K]V) error {
	panic("unimplemented")
}

// Set implements Map.
func (f *Map[K, V]) Set(newMap map[K]V) error {
	panic("unimplemented")
}

// SetAndDeferSwitchover implements Map.
func (f *Map[K, V]) SetAndDeferSwitchover(newMap map[K]V) (func(), error) {
	panic("unimplemented")
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		Map: make(map[K]V),
	}
}
