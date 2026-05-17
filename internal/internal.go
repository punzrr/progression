package internal

const (
	ErrZeroIndex = Sentinel("cannot operate with zero index")
	ErrNegIndex  = Sentinel("cannot operate with negative index")
)

// Sentinel struct represent [sentinel errors]
//
// [sentinel errors]: https://en.wikipedia.org/wiki/Exception_handling_(programming)
type Sentinel string

func (err Sentinel) Error() string {
	return string(err)
}

func (err Sentinel) Unwrap() error {
	return err
}
