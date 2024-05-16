package main

import (
	"api-gateway/config"
	"api-gateway/proto/blog"
	"api-gateway/proto/follower"
	"api-gateway/proto/stakeholder"
	"api-gateway/proto/tour"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Println("Api starting")

	cfg := config.GetConfig()
	log.Println("config dobar")
	blogConn, err := grpc.DialContext(
		context.Background(),
		cfg.BlogServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial blog service:", err)
	}
	log.Println("Api made first conn ")

	followerCon, err := grpc.DialContext(
		context.Background(),
		cfg.FollowerServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial follower service:", err)
	}

	stakeHolderCon, err := grpc.DialContext(
		context.Background(),
		cfg.StakeholderServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial stakeholder service:", err)
	}

	tourCon, err := grpc.DialContext(
		context.Background(),
		cfg.TourServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial tour service:", err)
	}

	log.Println("Connections made")

	gwmux := runtime.NewServeMux()
	// Register Blog
	blogClient := blog.NewBlogServiceClient(blogConn)
	err = blog.RegisterBlogServiceHandlerClient(
		context.Background(),
		gwmux,
		blogClient,
	)
	if err != nil {
		log.Fatalln("Failed to register blog gateway:", err)
	}

	log.Println("Blog client started ")

	// Register Follower
	followerClient := follower.NewFollowServiceClient(followerCon)
	err = follower.RegisterFollowServiceHandlerClient(
		context.Background(),
		gwmux,
		followerClient,
	)
	if err != nil {
		log.Fatalln("Failed to register follower gateway:", err)
	}

	log.Println("Follower client started ")

	// Register Other
	stakeholderClient := stakeholder.NewStakeholderServiceClient(stakeHolderCon)
	err = stakeholder.RegisterStakeholderServiceHandlerClient(
		context.Background(),
		gwmux,
		stakeholderClient,
	)
	if err != nil {
		log.Fatalln("Failed to register stakeholder gateway:", err)
	}
	log.Println("Stakeholder client started ")

	// Register Other
	tourClient := tour.NewTourServiceClient(tourCon)
	err = tour.RegisterTourServiceHandlerClient(
		context.Background(),
		gwmux,
		tourClient,
	)
	if err != nil {
		log.Fatalln("Failed to register tour gateway:", err)
	}
	/*
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

		tourEndpoint := "tour:8000"fmt.Sprintf("%s:%s", server.config.AccommodationReservationHost, server.config.AccommodationReservationPort)
		err := reservationGw.RegisterAccommodationReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
		if err != nil {
			panic(err)
		}

		accommodationEmdpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
		err = accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEmdpoint, opts)
		if err != nil {
			panic(err)
		}
	*/
	log.Println("Tour client started ")

	gwServer := &http.Server{
		Addr:    cfg.Address,
		Handler: gwmux,
	}

	log.Println("Api started")

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
