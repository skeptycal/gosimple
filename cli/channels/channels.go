package channels

type (
	None     = struct{}
	Done     GenericDoneChannel[None]
	DoneBool GenericDoneChannel[bool]

	GenericDoneChannel[T any] struct {
		done   chan T
		backup chan T
	}
)

func NewDoneChan() Done { return Done{make(chan None), nil} }

func (*GenericDoneChannel[T]) blank() T   { return *new(T) }
func (d *GenericDoneChannel[T]) Send()    { d.done <- d.blank() }
func (d *GenericDoneChannel[T]) Receive() { <-d.done }
func (d *GenericDoneChannel[T]) Close()   { close(d.done) }
func (d *GenericDoneChannel[T]) Pause()   { d.backup, d.done = d.done, nil }
func (d *GenericDoneChannel[T]) Resume()  { d.backup, d.done = d.done, d.backup }

var blank = None{}

func isDone(done chan None) {
	done <- blank
}
