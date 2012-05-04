package h

const (
	Post = "post"
	Get  = "get"
)

type Form struct {
	ID     string `h:"attr"`
	Class  string `h:"attr"`
	Style  string `h:"attr"`
	Action string `h:"attr"`
	Method string `h:"attr"`
	Target string `h:"attr"`
	Title  string `h:"attr"`
	Inner  HTML   `h:"inner"`
}

func (f *Form) HTML() (HTML, error) {
	return &ReflectNode{Tag: "form", Node: f}, nil
}
