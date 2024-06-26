package app

import (
	"context"
	"log"

	"github.com/kirillmc/platform_common/pkg/closer"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/platform_common/pkg/db/pg"
	"github.com/kirillmc/trainings-auth/internal/api/access"
	"github.com/kirillmc/trainings-auth/internal/api/auth"
	"github.com/kirillmc/trainings-auth/internal/api/user"
	"github.com/kirillmc/trainings-auth/internal/client/rpc"
	"github.com/kirillmc/trainings-auth/internal/config"
	"github.com/kirillmc/trainings-auth/internal/config/env"
	"github.com/kirillmc/trainings-auth/internal/interceptor"
	"github.com/kirillmc/trainings-auth/internal/repository"
	userRepo "github.com/kirillmc/trainings-auth/internal/repository/user"
	"github.com/kirillmc/trainings-auth/internal/service"
	accessService "github.com/kirillmc/trainings-auth/internal/service/access"
	authService "github.com/kirillmc/trainings-auth/internal/service/auth"
	userService "github.com/kirillmc/trainings-auth/internal/service/user"
)

// содержит все зависимости, необходимые в рамках приложения
type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	accessConfig  config.AccessConfig

	dbClient          db.Client
	accessClient      rpc.AccessClient
	interceptorClient *interceptor.Interceptor

	userRepository   repository.UserRepository
	accessRepository repository.AccessRepository
	authRepository   repository.AuthRepository

	userService   service.UserService
	accessService service.AccessService
	authService   service.AuthService

	userImpl   *user.Implementation
	authImpl   *auth.Implementation
	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// если GetUser - в GO GetUser НЕ УКАЗЫВАЮТ: НЕ GetPGConfig, A PGConfig
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		pgConfig, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err) // делаем Log.Fatalf, чтобы не обрабатывать ошибку в другом месте
			// + инициализация происходит при старте приложения, поэтому если ошибка - можно и сервер уронить
			// можно кинуть panic()
		}

		s.pgConfig = pgConfig
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpcConfig, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		s.grpcConfig = grpcConfig
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		httpConfig, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %v", err)
		}

		s.httpConfig = httpConfig
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		swaggerConfig, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %v", err)
		}

		s.swaggerConfig = swaggerConfig
	}

	return s.swaggerConfig
}

func (s *serviceProvider) AccessConfig() config.AccessConfig {
	if s.accessConfig == nil {
		accessConfig, err := env.NewAccessConfig()
		if err != nil {
			log.Fatalf("failed to get access configs: %v", err)
		}

		s.accessConfig = accessConfig
	}

	return s.accessConfig
}

//func (s *serviceProvider) AccessClient() rpc.AccessClient {
//	if s.accessClient == nil {
//		cfg := s.AccessConfig()
//
//		conn, err := grpc.Dial(
//			cfg.Address(),
//			grpc.WithTransportCredentials(insecure.NewCredentials()),
//		)
//		if err != nil {
//			log.Fatalf("failed to connect to access: %v", err)
//		}
//
//		s.accessClient = accessClient.NewAccessClient(descAccess.UnimplementedAccessV1Server{})
//	}
//
//	return s.accessClient
//}

func (s *serviceProvider) InterceptorClient(ctx context.Context) *interceptor.Interceptor {
	log.Printf("s.interceptorClient == nil: %v", s.interceptorClient == nil)
	if s.interceptorClient == nil {
		s.interceptorClient = &interceptor.Interceptor{
			Client: s.AccessService(ctx),
		}
		log.Printf("s.interceptorClient == nil: %v", s.interceptorClient == nil)
	}
	log.Printf("s.interceptorClient == nil: %v", s.interceptorClient == nil)
	return s.interceptorClient
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(cl.Close)
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = userRepo.NewAccessRepository(s.DBClient(ctx))
	}

	return s.accessRepository
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = userRepo.NewAuthRepository(s.DBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(s.AccessRepository(ctx))
	}

	return s.accessService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AccessImplementation(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}

func (s *serviceProvider) AuthImplementation(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}
