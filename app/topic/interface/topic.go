package ITopicService

import (
	m "github.com/og/goclub/app/model"
	vd "github.com/og/juice/validator"
)

type Interface interface {
	CreateTopic(req ReqCreateTopic) (topicID m.IDTopic, reject error)
	UpdateTopicIntro(req ReqUpdateTopicIntro) (reject error)
	UpdateTopicCoverPhoto(req ReqUpdateTopicCoverPhoto) (reject error)
}
type ReqCreateTopic struct {
	Title string `json:"title"`
	CoverPhoto string `json:"coverPhoto"`
	Intro string `json:"intro"`
}

func (v ReqCreateTopic) VD(r *vd.Rule) {
	r.String(v.Title, vd.StringSpec{
		Name: "标题",
		MinRuneLen: 1,
		MaxRuneLen: 50,
	})
	r.String(v.CoverPhoto, vd.StringSpec{
		Name: "图片",
		AllowEmpty: true,
	})
	r.String(v.Intro, vd.StringSpec{
		Name: "介绍",
		AllowEmpty: true,
	})
}

type ReqUpdateTopicIntro struct {
	TopicID m.IDTopic `json:"topicID"`
	Intro string `json:"intro"`
}

func (v ReqUpdateTopicIntro) VD(r *vd.Rule) {
	r.String(v.TopicID.String(), vd.StringSpec{})
	r.String(v.Intro, vd.StringSpec{
		Name: "介绍",
		AllowEmpty: true,
	})
}
type ReqUpdateTopicCoverPhoto struct {
	TopicID m.IDTopic `json:"topicID"`
	CoverPhoto string `json:"coverPhoto"`
}

func (v ReqUpdateTopicCoverPhoto) VD(r *vd.Rule) {
	r.String(v.TopicID.String(), vd.StringSpec{})
	r.String(v.CoverPhoto, vd.StringSpec{
		Name: "封面图片",
		AllowEmpty: true,
	})
}