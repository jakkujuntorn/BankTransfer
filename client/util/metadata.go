package util

import (
	"context"
	"fmt"
"strings"
	"google.golang.org/grpc/metadata"
	// "google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	Authorization              = "Authorization"
)

type Metadata struct {
	User_Token string
	// ClientIP  string
}

//  ดึงข้อมูลจาก header 
func Get_Metadata(ctx context.Context) *Metadata {
	fmt.Println("")
	mtdt := &Metadata{}

	// ดึง header *******************
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
	
		// ดึง authorization
		token := md.Get("authorization")
		mtdt.User_Token  = strings.TrimPrefix(token[0], "Bearer ")
	
	}


	// header := metadata.New(map[string]string{"authorization": ""})

	// metadata.FromIncomingContext(ctx)

	// metadata.NewIncomingContext()
	// ctxx := metadata.NewOutgoingContext(ctx, header)
	// _ = ctxx

	// 	md,ok:=metadata.FromIncomingContext(ctxx)
	// if ok  {
	// 	fmt.Println(ok)
	// }

	// lib "google.golang.org/grpc/metadata"
	//ทำอะไร ***************

	// 	if md, ok := metadata.FromIncomingContext(ctxx); ok {
	// fmt.Println("MD: ",md)
	// 		// get  3 ค่า  แต่ใสไปในที่เดียวกัน ******
	// 		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
	// 			mtdt.UserAgent = userAgents[0]
	// 		}

	// 		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
	// 			mtdt.UserAgent = userAgents[0]
	// 		}

	// 		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
	// 			mtdt.ClientIP = clientIPs[0]
	// 		}

	// 	 Auth := md.Get(Authorization)
	// 	 fmt.Println(Auth[0])

	// 	}

	// ทำอะไร *******************
	// ดึง addrss
	// if p, ok := peer.FromContext(ctx); ok {
	// 	mtdt.ClientIP = p.Addr.String()
	// }

	return mtdt
}
