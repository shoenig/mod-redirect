package store

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/shoenig/mod-redirect/internal/mods"
)

// StorageMock implements Storage
type StorageMock struct {
	t minimock.Tester

	funcGet          func(s1 string, s2 string) (rp1 *mods.Redirection, err error)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mStorageMockGet

	funcList          func() (rpa1 []*mods.Redirection, err error)
	afterListCounter  uint64
	beforeListCounter uint64
	ListMock          mStorageMockList

	funcSet          func(rp1 *mods.Redirection) (err error)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mStorageMockSet
}

// NewStorageMock returns a mock for Storage
func NewStorageMock(t minimock.Tester) *StorageMock {
	m := &StorageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetMock = mStorageMockGet{mock: m}
	m.GetMock.callArgs = []*StorageMockGetParams{}

	m.ListMock = mStorageMockList{mock: m}

	m.SetMock = mStorageMockSet{mock: m}
	m.SetMock.callArgs = []*StorageMockSetParams{}

	return m
}

type mStorageMockGet struct {
	mock               *StorageMock
	defaultExpectation *StorageMockGetExpectation
	expectations       []*StorageMockGetExpectation

	callArgs []*StorageMockGetParams
	mutex    sync.RWMutex
}

// StorageMockGetExpectation specifies expectation struct of the Storage.Get
type StorageMockGetExpectation struct {
	mock    *StorageMock
	params  *StorageMockGetParams
	results *StorageMockGetResults
	Counter uint64
}

// StorageMockGetParams contains parameters of the Storage.Get
type StorageMockGetParams struct {
	s1 string
	s2 string
}

// StorageMockGetResults contains results of the Storage.Get
type StorageMockGetResults struct {
	rp1 *mods.Redirection
	err error
}

// Expect sets up expected params for Storage.Get
func (mmGet *mStorageMockGet) Expect(s1 string, s2 string) *mStorageMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("StorageMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &StorageMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &StorageMockGetParams{s1, s2}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Return sets up results that will be returned by Storage.Get
func (mmGet *mStorageMockGet) Return(rp1 *mods.Redirection, err error) *StorageMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("StorageMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &StorageMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &StorageMockGetResults{rp1, err}
	return mmGet.mock
}

//Set uses given function f to mock the Storage.Get method
func (mmGet *mStorageMockGet) Set(f func(s1 string, s2 string) (rp1 *mods.Redirection, err error)) *StorageMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Storage.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Storage.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the Storage.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mStorageMockGet) When(s1 string, s2 string) *StorageMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("StorageMock.Get mock is already set by Set")
	}

	expectation := &StorageMockGetExpectation{
		mock:   mmGet.mock,
		params: &StorageMockGetParams{s1, s2},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up Storage.Get return parameters for the expectation previously defined by the When method
func (e *StorageMockGetExpectation) Then(rp1 *mods.Redirection, err error) *StorageMock {
	e.results = &StorageMockGetResults{rp1, err}
	return e.mock
}

// Get implements Storage
func (mmGet *StorageMock) Get(s1 string, s2 string) (rp1 *mods.Redirection, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	params := &StorageMockGetParams{s1, s2}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.rp1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		want := mmGet.GetMock.defaultExpectation.params
		got := StorageMockGetParams{s1, s2}
		if want != nil && !minimock.Equal(*want, got) {
			mmGet.t.Errorf("StorageMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmGet.GetMock.defaultExpectation.results
		if results == nil {
			mmGet.t.Fatal("No results are set for the StorageMock.Get")
		}
		return (*results).rp1, (*results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(s1, s2)
	}
	mmGet.t.Fatalf("Unexpected call to StorageMock.Get. %v %v", s1, s2)
	return
}

// GetAfterCounter returns a count of finished StorageMock.Get invocations
func (mmGet *StorageMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of StorageMock.Get invocations
func (mmGet *StorageMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to StorageMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mStorageMockGet) Calls() []*StorageMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*StorageMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *StorageMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StorageMock.Get")
		} else {
			m.t.Errorf("Expected call to StorageMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to StorageMock.Get")
	}
}

type mStorageMockList struct {
	mock               *StorageMock
	defaultExpectation *StorageMockListExpectation
	expectations       []*StorageMockListExpectation
}

// StorageMockListExpectation specifies expectation struct of the Storage.List
type StorageMockListExpectation struct {
	mock *StorageMock

	results *StorageMockListResults
	Counter uint64
}

// StorageMockListResults contains results of the Storage.List
type StorageMockListResults struct {
	rpa1 []*mods.Redirection
	err  error
}

// Expect sets up expected params for Storage.List
func (mmList *mStorageMockList) Expect() *mStorageMockList {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("StorageMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &StorageMockListExpectation{}
	}

	return mmList
}

// Return sets up results that will be returned by Storage.List
func (mmList *mStorageMockList) Return(rpa1 []*mods.Redirection, err error) *StorageMock {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("StorageMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &StorageMockListExpectation{mock: mmList.mock}
	}
	mmList.defaultExpectation.results = &StorageMockListResults{rpa1, err}
	return mmList.mock
}

//Set uses given function f to mock the Storage.List method
func (mmList *mStorageMockList) Set(f func() (rpa1 []*mods.Redirection, err error)) *StorageMock {
	if mmList.defaultExpectation != nil {
		mmList.mock.t.Fatalf("Default expectation is already set for the Storage.List method")
	}

	if len(mmList.expectations) > 0 {
		mmList.mock.t.Fatalf("Some expectations are already set for the Storage.List method")
	}

	mmList.mock.funcList = f
	return mmList.mock
}

// List implements Storage
func (mmList *StorageMock) List() (rpa1 []*mods.Redirection, err error) {
	mm_atomic.AddUint64(&mmList.beforeListCounter, 1)
	defer mm_atomic.AddUint64(&mmList.afterListCounter, 1)

	if mmList.ListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmList.ListMock.defaultExpectation.Counter, 1)

		results := mmList.ListMock.defaultExpectation.results
		if results == nil {
			mmList.t.Fatal("No results are set for the StorageMock.List")
		}
		return (*results).rpa1, (*results).err
	}
	if mmList.funcList != nil {
		return mmList.funcList()
	}
	mmList.t.Fatalf("Unexpected call to StorageMock.List.")
	return
}

// ListAfterCounter returns a count of finished StorageMock.List invocations
func (mmList *StorageMock) ListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.afterListCounter)
}

// ListBeforeCounter returns a count of StorageMock.List invocations
func (mmList *StorageMock) ListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.beforeListCounter)
}

// MinimockListDone returns true if the count of the List invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockListDone() bool {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	return true
}

// MinimockListInspect logs each unmet expectation
func (m *StorageMock) MinimockListInspect() {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to StorageMock.List")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		m.t.Error("Expected call to StorageMock.List")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		m.t.Error("Expected call to StorageMock.List")
	}
}

type mStorageMockSet struct {
	mock               *StorageMock
	defaultExpectation *StorageMockSetExpectation
	expectations       []*StorageMockSetExpectation

	callArgs []*StorageMockSetParams
	mutex    sync.RWMutex
}

// StorageMockSetExpectation specifies expectation struct of the Storage.Set
type StorageMockSetExpectation struct {
	mock    *StorageMock
	params  *StorageMockSetParams
	results *StorageMockSetResults
	Counter uint64
}

// StorageMockSetParams contains parameters of the Storage.Set
type StorageMockSetParams struct {
	rp1 *mods.Redirection
}

// StorageMockSetResults contains results of the Storage.Set
type StorageMockSetResults struct {
	err error
}

// Expect sets up expected params for Storage.Set
func (mmSet *mStorageMockSet) Expect(rp1 *mods.Redirection) *mStorageMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("StorageMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &StorageMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &StorageMockSetParams{rp1}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Return sets up results that will be returned by Storage.Set
func (mmSet *mStorageMockSet) Return(err error) *StorageMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("StorageMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &StorageMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &StorageMockSetResults{err}
	return mmSet.mock
}

//Set uses given function f to mock the Storage.Set method
func (mmSet *mStorageMockSet) Set(f func(rp1 *mods.Redirection) (err error)) *StorageMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the Storage.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the Storage.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the Storage.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mStorageMockSet) When(rp1 *mods.Redirection) *StorageMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("StorageMock.Set mock is already set by Set")
	}

	expectation := &StorageMockSetExpectation{
		mock:   mmSet.mock,
		params: &StorageMockSetParams{rp1},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up Storage.Set return parameters for the expectation previously defined by the When method
func (e *StorageMockSetExpectation) Then(err error) *StorageMock {
	e.results = &StorageMockSetResults{err}
	return e.mock
}

// Set implements Storage
func (mmSet *StorageMock) Set(rp1 *mods.Redirection) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	params := &StorageMockSetParams{rp1}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		want := mmSet.SetMock.defaultExpectation.params
		got := StorageMockSetParams{rp1}
		if want != nil && !minimock.Equal(*want, got) {
			mmSet.t.Errorf("StorageMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSet.SetMock.defaultExpectation.results
		if results == nil {
			mmSet.t.Fatal("No results are set for the StorageMock.Set")
		}
		return (*results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(rp1)
	}
	mmSet.t.Fatalf("Unexpected call to StorageMock.Set. %v", rp1)
	return
}

// SetAfterCounter returns a count of finished StorageMock.Set invocations
func (mmSet *StorageMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of StorageMock.Set invocations
func (mmSet *StorageMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to StorageMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mStorageMockSet) Calls() []*StorageMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*StorageMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *StorageMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *StorageMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StorageMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StorageMock.Set")
		} else {
			m.t.Errorf("Expected call to StorageMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to StorageMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StorageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetInspect()

		m.MinimockListInspect()

		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StorageMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *StorageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetDone() &&
		m.MinimockListDone() &&
		m.MinimockSetDone()
}
