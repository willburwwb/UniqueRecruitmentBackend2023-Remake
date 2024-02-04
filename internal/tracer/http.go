package tracer

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xylonx/zapx"
	"go.opentelemetry.io/otel/attribute"
)

func TracingMiddleware(ctx *gin.Context) {
	apmCtx, span := Tracer.Start(ctx.Request.Context(), ctx.Request.Method+" "+ctx.Request.RequestURI)
	defer span.End()

	zapx.WithContext(apmCtx).Info("start tracing")

	for key := range ctx.Request.Header {
		span.SetAttributes(attribute.String("http.request."+strings.ToLower(key), ctx.Request.Header.Get(key)))
	}

	ctx.Request = ctx.Request.WithContext(apmCtx)
	ctx.Next()

	for key := range ctx.Writer.Header() {
		span.SetAttributes(attribute.String("http.response."+strings.ToLower(key), ctx.Writer.Header().Get(key)))
	}
}
