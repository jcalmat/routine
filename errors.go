package routine

// First returns the first error from a slice of errors or nil if there is none
func (e Errors) First() error {
	for _, err := range e {
		return err
	}
	return nil
}
