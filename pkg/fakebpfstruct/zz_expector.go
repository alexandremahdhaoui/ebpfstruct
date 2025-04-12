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

import "fmt"

type expectation struct {
	method string
	err    error
}

type expector struct {
	ordered bool
	eList   []expectation
	eMap    map[string]error
}

func (e *expector) EXPECT(method string, err error) *expector {
	if e.ordered {
		e.appendExpectation(method, err)
	}
	e.addExpectation(method, err)
	return e
}

func (e *expector) appendExpectation(method string, err error) {
	e.eList = append(e.eList, expectation{
		method: method,
		err:    err,
	})
}

func (e *expector) addExpectation(method string, err error) {
	if e.eMap == nil {
		e.eMap = make(map[string]error)
	}
	e.eMap[method] = err
}

func (e *expector) checkExpectation(method string) error {
	if e.ordered {
		if len(e.eList) == 0 {
			panic("unexpected method call")
		}

		expected := e.eList[0]
		if expected.method != method {
			panic(fmt.Sprintf("unexpected method call; want: %s; got: %s", expected.method, method))
		}

		if len(e.eList) > 1 {
			e.eList = e.eList[1:]
		} else {
			e.eList = nil
		}

		return expected.err
	}

	if len(e.eMap) == 0 {
		panic("unexpected method call")
	}

	expectedErr, ok := e.eMap[method]
	if !ok {
		panic(fmt.Sprintf("unexpected method call; got: %s", method))
	}

	return expectedErr
}
