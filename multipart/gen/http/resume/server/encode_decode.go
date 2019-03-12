// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// resume HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/multipart/design -o
// $(GOPATH)/src/goa.design/examples/multipart

package server

import (
	"context"
	"net/http"

	resume "goa.design/examples/multipart/gen/resume"
	resumeviews "goa.design/examples/multipart/gen/resume/views"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// EncodeListResponse returns an encoder for responses returned by the resume
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(resumeviews.StoredResumeCollection)
		enc := encoder(ctx, w)
		body := NewStoredResumeResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeAddResponse returns an encoder for responses returned by the resume
// add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the resume add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload []*resume.Resume
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewResumeAddDecoder returns a decoder to decode the multipart request for
// the "resume" service "add" endpoint.
func NewResumeAddDecoder(mux goahttp.Muxer, resumeAddDecoderFn ResumeAddDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v interface{}) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(*[]*resume.Resume)
			if err := resumeAddDecoderFn(mr, p); err != nil {
				return err
			}
			return nil
		})
	}
}
