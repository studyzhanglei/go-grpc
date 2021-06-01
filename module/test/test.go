package test

import (
	"content-grpc/global"
	"content-grpc/model"
	"content-grpc/pb/search"
	"content-grpc/utils/errors"
	grpcError "content-grpc/pb/error"
	"encoding/json"
	"fmt"
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


	header, _ := json.Marshal(request.Header)
	global.LOG.Info(fmt.Sprintf("recive：%s", request))

	fmt.Println(string(request.Request), string(header))

	var user model.A
	err = global.DB.Where("name = ?", request.Request).First(&user).Error
	fmt.Println(user.Name)

	data, _ := global.REDIS.Get("a").Result()

	err = errors.NewFromCode(grpcError.SearchError_SEARCH_FAILED)

	err = srv.Send(&search.SearchResponse{
		Status: errors.GetResHeader(err),//一般情况下 异常了就不会有response了 这里是调试乱写的
		Response: user.Name + string(data),
	})

	if err != nil {
		global.LOG.Info(err.Error())
		return err
	}

	return nil
}
