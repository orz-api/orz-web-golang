package orzweb

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orz-api/orz-base-golang"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (def *ApiDef[Req, Rsp]) Gin(router gin.IRouter) {
	var handler = func(context *gin.Context) {
		ginProcessRequest(context, def)
	}
	if def.Query {
		router.PUT(def.buildPath(), handler)
	} else {
		router.POST(def.buildPath(), handler)
	}
}

func ginProcessRequest[Req any, Rsp any](context *gin.Context, definition *ApiDef[Req, Rsp]) {
	req := new(Req)
	if err := context.ShouldBind(req); err != nil {
		// TODO: log bad request error
		ginResponseHttpError(context, err, http.StatusBadRequest)
	} else {
		if rsp, err := definition.Handler(req); err != nil {
			var apiError *ApiError
			if isApiError := errors.As(err, &apiError); isApiError {
				if errorDef, errorDefExists := definition.Errors[apiError.Code]; errorDefExists {
					// TODO: log errorDef.Loggable is true
					// TODO: alarm errorDef.Alarm is true
					ginResponseApiError(context, err, NewErrorProtocol(nil, apiError.Code, errorDef.Notice), errorDef.Reason, nil)
				} else {
					// TODO: log undefined error
					reason := fmt.Sprintf("error undefined: code(`%s`)", apiError.Code)
					ginResponseApiError(context, err, NewErrorProtocol(nil, "", ""), reason, nil)
				}
			} else {
				// TODO: log unknown error
				ginResponseHttpError(context, err, http.StatusInternalServerError)
			}
		} else {
			ginResponseSuccess(context, rsp)
		}
	}
}

func ginResponseSuccess[Rsp any](context *gin.Context, rsp *Rsp) {
	context.Header(PropsObj.ResponseHeader.Version, strconv.Itoa(VERSION_CURRENT))
	context.JSON(http.StatusOK, rsp)
}

func ginResponseApiError(context *gin.Context, error error, protocol *ProtocolBo, reason string, extraTraces []*ErrorTraceTo) {
	var exposeReason = ""
	if PropsObj.ExposeErrorReason {
		exposeReason = reason
	}

	var exposeTraces []*ErrorTraceTo
	if PropsObj.ExposeErrorTraces {
		exposeTraces = ginGetTraces(context, error, extraTraces)
	}

	context.Header(PropsObj.ResponseHeader.Version, strconv.Itoa(VERSION_CURRENT))
	context.Header(PropsObj.ResponseHeader.Code, protocol.Code)
	if protocol.Notice != "" {
		context.Header(PropsObj.ResponseHeader.Notice, url.QueryEscape(protocol.Notice))
	}

	if exposeReason != "" || len(exposeTraces) != 0 {
		context.JSON(http.StatusOK, &ErrorRsp{
			Code:   protocol.Code,
			Reason: exposeReason,
			Traces: exposeTraces,
		})
	} else {
		context.Status(http.StatusOK)
	}
}

func ginResponseHttpError(context *gin.Context, error error, status int) {
	var exposeReason string
	if PropsObj.ExposeErrorReason {
		exposeReason = error.Error()
	}

	var exposeTraces []*ErrorTraceTo
	if PropsObj.ExposeErrorTraces {
		exposeTraces = ginGetTraces(context, error, nil)
	}

	if exposeReason != "" || len(exposeTraces) != 0 {
		context.JSON(status, &ErrorRsp{
			Reason: exposeReason,
			Traces: exposeTraces,
		})
	} else {
		context.Status(status)
	}
}

func ginGetEndpoint(context *gin.Context) string {
	endpoint := context.Request.RequestURI
	if PropsObj.BaseUri != "" && strings.HasPrefix(endpoint, PropsObj.BaseUri) {
		endpoint = endpoint[len(PropsObj.BaseUri):]
	}
	return endpoint
}

func ginGetTraces(context *gin.Context, error error, extraTraces []*ErrorTraceTo) []*ErrorTraceTo {
	var traces []*ErrorTraceTo
	traces = append(traces, &ErrorTraceTo{
		Service:  orzbase.PropsObj.Service,
		Endpoint: ginGetEndpoint(context),
		Details:  error.Error(),
	})
	if len(extraTraces) != 0 {
		traces = append(traces, extraTraces...)
	}
	return traces
}
