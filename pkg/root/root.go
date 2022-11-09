package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	treeCmd "github.com/victorsalaun/ohmyflux/pkg/cmd/tree"
	versionCmd "github.com/victorsalaun/ohmyflux/pkg/cmd/version"
	"github.com/victorsalaun/ohmyflux/pkg/cmdutil"
)

func NewCmdRoot(factory *cmdutil.Factory, buildVersion, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ohmyflux <command> <subcommand> [flags]",
		Short: "OhMyFlux CLI",

		SilenceErrors: true,
		SilenceUsage:  true,
		Example: heredoc.Doc(`
			$ ohmyflux version
		`),
		Annotations: map[string]string{
			"versionInfo": versionCmd.Format(buildVersion, buildDate),
		},
	}

	cmd.Flags().Bool("version", false, "Show ohmyflux version")

	// Child commands
	cmd.AddCommand(treeCmd.NewCmdTree(factory, nil))
	cmd.AddCommand(versionCmd.NewCmdVersion(factory, buildVersion, buildDate))

	return cmd
}
