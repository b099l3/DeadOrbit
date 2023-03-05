package main

import (
	"context"
	"net/http"
	"time"

	testv1 "github.com/b099l3/dead-orbit/deadorbit/proto/test/v1"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type thingServiceApi struct {
	Things    []*testv1.Thing
	CurrentId int32
}

func main() {

	thingsServer := testv1.NewThingServiceApiServer(&thingServiceApi{})

	mux := http.NewServeMux()
	mux.Handle(thingsServer.PathPrefix(), thingsServer)

	err := http.ListenAndServe(":8000", thingsServer)
	if err != nil {
		panic(err)
	}
}

func (s *thingServiceApi) CreateThing(ctx context.Context, params *testv1.CreateThingRequest) (*testv1.CreateThingResponse, error) {
	if len(params.Text) < 4 {
		return nil, twirp.InvalidArgument.Error("Text should be min 4 characters.")
	}

	thing := &testv1.Thing{
		Id:        s.CurrentId,
		Text:      params.Text,
		CreatedAt: timestamppb.New(time.Now()),
	}

	s.Things = append(s.Things, thing)

	s.CurrentId++

	return &testv1.CreateThingResponse{
		Thing: thing,
	}, nil
}

func (s *thingServiceApi) GetThings(ctx context.Context, params *testv1.GetThingsRequest) (*testv1.GetThingsResponse, error) {
	allThings := make([]*testv1.Thing, 0)
	allThings = append(allThings, s.Things...)

	return &testv1.GetThingsResponse{
		Things: allThings,
	}, nil
}
