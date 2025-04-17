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

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		a:         make(map[K]V),
		b:         make(map[K]V),
		activePtr: false,
		expector:  expector{},
		doneCh:    make(chan struct{}),
	}
}

type Map[K comparable, V any] struct {
	a, b      map[K]V
	activePtr bool
	doneCh    chan struct{}
	expector
}

// BatchDelete removes keys in batch from the active map.
func (m *Map[K, V]) BatchDelete(keys []K) error {
	if err := m.checkExpectation("BatchDelete"); err != nil {
		return err
	}
	activeMap := m.GetActiveMap()
	for _, k := range keys {
		delete(activeMap, k)
	}
	return nil
}

// BatchUpdate implements Map.
func (m *Map[K, V]) BatchUpdate(kv map[K]V) error {
	if err := m.checkExpectation("BatchUpdate"); err != nil {
		return err
	}
	activeMap := m.GetActiveMap()
	for k, v := range kv {
		activeMap[k] = v
	}
	return nil
}

// Set implements Map.
func (m *Map[K, V]) Set(newMap map[K]V) error {
	m.setPassiveMap(newMap)
	if err := m.checkExpectation("Set"); err != nil {
		return err
	}
	m.switchover()
	return nil
}

// SetAndDeferSwitchover implements Map.
func (m *Map[K, V]) SetAndDeferSwitchover(newMap map[K]V) (func(), error) {
	m.setPassiveMap(newMap)
	return m.switchover, m.checkExpectation("SetAndDeferSwitchover")
}

// -- GET ACTIVE

// It returns the actual state of the map in active state.
func (m *Map[K, V]) GetActiveMap() map[K]V {
	if m.activePtr {
		return m.a
	}
	return m.b
}

// -- DONE

func (m *Map[K, V]) Done() <-chan struct{} {
	return m.doneCh
}

// It will close the channel returned by Done(), notifying when closed
// that the work done on behalf of this Map[K,V] has been gracefully
// terminated.
func (m *Map[K, V]) CloseDoneChannel() {
	close(m.doneCh)
}

// -- HELPERS

func (m *Map[K, V]) setPassiveMap(newMap map[K]V) {
	if m.activePtr {
		m.b = newMap
	} else {
		m.a = newMap
	}
}

func (m *Map[K, V]) switchover() {
	m.activePtr = !m.activePtr
}
