package mock_greet

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/didiyudha/go-grpc-playground/unary/greeting/greet"

	"github.com/golang/mock/gomock"

	"github.com/golang/protobuf/proto"
)

type rpcMessage struct {
	msg proto.Message
}

func (rpm *rpcMessage) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, rpm.msg)
}

func (rpm *rpcMessage) String() string {
	return fmt.Sprintf("is %s", rpm.msg)
}

func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := NewMockGreeterClient(ctrl)
	greetRequest := greet.GreeetRequest{Name: "Didi"}
	mockGreeterClient.EXPECT().SayHello(gomock.Any(), &rpcMessage{msg: &greetRequest}).Return(&greet.GreeetResponse{Msg: "Hi, Didi nice to meet you"}, nil)
	testSayHello(t, mockGreeterClient)
}

func testSayHello(t *testing.T, client greet.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &greet.GreeetRequest{Name: "Didi"})
	if err != nil {
		t.Error(err)
	}
	if err != nil || r.Msg != "Hi, Didi nice to meet you" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply: ", r.Msg)
}
