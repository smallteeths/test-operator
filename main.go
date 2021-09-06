//go:generate go run pkg/codegen/cleanup/main.go
//go:generate go run pkg/codegen/main.go

package main

var (
	masterURL      string
	kubeconfigFile string
	leaderElect    bool
	lockName       string
	lockNamespace  string
	id             string
	qps            float64
	burst          int
)

func init() {

}

func main() {

}
