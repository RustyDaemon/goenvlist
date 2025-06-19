package core

type Options struct {
	Simple bool
	Path   bool
	Raw    bool
	Filter []string
}

func NewOptions() *Options {
	return &Options{
		Simple: false,
		Path:   false,
		Raw:    false,
		Filter: []string{},
	}
}
