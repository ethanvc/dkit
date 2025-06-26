package httpcli

import "context"

type HttpClient struct {
}

func (cli *HttpClient) Json(c context.Context, req *HttpReq)
