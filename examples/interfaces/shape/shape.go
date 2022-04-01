package shape

type Shape interface {
	Name() string
}

type Circle struct {
	Text string
}

func (c *Circle) Name() string {
	return c.Text
}
