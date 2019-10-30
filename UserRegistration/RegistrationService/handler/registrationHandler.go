package handler

import (
	"fmt"
	"context"
	RegistrationService "UserRegistration/RegistrationService/proto"
	dbconnectionApi "UserRegistration/RegistrationService/dbconnectionApi"
)
type RegistrationServiceServer struct {
}
func (reg *RegistrationServiceServer)Registration(ctx context.Context,in *RegistrationService.RegistrationRequest)(*RegistrationService.RegistrationResponse,error){
	db,err1 :=dbconnectionApi.GetConnection()
	if err1 !=nil{
		panic(err1.Error())
	}
	fmt.Print("Connected DB")
	result,err2 :=dbconnectionApi.RegistrationUser(db,in.Id,in.Email,in.Name,in.Password,in.Mobile,in.City)
	if err2 !=nil{
		panic(err2.Error())
	}
	return &RegistrationService.RegistrationResponse{
		Code:    200,
		Message: result,
		RegistrationRequest:in,
		},nil

}
