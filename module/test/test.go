package test

import (
	"content-grpc/global"
	"content-grpc/model"
	"content-grpc/utils/errors"
	"context"
	"encoding/json"
	"fmt"
	"github.com/studyzhanglei/grpc-proto/pb/search"
	"github.com/studyzhanglei/grpc-proto/pb/exception"
	"google.golang.org/grpc"
)

func Init(g *grpc.Server) {
	s := &SearchServiceServer{}
	search.RegisterSearchServiceServer(g, s)
}


type SearchServiceServer struct{}

func (s *SearchServiceServer) Search(srv search.SearchService_SearchServer) error {
	request, err := srv.Recv()
	if err != nil {
		global.LOG.Info(err.Error())
		return err
	}


	ctx := context.Background()

	header, _ := json.Marshal(request.Header)
	global.LOG.Info(fmt.Sprintf("recive：%s", request))

	fmt.Println(string(request.Request), string(header), 3333)

	fmt.Println(global.CONFIG.Mysql.Password, 7777)

	var user model.A
	err = global.DB.Where("name = ?", request.Request).First(&user).Error

	if err != nil {
		err = errors.NewFromCode(exception.SearchError_SEARCH_FAILED)

		err = srv.Send(&search.SearchResponse{
			Status: errors.GetResHeader(err),
		})

		return err
	}

	tx := global.DB.Begin()
	data, _ := global.REDIS.Get(ctx, "a").Result()
	fmt.Println(data)

	//time.Sleep(2 * time.Second)
	tx.Commit()

	str, err := json.Marshal(user)

	err = srv.Send(&search.SearchResponse{
		Status: errors.GetResHeader(err),//一般情况下 异常了就不会有response了 这里是调试乱写的
		Response: string(str),
	})

	if err != nil {
		global.LOG.Info(err.Error())
		return err
	}

	return nil
}


func (s *SearchServiceServer) GetUserInfo(srv search.SearchService_GetUserInfoServer) error {
	request, err := srv.Recv()
	if err != nil {
		global.LOG.Info(err.Error())
		return err
	}


	global.LOG.Info(fmt.Sprintf("recive：%s", request))


	var user model.A
	err = global.DB.Where("id = ?", request.Uid).First(&user).Error

	if err != nil {
		return err
	}

	err = srv.Send(&search.UserInfoResponse{
		Status: errors.GetResHeader(err),
		Username: user.Name,
		Ud: uint64(user.ID),
	})

	return err
}