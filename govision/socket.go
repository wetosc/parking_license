// package main

// import (
// 	"fmt"
// 	"net"

// 	"./tcpclient"
// )

// func main() {
// 	fmt.Printf("\nStarted socket server\n")

// 	tcpClient.StartServer(":8123", onConnect)
// }

// func onConnect(conn net.Conn) {
// 	client := tcpClient.NewClient(conn)
// 	client.ReadFileAsync(onImage)
// }

// func onImage(c *tcpClient.Client, data []byte) {
// 	if len(data) == 0 {
// 		return
// 	}
// 	fmt.Printf("%v", len(data))
// }

// func checkError(err error, info string) {
// 	if err != nil {
// 		fmt.Printf("\n%v : %v", info, err)
// 	}
// }
