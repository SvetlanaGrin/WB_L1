package main

type Target interface {
	Request() string
}

type Adaptee struct {
}

func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

type Adapter struct {
	*Adaptee
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
