// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/markturansky/sentry-operator/pkg/apis/sentryoperator/v1.SentryOperator":       schema_pkg_apis_sentryoperator_v1_SentryOperator(ref),
		"github.com/markturansky/sentry-operator/pkg/apis/sentryoperator/v1.SentryOperatorSpec":   schema_pkg_apis_sentryoperator_v1_SentryOperatorSpec(ref),
		"github.com/markturansky/sentry-operator/pkg/apis/sentryoperator/v1.SentryOperatorStatus": schema_pkg_apis_sentryoperator_v1_SentryOperatorStatus(ref),
	}
}

func schema_pkg_apis_sentryoperator_v1_SentryOperator(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SentryOperator is the Schema for the sentryoperators API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/markturansky/sentry-operator/pkg/apis/sentryoperator/v1.SentryOperatorSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/markturansky/sentry-operator/pkg/apis/sentryoperator/v1.SentryOperatorStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/markturansky/sentry-operator/pkg/apis/sentryoperator/v1.SentryOperatorSpec", "github.com/markturansky/sentry-operator/pkg/apis/sentryoperator/v1.SentryOperatorStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sentryoperator_v1_SentryOperatorSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SentryOperatorSpec defines the desired state of SentryOperator",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"desc": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "desc"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sentryoperator_v1_SentryOperatorStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SentryOperatorStatus defines the observed state of SentryOperator",
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"status"},
			},
		},
		Dependencies: []string{},
	}
}
