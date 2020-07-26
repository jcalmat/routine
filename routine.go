package sideroutine

import (
	"sync"
)

// R contains the required informations to abstract goroutine concurrency logic
type R struct {
	wg      *sync.WaitGroup
	methods []func()
	result  chan result
}

// Result represents a sideroutine.goroutine response type
type result struct {
	result interface{}
	errors error
}

// NewRoutine initializes a new Routine
func NewRoutine() *R {
	return &R{
		wg:      new(sync.WaitGroup),
		result:  make(chan result),
		methods: make([]func(), 0),
	}
}

// Add inserts methods to the Routine after adding them the goroutines logic
func (r *R) Add(methods ...func() (Interface, error)) {
	for _, m := range methods {
		m := m
		r.wg.Add(1)
		r.methods = append(r.methods, func() {
			go func() {
				defer r.wg.Done()
				res, err := m()
				r.result <- result{
					result: res,
					errors: err,
				}
			}()
		})
	}
}

// Run executes all the routines stacked in r
func (r *R) Run() {
	for _, m := range r.methods {
		m()
	}
}

// Wait waits for all the goroutines to be done before closing the result channel
func (r *R) Wait() {
	go func() {
		r.wg.Wait()
		close(r.result)
	}()
}

// Extract extracts the datas stored in the result channel
func (r *R) Extract() ([]interface{}, Errors) {
	var result []interface{}
	var errors Errors

	for res := range r.result {
		if res.result != nil {
			result = append(result, res.result)
		}
		if res.errors != nil {
			errors = append(errors, res.errors)
		}
	}
	return result, errors
}

// Flush flushes any channel and buffered methods to the underlying routine
func (r *R) Flush() {
	r.methods = make([]func(), 0)
	for len(r.result) > 0 {
		<-r.result
	}
}
