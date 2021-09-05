package main

import (
	"context"
	"net/http"

	"../api/grpc/server"
	"../pkg/api"
)

func main() {
	ctx := context.Background()

	// 起動前にいろんな設定をしている
	// ログ,、時刻、SpannerClient 、SQLite
	// Interceptor
	//
	grpcServer := grpc.NewServer()
	opts := []grpc.ServerOption{
		grpc.WithInsecure(),
	}

	durianServer := server.NewDurianServer()
	// grpc serverのポインタとgrpcサーバが持っているgrpcリクエスト用のIFを渡す
	// grpcServerにprotoで定義されたメソッドやハンドラーの組み合わせを登録する
	api.RegisterDurianServer(grpcServer, durianServer)

	// Server構造体にgRPC ServerとかDB/SpannerClient/Configを突っ込んでいる
	// redisもここに入れるべきか。
	server := &server.Server{grpcServer: *grpc.Server}

	defer server.Close()

	// サーバ起動
	if err := server.StartApiServer(); err != http.ErrServerClosed {
		panic(err)
	}
}
