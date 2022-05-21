package main

// Interfaces are usually named with “er” endings. We’ve already seen fmt.Stringer, but there are many more, including io.Reader, io.Closer, io.ReadCloser, json.Marshaler, and http.Handler.
type Stringer interface {
	String() string
}

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return data
}

//--------------------------------------------------------

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	c.L.Process("string")
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
