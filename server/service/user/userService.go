package user

import (
	"jsonrpc/server/service/user/dao"
	"jsonrpc"
	"github.com/tsenart/nap"
	"jsonrpc/registry"
)

var db *nap.DB
var serviceAddress string
const ServiceName = "User"

type User struct {
}

type UserService interface {
	GetUser()
}

func (u *User) GetUser(x, reply *string) error {
	*reply = dao.GetUser(db)
	return nil
}

func Register(DB *nap.DB) {
	jsonrpc.Register(new(User))
	registry.RegisterService(serviceAddress, ServiceName)
	db = DB
}

func InjectSA(ServiceAddress string) {
	serviceAddress = ServiceAddress
}