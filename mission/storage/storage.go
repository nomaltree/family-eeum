package storage

import (
	"github.com/nomaltree/family-eeum/mission/types"
	"log"
)

//mission, missionmodel, fammission, missioncomment에 관련된 쿼리함수 작성 페이지

// 미션 모델테이블에서 미션들을 가져오는 메소드(미션 추천 메소드)
func (s *Storage) GetMissionModels(missionTypes []string) ([]types.MissionModel, error) {
	var model types.MissionModel
	var models []types.MissionModel
	for _, missionType := range missionTypes {
		iter, err := s.db.Query(`SELECT model_id, mission_name, mission_content, mission_cycle, mission_term, mission_type, model_image, user_id FROM mission_model_view where mission_type=$1 AND del=false order by random() limit 5`, missionType)
		if err != nil {
			log.Fatal(err)
		}
		for iter.Next() {
			if err := iter.Scan(&model.ModelId, &model.MissionName, &model.MissionContent, &model.MissionCycle, &model.MissionTerm, &model.MissionType, &model.ImageName, &model.UserId); err != nil {
				log.Fatal("Error while trying to get DB: ", err)
			}
			user, err := s.GetOneUser(model.UserId)
			if err != nil {
				log.Fatal(err)
			}
			model.User = user
			models = append(models, model)
		}
		if err = iter.Close(); err != nil {
			log.Fatal(err)
		}
	}
	return models, nil
}

// 특정 미션에 참여하는 가족 구성원 정보를 fammission테이블에 등록하는 메소드
func (s *Storage) InsertFamMission(familyId int, userId int, missionId int) (int, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO fam_mission (user_id, family_id, mission_id) VALUES ($1, $2, $3) RETURNING fam_id`, userId, familyId, missionId).Scan(&id)
	if err != nil {
		log.Fatal("Error while trying to save to DB: ", err)
	}
	return id, nil
}

// 특정 미션에 참여하는 가족 구성원 정보를 조회하는 메소드
func (s *Storage) GetFamMission(familyId int, userId int, missionId int) ([]int, error) {
	var family int
	var familys []int
	iter, err := s.db.Query(`SELECT user_Id FROM fam_mission where (mission_id=$1 AND family_id=$2) AND (user_id!=$3 AND del=false) order by regdate desc`, missionId, familyId, userId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&family); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		familys = append(familys, family)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return familys, nil
}

// 사용자가 미션 모델을 담아가는 메소드
func (s *Storage) GetModelMission(modelId int) (types.MissionModel, error) {
	var model types.MissionModel
	iter, err := s.db.Query(`SELECT model_id, mission_name, mission_content, mission_cycle, mission_term, mission_type, model_image, user_id FROM mission_model_view where model_id=$1 AND del=false`, modelId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&model.ModelId, &model.MissionName, &model.MissionContent, &model.MissionCycle, &model.MissionTerm, &model.MissionType, &model.ImageName, &model.UserId); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		user, err := s.GetOneUser(model.UserId)
		if err != nil {
			log.Fatal(err)
		}
		model.User = user
	}
	return model, nil
}

// 사용자가 다른 사용자의 미션을 조회하는 메소드(이 경우 공개로 설정된 것만 가져옴)
func (s *Storage) GetAllOtherMissions(userId int) ([]types.Mission, error) {
	var mission types.Mission
	var missions []types.Mission
	iter, err := s.db.Query("SELECT mission_id, user_id, family_id, mission_cycle, mission_type, mission_term, mission_name, mission_content, mission_start, mission_end, regdate, views, comments FROM mission_view where (user_id=$1 AND pub=true) AND del=false order by regdate desc", userId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&mission.MissionId, &mission.UserId, &mission.FamilyId, &mission.MissionCycle, &mission.MissionType, &mission.MissionTerm, &mission.MissionName, &mission.MissionContent, &mission.MissionStart, &mission.MissionEnd, &mission.Regdate, &mission.Views, &mission.Comments); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		model, err := s.GetModelMission(mission.ModelId)
		if err != nil {
			log.Fatal(err)
		}
		mission.Model = model
		users, err := s.GetFamMission(mission.FamilyId, mission.UserId, mission.MissionId)
		if err != nil {
			log.Fatal(err)
		}
		var candidates []types.User
		for _, user := range users {
			candidate, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			candidates = append(candidates, candidate)
		}
		mission.Candidates = candidates
		missions = append(missions, mission)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return missions, nil
}

// 내가 참여하는 모든 미션 조회 메소드
func (s *Storage) GetAllUserMissions(userId int, missionType string) ([]types.Mission, error) {
	var mission = types.Mission{}
	var missions = []types.Mission{}
	missionType = "%" + missionType + "%"
	iter, err := s.db.Query("SELECT mission_id, user_id, family_id, mission_cycle, mission_type, mission_term, mission_name, mission_content, mission_start, mission_end, regdate, views, comments FROM mission_view where (user_id=$1 AND mission_type like $2) AND del=false order by regdate desc", userId, missionType)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&mission.MissionId, &mission.UserId, &mission.FamilyId, &mission.MissionCycle, &mission.MissionType, &mission.MissionTerm, &mission.MissionName, &mission.MissionContent, &mission.MissionStart, &mission.MissionEnd, &mission.Regdate, &mission.Views, &mission.Comments); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		model, err := s.GetModelMission(mission.ModelId)
		if err != nil {
			log.Fatal(err)
		}
		mission.Model = model
		users, err := s.GetFamMission(mission.FamilyId, mission.UserId, mission.MissionId)
		if err != nil {
			log.Fatal(err)
		}
		var candidates []types.User
		for _, user := range users {
			candidate, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			candidates = append(candidates, candidate)
		}
		mission.Candidates = candidates
		missions = append(missions, mission)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return missions, nil
}

// 미션을 등록하는 메소드
func (s *Storage) InsertMission(mission types.Mission) (int, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO mission (user_id, family_id, mission_cycle, mission_type, mission_term, mission_image, mission_name, mission_content, mission_start, mission_end, pub, model_id) VALUES
    ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING mission_id`, mission.UserId, mission.FamilyId, mission.MissionCycleId, mission.MissionTypeId, mission.MissionTermId, mission.MissionImage, mission.MissionName, mission.MissionContent, mission.MissionStart, mission.MissionEnd, mission.Pub, mission.ModelId).Scan(&id)
	if err != nil {
		log.Fatal("Error while trying to save to DB: ", err)
	}
	return id, nil
}

// 완료한 미션을 완료상태로 바꾸는 메소드
func (s *Storage) UpdateMissionComplete(missionId int) error {
	_, err := s.db.Query(`UPDATE mission SET complete=true where mission_id=$1`, missionId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 나를 제외한 내 가족이 등록한 미션을 조회하는 쿼리(가족 미션 조회 쿼리)
func (s *Storage) GetFamilyMission(familyId int, userId int) ([]types.Mission, error) {
	var mission types.Mission
	var missions []types.Mission
	iter, err := s.db.Query("SELECT mission_id, user_id, family_id, mission_cycle, mission_type, mission_term, mission_name, mission_content, mission_start, mission_end, regdate, views, comments FROM mission_view where (family_id = $1 AND user_id != $2) AND del=false order by regdate desc", familyId, userId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&mission.MissionId, &mission.UserId, &mission.FamilyId, &mission.MissionCycle, &mission.MissionType, &mission.MissionTerm, &mission.MissionName, &mission.MissionContent, &mission.MissionStart, &mission.MissionEnd, &mission.Regdate, &mission.Views, &mission.Comments); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		model, err := s.GetModelMission(mission.ModelId)
		if err != nil {
			log.Fatal(err)
		}
		mission.Model = model
		users, err := s.GetFamMission(mission.FamilyId, mission.UserId, mission.MissionId)
		if err != nil {
			log.Fatal(err)
		}
		var candidates []types.User
		for _, user := range users {
			candidate, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			candidates = append(candidates, candidate)
		}
		mission.Candidates = candidates
		missions = append(missions, mission)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return missions, nil
}

// 참여 대기 상태인 미션을 삭제하는 메소드
func (s *Storage) DeleteMission(missionId int) error {
	//mission 테이블에서 삭제
	_, err := s.db.Query(`UPDATE mission SET del=true where mission_id=$1`, missionId)
	if err != nil {
		log.Fatal("Error while trying to delete to DB 1: ", err)
	}
	err = s.DeleteFamMission(missionId)
	if err != nil {
		log.Fatal("Error while trying to delete to DB 2: ", err)
	}
	return nil
}

// 참여 대기 상태 미션 삭제시 해당 미션에 대한 정보를 fammission테이블에서 삭제하는 메소드
func (s *Storage) DeleteFamMission(missionId int) error {
	_, err := s.db.Query(`UPDATE fam_mission SET del=true where mission_id=$1`, missionId)
	if err != nil {
		log.Fatal("Error while trying to delete to DB: ", err)
	}
	return nil
}

// 특정 미션에 특정 구성원이 탈퇴하는 메소드
func (s *Storage) DeleteOneFamMission(missionId int, userId int) error {
	_, err := s.db.Query(`UPDATE fam_mission SET del=true where mission_id=$1 AND user_id=$2`, missionId, userId)
	if err != nil {
		log.Fatal("Error while trying to delete to DB: ", err)
	}
	return nil
}

// id로 특정 미션 정보를 조회하는 메소드
func (s *Storage) GetOneMission(missionId int) (types.Mission, error) {
	var mission types.Mission
	iter, err := s.db.Query(`SELECT mission_id, user_id, family_id, mission_cycle, mission_type, mission_term, mission_name, mission_content, mission_start, mission_end, regdate, views, comments FROM mission_view where (mission_id=$1 AND del=false)`, missionId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err = iter.Scan(&mission.MissionId, &mission.UserId, &mission.FamilyId, &mission.MissionCycle, &mission.MissionType, &mission.MissionTerm, &mission.MissionName, &mission.MissionContent, &mission.MissionStart, &mission.MissionEnd, &mission.Regdate, &mission.Views, &mission.Comments); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		model, err := s.GetModelMission(mission.ModelId)
		if err != nil {
			log.Fatal(err)
		}
		mission.Model = model
		users, err := s.GetFamMission(mission.FamilyId, mission.UserId, mission.MissionId)
		if err != nil {
			log.Fatal(err)
		}
		var candidates []types.User
		for _, user := range users {
			candidate, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			candidates = append(candidates, candidate)
		}
		mission.Candidates = candidates
	}
	return mission, nil
}

// 가족 미션들을 조회하는 메소드
func (s *Storage) GetAllFamilyMission(familyId int) ([]types.Mission, error) {
	var mission types.Mission
	var missions []types.Mission
	iter, err := s.db.Query(`SELECT mission_id, user_id, family_id, mission_cycle, mission_type, mission_term, mission_name, mission_content, mission_start, mission_end, regdate, views, comments FROM mission_view where family_id=$1 AND del=false order by regdate desc`, familyId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&mission.MissionId, &mission.UserId, &mission.FamilyId, &mission.MissionCycle, &mission.MissionType, &mission.MissionTerm, &mission.MissionName, &mission.MissionContent, &mission.MissionStart, &mission.MissionEnd, &mission.Regdate, &mission.Views, &mission.Comments); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		model, err := s.GetModelMission(mission.ModelId)
		if err != nil {
			log.Fatal(err)
		}
		mission.Model = model
		users, err := s.GetFamMission(mission.FamilyId, mission.UserId, mission.MissionId)
		if err != nil {
			log.Fatal(err)
		}
		var candidates []types.User
		for _, user := range users {
			candidate, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			candidates = append(candidates, candidate)
		}
		mission.Candidates = candidates
		missions = append(missions, mission)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return missions, nil
}

// 수정된 미션의 정보를 업데이트 하는 메소드
func (s *Storage) UpdateMission(mission types.Mission) error {
	_, err := s.db.Query(`UPDATE mission SET mission_cycle=$1, mission_type=$2, mission_image=$3, mission_name=$4, mission_start=$5, mission_end=$6, complete=false, pub=$7, mission_content=$8, mission_term=$9, model_id=$10
                                where mission_id=$11`, mission.MissionCycleId, mission.MissionTypeId, mission.MissionImage, mission.MissionName,
		mission.MissionStart, mission.MissionEnd, mission.Pub, mission.MissionContent, mission.MissionTermId, mission.ModelId, mission.MissionId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 미션 검색 메소드
func (s *Storage) SearchMissions(search string, missionType string) ([]types.Mission, error) {
	var mission types.Mission
	var missions []types.Mission
	var searchs = "%" + search + "%"
	if search == "" {
		searchs = ""
	}
	missionType = "%" + missionType + "%"
	iter, err := s.db.Query(`SELECT mission_id, user_id, family_id, mission_cycle, mission_type, mission_term, mission_name, mission_content, mission_start, mission_end, regdate, views, comments FROM mission_view where (((mission_content like $1 or mission_name like $2) AND (pub=true AND mission_type like $3)) AND del=false) order by regdate desc`, searchs, searchs, missionType)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&mission.MissionId, &mission.UserId, &mission.FamilyId, &mission.MissionCycle, &mission.MissionType, &mission.MissionTerm, &mission.MissionName, &mission.MissionContent, &mission.MissionStart, &mission.MissionEnd, &mission.Regdate, &mission.Views, &mission.Comments); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		model, err := s.GetModelMission(mission.ModelId)
		if err != nil {
			log.Fatal(err)
		}
		mission.Model = model
		users, err := s.GetFamMission(mission.FamilyId, mission.UserId, mission.MissionId)
		if err != nil {
			log.Fatal(err)
		}
		var candidates []types.User
		for _, user := range users {
			candidate, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			candidates = append(candidates, candidate)
		}
		mission.Candidates = candidates
		missions = append(missions, mission)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return missions, nil
}

// 댓글 수 증가 메소드
func (s *Storage) UpComments(missionId int) error {
	var comments int
	iter := s.db.QueryRow(`SELECT comments from mission where mission_id=$1`, missionId)
	iter.Scan(&comments)
	_, err := s.db.Query(`UPDATE mission SET comments=$1 where mission_id=$2`, comments+1, missionId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 댓글 수 감소 메소드
func (s *Storage) DownComments(missionId int) error {
	var comments int
	iter := s.db.QueryRow(`SELECT comments from mission where mission_id=$1`, missionId)
	iter.Scan(&comments)
	_, err := s.db.Query(`UPDATE mission SET comments=$1 where mission_id=$2`, comments-1, missionId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 조회 수 증가 메소드
func (s *Storage) UpViews(missionId int) error {
	var views int
	iter := s.db.QueryRow(`SELECT views from mission where mission_id=$1`, missionId)
	iter.Scan(&views)
	_, err := s.db.Query(`UPDATE mission SET views=$1 where mission_id=$2`, views+1, missionId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// missionmodel 테이블의 미션 담아가기 수 증가 메소드
func (s *Storage) UpMissioncopy(modelId int) error {
	var missioncopy int
	iter := s.db.QueryRow(`SELECT mission_copy from mission_model where mission_id=$1`, modelId)
	iter.Scan(&missioncopy)
	missioncopy = missioncopy + 1
	_, err := s.db.Query(`UPDATE mission_model SET mission_copy=$1 where model_id=$2`, missioncopy, modelId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 미션을 비공개로 설정하는 메소드
func (s *Storage) IsNotPub(missionId int) error {
	_, err := s.db.Query(`UPDATE mission SET pub=$1 where mission_id=$2`, false, missionId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 미션을 공개로 설정하는 메소드
func (s *Storage) IsPub(missionId int) error {
	_, err := s.db.Query(`UPDATE mission SET pub=$1 where mission_id=$2`, true, missionId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// ----------------------------------------------
// 아래서부터는 missioncomment 테이블에 관한 쿼리 메소드
// ----------------------------------------------

// 미션 댓글을 등록하는 메소드
func (s *Storage) InsertMissionComment(comment types.MissionComment) (int, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO mission_comment (user_id, mission_id, content) VALUES ($1,$2,$3) RETURNING mission_comment_id`,
		comment.UserId, comment.MissionId, comment.Content).Scan(&id)
	if err != nil {
		log.Fatal("Error while trying to save to DB: ", err)
	}
	return id, nil
}

// 특정 미션 댓글을 불러오는 메소드
func (s *Storage) GetOneMissionComment(missionCommentId int) (types.MissionComment, error) {
	var comment types.MissionComment
	iter, err := s.db.Query(`SELECT * FROM mission_comment where mission_comment_id=$1 AND del=false`, missionCommentId)
	if err != nil {
		log.Fatal(err)
	}
	err = iter.Scan(&comment)
	if err != nil {
		log.Fatal("Error while trying to get to DB: ", err)
	}
	return comment, nil
}

// 미션 댓글을 삭제하는 메소드
func (s *Storage) DeleteMissionComment(missionCommentId int) error {
	_, err := s.db.Query(`UPDATE mission_comment SET del=true where mission_comment_id=$1`, missionCommentId)
	if err != nil {
		log.Fatal("Error while trying to delete from DB: ", err)
	}
	return nil
}

// 해당 미션의 댓글을 불러오는 메소드
func (s *Storage) GetOneMissionComments(missionId int) ([]types.MissionComment, error) {
	var comment types.MissionComment
	var comments []types.MissionComment
	iter, err := s.db.Query(`SELECT mission_comment_id, user_id, mission_id, content, regdate FROM mission_comment where mission_id=$1 AND del=false order by regdate desc`, missionId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&comment.MissionCommentId, &comment.UserId, &comment.MissionId, &comment.Content, &comment.Regdate); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		comments = append(comments, comment)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return comments, nil
}

// 사용자가 쓴 미션 댓글들을 불러오는 메소드
func (s *Storage) GetOneUserMissionComments(userId int) ([]types.MissionComment, error) {
	var comment types.MissionComment
	var comments []types.MissionComment
	iter, err := s.db.Query(`SELECT mission_comment_id, user_id, mission_id, content, regdate FROM mission_comment where user_id=$1 AND del=false order by regdate desc`, userId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&comment.MissionCommentId, &comment.UserId, &comment.MissionId, &comment.Content, &comment.Regdate); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		comments = append(comments, comment)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return comments, nil
}

// 특정 유저의 정보를 조회하는 메소드
func (s *Storage) GetOneUser(userId int) (types.User, error) {
	var user types.User
	iter, err := s.db.Query(`SELECT id, name, birth, role, intro, profile_image, regdate FROM users WHERE id=$1`, userId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&user.UserId, &user.Name, &user.Birth, &user.Role, &user.Intro, &user.ProfileName, &user.Regdate); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return user, nil
}
