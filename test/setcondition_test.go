package test

import (
	"testing"
	"net/http"
	"net/url"
	"github.com/cjwddz/fast-model"
	"github.com/bmizerany/assert"
)

func TestSetCondition(t *testing.T){
	r:=&http.Request{}
	r.PostForm = url.Values{}
	r.PostForm.Add("name","stupi")
	r.PostForm.Add("id","123")
	r.PostForm.Add("valid","TrUe")
	r.PostForm.Add("pos","10")
	r.PostForm.Add("len","50")

	cond:= model.DbSetCondition{}
	assert.Equal(t,cond.GetSetCondStr(),"")
	assert.Equal(t,len(cond.GetSetCondParams()),0)

	cond=cond.And("like","name","cjwddz").Limit(11,2)
	cond=cond.Set2(r,"s_name").Set("valid",false)
	cstr:=cond.GetSetCondStr()
	assert.Equal(t,cstr," SET name=$1,valid=$2 WHERE name like $3  limit $4 offset $5")
	assert.Equal(t,cond.GetSetCondParams()[4],2)
	cond=cond.Reset()
	assert.Equal(t,cond.GetSetCondStr(),"")
	assert.Equal(t,len(cond.GetSetCondParams()),0)
}
