package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	widgets "github.com/nzoschke/gomesh-proto/gen/go/widgets/v2"
	"github.com/segmentio/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type config struct {
	Port      int    `conf:"p" help:"Port to listen"`
	MongoAddr string `conf:"m" help:"Mongo address"`
}

func main() {
	config := config{
		Port:      8000,
		MongoAddr: "mongo:27017",
	}
	conf.Load(&config)

	if err := serve(config); err != nil {
		panic(err)
	}
}

func serve(config config) error {
	session, err := mgo.Dial(config.MongoAddr)
	if err != nil {
		return err
	}
	defer session.Close()

	// TODO: db.widgets.createIndex({"name": 1}, {"unique": true})

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)
	widgets.RegisterWidgetsServer(s, &Server{
		DB: session.DB("gomesh"),
	})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}

// Server implements the widgets/v2 interface
type Server struct {
	DB *mgo.Database
}

// Widget is a Mongo widget
type Widget struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Parent      string
	Name        string
	DisplayName string
	CreateTime  time.Time
}

func fromWidget(w *Widget) (*widgets.Widget, error) {
	t, err := ptypes.TimestampProto(w.CreateTime)
	if err != nil {
		return nil, err
	}

	return &widgets.Widget{
		Parent:      w.Parent,
		Name:        w.Name,
		DisplayName: w.DisplayName,
		CreateTime:  t,
	}, nil
}

func toWidget(w *widgets.Widget) (*Widget, error) {
	t, err := ptypes.Timestamp(w.CreateTime)
	if err != nil {
		return nil, err
	}

	return &Widget{
		Parent:      w.Parent,
		Name:        w.Name,
		DisplayName: w.DisplayName,
		CreateTime:  t,
	}, nil
}

func handleError(err error) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

// BatchGet Widgets by names
func (s *Server) BatchGet(ctx context.Context, r *widgets.BatchGetRequest) (*widgets.BatchGetResponse, error) {
	ws := []Widget{}
	if err := s.DB.C("widgets").Find(bson.M{"parent": r.Parent, "name": bson.M{"$in": r.Names}}).All(&ws); err != nil {
		return nil, err
	}

	pws := []*widgets.Widget{}
	for _, w := range ws {
		pw, err := fromWidget(&w)
		if err != nil {
			return nil, err
		}

		pws = append(pws, pw)
	}

	return &widgets.BatchGetResponse{
		Widgets: pws,
	}, nil
}

// Create Widget
func (s *Server) Create(ctx context.Context, r *widgets.CreateRequest) (*widgets.Widget, error) {
	// generate read-only values
	r.Widget.CreateTime = ptypes.TimestampNow()
	r.Widget.Name = fmt.Sprintf("%s/widgets/%s", r.Parent, r.Id)
	r.Widget.Parent = r.Parent

	w, err := toWidget(r.Widget)
	if err != nil {
		return nil, err
	}

	if err := s.DB.C("widgets").Insert(w); err != nil {
		return nil, err
	}
	return r.Widget, nil
}

// Delete Widget
func (s *Server) Delete(ctx context.Context, r *widgets.DeleteRequest) (*empty.Empty, error) {
	if err := s.DB.C("widgets").Remove(bson.M{"name": r.Name}); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// Get Widgets
func (s *Server) Get(ctx context.Context, r *widgets.GetRequest) (*widgets.Widget, error) {
	w := Widget{}
	if err := s.DB.C("widgets").Find(bson.M{"name": r.Name}).One(&w); err != nil {
		return nil, err
	}

	pw, err := fromWidget(&w)
	if err != nil {
		return nil, err
	}

	return pw, nil
}

// List Widgets
// TODO: Pagination
func (s *Server) List(ctx context.Context, r *widgets.ListRequest) (*widgets.ListResponse, error) {
	ws := []Widget{}
	if err := s.DB.C("widgets").Find(bson.M{"parent": r.Parent}).All(&ws); err != nil {
		return nil, err
	}

	pws := []*widgets.Widget{}
	for _, w := range ws {
		pw, err := fromWidget(&w)
		if err != nil {
			return nil, err
		}

		pws = append(pws, pw)
	}

	return &widgets.ListResponse{
		Widgets: pws,
	}, nil
}

// Update Widget
// TODO: update_mask
func (s *Server) Update(ctx context.Context, r *widgets.UpdateRequest) (*widgets.Widget, error) {
	err := s.DB.C("widgets").Update(
		bson.M{"name": r.Widget.Name},
		bson.M{"$set": bson.M{"displayname": r.Widget.DisplayName}},
	)
	if err != nil {
		return nil, err
	}

	return s.Get(ctx, &widgets.GetRequest{Name: r.Widget.Name})
}
