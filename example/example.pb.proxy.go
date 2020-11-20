// Code generated by aws-grpc-service-proxy. DO NOT EDIT.

package proto

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"math"
	"net"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/golang/protobuf/jsonpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	o "github.com/katallaxie/protoc-gen-cloud-proxy/pkg/opts"
	"github.com/katallaxie/protoc-gen-cloud-proxy/pkg/proxy"
)

// SongJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SongJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var SongJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Song) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := SongJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Song)(nil)

// SongJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Song. This struct is safe to replace or modify but
// should not be done so concurrently.
var SongJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Song) UnmarshalJSON(b []byte) error {
	return SongJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Song)(nil)

// InsertJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of InsertJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var InsertJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := InsertJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert)(nil)

// InsertJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert. This struct is safe to replace or modify but
// should not be done so concurrently.
var InsertJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert) UnmarshalJSON(b []byte) error {
	return InsertJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert)(nil)

// Insert_RequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Insert_RequestJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_RequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert_Request) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Insert_RequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert_Request)(nil)

// Insert_RequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert_Request. This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_RequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert_Request) UnmarshalJSON(b []byte) error {
	return Insert_RequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert_Request)(nil)

// Insert_ResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Insert_ResponseJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_ResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert_Response) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Insert_ResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert_Response)(nil)

// Insert_ResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert_Response. This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_ResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert_Response) UnmarshalJSON(b []byte) error {
	return Insert_ResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert_Response)(nil)

// UpdateJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of UpdateJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var UpdateJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Update) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := UpdateJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Update)(nil)

// UpdateJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Update. This struct is safe to replace or modify but
// should not be done so concurrently.
var UpdateJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Update) UnmarshalJSON(b []byte) error {
	return UpdateJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Update)(nil)

// Update_RequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Update_RequestJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Update_RequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Update_Request) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Update_RequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Update_Request)(nil)

// Update_RequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Update_Request. This struct is safe to replace or modify but
// should not be done so concurrently.
var Update_RequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Update_Request) UnmarshalJSON(b []byte) error {
	return Update_RequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Update_Request)(nil)

// Update_ResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Update_ResponseJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Update_ResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Update_Response) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Update_ResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Update_Response)(nil)

// Update_ResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Update_Response. This struct is safe to replace or modify but
// should not be done so concurrently.
var Update_ResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Update_Response) UnmarshalJSON(b []byte) error {
	return Update_ResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Update_Response)(nil)

// ListSongsJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSongsJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongsJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSongs) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := ListSongsJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSongs)(nil)

// ListSongsJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSongs. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongsJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSongs) UnmarshalJSON(b []byte) error {
	return ListSongsJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSongs)(nil)

// ListSongs_RequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSongs_RequestJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongs_RequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSongs_Request) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := ListSongs_RequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSongs_Request)(nil)

// ListSongs_RequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSongs_Request. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongs_RequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSongs_Request) UnmarshalJSON(b []byte) error {
	return ListSongs_RequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSongs_Request)(nil)

// ListSongs_ResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSongs_ResponseJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongs_ResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSongs_Response) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := ListSongs_ResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSongs_Response)(nil)

// ListSongs_ResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSongs_Response. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongs_ResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSongs_Response) UnmarshalJSON(b []byte) error {
	return ListSongs_ResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSongs_Response)(nil)

// EmptyJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of EmptyJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var EmptyJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Empty) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := EmptyJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Empty)(nil)

// EmptyJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Empty. This struct is safe to replace or modify but
// should not be done so concurrently.
var EmptyJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Empty) UnmarshalJSON(b []byte) error {
	return EmptyJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Empty)(nil)

type srv struct {
	opts *o.Opts
}

type service struct {
	tlsCfg *tls.Config
	UnimplementedExampleServer
}

func New(opts *o.Opts) proxy.Listener {
	s := new(srv)
	s.opts = opts

	return s
}

func NewProxy(opts *o.Opts) proxy.Proxy {
	s := New(opts)
	p := proxy.New(s, opts)

	return p
}

func (s *srv) Start(ctx context.Context, ready func()) func() error {
	return func() error {
		lis, err := net.Listen("tcp", s.opts.Addr)
		if err != nil {
			return err
		}

		ll := s.opts.Logger.With(zap.String("addr", s.opts.Addr))
		srv := &service{}

		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
		srv.tlsCfg = tlsConfig

		var kaep = keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
			PermitWithoutStream: true,            // Allow pings even when there are no active streams
		}

		var kasp = keepalive.ServerParameters{
			MaxConnectionIdle:     time.Duration(math.MaxInt64), // If a client is idle for 15 seconds, send a GOAWAY
			MaxConnectionAge:      time.Duration(math.MaxInt64), // If any connection is alive for more than 30 seconds, send a GOAWAY
			MaxConnectionAgeGrace: 5 * time.Second,              // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
			Time:                  5 * time.Second,              // Ping the client if it is idle for 5 seconds to ensure the connection is still active
			Timeout:               1 * time.Second,              // Wait 1 second for the ping ack before assuming the connection is dead
		}

		ss := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
		RegisterExampleServer(ss, srv)
		// health.RegisterHealthServer(ss, srv)

		ready()

		ll.Info("start listening")

		if err := ss.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

// Here goes a message Insert

func (s *service) Insert(ctx context.Context, req *Insert_Request) (*Insert_Response, error) {
	b, err := req.MarshalJSON()
	if err != nil {
		return nil, err
	}

	svc := lambda.New(session.New())
	input := &lambda.InvokeInput{
		FunctionName: aws.String("arn:aws:lambda:eu-west-1:291339088935:function:Platform-testfunction5B23D3B0-4MTGLLV63WWF"),
		Payload:      b,
		Qualifier:    aws.String("$LATEST"),
	}

	result, err := svc.InvokeWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	var payload Insert_Response
	if err := payload.UnmarshalJSON(result.Payload); err != nil {
		return nil, err
	}

	return &payload, nil

}

// Here goes a message Update

func (s *service) Update(ctx context.Context, req *Update_Request) (*Update_Response, error) {

	svc := dynamodb.New(session.New())

	input := &dynamodb.UpdateItemInput{
		ReturnValues:     aws.String("ALL_NEW"),
		TableName:        aws.String("Music"),
		UpdateExpression: aws.String("SET #Y = :y, #AT = :t"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		return nil, err
	}

	// DynamoDB
	return nil, nil

}
