package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/lab-icn/water-potability-sensor-service/internal/infra"
	"github.com/lab-icn/water-potability-sensor-service/internal/water_potability/interface/mqtt"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/water_potability/interface/rpc"
	"github.com/lab-icn/water-potability-sensor-service/internal/water_potability/repository"
	"github.com/lab-icn/water-potability-sensor-service/internal/water_potability/service"
	"google.golang.org/grpc"
)

func mux() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	v1 := chi.NewRouter()
	v1.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy"))
	})
	r.Mount("/api/v1", v1)

	return r
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	conn, err := grpc.NewClient(os.Getenv("GRPC_ADDR"), opts...)
	if err != nil {
		log.Fatalf("failed to dial grpc: %v", err)
		return
	}

	defer conn.Close()

	wpGrpcClient := pb.NewWaterPotabilityServiceClient(conn)

	influxDB := infra.NewInfluxDB(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))
	wpRepository := repository.NewWaterPotabilityRepository(influxDB)
	wpService := service.NewWaterPotabilityService(wpRepository, wpGrpcClient)

	mqtt.NewMQTT(wpService)

	// server := &http.Server{Addr: "0.0.0.0:8080", Handler: mux()}
	// ctx, cancel := context.WithCancel(context.Background())
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// go func() {
	// 	<-sig
	// 	shutdownCtx, cancelShutdownCtx := context.WithTimeout(ctx, 30*time.Second)

	// 	go func() {
	// 		<-shutdownCtx.Done()
	// 		if shutdownCtx.Err() == context.DeadlineExceeded {
	// 			log.Fatal("graceful shutdown timed out.. forcing exit.")
	// 		}
	// 	}()

	// 	err := server.Shutdown(shutdownCtx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	cancelShutdownCtx()
	// 	cancel()
	// }()

	// fmt.Printf("server running on %s\n", server.Addr)
	// err := server.ListenAndServe()
	// if err != nil && err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }

	// <-ctx.Done()
}
