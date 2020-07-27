# Routine

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jcalmat/routine)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Linter](https://github.com/jcalmat/routine/workflows/golangci-lint/badge.svg)
![Unit tests](https://github.com/jcalmat/routine/workflows/tests/badge.svg)

Routine offers a convenient way to use basic goroutines without dealing with channels or waitgroups.

## When would you use it?

If you need to do some heavy computations repetitively typically inside a loop, you might want to use golang concurrency patterns to increase your performances. Unfortunately sometimes using the goroutines, channels and waitgroup syntax could be a bit messy, that's what Routine is for.
Routine is a very simple abstraction of the concurrency pattern which allows you to make your repetitive pieces of code concurrent.

### Example

If you're working in a microservice architecture and have to loop through a list of IDs to fetch some data from other microservices, you'll may want to make a loop fetching each data individually. Depending on your set of IDs' length, this small code could take a while to execute.

```go
type Data struct {
	Name  string
	Email string
}

func FetchDatas(userIDs ...string) ([]Data, error) {
	datas := make([]Data, 0)

	for _, id := range userIDs {
        id := id //shadow

        // perform the calls one by one
		usersResp, err := users.Fetch(id)
		if err != nil {
			return nil, err
		}
		datas = append(datas, Data{
			Name: usersResp.Name,
			Email: usersResp.Email,
		})
	}

	return datas, nil
}
```

In this case, in order to increase the performances, you could use Routine to make all the calls concurrent

```go
import "github.com/jcalmat/routine"

func FetchDatas(userIDs ...string) ([]Data, error) {
	// init the routine
	r := routine.NewRoutine()

	for _, id := range userIDs {
		id := id // shadow

		// add each method without executing the code yet
		r.Add(func() (routine.Interface, error) {
			usersResp, err := users.Fetch(id)
			if err != nil {
				return nil, err
			}
			return Data{
				Name: usersResp.Name,
				Email: usersResp.Email,
			}, nil
		})
	}

	// run the routines and extract the values
	res, errs := r.Run()
	datas := make([]Data, 0)

	// if you need to assert the interface array to a specific type, consider looping through the result and assert the values as needed
	for _, v := range res {
		datas = append(datas, v.(Data))
	}

	// return
    return datas, errs.First()
}
```
