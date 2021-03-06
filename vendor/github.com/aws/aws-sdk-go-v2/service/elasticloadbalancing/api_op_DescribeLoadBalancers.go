// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeLoadBalancers.
type DescribeLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	// The names of the load balancers.
	LoadBalancerNames []string `type:"list"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `type:"string"`

	// The maximum number of results to return with this call (a number from 1 to
	// 400). The default is 400.
	PageSize *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s DescribeLoadBalancersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLoadBalancersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLoadBalancersInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the parameters for DescribeLoadBalancers.
type DescribeLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	// Information about the load balancers.
	LoadBalancerDescriptions []LoadBalancerDescription `type:"list"`

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s DescribeLoadBalancersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLoadBalancers = "DescribeLoadBalancers"

// DescribeLoadBalancersRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the specified the load balancers. If no load balancers are specified,
// the call describes all of your load balancers.
//
//    // Example sending a request using DescribeLoadBalancersRequest.
//    req := client.DescribeLoadBalancersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeLoadBalancers
func (c *Client) DescribeLoadBalancersRequest(input *DescribeLoadBalancersInput) DescribeLoadBalancersRequest {
	op := &aws.Operation{
		Name:       opDescribeLoadBalancers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeLoadBalancersInput{}
	}

	req := c.newRequest(op, input, &DescribeLoadBalancersOutput{})
	return DescribeLoadBalancersRequest{Request: req, Input: input, Copy: c.DescribeLoadBalancersRequest}
}

// DescribeLoadBalancersRequest is the request type for the
// DescribeLoadBalancers API operation.
type DescribeLoadBalancersRequest struct {
	*aws.Request
	Input *DescribeLoadBalancersInput
	Copy  func(*DescribeLoadBalancersInput) DescribeLoadBalancersRequest
}

// Send marshals and sends the DescribeLoadBalancers API request.
func (r DescribeLoadBalancersRequest) Send(ctx context.Context) (*DescribeLoadBalancersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLoadBalancersResponse{
		DescribeLoadBalancersOutput: r.Request.Data.(*DescribeLoadBalancersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeLoadBalancersRequestPaginator returns a paginator for DescribeLoadBalancers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeLoadBalancersRequest(input)
//   p := elasticloadbalancing.NewDescribeLoadBalancersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeLoadBalancersPaginator(req DescribeLoadBalancersRequest) DescribeLoadBalancersPaginator {
	return DescribeLoadBalancersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeLoadBalancersInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeLoadBalancersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeLoadBalancersPaginator struct {
	aws.Pager
}

func (p *DescribeLoadBalancersPaginator) CurrentPage() *DescribeLoadBalancersOutput {
	return p.Pager.CurrentPage().(*DescribeLoadBalancersOutput)
}

// DescribeLoadBalancersResponse is the response type for the
// DescribeLoadBalancers API operation.
type DescribeLoadBalancersResponse struct {
	*DescribeLoadBalancersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLoadBalancers request.
func (r *DescribeLoadBalancersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
