package hardrive

import architecture "github.com/rcondongfox/golang-architecture"

type hd map[int]architecture.User

func (hd hd) Save(n int, u architecture.User) {
	hd[n] = u
}

func (hd hd) Retrieve(n int) architecture.User {
	return hd[n]
}