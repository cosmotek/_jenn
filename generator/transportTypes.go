package generator

type TransportType string

const (
	GRPC    TransportType = "grpc"
	GRPCWeb TransportType = "grpc_web"
	REST    TransportType = "rest"
	GraphQL TransportType = "graphql"
)
