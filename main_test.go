package main

import (
	"testing"

	pb "github.com/ffjlabo-playground/protobuf-sample/gen/proto/go/protobuf_sample/v1"
	"google.golang.org/protobuf/proto"
)

func TestUser(t *testing.T) {

	u := pb.User{
		Id: "user",
	}

	_, _ = proto.Marshal(&u)
}
