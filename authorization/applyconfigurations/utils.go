// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/openshift/api/authorization/v1"
	authorizationv1 "github.com/openshift/client-go/authorization/applyconfigurations/authorization/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=authorization.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("Action"):
		return &authorizationv1.ActionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterRole"):
		return &authorizationv1.ClusterRoleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterRoleBinding"):
		return &authorizationv1.ClusterRoleBindingApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GroupRestriction"):
		return &authorizationv1.GroupRestrictionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LocalResourceAccessReview"):
		return &authorizationv1.LocalResourceAccessReviewApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LocalSubjectAccessReview"):
		return &authorizationv1.LocalSubjectAccessReviewApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PolicyRule"):
		return &authorizationv1.PolicyRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ResourceAccessReview"):
		return &authorizationv1.ResourceAccessReviewApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Role"):
		return &authorizationv1.RoleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RoleBinding"):
		return &authorizationv1.RoleBindingApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RoleBindingRestriction"):
		return &authorizationv1.RoleBindingRestrictionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RoleBindingRestrictionSpec"):
		return &authorizationv1.RoleBindingRestrictionSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceAccountReference"):
		return &authorizationv1.ServiceAccountReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceAccountRestriction"):
		return &authorizationv1.ServiceAccountRestrictionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SubjectAccessReview"):
		return &authorizationv1.SubjectAccessReviewApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("UserRestriction"):
		return &authorizationv1.UserRestrictionApplyConfiguration{}

	}
	return nil
}