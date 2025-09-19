package mall

import (
	"fmt"
	"mall/api/router"
	"mall/internal/core"
	"mall/internal/log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

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

	if err := InitLogger(); err != nil {
		fmt.Printf("Error initializing log: %v\n", err)
		os.Exit(1)
	}

	Logger.Info("Starting web server", zap.String("address", core.GlobalConfig.Server.Addr))
	fmt.Printf("Starting web server on %s...\n", core.GlobalConfig.Server.Addr)

	engine := gin.New()
	router.RegisterRouter(engine)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		Logger.Debug("Ping request received", zap.String("path", r.URL.Path))
		_, _ = w.Write([]byte(r.URL.Path + " > ping response"))
	})

	server := initServer(engine)
	Logger.Info("Server starting to listen", zap.String("address", server.Addr))
	if err := server.ListenAndServe(); err != nil {
		Logger.Error("Error starting server", zap.Error(err))
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
	logWriter, err := log.GetWriter(core.GlobalConfig.Logger.LogFile)
	if err != nil {
		return err
	}

	c := zapcore.NewCore(log.GetEncoder(), zapcore.AddSync(logWriter), log.LogLevel(core.GlobalConfig.Logger.LogLevel))

	Logger = zap.New(c, zap.AddCaller())

	// 设置全局log，这样其他包也能使用
	zap.ReplaceGlobals(Logger)

	Logger.Info("Logger initialized successfully")
	return nil
}
