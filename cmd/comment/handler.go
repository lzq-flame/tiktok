package main

import (
	"context"
	"example/cmd/comment/kitex_gen/comment"
	"example/cmd/comment/pack"
	"example/cmd/comment/service"
	"example/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// AddComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) AddComment(ctx context.Context, req *comment.AddCommentRequest) (resp *comment.AddCommentResp, err error) {
	resp = new(comment.AddCommentResp)
	err = service.NewAddCommentService(ctx).AddCommentService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DelComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DelComment(ctx context.Context, req *comment.DelCommentRequest) (resp *comment.DelCommentResp, err error) {
	resp = new(comment.DelCommentResp)
	err = service.NewDelCommentService(ctx).DelCommentService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetCommentCount implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentCount(ctx context.Context, req *comment.GetCommentCountRequest) (resp *comment.GetCommentCountResp, err error) {
	resp = new(comment.GetCommentCountResp)
	count := service.NewCommentCountService(ctx).CommentCountService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.CommentCount = count
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *comment.GetCommentListRequest) (resp *comment.GetCommentListResp, err error) {
	resp = new(comment.GetCommentListResp)
	commentList, err := service.NewGetCommentListService(ctx).GetCommentList(req)
	if err != nil {
		resp.CommentList = nil
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.CommentList = commentList
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
