//go:build wireinject

package adapter

import "github.com/google/wire"

// INFRASTRUCTURE
var infraProviderSet = wire.NewSet(
	provideTracer,
	provideGormDB,
	provideExternalApiClient,
	provideFsManager,
	provideKeycloak,
	provideGrpcClient,
	wire.Struct(new(InfrastructureParams), "*"),
)

// DATASOURCE
var datasourceProviderSet = wire.NewSet(
	provideExampleDataSource,
	provideSqlDataSource,
	wire.Struct(new(DatasourcesParams), "*"),
)

// REPOSITORY
var repositoryProviderSet = wire.NewSet(
	provideExampleRepository,
	provideUserRepository,

	wire.Struct(new(RepositoryParams), "*"),
)

// USE_CASE
var useCaseSet = wire.NewSet(
	provideExampleUseCase,
	provideUserUseCase,

	wire.Struct(new(UseCaseParams), "*"),
)

// TRANSPORT
var transportSet = wire.NewSet(
	provideREST,
	provideGRPC,
	provideCRON,
	provideKafka,
	wire.Struct(new(TransportParams), "*"),
)

// APP COMPONENT
var adapterSet = wire.NewSet(
	wire.Struct(new(Adapters), "*"),
)

// APP
var appSet = wire.NewSet(
	provideInfrastructures,
	provideDatasources,
	provideRepositories,
	provideUseCases,
	provideTransports,
)