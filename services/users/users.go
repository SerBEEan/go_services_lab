package users

import (
	"github.com/patrickmn/go-cache"
)

const USERS_LIST_KEY = "USERS_LIST_KEY"
const LAST_USER_ID_KEY = "LAST_USER_ID_KEY"

type users struct {
	cache *cache.Cache
}

type userData struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type userInfo struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Users interface {
	GetUsers() []userData
	CreateUser(info userInfo) int
}

func (u *users) GetUsers() []userData {
	var users = []userData{}

	value, isFound := u.cache.Get(USERS_LIST_KEY)
	if isFound {
		users = value.([]userData)
	}

	return users
}

func (u *users) CreateUser(info userInfo) int {
	var lastId int = -1
	var users []userData

	id, isIdFound := u.cache.Get(LAST_USER_ID_KEY)
	if isIdFound {
		lastId = id.(int)
	}

	users = u.GetUsers()

	lastId++
	users = append(users, userData{
		Id:       lastId,
		Name:     info.Name,
		Login:    info.Login,
		Password: info.Password,
	})

	u.cache.Set(USERS_LIST_KEY, users, cache.NoExpiration)
	u.cache.Set(LAST_USER_ID_KEY, lastId, cache.NoExpiration)

	return lastId
}

func NewUsers() Users {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	var list = []userData{}

	c.Set(USERS_LIST_KEY, list, cache.NoExpiration)
	c.Set(LAST_USER_ID_KEY, -1, cache.NoExpiration)

	return &users{
		cache: c,
	}
}
