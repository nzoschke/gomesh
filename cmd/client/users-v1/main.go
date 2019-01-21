package main

import (
	"context"
	"fmt"
	"log"

	users "github.com/nzoschke/gomesh-interface/gen/go/users/v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	c := users.NewUsersClient(conn)

	ctx := context.Background()
	u, err := c.Create(ctx, &users.CreateRequest{
		Parent: "orgs/myorg",
		UserId: "myusername",
		User: &users.User{
			DisplayName: "My Full Name",
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("USER: %+v\n", u)

	u, err = c.Get(ctx, &users.GetRequest{
		Name: "users/foo",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("USER: %+v\n", u)
}
