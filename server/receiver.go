package server

import (
	"context"

	"github.com/bookoo-billy/jukebox/db"
	v1 "github.com/bookoo-billy/jukebox/gen/api/v1"
	"github.com/golang/glog"
)

type ReceiverServer struct {
	v1.UnimplementedReceiverServiceServer
	dao *db.ReceiverDAO
}

func NewReceiverServer(dao *db.ReceiverDAO) v1.ReceiverServiceServer {
	return &ReceiverServer{dao: dao}
}

func (s *ReceiverServer) List(ctx context.Context, query *v1.ReceiverQuery) (*v1.ReceiverList, error) {
	return s.dao.List(ctx, query)
}

func (s *ReceiverServer) Create(ctx context.Context, request *v1.ReceiverCreateRequest) (*v1.Receiver, error) {
	return s.dao.Create(ctx, request.Receiver)
}

func (s *ReceiverServer) Get(ctx context.Context, query *v1.ReceiverQuery) (*v1.Receiver, error) {
	return s.dao.Get(ctx, query)
}

func (s *ReceiverServer) Update(ctx context.Context, request *v1.ReceiverUpdateRequest) (*v1.Receiver, error) {
	return s.dao.Update(ctx, request.Receiver)
}

func (s *ReceiverServer) Delete(ctx context.Context, query *v1.ReceiverQuery) (*v1.Receiver, error) {
	return s.dao.Delete(ctx, query)
}

func (s *ReceiverServer) Play(ctx context.Context, request *v1.ReceiversPlayRequest) (*v1.ReceiversPlayResponse, error) {
	receivers, err := s.dao.List(ctx, &v1.ReceiverQuery{})
	if err != nil {
		return nil, err
	}

	for _, receiver := range receivers.Items {
		glog.Info("Sending play song request to receiver %s at address %s", receiver.Name, receiver.Url)
	}

	return &v1.ReceiversPlayResponse{Song: request.Song, Receivers: receivers}, nil
}

func (s *ReceiverServer) Stop(ctx context.Context, request *v1.ReceiversStopRequest) (*v1.ReceiversStopResponse, error) {
	receivers, err := s.dao.List(ctx, &v1.ReceiverQuery{})
	if err != nil {
		return nil, err
	}

	for _, receiver := range receivers.Items {
		glog.Info("Sending stop song request to receiver %s at address %s", receiver.Name, receiver.Url)
	}

	return &v1.ReceiversStopResponse{Receivers: receivers}, nil
}

func (s *ReceiverServer) ReceiverChat(chatSrv v1.ReceiverService_ReceiverChatServer) error {
	for {
		glog.Info("Sending message to receiver")
		chatSrv.Send(&v1.ReceiverCommandRequest{})
		glog.Info("Sent message to receiver, awaiting response....")

		receiverResp, err := chatSrv.Recv()
		if err != nil {
			glog.Errorln("Failed while receiving response from receiver")
			glog.Error(err)
			continue
		}
		glog.Info("Received message from receiver")
		glog.Info(receiverResp)
	}
}
