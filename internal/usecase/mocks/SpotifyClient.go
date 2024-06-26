// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	spotify "github.com/zmb3/spotify/v2"
)

// SpotifyClient is an autogenerated mock type for the SpotifyClient type
type SpotifyClient struct {
	mock.Mock
}

// Search provides a mock function with given fields: ctx, query, t, opts
func (_m *SpotifyClient) Search(ctx context.Context, query string, t spotify.SearchType, opts ...spotify.RequestOption) (*spotify.SearchResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, query, t)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 *spotify.SearchResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, spotify.SearchType, ...spotify.RequestOption) (*spotify.SearchResult, error)); ok {
		return rf(ctx, query, t, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, spotify.SearchType, ...spotify.RequestOption) *spotify.SearchResult); ok {
		r0 = rf(ctx, query, t, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spotify.SearchResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, spotify.SearchType, ...spotify.RequestOption) error); ok {
		r1 = rf(ctx, query, t, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSpotifyClient creates a new instance of SpotifyClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSpotifyClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SpotifyClient {
	mock := &SpotifyClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
