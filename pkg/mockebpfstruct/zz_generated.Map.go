// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify
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

package mockebpfstruct

import (
	mock "github.com/stretchr/testify/mock"
)

// NewMockMap creates a new instance of MockMap. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMap[K comparable, V any](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMap[K, V] {
	mock := &MockMap[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockMap is an autogenerated mock type for the Map type
type MockMap[K comparable, V any] struct {
	mock.Mock
}

type MockMap_Expecter[K comparable, V any] struct {
	mock *mock.Mock
}

func (_m *MockMap[K, V]) EXPECT() *MockMap_Expecter[K, V] {
	return &MockMap_Expecter[K, V]{mock: &_m.Mock}
}

// BatchDelete provides a mock function for the type MockMap
func (_mock *MockMap[K, V]) BatchDelete(vs []K) error {
	ret := _mock.Called(vs)

	if len(ret) == 0 {
		panic("no return value specified for BatchDelete")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func([]K) error); ok {
		r0 = returnFunc(vs)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockMap_BatchDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchDelete'
type MockMap_BatchDelete_Call[K comparable, V any] struct {
	*mock.Call
}

// BatchDelete is a helper method to define mock.On call
//   - vs
func (_e *MockMap_Expecter[K, V]) BatchDelete(vs interface{}) *MockMap_BatchDelete_Call[K, V] {
	return &MockMap_BatchDelete_Call[K, V]{Call: _e.mock.On("BatchDelete", vs)}
}

func (_c *MockMap_BatchDelete_Call[K, V]) Run(run func(vs []K)) *MockMap_BatchDelete_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]K))
	})
	return _c
}

func (_c *MockMap_BatchDelete_Call[K, V]) Return(err error) *MockMap_BatchDelete_Call[K, V] {
	_c.Call.Return(err)
	return _c
}

func (_c *MockMap_BatchDelete_Call[K, V]) RunAndReturn(run func(vs []K) error) *MockMap_BatchDelete_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// BatchUpdate provides a mock function for the type MockMap
func (_mock *MockMap[K, V]) BatchUpdate(kv map[K]V) error {
	ret := _mock.Called(kv)

	if len(ret) == 0 {
		panic("no return value specified for BatchUpdate")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(map[K]V) error); ok {
		r0 = returnFunc(kv)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockMap_BatchUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchUpdate'
type MockMap_BatchUpdate_Call[K comparable, V any] struct {
	*mock.Call
}

// BatchUpdate is a helper method to define mock.On call
//   - kv
func (_e *MockMap_Expecter[K, V]) BatchUpdate(kv interface{}) *MockMap_BatchUpdate_Call[K, V] {
	return &MockMap_BatchUpdate_Call[K, V]{Call: _e.mock.On("BatchUpdate", kv)}
}

func (_c *MockMap_BatchUpdate_Call[K, V]) Run(run func(kv map[K]V)) *MockMap_BatchUpdate_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[K]V))
	})
	return _c
}

func (_c *MockMap_BatchUpdate_Call[K, V]) Return(err error) *MockMap_BatchUpdate_Call[K, V] {
	_c.Call.Return(err)
	return _c
}

func (_c *MockMap_BatchUpdate_Call[K, V]) RunAndReturn(run func(kv map[K]V) error) *MockMap_BatchUpdate_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function for the type MockMap
func (_mock *MockMap[K, V]) Set(newMap map[K]V) error {
	ret := _mock.Called(newMap)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(map[K]V) error); ok {
		r0 = returnFunc(newMap)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockMap_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockMap_Set_Call[K comparable, V any] struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - newMap
func (_e *MockMap_Expecter[K, V]) Set(newMap interface{}) *MockMap_Set_Call[K, V] {
	return &MockMap_Set_Call[K, V]{Call: _e.mock.On("Set", newMap)}
}

func (_c *MockMap_Set_Call[K, V]) Run(run func(newMap map[K]V)) *MockMap_Set_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[K]V))
	})
	return _c
}

func (_c *MockMap_Set_Call[K, V]) Return(err error) *MockMap_Set_Call[K, V] {
	_c.Call.Return(err)
	return _c
}

func (_c *MockMap_Set_Call[K, V]) RunAndReturn(run func(newMap map[K]V) error) *MockMap_Set_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// SetAndDeferSwitchover provides a mock function for the type MockMap
func (_mock *MockMap[K, V]) SetAndDeferSwitchover(newMap map[K]V) (func(), error) {
	ret := _mock.Called(newMap)

	if len(ret) == 0 {
		panic("no return value specified for SetAndDeferSwitchover")
	}

	var r0 func()
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(map[K]V) (func(), error)); ok {
		return returnFunc(newMap)
	}
	if returnFunc, ok := ret.Get(0).(func(map[K]V) func()); ok {
		r0 = returnFunc(newMap)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}
	if returnFunc, ok := ret.Get(1).(func(map[K]V) error); ok {
		r1 = returnFunc(newMap)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockMap_SetAndDeferSwitchover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAndDeferSwitchover'
type MockMap_SetAndDeferSwitchover_Call[K comparable, V any] struct {
	*mock.Call
}

// SetAndDeferSwitchover is a helper method to define mock.On call
//   - newMap
func (_e *MockMap_Expecter[K, V]) SetAndDeferSwitchover(newMap interface{}) *MockMap_SetAndDeferSwitchover_Call[K, V] {
	return &MockMap_SetAndDeferSwitchover_Call[K, V]{Call: _e.mock.On("SetAndDeferSwitchover", newMap)}
}

func (_c *MockMap_SetAndDeferSwitchover_Call[K, V]) Run(run func(newMap map[K]V)) *MockMap_SetAndDeferSwitchover_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[K]V))
	})
	return _c
}

func (_c *MockMap_SetAndDeferSwitchover_Call[K, V]) Return(fn func(), err error) *MockMap_SetAndDeferSwitchover_Call[K, V] {
	_c.Call.Return(fn, err)
	return _c
}

func (_c *MockMap_SetAndDeferSwitchover_Call[K, V]) RunAndReturn(run func(newMap map[K]V) (func(), error)) *MockMap_SetAndDeferSwitchover_Call[K, V] {
	_c.Call.Return(run)
	return _c
}
