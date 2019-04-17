package main

func CheckAccount() {
	check := false

	err := RPCClient.Call("Listener.Check", UserAccount, &check)
	if err != nil {
		panic(err)
	}

	if !check {
		panic("User not checked")
	}
}
