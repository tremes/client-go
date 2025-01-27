// +build tools

// go mod won't pull in code that isn't depended upon, but we have some code we don't depend on from code that must be included
// for our build to work.
package dependencymagnet

import (
	_ "github.com/spf13/pflag"
	_ "github.com/openshift/api/openapi"
	_ "k8s.io/code-generator"
	_ "k8s.io/code-generator/cmd/applyconfiguration-gen"
)
