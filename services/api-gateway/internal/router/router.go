package router

import (
	"net/http"

	"cec-project/delivery-app/api-gateway/internal/grpc_client"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewRouter(clients *grpc_client.Clients) *http.ServeMux {
	mux := http.NewServeMux()
	gwMux := runtime.NewServeMux()

	// Register handlers for each service
	registerAuthHandlers(gwMux, clients)
	registerProductHandlers(gwMux, clients)
	registerOrderHandlers(gwMux, clients)
	registerDeliveryHandlers(gwMux, clients)

	mux.Handle("/", gwMux)
	return mux
}

func registerAuthHandlers(mux *runtime.ServeMux, clients *grpc_client.Clients) {
	// Auto-generated registration code from protobuf
	// authpb.RegisterAuthServiceHandlerServer(ctx, mux, server)
	// For custom handling:
	mux.HandlePath("POST", "/api/v1/auth/login", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		// Handle login request
	})

	mux.HandlePath("POST", "/api/v1/auth/register", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		// Handle registration
	})
}
