package main

import "github.com/schweigert/mga/libraries/randomizer"

func moveCharacter() {
	UserCharacter.DirectionX = randomizer.Int(3) - 1
	UserCharacter.DirectionY = randomizer.Int(3) - 1
	err := RPCGHClient.Call("Listener.MoveCharacter", UserCharacter, &UserCharacter)
	if err != nil {
		panic(err)
	}

	UserCharacter.Move()
}
