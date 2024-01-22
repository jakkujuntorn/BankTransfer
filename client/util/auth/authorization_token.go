package util

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)



type Metadata struct {
	User_Token string
}

// แกะ token ************
func (signature *JWTMaker)AuthorizeUser(ctx context.Context) (*Payload_Claims, error) {
	mtdt := &Metadata{}

	// ดึง header *******************
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		// ดึง authorization

		fmt.Println("")
		// fmt.Println(md)

		token := md.Get("authorization")
		mtdt.User_Token  = strings.TrimPrefix(token[0], "Bearer ")
	}

	

	payload,err:=signature.verifyToken(mtdt.User_Token)
	if err != nil {
		return &Payload_Claims{}, err
	}

	return payload, nil
}
