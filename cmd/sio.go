/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/dk-sirius/ztp/pkg/logs"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// sioCmd represents the sio command
var sioCmd = &cobra.Command{
	Use:   "sio",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	router := gin.New()
	srv := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
				Client: &http.Client{
					Timeout: time.Minute,
				},
			},
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	})
	defer func(srv *socketio.Server) {
		err := srv.Close()
		if err != nil {
			logs.CError(err.Error())
		}
	}(srv)

	srv.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		logs.CInfo(conn.ID())
		return nil
	})

	srv.OnDisconnect("/", func(conn socketio.Conn, s string) {
		logs.CInfo(s)
	})

	srv.OnEvent("/", "message", func(s socketio.Conn, content interface{}) {
		fmt.Println(content)
	})
	go func() {
		if err := srv.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	router.GET("/socket.io/*any", gin.WrapH(srv))
	router.POST("/socket.io/*any", gin.WrapH(srv))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

func init() {
	rootCmd.AddCommand(sioCmd)
}
