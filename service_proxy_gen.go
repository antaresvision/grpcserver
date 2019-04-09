//go:generate protoc -I. --grpc-gateway_out=logtostderr=true:. service.proto

package main
