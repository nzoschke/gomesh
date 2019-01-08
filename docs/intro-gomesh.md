# Intro to GoMesh
### Why a Service Mesh matters and why Go, gRPC and Envoy are well-suited

## Intro to Service Mesh

A service-oriented architecture (SOA) is one of the latest patterns in software-as-a-service (SaaS). Most modern services are built op top of other services. To start, systems generally consume cloud Infrastructure-as-a-Service (IaaS) for persistence like S3, DynamoDB or RDS. Then they generally consume SaaS for undifferentiated functionality like Stripe for credit cards or Twilio for communications.

An emerging best practice is to think of internal systems the same way, and write them as "microservices" that other systems consume as a service with a client and API. For example custom authorization, billing and notifications may be stand-alone services, built and operated by isolated teams, that support multiple product services like a web app, API, and job scheduler.

The world of microservices built on top of microservices built on top of other services can be complicated compared to monolithic software. To handle this complexity there is a shift in programming to move common concerns out of software and into the layer between services. This layer is the "service mesh". [Envoy](https://www.envoyproxy.io/) is one of the leading approaches for building a service mesh.

Envoy represents a shift in programming. Before Envoy, an SOA required an HTTP load balancer in front of every service. While AWS ALBs or ELBs are solid infrastructure components, they are black boxes that require lots of inefficient HTTP/1 request routing, and emit simplistic statistics. Developers have been responsible for building servers that speak HTTP/1, and implement custom logging, metrics, rate-limiting and retries.

Envoy, with its configurable service discovery and load balancing enables efficient, direct, zone-aware HTTP/2 request routing. Envoy, with out out-of-process design, moves observability and fault tolerance from application code to application agnostic configuration.

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
> Before: ALB

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
> After: Service Mesh

## Why Envoy?

In this world of services built on top of services built on top of services, the layer between services is critically important. An Envoy service mesh handles:

* Discovery -- how services find each other
* Load balancing -- how to distribute requests across many backends
* Protocols -- what data flows over the network between services and how
* Observability -- how operators understand the requests
* Fault-tolerance -- how we can automatically handle expected faiures like networking and load

Getting this layer right results in cost-efficient and reliable systems.

Moving these responsibilities out of software results in developer productivity and standardization.

Handling this with open-source software offers provider portability.

## Why gRPC?

**interoperability**

**efficiency**

**conventions**

**reliability**

## Why Go?

With a service mesh we focus entirely on writing service methods. We could do this in any language, but some features of Go make it particularly well-suited for a gRPC SOA.

Go's release management offers **stability**. With the 1.x family of Go, code that is 7 years old and code that will be written a year from now will all work. On the opposite end of the spectrum we have JavaScript, where the Node.JS runtime and programming style is constantly evolving.

Go's type system and error handling offers **correctness**. Our service methods take messages in and return messages (or an error).

Go's cross-compiler solves **packaging**. Every laptop with the Go tool chain can build Linux binaries ready for deployment. We don't need Docker or Linux build services to produce a suitable service binary.

Go's binary program format offers **speed**. Because Lambda runs new versions of our program on-demand, slow boot times can turn into a real problem. Go's binaries have very little overhead to start, compared to a dynamic VM like Python or Java.

Go's context pattern offers **observability**. Google sponsored the development of Go with a big goal in mind - to make large scale distributed systems programming easier. Thanks to the guidance and expertise of the Googler's behind Go, it comes with important observability primitives out of the box.

## Tradeoffs and Alternatives

Complexity


## Summary
