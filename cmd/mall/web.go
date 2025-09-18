package mall

import (
	"fmt"
	"mall/internal/core"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use: "web",
	Run: startWebServer,
}

var config string

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
}

func startWebServer(cmd *cobra.Command, args []string) {
	if err := core.InitConfig(config); err != nil {
		fmt.Printf("Error initializing config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Starting web server on %s...\n", core.GlobalConfig.Server.Addr)
	server := &http.Server{
		Addr: core.GlobalConfig.Server.Addr,
	}
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(r.URL.Path + " > ping response"))
	})

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
