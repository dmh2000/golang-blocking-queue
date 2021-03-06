package queue

import (
	"testing"
)

// test variables
const sqsize int = 8

// test an instance of a SynchronizedQueue
func sync1(t *testing.T, q SynchronizedQueue) {
	var err error

	if q == nil {
		t.Error("q should not be nil")
	}

	// check length
	if q.Len() != 0 {
		t.Error("length should be 0",q.Len())
	}

	// check capacity
	if q.Cap() != sqsize {
		t.Error("capacity should == sqsize",q.Cap(),sqsize)
	}

	// fill the queue with ints
	for i:=0;i<q.Cap();i++ {
		q.TryPut(i)
		//length should be == i at this point
		if q.Len() != (i+1) {
			t.Error("length should == i+1",q.Len(),i+1)
		}
	}

	// check the length, should be == capacity
	if q.Len() != q.Cap() {
		t.Error("length should == capacity")
	}

	// try to add one more
	err = q.TryPut(99)
	if err == nil {
		t.Error("err should be nil")
	}
	// check length is unchanged
	if q.Len() != sqsize {
		t.Error("length should == sqsize")
	}

	// cleanup 
	// for channels, this closes it for further Puts
	// any remaing data is still available for Gets
	// it is a noop for the mutex/condition variable methods
	q.Close()

	// remove all items. no need to block
	j := q.Len() - 1
	for i:=0;i<q.Cap();i++ {
		value, err := q.TryGet()
		if err != nil {
			t.Error(err)
		}
		// convert to int
		v := value.(int)
		if v != i {
			t.Error("v should == i",v,i)
		}
		// length should decrease
		if q.Len() != j {
			t.Error("length should == i",q.Len(),j)
		}
		j--
	}

	// check the length == 0
	if q.Len() != 0 {
		t.Error("length should == 0",q.Len())
	}
}


// test an instance of a SynchronizedQueue
func sync2(t *testing.T, q *NativeIntQueue) {
	var err error

	if q == nil {
		t.Error("q should not be nil")
	}

	// check length
	if q.Len() != 0 {
		t.Error("length should be 0",q.Len())
	}

	// check capacity
	if q.Cap() != sqsize {
		t.Error("capacity should == sqsize",q.Cap(),sqsize)
	}

	// fill the queue with ints
	for i:=0;i<q.Cap();i++ {
		q.Put(i)
		//length should be == i at this point
		if q.Len() != (i+1) {
			t.Error("length should == i+1",q.Len(),i+1)
		}
	}

	// check the length, should be == capacity
	if q.Len() != q.Cap() {
		t.Error("length should == capacity")
	}

	// try to add one more
	err = q.TryPut(99)
	if err == nil {
		t.Error("err should be nil")
	}
	// check length is unchanged
	if q.Len() != sqsize {
		t.Error("length should == sqsize")
	}

	// cleanup 
	// for channels, this closes it for further Puts
	// any remaing data is still available for Gets
	// it is a noop for the mutex/condition variable methods
	q.Close()

	// remove all items
	j := q.Len() - 1
	for i:=0;i<q.Cap();i++ {
		v, err := q.TryGet()
		if err != nil {
			t.Error(err)
		}
		if v != i {
			t.Error("v should == i",v,i)
		}
		// length should decrease
		if q.Len() != j {
			t.Error("length should == i",q.Len(),j)
		}
		j--
	}

	// check the length == 0
	if q.Len() != 0 {
		t.Error("length should == 0",q.Len())
	}
}

// ====================
// SYNCHRONOUS TESTS
// ====================

// CHANNEL
func TestChannelSync(t *testing.T) {
	// using channel
	sync1(t,NewChannelQueue(sqsize))
}

// NATIVE QUEUE
func TestQueueNativeSync(t *testing.T) {
	// using condition variable queue
	sync2(t,NewNativeQueue(sqsize))
}


// CIRCULAR BUFFER using SynchronizedQueue wrapper
func TestCircularQueueSync(t *testing.T) {
	// using condition variable queue
	sync1(t,NewSyncCircular(sqsize))
}

// LIST QUEUE using SynchronizedQueue wrapper
func TestListSync(t *testing.T) {
	// using condition variable queue
	sync1(t,NewSyncList(sqsize))
}

// RING QUEUE using SynchronizedQueue wrapper
func TestRingSync(t *testing.T) {
	// using condition variable queue
	sync1(t,NewSyncRing(sqsize))
}

// SLICE QUEUE using SynchronizedQueue wrapper
func TestSliceSync(t *testing.T) {
	// using condition variable queue
	sync1(t,NewSyncSlice(sqsize))
}

// Strings
func TestStringsSync(t *testing.T) {
	var q SynchronizedQueue
	q = NewChannelQueue(sqsize)
	q.Put(1)
	t.Log(q.String())

	q = NewSyncCircular(sqsize)
	q.Put(1)
	t.Log(q.String())

	q = NewSyncList(sqsize)
	q.Put(1)
	t.Log(q.String())

	q = NewSyncRing(sqsize)
	q.Put(1)
	t.Log(q.String())

	q = NewSyncPriority(sqsize)
	q.Put(PriorityItem{1,1})
	t.Log(q.String())

	q = NewSyncSlice(sqsize)
	q.Put(1)
	t.Log(q.String())

	// native queue isn't wrapped
	nq := NewNativeQueue(sqsize)
	nq.Put(1)
	t.Log(nq.String())
}
