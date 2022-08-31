package service

import (
	"errors"
	"strings"
)

const StrMaxSize = 20

type StringRequest struct {
	A string
	B string
}

type Service interface {
	Concat(req StringRequest, ret *string) error
	Diff(req StringRequest, ret *string) error
}

type StringService struct {
}

func (s StringService) Concat(req StringRequest, ret *string) error {
	if len(req.A)+len(req.B) > StrMaxSize {
		*ret = ""
		return errors.New("MaxSize error")
	}
	*ret = req.A + req.B
	return nil
}

func (s StringService) Diff(req StringRequest, ret *string) error {
	if len(req.A) < 1 || len(req.B) < 1 {
		*ret = ""
		return nil
	}
	res := ""
	if len(req.A) < len(req.B) {
		req.A, req.B = req.B, req.A
	}
	for _, char := range req.B {
		if strings.Contains(req.A, string(char)) {
			res += string(char)
		}
	}
	*ret = res
	return nil
}
