package cmd

import (
	"fmt"
	_ "strings"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	trafficsplitExample = `
	# create trafficsplit to service in the specified namespace
	%[1]s ts reivew --split review-v1=1 --split reivew-v2=500m -n namespace

	# list trafficsplit to service in the specified namespace
	%[1]s ts --list -n namespace
`
)

// TrafficSplitOptions provides information required to trafficsplit.
type TrafficSplitOptions struct {
	configFlags *genericclioptions.ConfigFlags

	setTrafficsplit []string

	listTrafficsplit bool

	genericclioptions.IOStreams
}

// NewTrafficSplitOptions provides an instance of TrafficSplitOptions with default values
func NewTrafficSplitOptions(streams genericclioptions.IOStreams) *TrafficSplitOptions {
	return &TrafficSplitOptions{
		configFlags: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

// NewTrafficSplit provides a cobra command wrapping TrafficSplitOptions
func NewCmdTrafficSplit(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewTrafficSplitOptions(streams)

	cmd := &cobra.Command{
		Use:          "ts service [--split service=1] [-n namespace]",
		Short:        "Create or list trafficsplit",
		Example:      fmt.Sprintf(trafficsplitExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&o.listTrafficsplit, "list", o.listTrafficsplit, "if true, print the list of all trafficsplit in specified namespace.")
	cmd.Flags().StringSliceVar(&o.setTrafficsplit, "split", o.setTrafficsplit, "Set trafficsplit to service in specified namespace.")
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

func (o *TrafficSplitOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (o *TrafficSplitOptions) Validate() error {
	return nil
}

func (o *TrafficSplitOptions) Run() error {
	return nil
}
