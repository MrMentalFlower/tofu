/*
    ___       ___       ___       ___
   /\  \     /\  \     /\  \     /\__\
   \:\  \   /::\  \   /::\  \   /:/ _/_
   /::\__\ /:/\:\__\ /::\:\__\ /:/_/\__\
  /:/\/__/ \:\/:/  / \/\:\/__/ \:\/:/  /
  \/__/     \::/  /     \/__/   \::/  /
             \/__/               \/__/

Date: 2023 1-4
Creator: flower
Description:

*/

package main

import (
	"fmt"

	"github.com/MrMentalFlower/tofu/pkg/cmd/client"
	"github.com/MrMentalFlower/tofu/pkg/cmd/server"
)

func main() {
	fmt.Println("start")
	go server.Open()
	client.Open()

}
