package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/redis.v3"
)

type Session interface {
	Create(value string) string
	Get(id string) (string, error)
	// Delete(id string) error
}

var (
	SessionFoundError = errors.New("Session not found")
	letterRunes       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

//세션값 랜덤생성
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	// fmt.Print(b)
	return string(b)
}

//메모리 기반 세션 매니저
type MemSession struct {
	DB map[string]string
}

func MemSessionNew() *MemSession {
	sess := &MemSession{DB: make(map[string]string)}
	return sess
}

func (m *MemSession) Create(v string) string {
	randStr := RandStringRunes(32)
	m.DB[randStr] = v
	return randStr
}

func (m *MemSession) Get(id string) (string, error) {
	if v, ok := m.DB[id]; ok {
		return v, nil
	}
	return "", SessionFoundError
}

//REDIS 기반 세션 매니저
type RedisSession struct {
	cli *redis.Client
}

func RedisSessionNew() *RedisSession {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return &RedisSession{cli: client}
}

func (r *RedisSession) Create(v string) string {
	randStr := RandStringRunes(32)
	r.cli.Set(randStr, v, 0)
	return randStr
}

func (r *RedisSession) Get(id string) (string, error) {
	v, err := r.cli.Get(id).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

//session provider
var SessionStore = make(map[string]Session)

func Register(name string, sess Session) error {
	if sess == nil {
		return errors.New("Session not found")
	}
	_, ok := SessionStore[name]
	if ok {
		return errors.New("Session exists")
	}
	SessionStore[name] = sess
	return nil
}

func GetSessionStore(name string) (Session, error) {
	sess, ok := SessionStore[name]
	if !ok {
		return nil, errors.New("Session store not found :" + name)
	}
	return sess, nil
}

func main() {
	sess := Session(MemSessionNew())
	sid := sess.Create("hello world")
	v, _ := sess.Get(sid)
	fmt.Println(sid, "->", v)

	redisSess := Session(RedisSessionNew())
	sid = redisSess.Create("myungsworld")
	v, err := redisSess.Get(sid)
	if err != nil {
		panic(err)
	}
	fmt.Println(sid, "->", v)
}
