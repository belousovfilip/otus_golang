package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users []User

func getUsers(r io.Reader) ([]User, error) {
	result := make([]User, 0, 100_000)
	reader := bufio.NewReader(r)
	var p fastjson.Parser
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		if v, e := p.Parse(line); e == nil {
			result = append(result, User{
				Email: string(v.GetStringBytes("Email")),
			})
		}
		if err == io.EOF {
			return result, nil
		}
	}
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	domain = "." + domain
	for _, user := range u {
		if !strings.HasSuffix(user.Email, domain) {
			continue
		}
		i := strings.Index(user.Email, "@")
		if i == -1 {
			continue
		}
		d := user.Email[i+1:]
		d = strings.ToLower(d)
		result[d]++
	}
	return result, nil
}
