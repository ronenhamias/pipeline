// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package processdriver

import (
	"context"
	"errors"
	"github.com/banzaicloud/pipeline/internal/app/pipeline/process"
	auth "github.com/banzaicloud/pipeline/src/auth"
	"github.com/go-kit/kit/endpoint"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetProcess    endpoint.Endpoint
	ListProcesses endpoint.Endpoint
	Log           endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service process.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{
		GetProcess:    kitxendpoint.OperationNameMiddleware("process.GetProcess")(mw(MakeGetProcessEndpoint(service))),
		ListProcesses: kitxendpoint.OperationNameMiddleware("process.ListProcesses")(mw(MakeListProcessesEndpoint(service))),
		Log:           kitxendpoint.OperationNameMiddleware("process.Log")(mw(MakeLogEndpoint(service))),
	}
}

// GetProcessRequest is a request struct for GetProcess endpoint.
type GetProcessRequest struct {
	Org auth.Organization
	Id  string
}

// GetProcessResponse is a response struct for GetProcess endpoint.
type GetProcessResponse struct {
	Process process.Process
	Err     error
}

func (r GetProcessResponse) Failed() error {
	return r.Err
}

// MakeGetProcessEndpoint returns an endpoint for the matching method of the underlying service.
func MakeGetProcessEndpoint(service process.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProcessRequest)

		process, err := service.GetProcess(ctx, req.Org, req.Id)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return GetProcessResponse{
					Err:     err,
					Process: process,
				}, nil
			}

			return GetProcessResponse{
				Err:     err,
				Process: process,
			}, err
		}

		return GetProcessResponse{Process: process}, nil
	}
}

// ListProcessesRequest is a request struct for ListProcesses endpoint.
type ListProcessesRequest struct {
	Org   auth.Organization
	Query map[string]string
}

// ListProcessesResponse is a response struct for ListProcesses endpoint.
type ListProcessesResponse struct {
	Processes []process.Process
	Err       error
}

func (r ListProcessesResponse) Failed() error {
	return r.Err
}

// MakeListProcessesEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListProcessesEndpoint(service process.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListProcessesRequest)

		processes, err := service.ListProcesses(ctx, req.Org, req.Query)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListProcessesResponse{
					Err:       err,
					Processes: processes,
				}, nil
			}

			return ListProcessesResponse{
				Err:       err,
				Processes: processes,
			}, err
		}

		return ListProcessesResponse{Processes: processes}, nil
	}
}

// LogRequest is a request struct for Log endpoint.
type LogRequest struct {
	Proc process.Process
}

// LogResponse is a response struct for Log endpoint.
type LogResponse struct {
	Process process.Process
	Err     error
}

func (r LogResponse) Failed() error {
	return r.Err
}

// MakeLogEndpoint returns an endpoint for the matching method of the underlying service.
func MakeLogEndpoint(service process.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LogRequest)

		process, err := service.Log(ctx, req.Proc)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return LogResponse{
					Err:     err,
					Process: process,
				}, nil
			}

			return LogResponse{
				Err:     err,
				Process: process,
			}, err
		}

		return LogResponse{Process: process}, nil
	}
}
