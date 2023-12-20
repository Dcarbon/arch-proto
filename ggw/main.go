package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Dcarbon/arch-proto/pbg"
	"github.com/Dcarbon/go-shared/gutils"
	"github.com/Dcarbon/go-shared/libs/aidh"
	"github.com/Dcarbon/go-shared/libs/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

var cateHost = "localhost:2030"

type Serve struct {
	*runtime.ServeMux
}

func NewServeMux() *Serve {
	var mux = &Serve{
		ServeMux: runtime.NewServeMux(),
	}
	return mux
}

func (s *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Path: ", r.URL.Path)
	s.ServeMux.ServeHTTP(w, r)
}

var iotHost = "localhost:9035"

func main() {
	mux := NewServeMux()

	ccIot, err := gutils.GetCCTimeout(iotHost, 5*time.Second)
	utils.PanicError("", err)

	// Register generated routes to mux
	err = pbg.RegisterIotServiceHandler(context.TODO(), mux.ServeMux, ccIot)
	utils.PanicError("", err)

	mux.HandlePath(http.MethodGet, "/api/v1/iots/geojson", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		client := pbg.NewIotServiceClient(ccIot)
		data, err := client.GetIotPositions(context.TODO(), &pbg.RIotGetList{})
		if nil != err {
			w.WriteHeader(400)
			aidh.SendJSON(w, 500, err)
			return
		}

		var featureCollection = geojson.NewFeatureCollection()
		for _, loc := range data.Data {
			var feature = geojson.NewFeature(&orb.Point{loc.Position.Longitude, loc.Position.Latitude})
			feature.Properties = make(geojson.Properties)
			feature.Properties["id"] = loc.Id
			featureCollection.Append(feature)
		}
		aidh.SendJSON(w, 200, featureCollection)
	})

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
