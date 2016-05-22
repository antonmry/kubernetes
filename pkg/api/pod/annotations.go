/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package pod

const (
	// The prefix of an annotation denoting the last port of the range. The annotation ends
	// with the name of the port and the value is the last port of the range.
	// @antonmry: move to annotations.go?
	PortRangeEndAnnotation = "portrange.alpha.kubernetes.io/port-end-"
)

