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
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"log/slog"
	"sync"

	"github.com/alexandremahdhaoui/tooling/pkg/flaterrors"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/ringbuf"
)

// FIFO[T] can be used to subscribe to structured notifications produced
// by a bpf program.
//
// Generics constraints:
// - T must be a **struct**.
// - T must not be a pointer.
// - T must not be an interface.
//
// Notes:
// - FIFO is thread-safe.
// - Subscribe() can be called only once.
type FIFO[T any] interface {
	// Done returns a channel that's closed when work done on behalf of this
	// interface has been gracefully terminated.
	//
	// While the channel from Subscribe() could be used to know if the FIFO
	// interface has been closed, using the channel from this interface is
	// more idiomatic.
	Done() <-chan struct{}

	// Subscribe returns a receiver channel of T or an error.
	Subscribe() (<-chan T, error)
}

type bpfFifo[T any] struct {
	rb    *ringbuf.Reader
	mu    *sync.Mutex
	inUse bool
	// doneCh is a channel used to notify the bpf data structures or bpf
	// program has been closed and they can no longer be used.
	doneCh <-chan struct{}
}

// Generics constraints:
// - T must be a **struct**.
// - T must not be a pointer.
// - T must not be an interface.
//
// doneCh is a channel used to notify the bpf data structures or bpf
// program has been closed and they can no longer be used.
func NewFIFO[T any](ringbufMap *ebpf.Map, doneCh <-chan struct{}) (FIFO[T], error) {
	rb, err := ringbuf.NewReader(ringbufMap)
	if err != nil {
		return nil, err
	}

	return &bpfFifo[T]{
		rb:     rb,
		mu:     &sync.Mutex{},
		inUse:  false,
		doneCh: doneCh,
	}, nil
}

func (f *bpfFifo[T]) Done() <-chan struct{} {
	return f.doneCh
}

var (
	ErrAnotherProcessAlreadySubscribed = errors.New("another process already subscribed")
	ErrSubscribingToFIFO               = errors.New("subscribing to fifo")
)

func (f *bpfFifo[T]) Subscribe() (<-chan T, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.inUse {
		return nil, flaterrors.Join(ErrAnotherProcessAlreadySubscribed, ErrSubscribingToFIFO)
	}
	f.inUse = true

	ch := make(chan T) // TODO: buffer it.

	go func() {
		for {
			rec, err := f.rb.Read()
			if err != nil {
				slog.ErrorContext(
					context.TODO(),
					"an unexpected error occured reading from bpf ring buffer",
					"err",
					err.Error(),
				)
			}

			v := new(T)

			if err := binary.Read(bytes.NewReader(rec.RawSample), binary.NativeEndian, v); err != nil {
				slog.ErrorContext(
					context.TODO(),
					"an error occured decoding record from bpf ring buffer",
					"err",
					err.Error(),
				)
			}

			ch <- *v
		}
	}()

	return ch, nil
}
