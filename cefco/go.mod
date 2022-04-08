module github.com/inspursoft/cefco

go 1.15

require (
	github.com/nats-io/jsm.go v0.0.20
	github.com/nats-io/nats.go v1.10.1-0.20201111151633-9e1f4a0d80d8
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
	k8s.io/code-generator v0.21.0
	k8s.io/klog/v2 v2.8.0
)

// Taken from kubernetes/sample-controller@442b3218b3d3eecdd9e55bffcb2c6b135f3084b7
replace (
	k8s.io/api => k8s.io/api v0.0.0-20210411031832-6eed676bc189
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20210411031630-d8cdd62f18c3
	k8s.io/client-go => k8s.io/client-go v0.0.0-20210411032117-0bb6464b1348
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20210411031433-bdc239664504
)
