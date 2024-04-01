package foo

type Impl struct{}

func init() {
	instance = &Impl{}
}

//export fermyon:spin/foo@2.0.0#greet
func Greet2(i *Impl, s *string) {
	i.Greet(s)
}

func (i *Impl) Greet(s *string) {
	*s = "hello from greet"
}
