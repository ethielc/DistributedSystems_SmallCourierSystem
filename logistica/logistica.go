package main

import (
        "fmt"
        "log"
        "net"
        "os"
        "github.com/432i/T1SisDistribuidos/logistica/chat"
        "google.golang.org/grpc"
)
func crearRegistro(){
        archivo, err := os.Create("registro.csv")
        if err != nil{
                log.Println(err)
        }
        defer archivo.Close()
})
func main() {

        fmt.Println("Go gRPC Beginners Tutorial!")

        lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        crearRegistro()
        

        s := chat.Server{}

        grpcServer := grpc.NewServer()

        chat.RegisterChatServiceServer(grpcServer, &s)

        if err := grpcServer.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %s", err)
        }
        fmt.Println("holahola")
}
