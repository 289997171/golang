package model

import "strconv"

type User struct {
	username string
	password string
	age      int32
}

func (this *User) String() string {
	return "username: " + this.username + " password: " + this.password + " age:" + strconv.Itoa(int(this.age))
}

type str32 [32]byte

type User2 struct {
	Username str32
	Password str32
}

func strcpy(dest *str32, src string, n int) {
	maxn := len(dest) - 1
	if n > maxn {
		n = maxn
	}

	for i := 0; i < n; i++ {
		dest[i] = src[i]
	}
}

func (this *str32) StrcpyN(src string, n int) {
	strcpy(this, src, n)
}

func (this *str32) Strcpy(src string) {
	strcpy(this, src, len(src))
}

func (this *str32) String() string {
	return string(this[:])
}

