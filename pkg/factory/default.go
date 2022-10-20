package factory

import (
	"github.com/victorsalaun/ohmyflux/pkg/cmdutil"
	"github.com/victorsalaun/ohmyflux/pkg/iostreams"
)

func New(appVersion string) *cmdutil.Factory {
	f := &cmdutil.Factory{
		ExecutableName: "ohmyflux",
	}

	f.IOStreams = ioStreams(f) // Depends on Config
	return f
}

func ioStreams(f *cmdutil.Factory) *iostreams.IOStreams {
	io := iostreams.System()
	return io
}
