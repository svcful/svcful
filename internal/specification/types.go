package specification

type Type string

const (
	// TypeDatabase defines a persistent data store that exists as a separate service
	TypeDatabase Type = "database"
	// TypeExternal defines a service which is exposed publicly but directly
	TypeExternal Type = "external"
	// TypeInternal defines a service which is exposed internally to other services
	TypeInternal Type = "internal"
	// TypeIsolated defines a service which is not exposed at all
	TypeIsolated Type = "isolated"
	// TypeProvider defines a third party service provider
	TypeProvider Type = "provider"
	// TypePublic defines a service which is exposed publicly AND directly
	TypePublic Type = "public"
)

type DatabaseName string

const (
	DatabaseCassandra     DatabaseName = "cassandra"
	DatabaseCockroach     DatabaseName = "cockroach"
	DatabaseCouch         DatabaseName = "couch"
	DatabaseElasticSearch DatabaseName = "elasticsearch"
	DatabaseMaria         DatabaseName = "maria"
	DatabaseMongo         DatabaseName = "mongo"
	DatabaseMySQL         DatabaseName = "mysql"
	DatabasePostgres      DatabaseName = "postgres"
	DatabaseRedis         DatabaseName = "redis"
	DatabaseSqlite        DatabaseName = "sqlite"
)

type DataFormatName string

const (
	DataFormatJson     DataFormatName = "json"
	DataFormatProtobuf DataFormatName = "protobuf"
	DataFormatXml      DataFormatName = "xml"
)

const DefaultDataFormatName = DataFormatJson

type LifecycleType string

const (
	LifecycleTypeCronJob LifecycleType = "cronjob"
	LifecycleTypeJob     LifecycleType = "job"
	LifecycleTypeServer  LifecycleType = "server"
)

const DefaultLifecycleType = LifecycleTypeServer

type InterfaceType string

const (
	InterfaceAsync   InterfaceType = "async"
	InterfaceGraphQL InterfaceType = "graphql"
	InterfaceGrpc    InterfaceType = "grpc"
	InterfaceRest    InterfaceType = "rest"
	InterfaceSoap    InterfaceType = "soap"
	InterfaceRpc     InterfaceType = "rpc"
)

const DefaultInterfaceType = InterfaceRest

type TargetName string

const (
	// TargetDockerCompose indicates the entity will be deployed into a Docker Swarm
	TargetDockerCompose TargetName = "docker-compose"
	// TargetKubernetes indicates the entity will be deployed into a Kubernetes cluster
	TargetKubernetes TargetName = "kubernetes"
	// TargetNomad indicates the entity will be deployed into a Nomad cluster
	TargetNomad TargetName = "nomad"
	// TargetVirtualMachine indicates the entity will be deployed into a virtual machine
	TargetVirtualMachine TargetName = "vm"
)

const DefaultTargetname = TargetKubernetes
