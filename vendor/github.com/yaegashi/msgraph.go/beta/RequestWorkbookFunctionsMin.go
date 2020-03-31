// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsMinRequestBuilder struct{ BaseRequestBuilder }

// Min action undocumented
func (b *WorkbookFunctionsRequestBuilder) Min(reqObj *WorkbookFunctionsMinRequestParameter) *WorkbookFunctionsMinRequestBuilder {
	bb := &WorkbookFunctionsMinRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/min"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsMinRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsMinRequestBuilder) Request() *WorkbookFunctionsMinRequest {
	return &WorkbookFunctionsMinRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsMinRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsMinARequestBuilder struct{ BaseRequestBuilder }

// MinA action undocumented
func (b *WorkbookFunctionsRequestBuilder) MinA(reqObj *WorkbookFunctionsMinARequestParameter) *WorkbookFunctionsMinARequestBuilder {
	bb := &WorkbookFunctionsMinARequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/minA"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsMinARequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsMinARequestBuilder) Request() *WorkbookFunctionsMinARequest {
	return &WorkbookFunctionsMinARequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsMinARequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}