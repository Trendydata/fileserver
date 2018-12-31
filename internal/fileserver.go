package internal

import (
	"context"
	"io"
	"io/ioutil"
	"os"

	"github.com/Trendydata/fileserver/pkg"
)

type FileServer struct {
}

func (s *FileServer) Upload(ctx context.Context, chunk *pkg.Chunk) (*pkg.UploadResponse, error) {
	err := ioutil.WriteFile(chunk.GetMeta().GetName(), chunk.Data, os.ModePerm|os.ModeAppend)
	return &pkg.UploadResponse{}, err
}

func (s *FileServer) UploadStream(stream pkg.File_UploadStreamServer) error {
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			err := stream.SendAndClose(&pkg.UploadResponse{})
			return err
		}
		if err != nil {
			return err
		}
		file, err := os.OpenFile(chunk.GetMeta().GetName(), os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.Write(chunk.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
