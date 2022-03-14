package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	pb "simple-api/gen/proto"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)

	}

	client := pb.NewTestApiClient(conn)
	g := gin.Default()

	g.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")

		req := &pb.ResponseRequest{Nsg: name}
		if resp, err := client.Echo(context.Background(), req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint("Hello" + "," + resp.Nsg),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
