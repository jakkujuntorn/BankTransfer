package util

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

// ทำงานอย่างไร
//************ ในนี้ แกะค่าอะไรสักอย่างจาก Ctx *************
func extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	// lib "google.golang.org/grpc/metadata"
	//ทำอะไร ***************
	if md, ok := metadata.FromIncomingContext(ctx); ok {

		// get  3 ค่า  แต่ใสไปในที่เดียวกัน ******
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtdt.ClientIP = clientIPs[0]
		}

	}

	// ทำอะไร *******************
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}