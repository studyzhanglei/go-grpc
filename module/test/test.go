package test

import (
	"content-grpc/global"
	"content-grpc/model"
	"content-grpc/pb/search"
	"content-grpc/utils/errors"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	grpcError "content-grpc/pb/error"
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
		err = errors.NewFromCode(grpcError.SearchError_SEARCH_FAILED)

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
