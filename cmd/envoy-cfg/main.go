package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	accesslogconfig "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v2"
	bootstrap "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	accesslog "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	hcm "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/proto"
)

type config struct {
	PortAdmin   uint32
	PortEgress  uint32
	PortIngress uint32

	VhostsEgress  []vhost
	VhostsIngress []vhost
}

type vhost struct {
	cluster  string
	domains  []string
	prefix   string
	wildcard bool
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: envoy-cfg output.yaml")
		os.Exit(2)
	}

	c := config{
		PortAdmin:   9901,
		PortEgress:  9100,
		PortIngress: 9000,

		VhostsEgress: []vhost{
			vhost{cluster: "widgets-v1"},
		},
		VhostsIngress: []vhost{
			vhost{cluster: "users-v2", wildcard: true},
		},
	}

	err := writeConfig(c)
	if err != nil {
		panic(err)
	}
}

func writeConfig(c config) error {
	bs := Boostrap(c)
	if err := bs.Validate(); err != nil {
		return err
	}

	m := jsonpb.Marshaler{}
	j, err := m.MarshalToString(&bs)
	if err != nil {
		return err
	}

	y, err := yaml.JSONToYAML([]byte(j))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(os.Args[1], y, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Boostrap(c config) bootstrap.Bootstrap {
	return bootstrap.Bootstrap{
		Admin: bootstrap.Admin{
			AccessLogPath: "/dev/stdout",
			Address:       Address("0.0.0.0", c.PortAdmin),
		},
		StaticResources: &bootstrap.Bootstrap_StaticResources{
			Clusters: Clusters(c),
			Listeners: []api.Listener{
				Listener("egress", c.PortEgress, c.VhostsEgress),
				Listener("ingress", c.PortIngress, c.VhostsIngress),
			},
		},
	}
}

func Address(address string, port uint32) *core.Address {
	return &core.Address{
		Address: &core.Address_SocketAddress{
			SocketAddress: &core.SocketAddress{
				Address: address,
				PortSpecifier: &core.SocketAddress_PortValue{
					PortValue: port,
				},
			},
		},
	}
}

func Cluster(v vhost) api.Cluster {
	return api.Cluster{
		Name:                 v.cluster,
		Http2ProtocolOptions: &core.Http2ProtocolOptions{},
		ConnectTimeout:       250 * time.Millisecond,
		Type:                 api.Cluster_STRICT_DNS,
		LbPolicy:             api.Cluster_ROUND_ROBIN,
		LoadAssignment: &api.ClusterLoadAssignment{
			ClusterName: v.cluster,
			Endpoints: []endpoint.LocalityLbEndpoints{
				endpoint.LocalityLbEndpoints{
					LbEndpoints: []endpoint.LbEndpoint{
						endpoint.LbEndpoint{
							Endpoint: &endpoint.Endpoint{
								Address: Address(v.cluster, 8000),
							},
						},
					},
				},
			},
		},
	}
}

func Clusters(c config) []api.Cluster {
	cs := []api.Cluster{}

	for _, v := range c.VhostsEgress {
		cs = append(cs, Cluster(v))
	}

	for _, v := range c.VhostsIngress {
		cs = append(cs, Cluster(v))
	}

	return cs
}

func Listener(name string, port uint32, vhosts []vhost) api.Listener {
	return api.Listener{
		Name:    name,
		Address: *Address("0.0.0.0", port),
		FilterChains: []listener.FilterChain{
			listener.FilterChain{
				Filters: []listener.Filter{
					listener.Filter{
						Name: "envoy.http_connection_manager",
						Config: typeStruct(&hcm.HttpConnectionManager{
							AccessLog: []*accesslog.AccessLog{
								&accesslog.AccessLog{
									Name: "envoy.file_access_log",
									Config: typeStruct(&accesslogconfig.FileAccessLog{
										Path: "/dev/stdout",
										// TODO: JSON logging https://github.com/envoyproxy/data-plane-api/commit/d3da61529d6b3b6c02de48c67bebb5cf432d9374
										Format: `[%START_TIME%] "%REQ(:METHOD)% %REQ(X-ENVOY-ORIGINAL-PATH?:PATH)% %PROTOCOL%" %RESPONSE_CODE% %RESPONSE_FLAGS% %BYTES_RECEIVED% %BYTES_SENT% %DURATION% %RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)% "%REQ(X-FORWARDED-FOR)%" "%REQ(USER-AGENT)%" "%REQ(X-REQUEST-ID)%" "%REQ(:AUTHORITY)%" "%UPSTREAM_HOST%"` + "\n",
									}),
								},
							},
							CodecType: hcm.AUTO,
							RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{
								RouteConfig: &api.RouteConfiguration{
									Name:         "local",
									VirtualHosts: virtualHosts(vhosts),
								},
							},
							StatPrefix: fmt.Sprintf("%s_http", name),
							HttpFilters: []*hcm.HttpFilter{
								&hcm.HttpFilter{
									Name: "envoy.router",
								},
							},
						}),
					},
				},
			},
		},
	}
}

func typeStruct(pb proto.Message) *types.Struct {
	m := jsonpb.Marshaler{}
	j, err := m.MarshalToString(pb)
	if err != nil {
		panic(err)
	}

	s := types.Struct{}
	err = jsonpb.UnmarshalString(j, &s)
	if err != nil {
		panic(err)
	}

	return &s
}

func virtualHosts(vhosts []vhost) []route.VirtualHost {
	vhs := []route.VirtualHost{}
	for _, v := range vhosts {
		d := v.cluster
		if v.wildcard {
			d = "*"
		}

		vhs = append(vhs, route.VirtualHost{
			Name:    v.cluster,
			Domains: []string{d},
			Routes: []route.Route{
				route.Route{
					Match: route.RouteMatch{
						PathSpecifier: &route.RouteMatch_Prefix{
							Prefix: "/",
						},
					},
					Action: &route.Route_Route{
						Route: &route.RouteAction{
							ClusterSpecifier: &route.RouteAction_Cluster{
								Cluster: v.cluster,
							},
						},
					},
				},
			},
		})
	}

	return vhs
}
