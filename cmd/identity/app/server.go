package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/code/codeimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/config"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/protocol/rest"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
	roleimpl "github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role/roleimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/staffimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user/userimpl"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/password"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	MongoDB    *mongodb.Database
	cfg        *config.Config
	RestServer *rest.Server
}

const serviceName string = "identity"

func NewServer() (*Server, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	mongoDB, err := mongodb.New(cfg.MongoDB.URI, cfg.MongoDB.DBName)
	if err != nil {
		return nil, err
	}
	codeStore := codeimpl.NewStore(mongoDB)
	userStore := userimpl.NewStore(mongoDB)
	staffStore := staffimpl.NewStore(mongoDB)
	roleStore := roleimpl.Newstore(mongoDB)

	initAccountAdminRoot(staffStore)

	orderClient, err := getOrderClient(cfg)
	if err != nil {
		return nil, err
	}

	userSrv := userimpl.NewService(userStore, codeStore, orderClient)
	staffSrv := staffimpl.NewService(staffStore, roleStore)
	roleSrv := roleimpl.NewService(roleStore, staffStore)

	restServer := rest.NewServer(&rest.Dependences{
		UserSrv:      userSrv,
		StaffSrv:     staffSrv,
		StaffRoleSrv: roleSrv,
	}, cfg)

	go func() {
		if err := restServer.Run(context.Background()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	return &Server{
		MongoDB: mongoDB,
		cfg:     cfg,
	}, nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.RestServer != nil {
		return s.RestServer.Shutdown(ctx)
	}
	return nil
}

func initAccountAdminRoot(staffStore staff.Store) {
	ctx := context.Background()
	total := staffStore.CountByCondition(ctx, bson.M{"isRoot": true})
	if total <= 0 {
		passwordHash, _ := password.HashPassword("123456")
		now := time.Now().UTC()
		doc := staff.Staff{
			ID:          primitive.NewObjectID(),
			Username:    "admin",
			Password:    passwordHash,
			Phone:       "0702118453",
			Active:      true,
			Role:        primitive.ObjectID{},
			CreatedAt:   now,
			UpdatedAt:   now,
			IsRoot:      true,
			Permissions: make([]string, 0),
		}

		staffStore.InsertOne(ctx, doc)
		fmt.Println(aurora.Green("*** Init account admin root: "))
	}
}
