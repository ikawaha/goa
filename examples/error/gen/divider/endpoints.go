// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// divider endpoints
//
// Command:
// $ goa gen goa.design/goa/examples/error/design -o
// $(GOPATH)/src/goa.design/goa/examples/error

package dividersvc

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "divider" service endpoints.
type Endpoints struct {
	IntegerDivide goa.Endpoint
	Divide        goa.Endpoint
}

// NewEndpoints wraps the methods of the "divider" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		IntegerDivide: NewIntegerDivideEndpoint(s),
		Divide:        NewDivideEndpoint(s),
	}
}

// Use applies the given middleware to all the "divider" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.IntegerDivide = m(e.IntegerDivide)
	e.Divide = m(e.Divide)
}

// NewIntegerDivideEndpoint returns an endpoint function that calls the method
// "integer_divide" of service "divider".
func NewIntegerDivideEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IntOperands)
		return s.IntegerDivide(ctx, p)
	}
}

// NewDivideEndpoint returns an endpoint function that calls the method
// "divide" of service "divider".
func NewDivideEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FloatOperands)
		return s.Divide(ctx, p)
	}
}