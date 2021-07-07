package skelplugin

import (
	"github.com/containernetworking/cni/pkg/skel"
)

type Plugin struct{}

func (p *Plugin) Add(args *skel.CmdArgs) error {

	return nil
}

func (p *Plugin) Del(args *skel.CmdArgs) error {

	return nil
}

func (p *Plugin) Check(args *skel.CmdArgs) error {

	return nil
}

// New create the plugin
func New() *Plugin {
	return &Plugin{}
}
