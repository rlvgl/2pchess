package main

import "github.com/rlvgl/2pchess/api"

func main() {

	s := api.NewAPIServer(":8080")
	s.Start()

	// g := game.NewGame()
	// g.Play()

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer conn.Close()

	// 	client := rpc.NewGameClient(conn)
	// 	res, err := client.Move(context.Background(), &rpc.MoveData{Move: "e2e4"})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(res)
	// }()

	// NewGRPCGameServer().Run()
}
