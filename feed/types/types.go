package types

import (
	"image"
	"time"
)

//familyeeum DB의 feed 관련 테이블 객체 타입 선언 페이지
//feed, feed comment, feed bookmark, feed_tag 총 4개 테이블에 대한 구조체 선언

// feed 테이블 객체
type Feed struct {
	Userid       int           `json:"userId"`
	User         User          `json:"user"`
	MissionId    int           `json:"missionId"`
	Familyid     int           `json:"familyId"`
	FeedId       int           `json:"feedId"`
	Views        int           `json:"views"`
	Comments     int           `json:"comments"`
	Content      string        `json:"content"`
	FeedImage    string        `json:"feedImage"`
	Pictures     []image.Image `json:"pictures"`
	Likes        int           `json:"likes"`
	LikesUser    []User        `json:"likesUser"`
	Regdate      time.Time     `json:"regdate"`
	Pub          int           `json:"pub"`
	FeedComments []FeedComment `json:"feedComments"`
	FeedTags     []string      `json:"feedTags"`
}

// feed_comment 테이블 객체
type FeedComment struct {
	Feedcommentid int       `json:"feedCommentId"`
	Userid        int       `json:"userId"`
	User          User      `json:"user"`
	Feedid        int       `json:"feedId"`
	Content       string    `json:"content"`
	Regdate       time.Time `json:"regdate"`
	MissionId     int       `json:"missionId"`
}

// feed_bookmark 테이블 객체
type FeedBookmark struct {
	FeedBookmarkId int       `json:"feedBookmarkId"`
	UserId         int       `json:"userId"`
	FeedId         int       `json:"feedId"`
	Regdate        time.Time `json:"regdate"`
}

// feed_tag 테이블 객체
type FeedTag struct {
	TagId   int    `json:"tagId"`
	FeedId  int    `json:"feedId"`
	Content string `json:"content"`
}

type User struct {
	UserId       int         `json:"userId"`
	Name         string      `json:"name"`
	ProfileName  string      `json:"profileName"`
	ProfileImage image.Image `json:"profileImage"`
	Intro        string      `json:"intro"`
	Regdate      time.Time   `json:"regdate"`
	Birth        time.Time   `json:"birth"`
	Role         string      `json:"role"`
}

type Mission struct {
	MissionId      int          `json:"missionId"`
	UserId         int          `json:"userId"`
	User           User         `json:"user"`
	FamilyId       int          `json:"familyId"`
	Candidates     []User       `json:"candidates"`
	MissionCycle   string       `json:"missionCycle"`
	MissionType    string       `json:"missionType"`
	MissionTypeId  int          `json:"missionTypeId"`
	MissionCycleId int          `json:"missionCycleId"`
	MissionTermId  int          `json:"missionTermId"`
	MissionName    string       `json:"missionName"`
	MissionImage   string       `json:"missionImage"`
	MissionTerm    string       `json:"missionTerm"`
	MissionStart   string       `json:"missionStart"`
	MissionEnd     string       `json:"missionEnd"`
	Regdate        time.Time    `json:"regdate"`
	Complete       bool         `json:"complete"`
	MissionContent string       `json:"missionContent"`
	Pub            bool         `json:"pub"`
	Views          int          `json:"views"`
	Comments       int          `json:"comments"`
	Model          MissionModel `json:"model"`
}

type MissionModel struct {
	UserId         int         `json:"userId"`
	User           User        `json:"user"`
	ModelId        int         `json:"modelId"`
	MissionName    string      `json:"missionName"`
	MissionCycle   string      `json:"missionCycle"`
	MissionType    string      `json:"missionType"`
	MissionImage   string      `json:"missionImage"`
	MissionTerm    string      `json:"missionTerm"`
	MissionContent string      `json:"missionContent"`
	MissionCopy    int         `json:"missionCopy"`
	ImageName      string      `json:"imageName"`
	Image          image.Image `json:"image"`
}
