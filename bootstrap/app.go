/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-07-11 13:11:23
 * @Description: app
 */

package bootstrap

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liziwei01/go-exec-upload/library/conf"
	"github.com/liziwei01/go-exec-upload/library/env"
)

var DefaultWriter io.Writer = os.Stdout

// Config app's conf
// default conf/app.toml
type Config struct {
	APPName string
	RunMode string

	Env env.AppEnv

	// conf of http service
	HTTPServer struct {
		Listen       string
		ReadTimeout  int // ms
		WriteTimeout int // ms
		IdleTimeout  int // ms
	}
}

// ParserAppConfig
func ParserAppConfig(filePath string) (*Config, error) {
	confPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}
	var c *Config
	if err := conf.Parse(confPath, &c); err != nil {
		return nil, err
	}
	// parse and set global conf
	rootDir := filepath.Dir(filepath.Dir(confPath))
	opt := env.Option{
		AppName: c.APPName,
		RunMode: c.RunMode,
		RootDir: rootDir,
		DataDir: filepath.Join(rootDir, "data"),
		LogDir:  filepath.Join(rootDir, "log"),
		ConfDir: filepath.Join(rootDir, filepath.Base(filepath.Dir(confPath))),
	}
	c.Env = env.New(opt)
	return c, nil
}

// App application
type App struct {
	ctx    context.Context
	config *Config
	server *http.Server
	close  func()
}

// NewApp establish an APP
func NewApp(ctx context.Context, c *Config, handler *gin.Engine) *App {
	ctxRet, cancel := context.WithCancel(ctx)
	app := &App{
		ctx:    ctxRet,
		config: c,
		close:  cancel,
	}
	app.initHTTPServer(handler)
	return app
}

func (app *App) initHTTPServer(handler *gin.Engine) {
	ser := &http.Server{
		Addr:         app.config.HTTPServer.Listen,
		Handler:      handler,
		ReadTimeout:  time.Millisecond * time.Duration(app.config.HTTPServer.ReadTimeout),
		WriteTimeout: time.Millisecond * time.Duration(app.config.HTTPServer.WriteTimeout),
		IdleTimeout:  time.Millisecond * time.Duration(app.config.HTTPServer.IdleTimeout),
	}
	app.server = ser
}

// Start start the service
func (app *App) Start() error {
	// start listening to port
	fmt.Fprintf(DefaultWriter, "[APP START] Listening and serving HTTP on %s\n", app.config.HTTPServer.Listen)
	// start distribute routers
	return app.server.ListenAndServe()
}
