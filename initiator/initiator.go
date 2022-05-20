package initiator

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"twof/blog-api/initiator/articleInit"
	"twof/blog-api/initiator/invoiceInit"
	"twof/blog-api/initiator/subscriptionInit"
	"twof/blog-api/initiator/userInit"

	"twof/blog-api/internal/glue"
	"twof/blog-api/internal/glue/auth"
	"twof/blog-api/internal/glue/enforcer"
	zaplog "twof/blog-api/internal/log"
)

func Init() {

	zaplog.SugerLogger.Errorf("this is error in initiator")

	// db, err := gorm.Open(postgres.Open("postgres://root@cockroachdb:26257/firstproject?sslmode=allow"), &gorm.Config{})
	db, err := gorm.Open(postgres.Open("postgres://root@db:26257/defaultdb?sslmode=disable"), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open("postgres://max:roach@localhost:26257/firstproject?sslmode=allow"), &gorm.Config{})

	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	adapter, _ := gormadapter.NewAdapterByDB(db)

	enf, err := casbin.NewEnforcer("rbac_model.conf", adapter)

	if err != nil {
		fmt.Println(err)
	}

	if hasPolicy := enf.HasPolicy("admin", "/api/*", "POST"); !hasPolicy {
		enf.AddPolicy("admin", "/api/*", "POST")
	}

	if hasPolicy := enf.HasPolicy("admin", "/api/*", "PUT"); !hasPolicy {
		enf.AddPolicy("admin", "/api/*", "PUT")
	}

	if hasPolicy := enf.HasPolicy("admin", "/api/*", "GET"); !hasPolicy {
		enf.AddPolicy("admin", "/api/*", "GET")
	}

	if hasPolicy := enf.HasPolicy("admin", "/api/v1/articles/", "GET"); !hasPolicy {
		enf.AddPolicy("admin", "/api/v1/articles/", "GET")
	}

	if hasPolicy := enf.HasPolicy("admin", "/api/v1/articles", "GET"); !hasPolicy {
		enf.AddPolicy("admin", "/api/v1/articles", "GET")
	}

	if hasPolicy := enf.HasPolicy("author", "/api/v1/articles", "POST"); !hasPolicy {
		enf.AddPolicy("author", "/api/v1/articles", "POST")
	}

	if hasPolicy := enf.HasPolicy("user", "/api/v1/articles/", "GET"); !hasPolicy {
		enf.AddPolicy("user", "/api/v1/articles", "GET")
	}

	en := enforcer.CasbinInit(enf)

	r := gin.Default()

	g := r.Group("/api/v1")

	g.Use(auth.AuthorizeJWT(
		glue.AllowPathPrefixSkipper("/api/v1/auth/"),
	))

	g.Use(auth.Authorize(enf,
		glue.AllowPathPrefixSkipper("/api/v1/auth"),
	))
	userInit.Init(en, g, db)
	articleInit.Init(g, db)
	subscriptionInit.Init(g, db)
	invoiceInit.Init(db)

	fmt.Println("runing on Port  8088")
	if err := r.Run("127.0.0.1:8088"); err != nil {
		log.Fatal(err)
	}
}
