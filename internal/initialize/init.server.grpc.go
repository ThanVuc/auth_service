package initialize

import (
	"auth_service/global"
	"auth_service/internal/controller"
	v1Auth "auth_service/internal/grpc/auth.v1"
	v1Permission "auth_service/internal/grpc/permission.v1"
	v1Role "auth_service/internal/grpc/role.v1"
	v1Token "auth_service/internal/grpc/token.v1"
	"auth_service/internal/wire"
	"auth_service/pkg/loggers"
	"auth_service/pkg/settings"
	"context"
	"fmt"
	"net"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AuthServer struct {
	authServiceServer       *controller.AuthController
	permissionServiceServer *controller.PermissionController
	roleServiceServer       *controller.RoleController
	tokenServiceServer      *controller.TokenController
	logger                  *loggers.LoggerZap
	config                  *settings.Server
}

func NewAuthService() *AuthServer {
	return &AuthServer{
		authServiceServer:       wire.InjectAuthWire(),
		permissionServiceServer: wire.InjectPermissionWire(),
		roleServiceServer:       wire.InjectRoleWire(),
		tokenServiceServer:      wire.InjectTokenWire(),
		logger:                  global.Logger,
		config:                  &global.Config.Server,
	}
}

func (as *AuthServer) RunServers(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go as.runServiceServer(ctx, wg)
}

func (as *AuthServer) runServiceServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := as.createListener()
	if err != nil {
		as.logger.ErrorString("Failed to create listener",
			zap.Error(err),
		)
		return
	}

	// Create a new gRPC server instance
	server := as.createServer()

	// Gracefully handle server shutdown
	go as.gracefullyShutdownServer(ctx, server)

	// Server listening on the specified port
	as.serverListening(server, lis)
}

func (as *AuthServer) gracefullyShutdownServer(ctx context.Context, server *grpc.Server) {
	<-ctx.Done()
	as.logger.InfoString("gRPC server is shutting down...")
	server.GracefulStop()
	as.logger.InfoString("gRPC server stopped gracefully!")
}

func (as *AuthServer) serverListening(server *grpc.Server, lis net.Listener) {
	as.logger.InfoString(fmt.Sprintf("gRPC server listening on %s:%d", as.config.Host, lis.Addr().(*net.TCPAddr).Port))
	if err := server.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			as.logger.InfoString("gRPC server exited normally")
		} else {
			as.logger.ErrorString("Failed to serve gRPC server",
				zap.Error(err),
			)
		}
	}
}

// create server factory
func (as *AuthServer) createServer() *grpc.Server {
	server := grpc.NewServer()
	authServer := wire.InjectAuthWire()
	permissionServer := wire.InjectPermissionWire()
	roleServer := wire.InjectRoleWire()
	tokenServer := wire.InjectTokenWire()

	v1Auth.RegisterAuthServiceServer(server, authServer)
	v1Permission.RegisterPermissionServiceServer(server, permissionServer)
	v1Role.RegisterRoleServiceServer(server, roleServer)
	v1Token.RegisterTokenServiceServer(server, tokenServer)

	return server
}

func (as *AuthServer) createListener() (net.Listener, error) {
	err := error(nil)
	lis := net.Listener(nil)
	lis, err = net.Listen("tcp", fmt.Sprintf("%s:%d", as.config.Host, as.config.AuthPort))
	if err != nil {
		as.logger.ErrorString("Failed to listen: %v", zap.String("error", err.Error()))
		return nil, err
	}

	return lis, nil
}
