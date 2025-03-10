module github.com/envoyproxy/ratelimit

go 1.14

require (
	github.com/alicebob/miniredis/v2 v2.11.4
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/coocood/freecache v1.1.0
	github.com/envoyproxy/go-control-plane v0.9.9-0.20201210154907-fd9021fe5dad
	//github.com/patramsey/go-control-plane v0.9.10
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/golang/mock v1.4.1
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.7.4-0.20191121170500-49c01487a141
	github.com/kavu/go_reuseport v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lyft/goruntime v0.2.5
	github.com/lyft/gostats v0.4.0
	github.com/mediocregopher/radix/v3 v3.5.1
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.5.1
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/text v0.3.3-0.20191122225017-cbf43d21aaeb // indirect
	google.golang.org/grpc v1.36.0
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/envoyproxy/ratelimit => github.com/patramsey/ratelimit v1.4.0

replace github.com/envoyproxy/go-control-plane v0.9.9-0.20201210154907-fd9021fe5dad => github.com/patramsey/go-control-plane v0.9.12
