package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/config"
	"github.com/liaozzzzzz/code-push-server/internal/database"
	"github.com/liaozzzzzz/code-push-server/internal/router"
)

// Server HTTP服务器
type Server struct {
	httpServer *http.Server
	router     *gin.Engine
}

// NewServer 创建新的服务器实例
func NewServer() *Server {
	return &Server{}
}

// Start 启动服务器
func (s *Server) Start() error {
	// 初始化数据库
	if err := database.Initialize(); err != nil {
		return fmt.Errorf("初始化数据库失败: %w", err)
	}

	// 设置Gin模式
	if !config.C.General.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// 设置路由
	s.router = router.SetupRouter()

	// 创建HTTP服务器
	s.httpServer = &http.Server{
		Addr:           config.C.General.HTTP.Addr,
		Handler:        s.router,
		ReadTimeout:    time.Duration(config.C.General.HTTP.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.C.General.HTTP.WriteTimeout) * time.Second,
		IdleTimeout:    time.Duration(config.C.General.HTTP.IdleTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// 启动服务器
	go func() {
		log.Printf("🚀 服务器启动在 %s", config.C.General.HTTP.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号
	s.waitForShutdown()

	return nil
}

// waitForShutdown 等待关闭信号
func (s *Server) waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 正在关闭服务器...")

	// 创建超时上下文
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.C.General.HTTP.ShutdownTimeout)*time.Second)
	defer cancel()

	// 关闭HTTP服务器
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("服务器关闭失败: %v", err)
	}

	// 关闭数据库连接
	if err := database.Close(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	}

	log.Println("✅ 服务器已关闭")
}

// Stop 停止服务器
func (s *Server) Stop() error {
	if s.httpServer != nil {
		return s.httpServer.Close()
	}
	return nil
}
