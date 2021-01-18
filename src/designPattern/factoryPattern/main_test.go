package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionStore(t *testing.T) {
	err := Register("mem", MemSessionNew())
	assert.Nil(t, err)
	err = Register("redis", RedisSessionNew())
	assert.Nil(t, err)

	// 존재하지 않는 프로바이더 호출로 인한 에러 출력
	_, err = GetSessionStore("mongo")
	assert.NotNil(t, err)
	sess, err := GetSessionStore("mem")
	assert.Nil(t, err)

	sid := sess.Create("hello world")
	v, err := sess.Get(sid)
	assert.Nil(t, err)
	assert.Equal(t, "hello world", v)
}
