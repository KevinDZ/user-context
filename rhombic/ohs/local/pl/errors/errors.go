package errors

import (
	"log"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BadRequest bad request
func BadRequest(errCode, message string, details ...interface{}) error {
	log.Printf("bad request: %v, %v", errCode, message)
	st := status.New(codes.InvalidArgument, message)
	errInfo := &errdetails.ErrorInfo{
		Reason: errCode,
	}
	ds, e := st.WithDetails(errInfo)
	if e != nil {
		return ds.Err()
	}
	return st.Err()
}
