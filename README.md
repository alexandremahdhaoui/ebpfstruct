# ebpfstruct

Useful data structure for eBPF programming in Go. They conviniently wrap ebpf data structures
in a safe-thread manner.

These data structures offering an abstract interface simplifies testing components.

## Getting started `ebpfstruct`

To get started using `ebpfstruct` please follow the 3 step presented below:

### 1. Get the package

```shell
go get github.com/alexandremahdhaoui/ebpfstruct
```

### 2. Import `ebpfstruct`

```go
package your_package

import "github.com/alexandremahdhaoui/ebpfstruct"

func yourFunc() {
    // [...]
    arr := ebpfstruct.NewArray(
	    a, b,
	    aLen, bLen,
	    activePointer,
    )
    // [...]
}
```

### 3. Learn the package

There is four way to learn how to use this package:

- Learn by doing:
  - With the examples in the section below.

- Learn by reading the documentation:
  - I keep documentation as close as possible from their source of truth.
  - Hence, you'll find relevant information close to the interface or struct
    you're studying.
  - If you need to learn even more, you can always dive deeper and check how the
  concrete implementation of an interface or their dependencies are implemented.
  - Please note that all interfaces are themselves briefly documented below.

- Learn by asking for help:
  - Maybe something is missing from the documentation or a piece of information 
    is unclear.
  - Feel free to open an issue and ask.

- Learn by reading the code:
  - That's the most straightforward, start with an interface and check how it's 
    implemented.
  - All files are similarly structured: they start with an interface declaration 
    and its concrete implementation lays below.

## Examples

This example is taken from [https://github.com/alexandremahdhaoui/udplb](https://github.com/alexandremahdhaoui/udplb)

#### Introduction

You're writing a loadbalancer and needs to maintain the following data structures.

Let `bpfBackendListT` the ebpf-go representation of a BPF map of type array. This
data structure contains a list of available backend for your load balancer 
program.

Let `bpfLookupTableT` the ebpf-go representation of another BPF map of type array.
The data structure is a lookup table, that maps each entry to the index of an
available backend in the map `bpfBackendListT`

#### The problem

- If a packet is received while both data structures are not updated, then we will
  forward a packet to a potentially wrong backend.
- If a large array is being udpated we do not want to delay packet processing or degrade performance e.g. because of spinlocks.
- If an error occured while updating the second data structure (e.g. the lookup table), how 
  can we revert the changes that affected the fist data structure (e.g. the backend list).

#### The solution

Excerpt of a comment in the code of the Array[T] concrete implementation:

	// The idea is to internally use 2 BPF maps for each data structure and
	// when updating 1 Array[T], we update the internal BPF map that is not 
	// currently being used by the BPF program.
	//
	// The BPF program will check a variable that will let it know which map
	// to read from.
	//
	// This is a sort of a blue/green deployment of the new data structure.
	// This solution simplifies error handling as it ensures the whole data
	// structure is atomically updated from the bpf program point of view.
	//
	// a & b are the internal maps, either one or the other is in the active
	// state while the other one is in passive state.
	// The userland program will update data structures in passive states and
	// perform a switchover to "notify" the BPF program which map is in active
	// state.

#### The example (finally)

```go
func main() {
    objs := &bpfObjects{}
	if err = spec.LoadAndAssign(lb.objs, nil); err != nil {
        slog.Error(err.Error)
        os.Exit(1)
	}

    // [...]

	backendList, err := ebpfstruct.NewArray[*bpfBackendListT](
		objs.BackendListA,
		objs.BackendListB,
		objs.BackendListA_len,
		objs.BackendListB_len,
		objs.ActivePointer,
	)
    if err != nil {
        slog.Error(err.Error)
        os.Exit(1)
    }

	lookupTable, err := ebpfstruct.NewArray[uint32](
		objs.LookupTableA,
		objs.LookupTableB,
		objs.LookupTableA_len,
		objs.LookupTableB_len,
		objs.ActivePointer,
	)
    if err != nil {
        slog.Error(err.Error)
        os.Exit(1)
    }

    // [...]

    for retryCount := 0; retryCount < 3;{
        // You received an event that updates the status of a backend. 
        // One of the backend became unavailable and you need to update
        // both backendList and lookupTable data structures.
        newBackendList := <- eventCh
        newLookupTable := recomputeLup(newBackendList)

        // Updates the passive backend list.
        backendSwitchover, err := backendList.SetAndDeferSwitchover()
        if err != nil { // retry on failure
            // -> because we defer the switchover after all data structures
            //    are updated, we can simply retry on failure.
            eventCh <- newBackendList
            retryCount += 1
            slog.Error(err.Error)
            continue
        }

        // Updates the passive lookup table.
        lupSwitchover, err := lookupTable.SetAndDeferSwitchover()
        if err != nil { // retry on failure
            // -> because we defer the switchover after all data structures
            //    are updated, we can simply retry on failure.
            eventCh <- newBackendList
            retryCount += 1
            slog.Error(err.Error)
            continue
        }

        // Perform the switchover:
        // -> The switchover is done in sync for all involved data structures,
        //    as soon as the first deferable switchover function is called.
        // -> However, we MUST call all switchover function in order to update
        //    the internal state of all go data structures.
        backendSwitchover()
        lupSwitchover()
        retryCount = 0
    }

    slog.Error("program failed 3 times to update bpf data structures")
    os.Exit(1)
}
```

## Interfaces

[//] # TODO: generate interfaces from code.

## Fakes

Fakes can be injected into components for testing purposes.

```go
package your_test

import (
    "github.com/alexandremahdhaoui/ebpfstruct/pkg/fakebpfstruct"
    "github.com/stretchr/testify/assert"
) 

func TestSomething(t *testing.T) {
    // ...
    expectedArrayInternalValue := [0, 1, 2, 3]
    arr := fakebpfstruct.NewArray[T]()
    arr.
        ORDERED().
        EXPECT("Set", errors.New("a random error to see if your component handles unexpected errors"), 1).
        EXPECT(
            "SetAndDeferSwitchover", // expected method
            nil,                     // if the method returns an error, it will return this value.
            1,                       // the number of times it is expected.
        )
    yourComponent := NewComponentDependingOnArray(arr)
    err := yourComponent.RunWithRetry(expectedArrayInternalValue)
    assert.NoError(t, err)
    assert.Equal(t, arr.Array, expectedArrayInternalValue)
    // ...
}
```

## Mocks

Mocks can also be injected into components for testing purposes.

```go
package your_test

import "github.com/alexandremahdhaoui/ebpfstruct/pkg/mockebpfstruct"

func TestSomething(t *testing.T) {
    // ...
    arr := mockebpfstruct.NewArray[any](t *testing.T)
    arr.
        EXPECT().
        Set(/* expectations */).
        Times(1)
    yourComponent := NewComponentDependingOnArray(arr)
    // ...
}
```

