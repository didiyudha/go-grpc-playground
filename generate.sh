protoc -I  unary/greeting/proto unary/greeting/proto/greet.proto --go_out=plugins=grpc:unary/greeting/greet
protoc -I route_guide/proto route_guide/proto/route_guide.proto --go_out=plugins=grpc:route_guide/routeguide