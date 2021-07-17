package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello")

	sachin := &Person{
		Name: "Sachin",
		Age:  21,
		SocialFollowers: &SociaSocialFollowers{
			Youtube: 2321,
			Twitter: 54543,
		},
	}

	data, err := proto.Marshal(sachin)
	if err != nil {
		log.Fatal("Marshalling error:", err)
	}

	fmt.Println(data)

	newSachin := &Person{}
	err = proto.Unmarshal(data, newSachin)
	if err != nil {
		log.Fatal("unmarshalling error:", err)
	}

	fmt.Println(newSachin.GetAge())
	fmt.Println(newSachin.GetName())
	fmt.Println(newSachin.SocialFollowers.GetTwitter())
	fmt.Println(newSachin.SocialFollowers.GetYoutube())
}
