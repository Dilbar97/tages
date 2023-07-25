package upload

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rs/zerolog/log"
	"io"
	"sync/atomic"
	"tages/internal/server/service"
	"tages/internal/storage"
	pb "tages/protocol/media"
)

type Media struct {
	limit   atomic.Int32
	storage storage.Storage
	svc     service.MediaSvc
	pb.UnimplementedUploadServiceServer
}

func NewMediaServer(storage storage.Storage, svc service.MediaSvc) Media {
	return Media{storage: storage, svc: svc}
}

func (m *Media) Upload(stream pb.UploadService_UploadServer) error {
	inProgressCount := m.limit.Add(1)
	defer m.limit.Add(-1)

	if inProgressCount > 10 {
		return fmt.Errorf("upload limit hit")
	}

	file := storage.NewFile()

	for {
		req, err := stream.Recv()
		if file.GetName() == "" {
			file.SetName(req.FilePath)
		}
		if err == io.EOF {
			if err = m.storage.Store(file); err != nil {
				fmt.Println(fmt.Sprintf("store file err: %s", err.Error()))
				return err
			}

			if err = stream.SendAndClose(&pb.UploadResponse{Name: file.GetName()}); err != nil {
				fmt.Println(fmt.Sprintf("send res err: %s", err.Error()))
				return err
			}

			if _, err = m.svc.Upload(m.storage.Dir + file.GetName()); err != nil {
				fmt.Println(fmt.Sprintf("save file in db err: %s", err.Error()))
				return err
			}

			return nil
		}

		if err != nil {
			fmt.Println(fmt.Sprintf("receive chunk err: %s", err.Error()))
			return err
		}

		if err = file.Write(req.GetChunk()); err != nil {
			fmt.Println(fmt.Sprintf("write chunk to buffer err: %s", err.Error()))
			return err
		}
	}
}

func (m *Media) List(ctx context.Context, r *pb.ListRequest) (*pb.ListResponse, error) {
	var resp pb.ListResponse

	inProgressCount := m.limit.Add(1)
	defer m.limit.Add(-1)

	if inProgressCount > 100 {
		return &resp, fmt.Errorf("list limit hit")
	}

	mediaList, err := m.svc.List(ctx)
	if err != nil {
		log.Error().Msgf("server.Media.List: %v", err)
		resp.Error = "Error on files list"

		return &resp, err
	}

	for _, media := range mediaList {
		resp.Data = append(resp.Data, &pb.Media{
			Name:      media.Name,
			CreatedAt: &timestamp.Timestamp{Seconds: media.CreatedAt.Unix()},
			UpdatedAt: &timestamp.Timestamp{Seconds: media.UpdatedAt.Unix()},
		})
	}

	return &resp, nil
}

func (m *Media) GetMedia(ctx context.Context, r *pb.GetMediaRequest) (*pb.GetMediaResponse, error) {
	var resp pb.GetMediaResponse

	inProgressCount := m.limit.Add(1)
	defer m.limit.Add(-1)

	if inProgressCount > 10 {
		return &resp, fmt.Errorf("list limit hit")
	}

	mediaData, err := m.svc.GetMedia(ctx, r.Name)
	if err != nil {
		log.Error().Msgf("server.Media.GetMedia: %v", err)
		resp.Error = "Error on get file"

		return &resp, err
	}

	resp.Data = &pb.Media{
		Name:      mediaData.Name,
		CreatedAt: &timestamp.Timestamp{Seconds: mediaData.CreatedAt.Unix()},
		UpdatedAt: &timestamp.Timestamp{Seconds: mediaData.UpdatedAt.Unix()},
	}

	return &resp, nil
}
