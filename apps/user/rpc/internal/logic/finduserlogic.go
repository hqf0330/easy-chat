package logic

import (
	"context"
	"easy-chat/apps/user/models"
	"github.com/jinzhu/copier"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	var userEntities []*models.Users
	var err error

	if in.Phone != "" {
		userEntity, err := l.svcCtx.UserModel.FindByPhone(l.ctx, in.Phone)
		if err == nil {
			userEntities = append(userEntities, userEntity)
		}
	} else if len(in.Ids) > 0 {
		userEntities, err = l.svcCtx.UserModel.ListByIds(l.ctx, in.Ids)
	} else if in.Name != "" {
		userEntities, err = l.svcCtx.UserModel.ListByName(l.ctx, in.Name)
	}

	if err != nil {
		return nil, err
	}

	var resp []*user.UserEntity
	_ = copier.Copy(&resp, userEntities)

	return &user.FindUserResp{
		User: resp,
	}, nil
}
