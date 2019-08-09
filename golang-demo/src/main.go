package main

import (
	"flag"
	"fmt"
	"time"

	kpb "./kunpengBattle"
)

var ip string
var port int
var teamID int

func init() {
	flag.IntVar(&teamID, "teamID", 0, "TeamID")
	flag.StringVar(&ip, "ip", "127.0.0.1", "ServerIP")
	flag.IntVar(&port, "port", 6001, "Server Port")
}

func main() {
	flag.Parse()
	fmt.Printf("%v join KunPengBattle (%v, %v): \n", teamID, ip, port)

	strategy := new(ramdomStrategy)
	client := kpb.NewKunPengBattleClient(teamID, "Random", strategy)

	var err error
	for i := 0; i < 30; i++ {
		err = client.Connect(ip, port)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	if err != nil {
		fmt.Println("Connection Failed!!!")
		return
	}

	client.Start()

}
