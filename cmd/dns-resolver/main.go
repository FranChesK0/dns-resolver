package main

import (
	"fmt"
	"os"

	"github.com/FranChesK0/dns-resolver/internal/cli"
	"github.com/FranChesK0/dns-resolver/internal/packet"
	"github.com/FranChesK0/dns-resolver/internal/resolver"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dns-resolver <domain> [<domain> ...]",
	Short: cli.BoldText.Render("Simple DNS Resolver"),
	Run:   run,
}

func init() {
	rootCmd.Flags().StringP("name-server", "s", "77.240.157.30", "Choose your own name server for resolving.")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	var err error
	if len(args) < 1 {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	nameServer, err := cmd.Flags().GetString("name-server")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rsv := resolver.NewResolver(nameServer)
	for _, domain := range args {
		rsv.Resolve(domain, packet.TYPE_A)
	}
}
