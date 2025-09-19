package mall

import (
	"fmt"
	"mall/api/router"
	"mall/internal/core"
	"mall/logs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	engine := gin.New()
	router.RegisterRouter(engine)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(r.URL.Path + " > ping response"))
	})

	server := initServer(engine)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func initServer(handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:        core.GlobalConfig.Server.Addr,
		IdleTimeout: core.GlobalConfig.Server.IdleTimeout,
		Handler:     handler,
	}
	return server
}

func InitLogger() error {
	logWriter, _ := logs.GetWriter(core.GlobalConfig.Logger.LogFile)

	c := zapcore.NewCore(logs.GetEncoder(), zapcore.AddSync(logWriter), logs.LogLevel(core.GlobalConfig.Logger.LogLevel))

	log := zap.New(c, zap.AddCaller())
	_ = log.Sugar()

	return nil
}
