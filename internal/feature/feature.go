package feature

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
)

const (
	AutomatedDeployments featuregate.Feature = "AutomatedDeployments"
)

func init() {
	runtime.Must(MutableGates.Add(defaultEGFeatureGates))
}

// defaultEGFeatureGates consists of all known EG feature keys.
// To add a new feature, define a key for it above and add it here.
var defaultEGFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	// Every feature should be initiated here:
	AutomatedDeployments: {Default: false, PreRelease: featuregate.Alpha},
}
