package wrapper

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"homework_ipl/utils/errors"
	"homework_ipl/utils/httputils"
	"homework_ipl/utils/logger"

	"github.com/go-chi/chi/v5"
)

var (
	validationErr = errors.HttpError{
		Code:    http.StatusBadRequest,
		Message: "invalid request data",
	}

	decodingErr = errors.HttpError{
		Code:    http.StatusBadRequest,
		Message: "json decoding error",
	}

	encodingErr = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "json encoding error",
	}
)

type Wrapper[T Validator, Resp any] struct {
	ServeHTTP func(ctx context.Context, parsedRequest T) (Resp, error)
}

type Validator interface {
	Validate() error
}

func (w *Wrapper[T, Resp]) HandlerWrapper(resWriter http.ResponseWriter, httpReq *http.Request) {
	ctx := httpReq.Context()
	logger := logger.Logger()

	pathParams := GetPathParams(httpReq)
	queryParams := GetQueryParams(httpReq)
	ctx = SetPathParamsToCtx(ctx, pathParams)
	ctx = SetQueryParamsToCtx(ctx, queryParams)
	ctx = context.WithValue(ctx, httputils.ResponseWriterKey, resWriter)
	ctx = context.WithValue(ctx, httputils.HttpRequestKey, httpReq)

	limitedReader := io.LimitReader(httpReq.Body, 1_000_000)

	var requestData T
	if httpReq.ContentLength > 0 {
		err := json.NewDecoder(limitedReader).Decode(&requestData)
		if err != nil {
			logger.Error("Error decoding request body", "error", err)
			errors.WriteHttpError(decodingErr, resWriter)
			return
		}

		if err = requestData.Validate(); err != nil {
			logger.Error("Validation error", "error", err)
			errors.WriteHttpError(validationErr, resWriter)
			return
		}
	}

	response, httpErr := w.ServeHTTP(ctx, requestData)
	if httpErr != nil {
		logger.Error("Handler error", "error", httpErr)
		errors.WriteHttpError(httpErr, resWriter)
		return
	}

	rawJSON, err := json.Marshal(response)
	if err != nil {
		logger.Error("Error encoding response", "error", err)
		errors.WriteHttpError(encodingErr, resWriter)
		return
	}

	resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	resWriter.WriteHeader(http.StatusOK)
	_, _ = resWriter.Write(rawJSON)
}

func GetPathParams(r *http.Request) map[string]string {
	params := chi.RouteContext(r.Context()).URLParams
	pathParams := make(map[string]string)
	for k := len(params.Keys) - 1; k >= 0; k-- {
		key := params.Keys[k]
		value := params.Values[k]
		pathParams[key] = value
	}
	return pathParams
}

func GetQueryParams(r *http.Request) map[string]string {
	queryParams := r.URL.Query()
	params := make(map[string]string)

	for key, values := range queryParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	return params
}

func SetPathParamsToCtx(ctx context.Context, pathParams map[string]string) context.Context {
	return context.WithValue(ctx, httputils.RequestPathParamsKey, pathParams)
}

func SetQueryParamsToCtx(ctx context.Context, queryParams map[string]string) context.Context {
	return context.WithValue(ctx, httputils.RequestQueryParamsKey, queryParams)
}

func GetPathParamsFromCtx(ctx context.Context) map[string]string {
	pathParams, ok := ctx.Value(httputils.RequestPathParamsKey).(map[string]string)
	if !ok {
		return nil
	}
	return pathParams
}

func GetQueryParamsFromCtx(ctx context.Context) map[string]string {
	queryParams, ok := ctx.Value(httputils.RequestQueryParamsKey).(map[string]string)
	if !ok {
		return nil
	}
	return queryParams
}
