// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package types

import (
	"context"

	awscfg "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/go-logr/logr"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcfg "github.com/aws-controllers-k8s/runtime/pkg/config"
	ackmetrics "github.com/aws-controllers-k8s/runtime/pkg/metrics"
)

// AWSResourceManager is responsible for providing a consistent way to perform
// CRUD+L operations in a backend AWS service API for Kubernetes custom
// resources (CR) corresponding to those AWS service API resources.
//
// Use an AWSResourceManagerFactory to create an AWSResourceManager for a
// particular APIResource and AWS account.
type AWSResourceManager interface {
	ReferenceManager
	// ReadOne returns the currently-observed state of the supplied AWSResource
	// in the backend AWS service API.
	//
	// Implementers should return (nil, ackerrors.NotFound) when the backend
	// AWS service API cannot find the resource identified by the supplied
	// AWSResource's AWS identifier information.
	ReadOne(context.Context, AWSResource) (AWSResource, error)
	// Create attempts to create the supplied AWSResource in the backend AWS
	// service API, returning an AWSResource representing the newly-created
	// resource
	Create(context.Context, AWSResource) (AWSResource, error)
	// Update attempts to mutate the supplied desired AWSResource in the
	// backend AWS service API, returning an AWSResource representing the
	// newly-mutated resource.
	// Note for specialized logic implementers can check to see how the latest
	// observed resource differs from the supplied desired state. The
	// higher-level reconciler determines whether or not the desired differs
	// from the latest observed and decides whether to call the resource
	// manager's Update method
	Update(
		context.Context,
		AWSResource, /* desired */
		AWSResource, /* latest */
		*ackcompare.Delta,
	) (AWSResource, error)

	// Delete attempts to destroy the supplied AWSResource in the backend AWS
	// service API, returning an AWSResource representing the
	// resource being deleted (if delete is asynchronous and takes time)
	Delete(context.Context, AWSResource) (AWSResource, error)
	// ARNFromName returns an AWS Resource Name from a given string name. This
	// is useful for constructing ARNs for APIs that require ARNs in their
	// GetAttributes operations but all we have (for new CRs at least) is a
	// name for the resource
	ARNFromName(string) string
	// LateInitialize returns an AWS Resource after setting the late initialized
	// fields from the ReadOne call. This method will initialize the optional fields
	// which were not provided by the k8s user but were defaulted by the AWS service.
	// If there are no such fields to be initialized, the returned object is identical to
	// object passed in the parameter.
	// This method also adds/updates the ConditionTypeLateInitialized for the AWSResource.
	LateInitialize(context.Context, AWSResource) (AWSResource, error)
	// IsSynced returns true if a resource is synced.
	IsSynced(context.Context, AWSResource) (bool, error)
	// EnsureTags ensures that tags are present inside the AWSResource.
	// If the AWSResource does not have any existing resource tags, the 'tags'
	// field is initialized and the controller tags are added.
	// If the AWSResource has existing resource tags, then controller tags are
	// added to the existing resource tags without overriding them.
	// If the AWSResource does not support tags, only then the controller tags
	// will not be added to the AWSResource.
	EnsureTags(context.Context, AWSResource, ServiceControllerMetadata) error
	// FilterSystemTags ignores tags that are either injected by the controller
	// or by AWS. These tags have keys that start with "aws:" or "services.k8s.aws/"
	// and this function will remove them before adoption.
	// Eg. resources created with cloudformation have tags that cannot be
	//removed by an ACK controller
	FilterSystemTags(AWSResource)
}

// AWSResourceManagerFactory returns an AWSResourceManager that can be used to
// manage AWS resources for a particular AWS account
// TODO(jaypipes): Move AWSResourceManagerFactory into its own file
type AWSResourceManagerFactory interface {
	// ResourceDescriptor returns an AWSResourceDescriptor that can be used by
	// the upstream controller-runtime to introspect the CRs that the resource
	// manager will manage as well as produce Kubernetes runtime object
	// prototypes
	ResourceDescriptor() AWSResourceDescriptor
	// ManagerFor returns an AWSResourceManager that manages AWS resources on
	// behalf of a particular AWS account and in a specific AWS region
	ManagerFor(
		ackcfg.Config, // passed by-value to avoid mutation by consumers
		awscfg.Config,
		logr.Logger,
		*ackmetrics.Metrics,
		Reconciler,
		ackv1alpha1.AWSAccountID,
		ackv1alpha1.AWSRegion,
		ackv1alpha1.AWSResourceName,
	) (AWSResourceManager, error)
	// IsAdoptable returns true if the resource is able to be adopted
	IsAdoptable() bool
	// RequeueOnSuccessSeconds returns true if the resource should be requeued after specified seconds
	// Default is false which means resource will not be requeued after success.
	RequeueOnSuccessSeconds() int
}
