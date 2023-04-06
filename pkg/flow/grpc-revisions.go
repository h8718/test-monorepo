package flow

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/direktiv/direktiv/pkg/flow/bytedata"
	"github.com/direktiv/direktiv/pkg/flow/grpc"
)

func (flow *flow) Revisions(ctx context.Context, req *grpc.RevisionsRequest) (*grpc.RevisionsResponse, error) {
	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ns, err := flow.edb.NamespaceByName(ctx, req.GetNamespace())
	if err != nil {
		return nil, err
	}

	fStore, _, _, rollback, err := flow.beginSqlTx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(ctx)

	file, err := fStore.ForRootID(ns.ID).GetFile(ctx, req.GetPath())
	if err != nil {
		return nil, err
	}
	revs, err := fStore.ForFile(file).GetAllRevisions(ctx)
	if err != nil {
		return nil, err
	}

	revisions := []*grpc.Ref{}
	for idx, rev := range revs {
		if idx > 0 {
			revisions = append(revisions, &grpc.Ref{
				Name: rev.ID.String(),
			})
		}
	}

	resp := new(grpc.RevisionsResponse)
	resp.Namespace = ns.Name
	resp.Results = revisions
	resp.PageInfo = &grpc.PageInfo{
		Total: int32(len(resp.Results)),
	}
	resp.Node = bytedata.ConvertFileToGrpcNode(file)

	return resp, nil
}

func (flow *flow) RevisionsStream(req *grpc.RevisionsRequest, srv grpc.Flow_RevisionsStreamServer) error {
	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ctx := srv.Context()

	resp, err := flow.Revisions(ctx, req)
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

func (flow *flow) DeleteRevision(ctx context.Context, req *grpc.DeleteRevisionRequest) (*emptypb.Empty, error) {
	flow.sugar.Debugf("Handling gRPC request: %s", this())
	return nil, nil
	// TODO: yassir, implement
	/*
		tctx, tx, err := flow.database.Tx(ctx)
		if err != nil {
			return nil, err
		}
		defer rollback(tx)

		cached, err := flow.traverseToRef(tctx, req.GetNamespace(), req.GetPath(), req.GetRevision())
		if err != nil {
			return nil, err
		}

		if !cached.Ref.Immutable {
			return nil, errors.New("not a revision")
		}

		clients := flow.edb.Clients(tctx)

		query := clients.Ref.Query().Where(entref.HasRevisionWith(entrev.ID(cached.Revision.ID)), entref.Immutable(false))

		xrefs, err := query.All(tctx)
		if err != nil {
			return nil, err
		}

		if len(xrefs) > 1 || (len(xrefs) == 1 && xrefs[0].Name != "latest") {
			return nil, errors.New("cannot delete revision while refs to it exist")
		}

		if len(xrefs) == 1 && xrefs[0].Name == "latest" {
			err = flow.configureRouter(tctx, cached, rcfBreaking,
				func() error {
					err := clients.Ref.DeleteOneID(cached.Ref.ID).Exec(tctx)
					if err != nil {
						return err
					}

					return nil
				},
				tx.Commit,
			)
			if err != nil {
				return nil, err
			}
		} else {
			err = flow.configureRouter(tctx, cached, rcfBreaking,
				func() error {
					err := clients.Revision.DeleteOneID(cached.Revision.ID).Exec(tctx)
					if err != nil {
						return err
					}

					return nil
				},
				tx.Commit,
			)
			if err != nil {
				return nil, err
			}
		}

		flow.logger.Infof(ctx, cached.Workflow.ID, cached.GetAttributes(recipient.Workflow), "Deleted workflow revision: %s.", cached.Revision.ID.String())
		flow.pubsub.NotifyWorkflow(cached.Workflow)

		var resp emptypb.Empty

		return &resp, nil
	*/
}
