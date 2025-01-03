// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	config "github.com/aws-controllers-k8s/runtime/pkg/config"
	aws "github.com/aws/aws-sdk-go-v2/aws"

	logr "github.com/go-logr/logr"

	manager "sigs.k8s.io/controller-runtime/pkg/manager"

	mock "github.com/stretchr/testify/mock"

	prometheus "github.com/prometheus/client_golang/prometheus"

	schema "k8s.io/apimachinery/pkg/runtime/schema"

	types "github.com/aws-controllers-k8s/runtime/pkg/types"

	v1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
)

// ServiceController is an autogenerated mock type for the ServiceController type
type ServiceController struct {
	mock.Mock
}

// BindControllerManager provides a mock function with given fields: _a0, _a1
func (_m *ServiceController) BindControllerManager(_a0 manager.Manager, _a1 config.Config) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(manager.Manager, config.Config) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMetadata provides a mock function with given fields:
func (_m *ServiceController) GetMetadata() types.ServiceControllerMetadata {
	ret := _m.Called()

	var r0 types.ServiceControllerMetadata
	if rf, ok := ret.Get(0).(func() types.ServiceControllerMetadata); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.ServiceControllerMetadata)
	}

	return r0
}

// GetReconcilers provides a mock function with given fields:
func (_m *ServiceController) GetReconcilers() []types.AWSResourceReconciler {
	ret := _m.Called()

	var r0 []types.AWSResourceReconciler
	if rf, ok := ret.Get(0).(func() []types.AWSResourceReconciler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.AWSResourceReconciler)
		}
	}

	return r0
}

// GetResourceManagerFactories provides a mock function with given fields:
func (_m *ServiceController) GetResourceManagerFactories() map[string]types.AWSResourceManagerFactory {
	ret := _m.Called()

	var r0 map[string]types.AWSResourceManagerFactory
	if rf, ok := ret.Get(0).(func() map[string]types.AWSResourceManagerFactory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]types.AWSResourceManagerFactory)
		}
	}

	return r0
}

// NewConfig provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *ServiceController) NewConfig(_a0 context.Context, _a1 v1alpha1.AWSRegion, _a2 *string, _a3 v1alpha1.AWSResourceName, _a4 schema.GroupVersionKind) (aws.Config, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 aws.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1alpha1.AWSRegion, *string, v1alpha1.AWSResourceName, schema.GroupVersionKind) (aws.Config, error)); ok {
		return rf(_a0, _a1, _a2, _a3, _a4)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1alpha1.AWSRegion, *string, v1alpha1.AWSResourceName, schema.GroupVersionKind) aws.Config); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Get(0).(aws.Config)
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1alpha1.AWSRegion, *string, v1alpha1.AWSResourceName, schema.GroupVersionKind) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithLogger provides a mock function with given fields: _a0
func (_m *ServiceController) WithLogger(_a0 logr.Logger) types.ServiceController {
	ret := _m.Called(_a0)

	var r0 types.ServiceController
	if rf, ok := ret.Get(0).(func(logr.Logger) types.ServiceController); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ServiceController)
		}
	}

	return r0
}

// WithPrometheusRegistry provides a mock function with given fields: _a0
func (_m *ServiceController) WithPrometheusRegistry(_a0 prometheus.Registerer) types.ServiceController {
	ret := _m.Called(_a0)

	var r0 types.ServiceController
	if rf, ok := ret.Get(0).(func(prometheus.Registerer) types.ServiceController); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ServiceController)
		}
	}

	return r0
}

// WithResourceManagerFactories provides a mock function with given fields: _a0
func (_m *ServiceController) WithResourceManagerFactories(_a0 []types.AWSResourceManagerFactory) types.ServiceController {
	ret := _m.Called(_a0)

	var r0 types.ServiceController
	if rf, ok := ret.Get(0).(func([]types.AWSResourceManagerFactory) types.ServiceController); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ServiceController)
		}
	}

	return r0
}

// NewServiceController creates a new instance of ServiceController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceController(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceController {
	mock := &ServiceController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
