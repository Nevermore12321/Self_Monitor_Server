package global_middleware

import (
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
)

func CorssDomain(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("global middleware")
		//  获取 请求的 method 和 Headers 中的 Origin
		method := r.Method
		requestOrigin := r.Header.Get("Origin")
		logx.Info(method, requestOrigin)
		//  判断 请求的 来源 ， 请求 Headers 中的 Origin 表示请求来源
		//  如果有来源，则回response， 否则视为 不安全请求，不给予 跨域认证
		if requestOrigin != "" {
			//  取到 客户端 request 的Origin，也就是 客户端的地址，并且设置为 Access-Control-Allow-Origin 头部
			//  ctx.Header() ctx.Writer.Header().Set() 作用一一样
			w.Header().Set("Access-Control-Allow-Origin", requestOrigin)

			//  Access-Control-Allow-Method 表示跨域请求允许的方法
			w.Header().Set("Access-Control-Allow-Method", "POST, GET, PUT, DELETE, UPDATE, OPTIONS")

			//  Access-Control-Allow-Headers 表示跨域请求的Header中允许带的字段
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, Content-Type, X-Csrf-Token, X-Xsrf-Token,Token, Session")

			//  Access-Control-Expose-Headers  表示客户端（浏览器） 可以解析出来的头部Header
			w.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Origin, Content-Length, X-Csrf-Token")

			//  Access-Control-Allow-Credentials   表示允许客户端传递验证信息（COokies）
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			//  Access-Control-Max-Age  表示 设置缓存时间,单位秒 (24小时)
			w.Header().Set("Access-Control-Max-Age", "86400")
		}
		if method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
