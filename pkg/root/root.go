package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
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
	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		rootHelpFunc(factory, c, args)
	})
	cmd.SetUsageFunc(func(c *cobra.Command) error {
		return rootUsageFunc(factory.IOStreams.ErrOut, c)
	})
	cmd.SetFlagErrorFunc(rootFlagErrorFunc)

	// Child commands
	cmd.AddCommand(versionCmd.NewCmdVersion(factory, buildVersion, buildDate))

	return cmd
}
