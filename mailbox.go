package mailbox

import (
	"fmt"
	"net/http"

	"github.com/dynamicgo/config"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/goany/slf4go"
)

// APIServer .
type APIServer struct {
	engine *gin.Engine
	slf4go.Logger
	laddr string
	db    *xorm.Engine
}

// NewAPIServer .
func NewAPIServer(conf *config.Config) (*APIServer, error) {

	db, err := initXORM(conf)

	if err != nil {
		return nil, err
	}

	if !conf.GetBool("mailbox.debug", true) {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())

	if conf.GetBool("mailbox.debug", true) {
		engine.Use(gin.Logger())
	}

	server := &APIServer{
		engine: engine,
		Logger: slf4go.Get("mailbox"),
		laddr:  conf.GetString("mailbox.laddr", ":8000"),
		db:     db,
	}

	server.makeRouters()

	return server, nil
}

func initXORM(conf *config.Config) (*xorm.Engine, error) {
	username := conf.GetString("mailbox.db.username", "xxx")
	password := conf.GetString("mailbox.db.password", "xxx")
	port := conf.GetString("mailbox.db.port", "6543")
	host := conf.GetString("mailbox.db.host", "localhost")
	scheme := conf.GetString("mailbox.db.schema", "postgres")

	return xorm.NewEngine(
		"postgres",
		fmt.Sprintf(
			"user=%v password=%v host=%v dbname=%v port=%v sslmode=disable",
			username, password, host, scheme, port,
		),
	)
}

func (server *APIServer) makeRouters() {
	server.engine.POST("/user", func(ctx *gin.Context) {

		var user User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			server.ErrorF("parse user object error :%s", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, err := server.db.Insert(&user)

		if err != nil {
			server.ErrorF("save user int pg database error :%s", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// if err := server.createWallet(ctx.Param("address"), ctx.Param("userid")); err != nil {
		// 	server.ErrorF("create wallet error :%s", err)
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }
	})

	server.engine.POST("/user", func(ctx *gin.Context) {
	})
}

// Run run http service
func (server *APIServer) Run() error {
	return server.engine.Run(server.laddr)
}
