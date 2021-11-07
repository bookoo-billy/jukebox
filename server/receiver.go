package server

import (
	"context"
	"github.com/hajimehoshi/go-mp3"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"os"
	"time"

	"github.com/bookoo-billy/jukebox/db"
	v1 "github.com/bookoo-billy/jukebox/gen/api/v1"
)

const (
	Mp3SamplesPerChannel = 1152
	Mp3Channels = 2
	Mp3ByteDepth = 2
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
		logrus.Info("Sending play song request to receiver %s at address %s", receiver.Name, receiver.Url)
	}

	return &v1.ReceiversPlayResponse{Song: request.Song, Receivers: receivers}, nil
}

func (s *ReceiverServer) Stop(ctx context.Context, request *v1.ReceiversStopRequest) (*v1.ReceiversStopResponse, error) {
	receivers, err := s.dao.List(ctx, &v1.ReceiverQuery{})
	if err != nil {
		return nil, err
	}

	for _, receiver := range receivers.Items {
		logrus.Info("Sending stop song request to receiver %s at address %s", receiver.Name, receiver.Url)
	}

	return &v1.ReceiversStopResponse{Receivers: receivers}, nil
}

func (s *ReceiverServer) ReceiverChat(chatSrv v1.ReceiverService_ReceiverChatServer) error {

	for {
		err := s.sendSong(chatSrv)
		if err != nil {
			return err
		}
	}
}

func (s *ReceiverServer) sendSong(chatSrv v1.ReceiverService_ReceiverChatServer) error {
	logrus.Info("Decoding MP3 to PCM...")
	f, err := os.Open("test.mp3")
	if err != nil {
		logrus.WithError(err).Error("Failed to open mp3")
		return err
	}
	defer f.Close()

	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		logrus.WithError(err).Error("Failed to create decoder from mp3 file")
	}

	logrus.Info("Sending song header to receiver")
	err = chatSrv.Send(&v1.ReceiverCommandRequest{
		Command: &v1.ReceiverCommandRequest_PlaySongHeader{
			PlaySongHeader: &v1.PlaySongHeader{
				Song: &v1.Song{
					Name: "Our Father",
					Album: &v1.Album{
						Name: "III Sides to Every Story",
					},
					Artist: &v1.Artist{
						Name: "Extreme",
					},
				},
				Format:     v1.Format_JUKEBOX_S16,
				SampleRate: int32(decoder.SampleRate()),
				Channels:   uint32(Mp3Channels),
				Samples:    uint32(Mp3SamplesPerChannel),
			},
		},
	})
	if err != nil {
		logrus.WithError(err).Errorln("Receiver stopped accepting commands, shutting down stream")
		return err
	}

	_, err = chatSrv.Recv()
	if err != nil {
		logrus.WithError(err).Errorln("Failed to receive ack after sending song header to receiver")
		return err
	}

	logrus.Info("Sending song chunks to receiver...")
	buf := make([]byte, Mp3SamplesPerChannel * Mp3Channels * Mp3ByteDepth)

	startTime := time.Now()
	startTime.Add(2 * time.Second)
	count := 0

	for {
		length, decErr := decoder.Read(buf)
		playAt := startTime.Add(time.Duration((float64(Mp3SamplesPerChannel) / float64(decoder.SampleRate())) * float64(count)) * time.Second)

		playErr := chatSrv.Send(&v1.ReceiverCommandRequest{
			Command: &v1.ReceiverCommandRequest_PlaySongChunk{
				PlaySongChunk: &v1.PlaySongChunk{
					Chunk: buf,
					Size: int32(length),
					Timestamp: timestamppb.New(playAt),
				},
			},
		})
		if decErr != nil {
			if decErr != io.EOF {
				logrus.WithError(decErr).Error("Failed to read decoder buffer")
				return decErr
			} else {
				logrus.Info("Received EOF from mp3 decoder")
				break
			}
		}

		if playErr != nil {
			logrus.WithError(playErr).Errorln("Failed while writing chunk to receiver")
			return playErr
		}
		
		if length == 0 {
			logrus.Info("Received zero bytes from decoder, assuming finished?")
			break
		}
		count++
	}

	logrus.Info("Sending song trailer to receiver")
	err = chatSrv.Send(&v1.ReceiverCommandRequest{
		Command: &v1.ReceiverCommandRequest_PlaySongTrailer{
			PlaySongTrailer: &v1.PlaySongTrailer{},
		},
	})
	if err != nil {
		logrus.WithError(err).Errorln("Failed to send song trailer to receiver")
		return err
	}

	_, err = chatSrv.Recv()
	if err != nil {
		logrus.WithError(err).Errorln("Failed to receive ack while after sending song trailer to receiver")
		return err
	}

	return nil

}