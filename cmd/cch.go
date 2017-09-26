package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/techninja1008/chainchain/p2p"
)

var CCHCmd = &cobra.Command{
	Use:   "cch",
	Short: "Launch CCH",
	Run: func(cmd *cobra.Command, args []string) {
		stack := p2p.BootNode(httpAPIs, httpPort, dataDir)

		go func() {
			sigc := make(chan os.Signal, 1)
			signal.Notify(sigc, syscall.SIGTERM)
			defer signal.Stop(sigc)
			<-sigc
			log.Println("Got sigterm, shutting down...")
			stack.Stop()
		}()

		stack.Wait()
	},
}

var httpAPIs string
var dataDir string
var httpPort int

func init() {
	CCHCmd.Flags().StringVar(&httpAPIs, "http-api", "cch", "Comma-seperated list of HTTP RPC APIs to expose")
	CCHCmd.Flags().IntVarP(&httpPort, "http-port", "p", 8228, "Port to use for HTTP RPC APIs")
	CCHCmd.Flags().StringVarP(&dataDir, "data-dir", "d", "./cchdata", "Directory to store all data in")
}
