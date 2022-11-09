package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/victorsalaun/ohmyflux/internal/build"
	"github.com/victorsalaun/ohmyflux/internal/kube"
	"github.com/victorsalaun/ohmyflux/pkg/cmdutil"
	"github.com/victorsalaun/ohmyflux/pkg/factory"
	"github.com/victorsalaun/ohmyflux/pkg/root"
	"io"
	"k8s.io/client-go/rest"
	"net"
	"os"
	"strings"
)

type exitCode int

const (
	exitOk    exitCode = 0
	exitError exitCode = 1
)

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	config, _ := kube.GetKubeConfig()
	config.WarningHandler = rest.NoWarnings{}
	kubernetesClient := kube.GetKubernetesClient(config)
	kubernetesDynamicClient := kube.GetKubernetesDynamicClient(config)

	buildVersion := build.Version
	buildDate := build.Date

	cmdFactory := factory.New(kubernetesClient, kubernetesDynamicClient)
	stderr := cmdFactory.IOStreams.ErrOut

	rootCmd := root.NewCmdRoot(cmdFactory, buildVersion, buildDate)

	if cmd, err := rootCmd.ExecuteC(); err != nil {
		printError(stderr, err, cmd, false)
		return exitError
	}
	return exitOk
}

func printError(out io.Writer, err error, cmd *cobra.Command, debug bool) {
	var dnsError *net.DNSError
	if errors.As(err, &dnsError) {
		_, _ = fmt.Fprintf(out, "error connecting to %s\n", dnsError.Name)
		if debug {
			_, _ = fmt.Fprintln(out, dnsError)
		}
		_, _ = fmt.Fprintln(out, "check your internet connection")
		return
	}

	_, _ = fmt.Fprintln(out, err)

	var flagError *cmdutil.FlagError
	if errors.As(err, &flagError) || strings.HasPrefix(err.Error(), "unknown command ") {
		if !strings.HasSuffix(err.Error(), "\n") {
			_, _ = fmt.Fprintln(out)
		}
		_, _ = fmt.Fprintln(out, cmd.UsageString())
	}
}
