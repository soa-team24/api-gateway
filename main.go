package main

import (
	"api-gateway/config"
	"api-gateway/proto/blog"
	"api-gateway/proto/follower"
	"api-gateway/proto/stakeholder"
	"api-gateway/proto/tour"
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/spf13/viper"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}
func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	log.Println("Blog starting")
	blogEndpoint := fmt.Sprintf("%s:%s", "blog_service", server.config.BlogServiceAddress)
	err := blog.RegisterBlogServiceHandlerFromEndpoint(context.TODO(), server.mux, blogEndpoint, opts)
	if err != nil {
		panic(err)
	}
	log.Println("Blog started")

	log.Println("Follower starting")
	followerEndpoint := fmt.Sprintf("%s:%s", "follower_service", server.config.FollowerServiceAddress)
	errFollower := follower.RegisterFollowServiceHandlerFromEndpoint(context.TODO(), server.mux, followerEndpoint, opts)
	if errFollower != nil {
		panic(errFollower)
	}
	log.Println("Follower started")

	log.Println("Stakeholder starting")
	stakeholderEndpoint := fmt.Sprintf("%s:%s", "stakeholder_service", server.config.StakeholderServiceAddress)
	errStakeholder := stakeholder.RegisterStakeholderServiceHandlerFromEndpoint(context.TODO(), server.mux, stakeholderEndpoint, opts)
	if errStakeholder != nil {
		panic(errStakeholder)
	}
	log.Println("Stakeholder started")

	log.Println("Follower started")

	log.Println("Tour starting")
	tourEndpoint := fmt.Sprintf("%s:%s", "tour_service", server.config.TourServiceAddress)
	errTour := tour.RegisterTourServiceHandlerFromEndpoint(context.TODO(), server.mux, tourEndpoint, opts)
	if errTour != nil {
		panic(errTour)
	}
	log.Println("Tour started")

}

/*
func (server *Server) initHandlers() {

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		services := []struct {
			address   string
			register  func(context.Context, *runtime.ServeMux, interface{}) error
			service   string
			newClient func(*grpc.ClientConn) interface{}
		}{
			{
				address: server.config.BlogServiceAddress,
				register: func(ctx context.Context, mux *runtime.ServeMux, client interface{}) error {
					return blog.RegisterBlogServiceHandlerClient(ctx, mux, client.(blog.BlogServiceClient))
				},
				service: "blog_service",
				newClient: func(conn *grpc.ClientConn) interface{} {
					return blog.NewBlogServiceClient(conn)
				},
			},
			{
				address: server.config.FollowerServiceAddress,
				register: func(ctx context.Context, mux *runtime.ServeMux, client interface{}) error {
					return follower.RegisterFollowServiceHandlerClient(ctx, mux, client.(follower.FollowServiceClient))
				},
				service: "follower_service",
				newClient: func(conn *grpc.ClientConn) interface{} {
					return follower.NewFollowServiceClient(conn)
				},
			},
			{
				address: server.config.StakeholderServiceAddress,
				register: func(ctx context.Context, mux *runtime.ServeMux, client interface{}) error {
					return stakeholder.RegisterStakeholderServiceHandlerClient(ctx, mux, client.(stakeholder.StakeholderServiceClient))
				},
				service: "stakeholder_service",
				newClient: func(conn *grpc.ClientConn) interface{} {
					return stakeholder.NewStakeholderServiceClient(conn)
				},
			},

		}

		for _, svc := range services {
			conn, err := grpc.DialContext(ctx, svc.address, opts...)
			if err != nil {
				log.Fatalf("Failed to dial %s service: %v", svc.service, err)
			}
			defer conn.Close()

			client := svc.newClient(conn)
			if err := svc.register(ctx, server.mux, client); err != nil {
				log.Fatalf("Failed to register %s gateway: %v", svc.service, err)
			}
			log.Printf("%s client registered", svc.service)
		}
	}
*/
func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8000"), cors(server.mux)))
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("API starting")
	cfg := config.GetConfig()
	server := NewServer(&cfg)
	server.Start()
}
