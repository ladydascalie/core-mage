package targets

import (
	"os"
	"path"
)

const defaultGoProxy = "https://proxy.golang.org"

var (
	// ProjectType configures type of the project.
	ProjectType = "service"

	// ProjectName configures the name of the project.
	ProjectName = "service"

	// Environment describes the environment variables that should be sent with the target.
	Environment = CMDEnv{
		"GOPATH":    os.Getenv("GOPATH"),
		"GOPROXY":   os.Getenv("GOPROXY"),
		"GOMODPATH": os.Getenv("GOMODPATH"),
		"PWD":       os.Getenv("PWD"),
	}
)

// CMDEnv is used to wrap the command environment with convenience methods.
type CMDEnv map[string]string

// GoModPath derives the go module path from the environment.
func (e CMDEnv) GoModPath() string {
	var mod string
	if p := e["GOMODPATH"]; p != "" {
		mod = p
	}
	if p := e["GOPATH"]; p != "" && mod == "" {
		mod = path.Join(p, "mod")
	}
	if mod == "" {
		mod = "/go/pkg/mod"
	}
	return mod
}

// GoProxy will return the configured go proxy or provide a default.
func (e CMDEnv) GoProxy() string {
	if proxy := e["GOPROXY"]; proxy != "" {
		return proxy
	}
	return defaultGoProxy
}
