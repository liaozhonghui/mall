package mall

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use: "web",
	Run: startWebServer,
}

func init() {
	rootCmd.AddCommand(webCmd)
}

func startWebServer(cmd *cobra.Command, args []string) {
	fmt.Println("Starting web server...")
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(r.URL.Path + " > ping response"))
	})

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
