package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	docker "github.com/docker/docker/client"
	xds "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/segmentio/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port int `conf:"p" help:"Port to listen"`
}

// Server implements the Envoy xDS Cluster Discovery Service
type Server struct {
	DockerClient *docker.Client
	lastTime     time.Time
	lastVersion  int
}

func main() {
	config := config{
		Port: 8000,
	}
	conf.Load(&config)

	err := serve(config)
	if err != nil {
		panic(err)
	}
}

func serve(config config) error {
	c, err := docker.NewEnvClient() // unix:///var/run/docker.sock by default
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	xds.RegisterClusterDiscoveryServiceServer(s, &Server{
		DockerClient: c,
	})
	reflection.Register(s)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		return err
	}

	fmt.Printf("listening on :%d\n", config.Port)
	return s.Serve(l)
}

// FetchClusters implements the Envoy CDS REST endpoint
func (s *Server) FetchClusters(ctx context.Context, r *xds.DiscoveryRequest) (*xds.DiscoveryResponse, error) {
	fmt.Printf("FetchClusters R: %+v\n", r)
	return nil, fmt.Errorf("unimplemented")
}

// IncrementalClusters implements the Envoy CDS incremental endpoint
func (s *Server) IncrementalClusters(server xds.ClusterDiscoveryService_IncrementalClustersServer) error {
	for {
		r, err := server.Recv()
		if err != nil {
			return err
		}
		fmt.Printf("IncrementalClusters R: %+v\n", r)
		return fmt.Errorf("unimplemented")
	}
}

// StreamClusters implements the Envoy CDS streaming endpoint
func (s *Server) StreamClusters(server xds.ClusterDiscoveryService_StreamClustersServer) error {
	for {
		_, err := server.Recv()
		if err != nil {
			return err
		}

		// send current clusters
		err = s.sendClusters(server, time.Now())
		if err != nil {
			return err
		}

		// poll for container changes
		filter := filters.NewArgs()
		filter.Add("type", "container")
		filter.Add("event", "start")
		filter.Add("event", "die")
		filter.Add("label", "com.docker.compose.project=gomesh")

		events, errs := s.DockerClient.Events(context.Background(), types.EventsOptions{
			Since:   s.lastTime.Format(time.RFC3339Nano),
			Filters: filter,
		})

		for {
			select {
			case err := <-errs:
				return err
			case e := <-events:
				// FIXME: should send incremental update for specific event
				err = s.sendClusters(server, time.Unix(0, e.TimeNano))
				if err != nil {
					return err
				}
			}
		}
	}
}

func (s *Server) sendClusters(server xds.ClusterDiscoveryService_StreamClustersServer, at time.Time) error {
	cs, err := s.clusters(server.Context())
	if err != nil {
		return err
	}

	var as []ptypes.Any
	for _, c := range cs {
		a, err := ptypes.MarshalAny(&c)
		if err != nil {
			return err
		}
		as = append(as, *a)
	}

	// increase time (nonce) and version
	s.lastTime = at
	s.lastVersion = s.lastVersion + 1

	return server.Send(&xds.DiscoveryResponse{
		Resources:   as,
		Nonce:       string(s.lastTime.Nanosecond()),
		TypeUrl:     "type.googleapis.com/envoy.api.v2.Cluster",
		VersionInfo: strconv.Itoa(s.lastVersion),
	})
}

func (s *Server) clusters(ctx context.Context) ([]xds.Cluster, error) {
	cs, err := s.DockerClient.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	// collect service name -> list of ip:port
	addresses := map[string][]core.Address_SocketAddress{}

	for _, c := range cs {
		if c.Labels["com.docker.compose.project"] != "gomesh" {
			continue
		}

		i, err := s.DockerClient.ContainerInspect(ctx, c.ID)
		if err != nil {
			return nil, err
		}

		name := c.Labels["com.docker.compose.service"]
		if _, ok := addresses[name]; !ok {
			addresses[name] = []core.Address_SocketAddress{}
		}

		// assume single port specified
		port := 10000
		// for p := range i.Config.ExposedPorts {
		// 	port = p.Int()
		// }

		// assume single network specified
		ip := "0.0.0.0"
		for _, n := range i.NetworkSettings.Networks {
			ip = n.IPAddress
		}

		addresses[name] = append(addresses[name], address(ip, port))
	}

	clusters := []xds.Cluster{}
	for n, as := range addresses {
		clusters = append(clusters, cluster(n, as))
	}

	return clusters, nil
}

func address(ip string, port int) core.Address_SocketAddress {
	return core.Address_SocketAddress{
		SocketAddress: &core.SocketAddress{
			Address: ip,
			PortSpecifier: &core.SocketAddress_PortValue{
				PortValue: uint32(port),
			},
		},
	}
}

func cluster(name string, addresses []core.Address_SocketAddress) xds.Cluster {
	es := []endpoint.LbEndpoint{}
	for i := range addresses {
		es = append(es, endpoint.LbEndpoint{
			Endpoint: &endpoint.Endpoint{
				Address: &core.Address{
					Address: &addresses[i],
				},
			},
		})
	}

	return xds.Cluster{
		Name:                 name,
		Type:                 xds.Cluster_STRICT_DNS,
		ConnectTimeout:       250 * time.Millisecond,
		Http2ProtocolOptions: &core.Http2ProtocolOptions{},
		LoadAssignment: &xds.ClusterLoadAssignment{
			ClusterName: name,
			Endpoints: []endpoint.LocalityLbEndpoints{
				endpoint.LocalityLbEndpoints{
					LbEndpoints: es,
				},
			},
		},
	}
}
