package server

import "context"

type Server struct {
	// conf              *Config
	// logger            *zap.Logger
	// db                *gorm.DB
	grpcServer *grpc.Server
	// spannerClient     *spanner.Client
	// voicechatGRPCConn *grpc.ClientConn
}

type DurianServer struct {
	// TODO gRPCのメソッド群
	unaryEcho *unary_echo.UnaryEchoHandler
	// debugError *debug_error.DebugErrorHandler
}

func NewDurianServer() *DurianServer {

	s := &DurianServer{
		unaryEcho: unary_echo.NewUnaryEchoHandler(config),
		// debugError: debug_error.NewDebugErrorHandler(config),
	}

	return s
}

func (s *DurianServer) UnaryEcho(ctx context.Context, req *api.UnaryEchoRequest) (*api.UnaryEchoResponse, error) {
	return s.unaryEcho.UnaryEcho(ctx, req)
}

func (s *Server) StartApiServer() {

	//s.initCodec()
	// http serverをTLSで起動する

	// ListenAndServeTLSは引数handle関数を取るのでResponseWriter/requestを引数とする関数を渡す必要がある。


	return http.ListenAndServeTLS(
		addr,
		//cert
		//key

		// zapext.WithRequestIdは独自に定義した関数
		// http.handlerを受け取ってhandlerを返す、返却するhandlerでは
		zapext.WithNewRequestId(
			contenttype.Middleware(
				http.HandleFunc(s.handleApi())
			)
		),
	)

}


// func WithNewRequestId(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx := context.WithValue(r.Context(), RequestIdKey, newUnique())
// 		r = r.WithContext(ctx)
// 		h.ServeHTTP(w, r)
// 	})
// }
