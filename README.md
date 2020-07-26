# Routine

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jcalmat/routine)
![Coverage](https://img.shields.io/codecov/c/github/jcalmat/routine)


Routine offers a convenient way to use basic goroutines without dealing with channels or waitgroups.

## Usage

```go    
import "github.com/jcalmat/routine"

func main() {

	routine := routine.NewRoutine()

	sample := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for _, s := range sample {
		s := s
		routine.Add(func() (Interface, error) {
			// do some heavy computations/API calls/etc
			return s + 42, nil
		})
	}

	routine.Run()
	routine.Wait()
    
    fmt.Println(routine.Extract()) // [50 58 74 106 170 43 44 46] []
}
```

If you need to assert the interface array to a specific type, consider looping through the result and assert the values as needed

```go
	custom := make([]int, 0)
	for _, v := range res {
		custom = append(custom, v.(int))
	}
	fmt.Println(custom)
```