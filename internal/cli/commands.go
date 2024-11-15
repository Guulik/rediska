package cli

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "redis-cli",
	Short: "runs redis-cli in interactive mode",
}

var pingCmd = &cobra.Command{
	Use:   "PING",
	Short: "check server availability",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.sendCommand("PING")
	},
}

var getCmd = &cobra.Command{
	Use:   "GET key",
	Short: "Get value by provided key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.sendCommand("GET", args...)
	},
}

var setCmd = &cobra.Command{
	Use: "SET key value [NX | XX] [GET] [EX seconds | PX milliseconds |" +
		" EXAT unix-time-seconds | PXAT unix-time-milliseconds | KEEPTTL]",
	Short: "Set key to hold the string value. If key already holds a value, it is overwritten, regardless of its type." +
		" Any previous time to live associated with the key is discarded on successful SET operation.",
	RunE: func(cmd *cobra.Command, args []string) error {
		//TODO: send args as functionalOpts
		return nil
	},
}
