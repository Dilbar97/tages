package upload

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"os"

	pb "tages/protocol/media"
)

type Client struct {
	client pb.UploadServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{client: pb.NewUploadServiceClient(conn)}
}

func (c Client) Upload(ctx context.Context, file string) (string, error) {
	stream, err := c.client.Upload(ctx)
	if err != nil {
		return "", err
	}

	fil, err := os.Open(file)
	if err != nil {
		return "", err
	}

	buf := make([]byte, 1024)

	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if err = stream.Send(&pb.UploadRequest{Chunk: buf[:num], FilePath: fil.Name()}); err != nil {
			return "", err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return res.GetName(), nil
}
