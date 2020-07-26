package sideroutine

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutine(t *testing.T) {
	for _, data := range []struct {
		name   string
		input  []func() (Interface, error)
		res    interface{}
		errors Errors
	}{
		{
			name: "basic",
			input: []func() (Interface, error){
				func() (Interface, error) { return 42, nil },
			},
			res:    []interface{}{42},
			errors: nil,
		},
		{
			name: "basic with error",
			input: []func() (Interface, error){
				func() (Interface, error) { return nil, errors.New("an error occurred") },
			},
			res:    nil,
			errors: Errors{errors.New("an error occurred")},
		},
		{
			name: "multiple input methods",
			input: []func() (Interface, error){
				func() (Interface, error) { return 42, nil },
				func() (Interface, error) { return 21, nil },
				func() (Interface, error) { return 12, nil },
			},
			res:    []interface{}{42, 21, 12},
			errors: nil,
		},
		{
			name: "multiple input methods with one error",
			input: []func() (Interface, error){
				func() (Interface, error) { return 42, nil },
				func() (Interface, error) { return nil, errors.New("fail") },
				func() (Interface, error) { return 12, nil },
			},
			res:    []interface{}{42, 12},
			errors: Errors{errors.New("fail")},
		},
		{
			name: "only errors",
			input: []func() (Interface, error){
				func() (Interface, error) { return nil, errors.New("fail1") },
				func() (Interface, error) { return nil, errors.New("fail2") },
				func() (Interface, error) { return nil, errors.New("fail3") },
			},
			res:    nil,
			errors: Errors{errors.New("fail1"), errors.New("fail2"), errors.New("fail3")},
		},
	} {
		t.Run(data.name, func(t *testing.T) {
			r := NewRoutine()
			r.Add(data.input...)
			r.Run()
			r.Wait()
			res, errs := r.Extract()
			assert.ElementsMatch(t, data.res, res)
			assert.ElementsMatch(t, data.errors, errs)

			if len(errs) > 0 {
				assert.Equal(t, errs[0], errs.First())
			} else {
				assert.Equal(t, nil, errs.First())
			}
		})
	}
}
