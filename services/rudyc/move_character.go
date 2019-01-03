package main

import (
	"log"

	"github.com/schweigert/mga/libraries/randomizer"
)

func moveCharacter() {
	UserCharacter.DirectionX = randomizer.Int(3) - 1
	UserCharacter.DirectionY = randomizer.Int(3) - 1
	log.Println("Before:", UserCharacter.PositionX, UserCharacter.PositionY)
	err := RPCClient.Call("Listener.MoveCharacter", UserCharacter, &UserCharacter)
	if err != nil {
		panic(err)
	}
	log.Println("After:", UserCharacter.PositionX, UserCharacter.PositionY)
}
