package test

import (
	"testing"
	"fmt"
	h "libraries/mock/hello"
	"github.com/golang/mock/gomock"
	mock "libraries/mock/hello/mock"
)

func TestTalk(t *testing.T) {

	fmt.Println()
	p := h.NewPerson()
	c := h.NreComP(p)
	c.Metting(" l  ")

}

func TestMockTalk(t *testing.T) {
	ctl:=gomock.NewController(t)
	defer ctl.Finish()
	mocktalk:=mock.NewMockTalker(ctl)
	mocktalk.EXPECT().SayHello(gomock.Eq("zjx"))

	conp :=h.NreComP(mocktalk)

	t.Log( conp.Metting("zjx"))

}
