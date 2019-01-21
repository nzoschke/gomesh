# Intro to GoMesh
### What is a service mesh? And why use Go, gRPC and Envoy for mesh networking?

## Intro to a Service Mesh with Envoy

A service-oriented architecture (SOA) is one of the latest evolutions in software architecture. 

In the PHP era we built interactive web pages on a single application and database **server**. In the Rails era we built web apps as a decoupled front-end communicating with a **monolithic API** over AJAX. In the current services era we build web and mobile apps that talk to **many services** and are "reactive" to the responses.

These underlying services are extremely diverse. They may be cloud Infrastructure-as-a-Service (IaaS) for persistence like S3, DynamoDB or RDS. They may be Software-as-a-Service (SaaS) for undifferentiated functionality like Stripe for credit cards or Twilio for communications.

An emerging best practice is to think of internal systems the same way, and build and deploy them as "microservices" that other systems consume as a service with a client and API. For example custom authorization, billing and notifications may be stand-alone services, built and operated by isolated teams, that support multiple product systems like a web app, API and job scheduler.

The world of microservices built on top of microservices built on top of other services can be complicated compared to monolithic systems. [Envoy](https://www.envoyproxy.io/) represents an evolution in software architecture to handle this complexity.

The idea is to move common concerns for talking to services -- discovery, authorization, rate-limiting, retries, logs and metrics -- out of appliation code and into a common layer between services. This layer is the "service mesh" and Envoy is one of the leading approaches for building a service mesh.

Envoy is a proxy server, similar to HAProxy or NGINX, that you run next to services as a "sidecar". All incoming network requests are sent to Envoy, who handles service mesh functionality before forwarding the request onto the underlying service (or dropping it due to rate-limiting, etc.). All outgoing network requests may also be sent to Envoy to broker the external request/response.

One goal is to make building a service easier. When building a service without Envoy, lots of engineering effort goes into undiffrentiated code for logging, metrics, network error handling, etc. Failure to write this code results in services that are error prone and difficult to operate. With Envoy we can focus entirely on business logic, and delegate network concerns to Envoy.

```
                                                                            ┌─────────────────┐                  
                                                                            │      envoy      │                  
                                                                            │                 │                  
                                                                            │     sidecar     │                  
                                                                            ├────────┬────────┤        ┌ ─ ─ ─ ─ 
       ┌─────────────────┐                                             req  │ingress │ egress │         external│
       │                 │                                           ──────▶│        │        │──────▶ │         
       │                 │        ┌ ─ ─ ─ ─                          ◀─ ─ ─ ┤listener│listener│◀ ─ ─ ─  service │
  req  │    internal     │         external│                           resp └────────┴────────┘        └ ─ ─ ─ ─ 
──────▶│                 │──────▶ │                                              │▲      ▲│                      
◀─ ─ ─ ┤     service     │◀ ─ ─ ─  service │         network concerns            ││      │                       
  resp │                 │        └ ─ ─ ─ ─      ────────────────────────        │       ││                      
       │                 │                        business logic concerns        ▼│      │▼                      
       └─────────────────┘                                                  ┌─────────────────┐                  
                                                                            │    internal     │                  
                                                                            │                 │                  
                                                                            │     service     │                  
                                                                            └─────────────────┘                  
```
> Before: write service code (irr)responsible for everything
> After:  delegate network concerns to Envoy, focus on business logic

Another goal is to make running many services more efficient. Before Envoy, an SOA required an HTTP load balancer in front of every service. While AWS ALBs or ELBs are solid infrastructure components, they are black boxes that require lots of inefficient HTTP/1 and cross-AZ request routing, and offer only basic statistics. Developers have been responsible for building servers that speak HTTP/1 and talk to other services via their ALB hostname. This means gRPC is out of the question due to its use of HTTP/2.

Envoy, with its configurable service discovery and load balancing, enables efficient, direct, zone-aware HTTP/2 routing.

```
                                                                                  
         │       custom retry logic       no retry logic                          
         │          ┌───────────┐          ┌───────────┐                          
         │          │           │          │           │                          
         ▼          │           ▼          │           ▼                          
┌────────────────┐  │  ┌────────────────┐  │  ┌────────────────┐  basic metrics   
│      ALB       │  │  │      ALB       │  │  │      ALB       │─ ─ ─ ─ ─ ─ ─ ─ ─▶
└────────────────┘  │  └────────────────┘  │  └────────────────┘                  
         │          │           │          │           │                          
    ┌────┴────┐     │      ┌────┴────┐     │      ┌────┴────┐      custom logs    
    ▼         ▼     │      ▼         ▼     │      ▼         ▼        metrics      
┌──────┐  ┌──────┐  │  ┌──────┐  ┌──────┐  │  ┌──────┐  ┌──────┐      stats       
│svc a1│  │svc a2│  │  │svc b1│  │svc b2│  │  │svc c1│  │svc c2│─ ─ ─ ─ ─ ─ ─ ─ ─▶
└──────┘  └──────┘  │  └──────┘  └──────┘  │  └──────┘  └──────┘                  
              │     │      │               │                │                     
              │     │      │               │                │                     
              └─────┘      └───────────────┘                └────────────────────▶
```
> Before: run everything as an HTTP/1 service behind ALBs

```
                                                                                      
                           standard retry logic                                       
                  ┌───────────┐┌───────────────────────────────┐                      
                  │           ││                               │      standard logs   
                  │           ▼│                               ▼         metrics      
    ┌──────┐  ┌──────┐     ┌──────┐  ┌──────┐     ┌──────┐  ┌──────┐      stats       
 ─ ─│envoy ├ ─│envoy ├ ─ ─ ┤envoy │─ ┤envoy │─ ─ ─│envoy ├ ─│envoy │─ ─ ─ ─ ─ ─ ─ ─ ─▶
│   └──────┘  └──────┘     └──────┘  └──────┘     └──────┘  └──────┘                  
                 │▲           │▲                               │                      
│                ││           ││                               │                      
                 ▼│           ▼│                               ▼                      
│   ┌──────┐  ┌──────┐     ┌──────┐  ┌──────┐     ┌──────┐  ┌──────┐                  
    │svc a1│  │svc a2│     │svc b1│  │svc b2│     │svc c1│  │svc c2│                  
│   └──────┘  └──────┘     └──────┘  └──────┘     └──────┘  └──┬───┘                  
         │        │            │         │            │         │                     
│                                                              ││                     
         │        │            ▼         ▼            │         └────────────────────▶
│                          ┌────────────────┐                  │                      
 ─ ─ ─ ─ ┴ ─ ─ ─ ─└─ ─ ─ ─▶│ svc discovery  │◀─ ─ ─ ─ ┴ ─ ─ ─ ─                       
                           └────────────────┘                                         
```
> After: run everything as a mesh of HTTP/2 services

## Why Envoy?

We have seen how Envoy handles critical components of the layer between services:

* Discovery -- how services find each other
* Load balancing -- how to distribute requests across many backends
* Protocols -- what data flows over the network between services and how
* Observability -- how operators understand network flow
* Fault-tolerance -- how to automatically handle expected faiures like networking and load

Getting this layer right results in cost-efficient and reliable systems.

Moving these responsibilities out of software results in developer productivity and standardization.

Handling this with open-source software offers provider portability.

## Why Protocol Buffers and gRPC?

Envoy helps build a service mesh around virtually any TCP, HTTP or HTTP/2 server. But gRPC offers key aspects for how we write business logic services:

* Strict Interfaces -- define type-safe service definitions as .proto files
* Interoperability -- implement and consume services in virtually any programming language
* Efficiency -- communicate with efficient binary HTTP/2 requests and responses
* Standard Errors -- return canonical error codes that indicate fatal or retryable errors

Using Protocol Buffers results in strict definitions and interoperability between systems.

Using gRPC results in efficient services and standardization.

## Why Go?

gRPC supports virtually any programmign language. But Go offers key aspects for how we code gRPC services:

* Stability -- the 1.x family of Go has been consistent for 8+ years
* Correctness -- the type system and error handling matches the gRPC message in -> message out and error semantics
* Packaging -- the Go tool chain makes it easy to build binaries ready for deployment next to the Envoy sidecar binary
* Observability -- the context pattern serves Go's goal of making distriuted systems programming easy

Building software in Go and gRPC results in developer productivity.

## Tradeoffs

One big tradeoff is complexity. Building an API on top of a gRPC service mesh is far more complicated than, say, a monolithic Rails API. Learning the .proto syntax, toolchain and code generation pipeline is harder than slinging JSON around. Figuring out an initial configuration for Envoy, picking a sidecar deployment strategy, and operating a service discovery backend is tougher than deploying a Rails app.

Another tradeoff is how opinionated Envoy is. Moving logging and metrics out-of-process into an Envoy sidecar means you have to adopt its logging and metrics format. This almost certainly won't match existing approaches to logging and metrics.

## Summary

Envoy, gRPC and Go offer many benefits for building a service mesh and SOA:

* Define RPC methods, inputs and outputs
* Generate Go service interfaces
* Write business logic functions that return standard error codes
* Configure Envoy sidecars
* Configure a service discovery backend

And that's about it.

Many responsibilities are no longer our problems:

* Juggling poorly defined JSON requests / responses
* Writing code for logging and metrics
* Running load balancers
* Handling retries, backoff and circuit breaking

We get to write less code, pay less for infrastructure, get consistent observability and more reliability. No wonder there's so much interest in building a service mesh with Envoy and gRPC. And Go is a great choice for writing all the service handlers.
