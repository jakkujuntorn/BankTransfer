package auth

import (
	"context"
	"fmt"
	"strings"
	"google.golang.org/grpc/metadata"

	
)

const (
	Authorization = "authorization"
	Bearer = "bearer"
)

// แกะ token ************
func AuthorizeUser(ctx context.Context) (*Payload_Claims, error) {
	
	// ใช้ทำอะไร
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	// ทำงานคล้าย fiber.Get ที่รับ header authorization มาเช็ค token รึป่าว
	// แกะ authorization ****************
	values := md.Get(Authorization)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	// เเกะ bearer ***********
	authType := strings.ToLower(fields[0])
	if authType != Bearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	// ส่ง tokenเข้าไปเช็ค *****************
	payload, err := VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}
