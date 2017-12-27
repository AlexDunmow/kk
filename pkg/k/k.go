package k

import (
	"encoding/json"
	"errors"

	"github.com/urfave/cli"
)

var PodListHeaders = []string{"Namespace", "Name", "Restarts", "Age", "Ready", "Status"}
var NamespaceListHeaders = []string{"Name"}

type Kind string

func (k Kind) String() string {
	return string(k)
}

const (
	KindPods       Kind = "Pods"
	KindNamespaces Kind = "Namespaces"
	KindResources  Kind = "Resources"
)

type Screen string

const (
	ScreenTable Screen = "Table"
	ScreenModal Screen = "Modal"
	ScreenState Screen = "State"
	ScreenDebug Screen = "Debug"
)

func (s Screen) String() string {
	return string(s)
}

type State struct {
	UI       *UIReducer
	Entities *EntitiesReducer
}

func (s *State) JSONString() (string, error) {
	b, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ParsedFlags will contain the config for the app
type ParsedFlags struct {
	KubeConfigPath  string
	RefreshInterval int
	AutoRefresh     bool
	DebugFile       string
	PROD            bool
	DEBUG           bool
	TEST            bool
}

// Parse will parse the flags into a struct
func (flags *ParsedFlags) Parse(c *cli.Context) error {
	flags.KubeConfigPath = c.String("kubeconfig-path")
	flags.RefreshInterval = c.Int("refresh-interval")
	flags.PROD = c.Bool("production")
	flags.DEBUG = c.Bool("debug")
	flags.TEST = c.Bool("test")
	flags.DebugFile = c.String("debug-file")
	flags.AutoRefresh = c.Bool("auto-refresh")

	if flags.KubeConfigPath == "" {
		return errors.New("Error parsing flags")
	}
	return nil
}
