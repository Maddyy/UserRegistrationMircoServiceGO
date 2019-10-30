package main

import (
	"google.golang.org/grpc"
	"fmt"
	"context"
	registrationService "UserRegistration/RegistrationService/proto"
	
)
func main(){
 
	var conn *grpc.ClientConn
	conn,err :=grpc.Dial(":1111",grpc.WithInsecure())
	if err!=nil{
		fmt.Println(err)
	}

   defer conn.Close()

   regServ := registrationService.NewRegistrationServiceClient(conn)
   result,err2 := regServ.Registration(context.Background(),&registrationService.RegistrationRequest{
	Id:"1",
	Email:"sharad@gmail.com",
	Name:"sharad",
	Password:"sharad",
	Mobile:"98000001",
	City:"Delhi",
 })

   if err2 != nil {
	fmt.Println(err2)
   } else {
	fmt.Println("Result :",result.RegistrationRequest)
   }
}