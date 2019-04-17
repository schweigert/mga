package main

func roiCharacter() {
	err := RPCGHClient.Call("Listener.RoiCharacter", UserCharacter, &ROI)
	if err != nil {
		panic(err)
	}
}
