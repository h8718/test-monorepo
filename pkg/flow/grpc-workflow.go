package flow

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/direktiv/direktiv/pkg/refactor/core"

	"github.com/direktiv/direktiv/pkg/refactor/filestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/direktiv/direktiv/pkg/flow/bytedata"
	"github.com/direktiv/direktiv/pkg/flow/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (flow *flow) ResolveWorkflowUID(ctx context.Context, req *grpc.ResolveWorkflowUIDRequest) (*grpc.WorkflowResponse, error) {
	// TODO: yassir, question this controller.
	flow.sugar.Debugf("Handling gRPC request: %s", this())

	var resp *grpc.WorkflowResponse

	return resp, nil
}

func (flow *flow) Workflow(ctx context.Context, req *grpc.WorkflowRequest) (*grpc.WorkflowResponse, error) {
	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ns, err := flow.edb.NamespaceByName(ctx, req.GetNamespace())
	if err != nil {
		return nil, err
	}
	fStore, _, commit, rollback, err := flow.beginSqlTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(ctx)

	file, err := fStore.ForRootID(ns.ID).GetFile(ctx, req.GetPath())
	if err != nil {
		return nil, err
	}
	revision, err := fStore.ForFile(file).GetCurrentRevision(ctx)
	if err != nil {
		return nil, err
	}
	dataReader, err := fStore.ForRevision(revision).GetData(ctx)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(dataReader)
	if err != nil {
		return nil, err
	}
	if err = commit(ctx); err != nil {
		return nil, err
	}

	var resp *grpc.WorkflowResponse
	resp.Namespace = ns.Name
	resp.Node = bytedata.ConvertFileToGrpcNode(file)
	resp.Revision = bytedata.ConvertRevisionToGrpcRevision(revision)
	resp.Revision.Source = data
	resp.EventLogging = ""
	resp.Oid = file.ID.String()

	return resp, nil
}

func (flow *flow) WorkflowStream(req *grpc.WorkflowRequest, srv grpc.Flow_WorkflowStreamServer) error {
	flow.sugar.Debugf("Handling gRPC request: %s", this())
	ctx := srv.Context()

	resp, err := flow.Workflow(ctx, req)
	if err != nil {
		return err
	}
	// mock streaming response.
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err = srv.Send(resp)
			if err != nil {
				return err
			}
			time.Sleep(time.Second * 5)
		}
	}
}

func (flow *flow) CreateWorkflow(ctx context.Context, req *grpc.CreateWorkflowRequest) (*grpc.CreateWorkflowResponse, error) {
	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ns, err := flow.edb.NamespaceByName(ctx, req.GetNamespace())
	if err != nil {
		return nil, err
	}
	fStore, _, commit, rollback, err := flow.beginSqlTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(ctx)

	file, revision, err := fStore.ForRootID(ns.ID).CreateFile(ctx, req.GetPath(), filestore.FileTypeWorkflow, bytes.NewReader(req.GetSource()))
	if err != nil {
		return nil, err
	}
	dataReader, err := fStore.ForRevision(revision).GetData(ctx)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(dataReader)
	if err != nil {
		return nil, err
	}

	if err = commit(ctx); err != nil {
		return nil, err
	}

	// TODO: yassir, needs fix here.
	//metricsWf.WithLabelValues(cached.Namespace.Name, cached.Namespace.Name).Inc()
	//metricsWfUpdated.WithLabelValues(cached.Namespace.Name, path, cached.Namespace.Name).Inc()
	//
	//flow.logger.Infof(ctx, cached.Namespace.ID, cached.GetAttributes(recipient.Namespace), "Created workflow '%s'.", path)
	//
	//err = flow.BroadcastWorkflow(ctx, BroadcastEventTypeCreate,
	//	broadcastWorkflowInput{
	//		Name:   resp.Node.Name,
	//		Path:   resp.Node.Path,
	//		Parent: resp.Node.Parent,
	//		Live:   true,
	//	}, cached)
	//
	//if err != nil {
	//	return nil, err
	//}

	resp := &grpc.CreateWorkflowResponse{}
	resp.Namespace = ns.Name
	resp.Node = bytedata.ConvertFileToGrpcNode(file)
	resp.Revision = bytedata.ConvertRevisionToGrpcRevision(revision)
	resp.Revision.Source = data

	return resp, nil
}

func (flow *flow) UpdateWorkflow(ctx context.Context, req *grpc.UpdateWorkflowRequest) (*grpc.UpdateWorkflowResponse, error) {
	// This is being called by the frontend when a user changes a workflow via a UI and press save button.

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ns, err := flow.edb.NamespaceByName(ctx, req.GetNamespace())
	if err != nil {
		return nil, err
	}

	fStore, _, commit, rollback, err := flow.beginSqlTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(ctx)

	file, err := fStore.ForRootID(ns.ID).GetFile(ctx, req.GetPath())
	if err != nil {
		return nil, err
	}
	if file.Typ != filestore.FileTypeWorkflow {
		return nil, status.Error(codes.InvalidArgument, "file type is not workflow")
	}

	newRevision, err := fStore.ForFile(file).CreateRevision(ctx, "", bytes.NewReader(req.GetSource()))
	if err != nil {
		return nil, err
	}
	dataReader, err := fStore.ForRevision(newRevision).GetData(ctx)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(dataReader)
	if err != nil {
		return nil, err
	}

	if err = commit(ctx); err != nil {
		return nil, err
	}
	var resp grpc.UpdateWorkflowResponse

	resp.Namespace = ns.Name
	resp.Node = bytedata.ConvertFileToGrpcNode(file)
	resp.Revision = bytedata.ConvertRevisionToGrpcRevision(newRevision)
	resp.Revision.Source = data

	return &resp, nil
}

func (flow *flow) SaveHead(ctx context.Context, req *grpc.SaveHeadRequest) (*grpc.SaveHeadResponse, error) {
	// This is being called by the UI when a user clicks create revision button.

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ns, err := flow.edb.NamespaceByName(ctx, req.GetNamespace())
	if err != nil {
		return nil, err
	}

	fStore, _, commit, rollback, err := flow.beginSqlTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(ctx)

	file, err := fStore.ForRootID(ns.ID).GetFile(ctx, req.GetPath())
	if err != nil {
		return nil, err
	}
	if file.Typ != filestore.FileTypeWorkflow {
		return nil, status.Error(codes.InvalidArgument, "file type is not workflow")
	}
	revision, err := fStore.ForFile(file).GetCurrentRevision(ctx)
	if err != nil {
		return nil, err
	}
	dataReader, err := fStore.ForRevision(revision).GetData(ctx)
	if err != nil {
		return nil, err
	}
	_, err = fStore.ForFile(file).CreateRevision(ctx, "", dataReader)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(dataReader)
	if err != nil {
		return nil, err
	}

	if err = commit(ctx); err != nil {
		return nil, err
	}

	var resp grpc.SaveHeadResponse

	resp.Namespace = ns.Name
	resp.Node = bytedata.ConvertFileToGrpcNode(file)
	resp.Revision = bytedata.ConvertRevisionToGrpcRevision(revision)
	resp.Revision.Source = data

	return &resp, nil
}

func (flow *flow) DiscardHead(ctx context.Context, req *grpc.DiscardHeadRequest) (*grpc.DiscardHeadResponse, error) {
	// This is being called by the UI when a user clicks 'revert' button.

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ns, err := flow.edb.NamespaceByName(ctx, req.GetNamespace())
	if err != nil {
		return nil, err
	}

	fStore, _, commit, rollback, err := flow.beginSqlTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(ctx)

	file, err := fStore.ForRootID(ns.ID).GetFile(ctx, req.GetPath())
	if err != nil {
		return nil, err
	}
	if file.Typ != filestore.FileTypeWorkflow {
		return nil, status.Error(codes.InvalidArgument, "file type is not workflow")
	}
	revision, err := fStore.ForFile(file).GetCurrentRevision(ctx)
	if err != nil {
		return nil, err
	}
	dataReader, err := fStore.ForRevision(revision).GetData(ctx)
	if err != nil {
		return nil, err
	}
	//_, err = fStore.ForFile(file).RevertRevision(ctx, "")
	//if err != nil {
	//	return nil, err
	//}
	data, err := io.ReadAll(dataReader)
	if err != nil {
		return nil, err
	}

	if err = commit(ctx); err != nil {
		return nil, err
	}

	var resp grpc.DiscardHeadResponse

	resp.Namespace = ns.Name
	resp.Node = bytedata.ConvertFileToGrpcNode(file)
	resp.Revision = bytedata.ConvertRevisionToGrpcRevision(revision)
	resp.Revision.Source = data

	return &resp, nil
}

func (flow *flow) ToggleWorkflow(ctx context.Context, req *grpc.ToggleWorkflowRequest) (*emptypb.Empty, error) {
	// TODO: yassir, question this controller.
	flow.sugar.Debugf("Handling gRPC request: %s", this())

	return &emptypb.Empty{}, nil
}

func (flow *flow) SetWorkflowEventLogging(ctx context.Context, req *grpc.SetWorkflowEventLoggingRequest) (*emptypb.Empty, error) {
	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ns, err := flow.edb.NamespaceByName(ctx, req.GetNamespace())
	if err != nil {
		return nil, err
	}

	fStore, store, commit, rollback, err := flow.beginSqlTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(ctx)

	file, err := fStore.ForRootID(ns.ID).GetFile(ctx, req.GetPath())
	if err != nil {
		return nil, err
	}

	annotations, err := store.FileAnnotations().Get(ctx, file.ID)

	if err == core.ErrFileAnnotationsNotSet {
		annotations = &core.FileAnnotations{
			FileID: file.ID,
			Data:   nil,
		}
	} else if err != nil {
		return nil, err
	}

	annotations.Data = annotations.Data.SetEntry("workflow_log_event_key", req.GetLogger())

	err = store.FileAnnotations().Set(ctx, annotations)
	if err != nil {
		return nil, err
	}

	if err := commit(ctx); err != nil {
		return nil, err
	}
	var resp emptypb.Empty

	return &resp, nil
}
