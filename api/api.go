package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rlvgl/2pchess/game"
)

type APIServer struct {
	listenAddr string
}

var upgrader = websocket.Upgrader{}
var moveLog []string

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/", s.ReturnHTML)
	router.HandleFunc("/new-game", s.NewGame)
	router.HandleFunc("/game", s.Game)

	log.Printf(fmt.Sprintf("Starting server on localhost%s", s.listenAddr))
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) Game(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade failed")
		return
	}

	defer conn.Close()

	g := game.NewGame()

	for {
		mt, message, err := conn.ReadMessage()

		if err != nil {
			log.Printf("read failed")
		}
		move := string(message)
		g.Move(move)

		board := g.ReturnBoard()

		// converting [8][8]rune to []byte
		//! how to convert this back to rune??
		returnBoard := []byte{}
		for _, row := range board {
			for _, val := range row {
				bytes := []byte(strconv.Itoa(int(val)))
				for _, b := range bytes {
					returnBoard = append(returnBoard, b)
				}
			}
		}

		err = conn.WriteMessage(mt, returnBoard)
		if err != nil {
			log.Printf("write failed")
			return
		}

	}
}

func (s *APIServer) NewGame(w http.ResponseWriter, r *http.Request) {
	g := game.NewGame()
	WriteJSON(w, g.ReturnBoard())
}

func (s *APIServer) HandleHome(w http.ResponseWriter, r *http.Request) {
	// conn, err := upgrader.Upgrade(w, r, nil)

	// if err != nil {
	// 	log.Print("upgrade failed: ", err)
	// 	return
	// }

	// defer conn.Close()

	// for {

	// }

	WriteJSON(w, "Hello world")
}

func (s *APIServer) ReturnHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/game.html")
}

func WriteJSON(w http.ResponseWriter, v any) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}
