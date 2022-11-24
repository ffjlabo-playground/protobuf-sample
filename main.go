package main

import (
	"encoding/hex"
	"fmt"

	pb "github.com/ffjlabo-playground/protobuf-sample/gen/proto/go/protobuf_sample/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := pb.User{
		Id: "user",
	}

	b, err := proto.Marshal(&u)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(b))
}
