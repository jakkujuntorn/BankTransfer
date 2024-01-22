curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto

// ให้ gen ไปลงที่ไหน***********

// client
protoc *.proto --go_out=../../client/service_proto --go-grpc_out=../../client/service_proto


// server 
protoc *.proto --go_out=../../server/service_proto --go-grpc_out=../../server/service_proto

//******************************************

// client for account
protoc *.proto --go_out=../../client/controller_client --go-grpc_out=../../client/controller_client --grpc-gateway_out=../../client/controller_client


// server for account
protoc *.proto --go_out=../../server --go-grpc_out=../../server

// ********************************************
// client for transfer
protoc *.proto --go_out=../../client/controller_client --go-grpc_out=../../client/controller_client --grpc-gateway_out=../../client/controller_client


// server for transfer
protoc *.proto --go_out=../../server --go-grpc_out=../../server

//*********************************************
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > user/google/api/annotations.proto
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > user/google/api/http.proto

// client for user
protoc *.proto --go_out=../../client/controller_client --go-grpc_out=../../client/controller_client --grpc-gateway_out=../../client/controller_client


// server for user
protoc *.proto --go_out=../../server --go-grpc_out=../../server