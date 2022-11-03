/*
Copyright 2021 Red Hat, Inc.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComponentSrcType describes the type of
// the src for the Component.
// Only one of the following location type may be specified.
// +kubebuilder:validation:Enum=Git;Image
type ComponentSrcType string

const (
	GitComponentSrcType   ComponentSrcType = "Git"
	ImageComponentSrcType ComponentSrcType = "Image"
)

type GitSource struct {
	// If importing from git, the repository to create the component from
	URL string `json:"url"`

	// Specify a branch/tag/commit id. If not specified, default is `main`/`master`.
	Revision string `json:"revision,omitempty"`

	// A relative path inside the git repo containing the component
	Context string `json:"context,omitempty"`

	// If specified, the devfile at the URL will be used for the component.
	DevfileURL string `json:"devfileUrl,omitempty"`

	// If specified, the dockerfile at the URL will be used for the component.
	DockerfileURL string `json:"dockerfileUrl,omitempty"`
}

// ComponentSource describes the Component source
type ComponentSource struct {
	ComponentSourceUnion `json:",inline"`
}

// +union
type ComponentSourceUnion struct {
	// Git Source for a Component
	GitSource *GitSource `json:"git,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ComponentSpec defines the desired state of Component
type ComponentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	// +kubebuilder:validation:MaxLength=63
	// ComponentName is name of the component to be added to the HASApplication
	ComponentName string `json:"componentName"`

	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	// Application to add the component to
	Application string `json:"application"`

	// Secret describes the name of a Kubernetes secret containing either:
	// 1. A Personal Access Token to access the Component's git repostiory (if using a Git-source component) or
	// 2. An Image Pull Secret to access the Component's container image (if using an Image-source component).
	Secret string `json:"secret,omitempty"`

	// Source describes the Component source
	Source ComponentSource `json:"source,omitempty"`

	// Compute Resources required by this component
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// The number of replicas to deploy the component with
	Replicas int `json:"replicas,omitempty"`

	// The port to expose the component over
	TargetPort int `json:"targetPort,omitempty"`

	// The route to expose the component with
	Route string `json:"route,omitempty"`

	// An array of environment variables to add to the component
	Env []corev1.EnvVar `json:"env,omitempty"`

	// The container image to build or create the component from
	ContainerImage string `json:"containerImage,omitempty"`

	// Whether to bypass the generation of GitOps resources for the Component. Defaults to false.
	SkipGitOpsResourceGeneration bool `json:"skipGitOpsResourceGeneration,omitempty"`
}

// ComponentStatus contains the observed state of Component
type ComponentStatus struct {

	// Name is the name of the component.
	Name string `json:"name"`

	// GitOpsRepository contains the Git URL, path, branch, and most recent commit id for the component
	GitOpsRepository BindingComponentGitOpsRepository `json:"gitopsRepository"`
}

// GitOpsStatus contains GitOps repository-specific status for the component
type GitOpsStatus struct {
	// RepositoryURL is the gitops repository URL for the component
	RepositoryURL string `json:"repositoryURL,omitempty"`

	// Branch is the branch used for the gitops repository
	Branch string `json:"branch,omitempty"`

	// Context is the path within the gitops repository used for the gitops resources
	Context string `json:"context,omitempty"`

	// ResourceGenerationSkipped is whether or not GitOps resource generation was skipped for the component
	ResourceGenerationSkipped bool `json:"resourceGenerationSkipped,omitempty"`

	// CommitID is the most recent commit ID in the GitOps repository for this component
	CommitID string `json:"commitID,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Component is the Schema for the components API
// +kubebuilder:resource:path=components,shortName=hascmp;hc;comp
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComponentSpec   `json:"spec"`
	Status ComponentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ComponentList contains a list of Component
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Component `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Component{}, &ComponentList{})
}
