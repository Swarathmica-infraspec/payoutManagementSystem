package main

import (
	payee "payoutmanagementsystem/payee"
)

func Print(s string) string {
	return s
}

func main() {
	r := payee.SetupRouter()
	r.Run(":8080")
}
