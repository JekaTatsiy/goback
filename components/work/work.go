package work

type Work struct {
	Name   string
	Action func() error
}
