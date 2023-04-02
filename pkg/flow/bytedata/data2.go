package bytedata

import (
	"path/filepath"

	"github.com/direktiv/direktiv/pkg/flow/grpc"
	"github.com/direktiv/direktiv/pkg/refactor/filestore"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertFileToGrpcNode(f *filestore.File) *grpc.Node {
	return &grpc.Node{
		CreatedAt:  timestamppb.New(f.CreatedAt),
		UpdatedAt:  timestamppb.New(f.UpdatedAt),
		Name:       filepath.Base(f.Path),
		Path:       f.Path,
		Parent:     filepath.Dir(f.Path),
		Type:       string(f.Typ),
		Attributes: []string{},
		Oid:        f.ID.String(),
		ReadOnly:   false,
	}
}

func ConvertFilesToGrpcNodeList(list []*filestore.File) []*grpc.Node {
	var result []*grpc.Node
	for _, f := range list {
		result = append(result, ConvertFileToGrpcNode(f))
	}

	return result
}

func ConvertRevisionToGrpcRevision(rev *filestore.Revision) *grpc.Revision {
	return &grpc.Revision{
		Name: rev.ID.String(),
	}
}

func ConvertRevisionsToGrpcRevisionList(list []*filestore.Revision) []*grpc.Revision {
	var result []*grpc.Revision
	for _, f := range list {
		result = append(result, ConvertRevisionToGrpcRevision(f))
	}

	return result
}

func ConvertRevisionToGrpcRef(rev *filestore.Revision) *grpc.Ref {
	return &grpc.Ref{
		Name: rev.ID.String(),
	}
}

func ConvertRevisionsToGrpcRefList(list []*filestore.Revision) []*grpc.Ref {
	var result []*grpc.Ref
	for _, f := range list {
		result = append(result, ConvertRevisionToGrpcRef(f))
	}

	return result
}
