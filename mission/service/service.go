package service

import (
	"github.com/nomaltree/family-eeum/mission/types"
	"log"
)

//mission, missionmodel, fammission, missioncomment에 관련된 서비스함수 작성 페이지

type Service interface {
	InsertFamMission(familyId int, userId int, missionId int) (int, error)
	GetFamMission(familyId int, userId int, missionId int) ([]int, error)
	GetModelMission(missionId int) (types.MissionModel, error)
	GetAllOtherMissions(userId int) ([]types.Mission, error)
	GetAllUserMissions(userId int, missionType string) ([]types.Mission, error)
	InsertMission(mission types.Mission) (int, error)
	UpdateMissionComplete(missionId int) error
	GetFamilyMission(familyId int, userId int) ([]types.Mission, error)
	DeleteMission(missionId int) error
	DeleteOneFamMission(missionId int, userId int) error
	GetOneMission(missionId int) (types.Mission, error)
	GetAllFamilyMission(familyId int) ([]types.Mission, error)
	UpdateMission(mission types.Mission) error
	SearchMissions(search string, missionType string) ([]types.Mission, error)
	IsNotPub(missionId int) error
	IsPub(missionId int) error
	InsertMissionComment(comment types.MissionComment) (int, error)
	DeleteMissionComment(comment types.MissionComment) error
	GetOneMissionComments(missionId int) ([]types.MissionComment, error)
	GetOneUserMissionComments(userId int) ([]types.MissionComment, error)
	GetMissionModels(missiontypes []string) ([]types.MissionModel, error)
}

type Repository interface {
	InsertFamMission(familyId int, userId int, missionId int) (int, error)
	GetFamMission(familyId int, userId int, missionId int) ([]int, error)
	GetModelMission(missionId int) (types.MissionModel, error)
	GetAllOtherMissions(userId int) ([]types.Mission, error)
	GetAllUserMissions(userId int, missionType string) ([]types.Mission, error)
	InsertMission(mission types.Mission) (int, error)
	UpdateMissionComplete(missionId int) error
	GetFamilyMission(familyId int, userId int) ([]types.Mission, error)
	DeleteMission(missionId int) error
	DeleteFamMission(missionId int) error
	DeleteOneFamMission(missionId int, userId int) error
	GetOneMission(missionId int) (types.Mission, error)
	GetAllFamilyMission(familyId int) ([]types.Mission, error)
	UpdateMission(mission types.Mission) error
	SearchMissions(search string, missionType string) ([]types.Mission, error)
	UpComments(missionId int) error
	DownComments(missionId int) error
	UpViews(missionId int) error
	UpMissioncopy(missionId int) error
	IsNotPub(missionId int) error
	IsPub(missionId int) error
	InsertMissionComment(comment types.MissionComment) (int, error)
	DeleteMissionComment(missionCommentId int) error
	GetOneMissionComments(missionId int) ([]types.MissionComment, error)
	GetOneUserMissionComments(userId int) ([]types.MissionComment, error)
	GetOneMissionComment(missioncommentId int) (types.MissionComment, error)
	GetMissionModels(missiontypes []string) ([]types.MissionModel, error)
}

type service struct {
	r Repository
}

func NewMissionService(r Repository) Service {
	return &service{r}
}

// 미션 추천 알고리즘 메소드(각 가족구성원의 성향이 담긴 배열을 매개변수로 받음)
func (s *service) GetMissionModels(missionTypes []string) ([]types.MissionModel, error) {
	var D = 0 //주도형: 건강 추천
	var I = 0 //사교형: 일상 추천
	var S = 0 //안정형: 마음챙김 추천
	var C = 0 //신중형: 성취 추천
	for _, missionType := range missionTypes {
		switch missionType {
		case "D":
			D = D + 1
		case "I":
			I = I + 1
		case "S":
			S = S + 1
		case "C":
			C = C + 1
		default:
			continue
		}
	}
	var recommendTypes []string
	var DISC []int
	DISC = append(DISC, D)
	DISC = append(DISC, I)
	DISC = append(DISC, S)
	DISC = append(DISC, C)

	var manyDISC [4]bool
	var many = 0
	for i := 0; i < 3; i++ {
		if DISC[many] > DISC[i+1] {
			manyDISC[many] = true
		}
		if DISC[many] < DISC[i+1] {
			many = i + 1
			manyDISC[i+1] = true
			for j := 0; j < i+1; j++ {
				manyDISC[j] = false
			}
		}
		if DISC[many] == DISC[i+1] {
			manyDISC[many] = true
			manyDISC[i+1] = true
		}
	}
	if (D == 0 && I == 0) && (S == 0 && C == 0) {
		manyDISC[0] = false
		manyDISC[1] = false
		manyDISC[2] = false
		manyDISC[3] = false
	}
	if manyDISC[0] == true {
		recommendTypes = append(recommendTypes, "건강")
	}
	if manyDISC[1] == true {
		recommendTypes = append(recommendTypes, "일상")
	}
	if manyDISC[2] == true {
		recommendTypes = append(recommendTypes, "마음챙김")
	}
	if manyDISC[3] == true {
		recommendTypes = append(recommendTypes, "성취")
	}
	recommendTypes = append(recommendTypes, "취미") // 취미는 공통적으로 모두 추천
	models, err := s.r.GetMissionModels(recommendTypes)
	if err != nil {
		return nil, err
	}
	return models, nil
}

// 특정 미션에 참여하는 가족 구성원 정보를 fammission테이블에 등록하는 메소드
func (s *service) InsertFamMission(familyId int, userId int, missionId int) (int, error) {
	id, err := s.r.InsertFamMission(familyId, userId, missionId)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

// 특정 미션에 참여하는 가족 구성원 정보를 조회하는 메소드
func (s *service) GetFamMission(familyId int, userId int, missionId int) ([]int, error) {
	familys, err := s.r.GetFamMission(familyId, userId, missionId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return familys, nil
}

// 사용자가 미션 모델을 담아가는 메소드
func (s *service) GetModelMission(modelId int) (types.MissionModel, error) {
	model, err := s.r.GetModelMission(modelId)
	if err != nil {
		log.Fatal(err)
		return types.MissionModel{}, err
	}
	err = s.r.UpMissioncopy(modelId)
	if err != nil {
		log.Fatal(err)
		return types.MissionModel{}, err
	}
	return model, nil
}

// 사용자가 다른 사용자의 미션을 조회하는 메소드(이 경우 공개로 설정된 것만 가져옴)
func (s *service) GetAllOtherMissions(userId int) ([]types.Mission, error) {
	missions, err := s.r.GetAllOtherMissions(userId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return missions, nil
}

// 내가 참여하는 모든 미션 조회 메소드
func (s *service) GetAllUserMissions(userId int, missionType string) ([]types.Mission, error) {
	missions, err := s.r.GetAllUserMissions(userId, missionType)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return missions, nil
}

// 미션을 등록하는 메소드
func (s *service) InsertMission(mission types.Mission) (int, error) {
	id, err := s.r.InsertMission(mission)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

// 완료한 미션을 완료상태로 바꾸는 메소드
func (s *service) UpdateMissionComplete(missionId int) error {
	err := s.r.UpdateMissionComplete(missionId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 나를 제외한 내 가족이 등록한 미션을 조회하는 쿼리(가족 미션 조회 쿼리)
func (s *service) GetFamilyMission(familyId int, userId int) ([]types.Mission, error) {
	missions, err := s.r.GetFamilyMission(familyId, userId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return missions, nil
}

// 참여 대기 상태인 미션을 삭제하는 메소드
func (s *service) DeleteMission(missionId int) error {
	err := s.r.DeleteMission(missionId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = s.r.DeleteFamMission(missionId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 특정 미션에 특정 구성원이 탈퇴하는 메소드
func (s *service) DeleteOneFamMission(missionId int, userId int) error {
	err := s.r.DeleteOneFamMission(missionId, userId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// id로 특정 미션 정보를 조회하는 메소드
func (s *service) GetOneMission(missionId int) (types.Mission, error) {
	mission, err := s.r.GetOneMission(missionId)
	if err != nil {
		log.Fatal(err)
		return types.Mission{}, err
	}
	err = s.r.UpViews(missionId)
	if err != nil {
		log.Fatal(err)
	}
	return mission, nil
}

// 가족 미션들을 조회하는 메소드
func (s *service) GetAllFamilyMission(familyId int) ([]types.Mission, error) {
	missions, err := s.r.GetAllFamilyMission(familyId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return missions, nil
}

// 수정된 미션의 정보를 업데이트 하는 메소드
func (s *service) UpdateMission(mission types.Mission) error {
	err := s.r.UpdateMission(mission)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 미션 검색 메소드
func (s *service) SearchMissions(search string, missionType string) ([]types.Mission, error) {
	missions, err := s.r.SearchMissions(search, missionType)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return missions, nil
}

// 미션을 비공개로 설정하는 메소드
func (s *service) IsNotPub(missionId int) error {
	err := s.r.IsNotPub(missionId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 미션을 공개로 설정하는 메소드
func (s *service) IsPub(missionId int) error {
	err := s.r.IsPub(missionId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 미션 댓글을 등록하는 메소드
func (s *service) InsertMissionComment(comment types.MissionComment) (int, error) {
	id, err := s.r.InsertMissionComment(comment)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	err = s.r.UpComments(comment.MissionId)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

// 미션 댓글을 삭제하는 메소드
func (s *service) DeleteMissionComment(comment types.MissionComment) error {
	err := s.r.DownComments(comment.MissionId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = s.r.DeleteMissionComment(comment.MissionCommentId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 해당 미션의 댓글을 불러오는 메소드
func (s *service) GetOneMissionComments(missionId int) ([]types.MissionComment, error) {
	comments, err := s.r.GetOneMissionComments(missionId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return comments, nil
}

// 사용자가 쓴 미션 댓글들을 불러오는 메소드
func (s *service) GetOneUserMissionComments(userId int) ([]types.MissionComment, error) {
	comments, err := s.r.GetOneUserMissionComments(userId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return comments, nil
}
