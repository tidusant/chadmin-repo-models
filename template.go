package models

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Template struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Code         string        `bson:"code"`
	UserID       string        `bson:"userid"`
	Status       int           `bson:"status"` //-2: delete, -1: reject, 1: approved and publish, 2: pending, 3: approved but not publish
	Title        string        `bson:"title"`
	Description  string        `bson:"description"`
	Viewed       int           `bson:"viewed"`
	InstalledIDs []string      `bson:"installedid"`
	ActiveIDs    []string      `bson:"activedid"`
	Avatar       string        `bson:"avatar"`
	Created      time.Time     `bson:"created"`
	Modified     time.Time     `bson:"modified"`
	Content      string        `bson:"content"`
	Pages        string        `bson:"pages"`
	CSS          string        `bson:"css"`
	Script       string        `bson:"script"`
	Images       string        `bson:"images"`
	Fonts        string        `bson:"fonts"`
	Screenshot   string        `bson:"screenshot"`
	Configs      string        `bson:"configs"`
	Langs        string        `bson:"langs"`
}

//News ...
type TemplateConfig struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	TemplateCode string        `bson:"templatecode"`
	ShopID       string        `bson:"shopid"`
	Key          string        `bson:"key"`
	Type         string        `bson:"type"`
	Value        string        `bson:"value"`
}

type TemplateLang struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	TemplateCode string        `bson:"templatecode"`
	Lang         string        `bson:"lang"`
	ShopID       string        `bson:"shopid"`
	Key          string        `bson:"key"`
	Value        string        `bson:"value"`
}

type TemplateViewData struct {
	PageName     string
	Siteurl      string
	Data         map[string]json.RawMessage
	TemplatePath string
	Templateurl  string
	Imageurl     string
	Pages        map[string]string
	Resources    map[string]string
	Configs      map[string]string
}
