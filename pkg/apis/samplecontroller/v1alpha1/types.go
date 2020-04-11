/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

<<<<<<< HEAD
<<<<<<< HEAD
	Spec   FooSpec   `json:"spec"`
	Status FooStatus `json:"status"`

=======
	Spec     FooSpec     `json:"spec"`
>>>>>>> dd1d1f84c9c9b153b9771997521aa8395f56eac9
=======
	Spec     FooSpec     `json:"spec"`
>>>>>>> dd1d1f84c9c9b153b9771997521aa8395f56eac9
	Database FooDatabase `json:"database"`
}

// FooSpec is the spec for a Foo resource
type FooSpec struct {
	SecretName      string `json:"secretName"`
	TargetNamespace string `json:"targetNamespace"`
<<<<<<< HEAD
}

// FooStatus is the status for a Foo resource
type FooDatabase struct {
	DBName string `json:"dbName"`
	DBPass string `json:"dbPass"`
=======
>>>>>>> dd1d1f84c9c9b153b9771997521aa8395f56eac9
}

// FooStatus is the status for a Foo resource
type FooDatabase struct {
	DBName string `json:"dbName"`
	DBPass string `json:"dbPass"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Foo `json:"items"`
}
