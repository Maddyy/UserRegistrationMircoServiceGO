package main
import (
	"net"
	registrationService "UserRegistration/RegistrationService/proto"
	regHandlers  "UserRegistration/RegistrationService/handler"
   "fmt"
  "google.golang.org/grpc"
)
func main(){
	lis ,err := net.Listen("tcp",":1111")
  if err!=nil{
	  fmt.Print(err)
  }
  defer lis.Close()
  registrationServ := regHandlers.RegistrationServiceServer{}
  grpcServer := grpc.NewServer()
  registrationService.RegisterRegistrationServiceServer(grpcServer,&registrationServ)

  if err :=grpcServer.Serve(lis);err !=nil{
	  fmt.Println(err)
  }
}