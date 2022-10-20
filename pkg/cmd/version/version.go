package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/victorsalaun/ohmyflux/pkg/cmdutil"
	"strings"
)

func NewCmdVersion(f *cmdutil.Factory, buildVersion, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: false,
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprint(f.IOStreams.Out, cmd.Root().Annotations["versionInfo"])
		},
	}

	return cmd
}

func Format(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")
	return fmt.Sprintf("ohmyflux %s (%s)\n", version, buildDate)
}
