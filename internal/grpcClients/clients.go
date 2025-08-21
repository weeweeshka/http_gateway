package grpcClients

import (
	pbSSO "github.com/weeweeshka/sso_proto/gen/go/sso"
	pbTataisk "github.com/weeweeshka/tataisk_proto/gen/go/tataisk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SetupGateway() (pbSSO.SsoClient, pbTataisk.TataiskClient) {

	ssoConn, _ := grpc.NewClient("sso:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	tataiskConn, _ := grpc.NewClient("tataisk:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	ssoClient := pbSSO.NewSsoClient(ssoConn)
	tataiskClient := pbTataisk.NewTataiskClient(tataiskConn)

	return ssoClient, tataiskClient
}
