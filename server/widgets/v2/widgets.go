package users

import (
	"context"
	"fmt"

	empty "github.com/golang/protobuf/ptypes/empty"
	widgets "github.com/nzoschke/gomesh-interface/gen/go/widgets/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Interface assertion
var _ widgets.WidgetsServer = (*Server)(nil)

// Server implements the widgets/v2 interface
type Server struct{}

// BatchGet returns a batch of Widgets by name
func (s *Server) BatchGet(ctx context.Context, in *widgets.BatchGetRequest) (*widgets.BatchGetResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Create creates a Widget by parent and id
func (s *Server) Create(ctx context.Context, in *widgets.CreateRequest) (*widgets.Widget, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Delete deletes a Widget by name
func (s *Server) Delete(ctx context.Context, in *widgets.DeleteRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// Get returns a Widget
func (s *Server) Get(ctx context.Context, u *widgets.GetRequest) (*widgets.Widget, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

// List returns a page of Widgets
func (s *Server) List(ctx context.Context, in *widgets.ListRequest) (*widgets.ListResponse, error) {
	return &widgets.ListResponse{
		Widgets: []*widgets.Widget{
			&widgets.Widget{
				Color:       widgets.Widget_WIDGET_COLOR_BLUE,
				DisplayName: "A fine widget",
				Name:        fmt.Sprintf("%s/widgets/bar", in.Parent),
				Parent:      in.Parent,
			},
		},
		TotalSize: 1,
	}, nil
}

// Update updates a Widget
func (s *Server) Update(ctx context.Context, in *widgets.UpdateRequest) (*widgets.Widget, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
