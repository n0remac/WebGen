package main

//go:generate npx buf generate

import (
	"MyApp/gen/proto/user/userconnect"
	"MyApp/pkg/database"
	"MyApp/pkg/service"
	"MyApp/pkg/user"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewLogInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err != nil {
				// log.Error("connect error", "error", fmt.Sprintf("%+v", err))
				// TODO breadchris this should only be done for local dev
				fmt.Printf("%+v\n", err)
			}
			return resp, err
		}
	}
	return interceptor
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: method=%s, path=%s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	database.InitDB()
	imageServer := http.FileServer(http.Dir("./pkg/images"))

	interceptors := connect.WithInterceptors(NewLogInterceptor())

	apiRoot := http.NewServeMux()
	apiRoot.Handle("/images/", http.StripPrefix("/images/", imageServer))
	reflectorServices := service.SetupServices(apiRoot, interceptors)

	userService := &user.UserService{} //TODO generate this
	apiRoot.Handle(userconnect.NewUserServiceHandler(userService, interceptors))

	reflectorServices = append(reflectorServices, "user.userService")
	// TODO move reflector to serviceRegistry so generated code can access it
	reflector := grpcreflect.NewStaticReflector(
		reflectorServices...,
	)

	recoverCall := func(_ context.Context, spec connect.Spec, _ http.Header, p any) error {
		log.Printf("panic: %+v\n", p)
		if err, ok := p.(error); ok {
			return err
		}
		return fmt.Errorf("panic: %v", p)
	}
	apiRoot.Handle(grpcreflect.NewHandlerV1(reflector, connect.WithRecover(recoverCall)))
	// Many tools still expect the older version of the server reflection Service, so
	// most servers should mount both handlers.
	apiRoot.Handle(grpcreflect.NewHandlerV1Alpha(reflector, connect.WithRecover(recoverCall)))

	// Serve the frontend build
	fs := http.FileServer(http.Dir("./frontend/dist/site"))
	apiRoot.Handle("/static/", fs)

	// Handle fallback to index.html for SPA
	apiRoot.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling fallback for path: %s\n", r.URL.Path)
		if r.URL.Path != "/" {
			_, err := os.Stat("./frontend/dist/site" + r.URL.Path)
			if os.IsNotExist(err) {
				http.ServeFile(w, r, "./frontend/dist/site/index.html")
				return
			}
		}
		fs.ServeHTTP(w, r)
	})

	addr := fmt.Sprintf(":%d", 8080)

	log.Printf("starting http server on %s\n", addr)

	http.ListenAndServe(addr, h2c.NewHandler(corsMiddleware(apiRoot), &http2.Server{}))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO breadchris how bad is this? lol
		origin := r.Header.Get("Origin")

		// TODO breadchris this should only be done for local dev
		w.Header().Set("Access-Control-Allow-Origin", origin)

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, connect-protocol-version")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
