package mall

import (
	"fmt"
	"mall/api/router"
	"mall/internal/core"
	"mall/internal/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Generate images for web",
	Long:  `Generate images suitable for web usage using AI.`,
	Run:   StartwebServer,
}

var defaultConfig = "configs/config.yaml"
var config string

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
}
func initServer(hander http.Handler) *http.Server {
	server := &http.Server{
		Handler:      hander,
		Addr:         core.GlobalConfig.Server.Addr,
		ReadTimeout:  core.GlobalConfig.Server.ReadTimeout,
		WriteTimeout: core.GlobalConfig.Server.WriteTimeout,
		IdleTimeout:  core.GlobalConfig.Server.IdleTimeout,
	}
	return server
}

func StartwebServer(cmd *cobra.Command, args []string) {
	if config == "" {
		config = defaultConfig
	}
	if err := core.InitConfig(config); err != nil {
		fmt.Printf("Failed to initialize config: %vn", err)
		return
	}
	fmt.Println("Config loaded")

	if err := logger.InitLogger(); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}
	fmt.Println("Logger initialized")

	engine := gin.New()

	// Serve static files from web/static
	engine.Static("/static", "./web/static")

	// Load HTML templates from web/template
	// engine.LoadHTMLGlob("web/template/*")  // Removed because template directory is empty

	// Serve SPA from web/app
	engine.Static("/app", "./web/app")
	engine.GET("/", func(c *gin.Context) {
		c.File("./web/app/index.html")
	})
	engine.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") || strings.HasPrefix(c.Request.URL.Path, "/static") {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.File("./web/app/index.html")
	})

	router.RegisterRouter(engine)

	server := initServer(engine)

	fmt.Printf("Starting server on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
