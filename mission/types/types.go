package types

import (
	"image"
	"time"
)

// mission 테이블 객체
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
	ModelId        int          `json:"modelId"`
	Model          MissionModel `json:"model"`
}

// mission_model 테이블 객체
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

// mission_comment 테이블 객체
type MissionComment struct {
	MissionCommentId int       `json:"missionCommentId"`
	UserId           int       `json:"userId"`
	MissionId        int       `json:"missionId"`
	Content          string    `json:"content"`
	Regdate          time.Time `json:"regdate"`
}

// fam_mission 테이블 객체
type FamMission struct {
	FamId     int       `json:"famId"`
	MissionId int       `json:"missionId"`
	UserId    int       `json:"userId"`
	FamilyId  int       `json:"familyId"`
	Regdate   time.Time `json:"regdate"`
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
