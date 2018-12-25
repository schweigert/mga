package main

func roiCharacter() {
	err := RPCClient.Call("Listener.RoiCharacter", UserCharacter, &ROI)
	if err != nil {
		panic(err)
	}
}
