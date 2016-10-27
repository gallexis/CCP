package Payloads

type alert struct{
	message [4]byte
}

func V() string{
	return "alert"
}