package service

import (
	"github.com/nomaltree/family-eeum/feed/types"
	"log"
)

//feed, feedcomment, feedbookmark에 관련된 서비스함수 작성 페이지

type Service interface {
	InsertFeed(feed types.Feed) (int, error)
	DeleteFeed(feedId int) error
	GetFeed(feedId int) (types.Feed, error)
	SearchFeed(search string) ([]types.Feed, error)
	GetUserFeeds(userId int) ([]types.Feed, error)
	InsertFeedComment(comment types.FeedComment) (int, error)
	DeleteFeedComment(comment types.FeedComment) error
	GetOneFeedComments(feedId int) ([]types.FeedComment, error)
	GetOneUserFeedComments(userId int) ([]types.FeedComment, error)
	GetUserBookmarkFeed(userId int) ([]types.Feed, error)
	InsertFeedBookmark(userId int, feedId int) (int, error)
	DeleteFeedBookmark(userId int, feedId int) error
	InsertFeedTag(feedId int, tag string) (int, error)
	GetFeedTags(feedId int) ([]string, error)
	DeleteFeedTags(feedId int) error
	UpdateFeed(feed types.Feed) error
	GetAllFeeds() ([]types.Feed, error)
	GetFamilyFeeds(familyId int) ([]types.Feed, error)
}

type Repository interface {
	InsertFeed(feed types.Feed) (int, error)
	DeleteFeed(feedId int) error
	GetFeed(feedId int) (types.Feed, error)
	SearchFeed(search string) ([]types.Feed, error)
	GetUserFeeds(userId int) ([]types.Feed, error)
	UpFeedComments(feedId int) error
	DownFeedComments(feedId int) error
	UpFeedViews(feedId int) error
	UpFeedLikes(feedId int) error
	DownFeedLikes(feedId int) error
	InsertFeedComment(comment types.FeedComment) (int, error)
	DeleteFeedComment(feedcommentId int) error
	GetOneFeedComments(feedId int) ([]types.FeedComment, error)
	GetOneUserFeedComments(userId int) ([]types.FeedComment, error)
	GetUserBookmarkFeed(userId int) ([]types.Feed, error)
	InsertFeedBookmark(userId int, feedId int) (int, error)
	DeleteFeedBookmark(userId int, feedId int) error
	GetOneFeedComment(feedcommentId int) (types.FeedComment, error)
	InsertFeedTag(feedId int, tag string) (int, error)
	GetFeedTags(feedId int) ([]string, error)
	DeleteFeedTags(feedId int) error
	UpdateFeed(feed types.Feed) error
	GetAllFeeds() ([]types.Feed, error)
	GetFamilyFeeds(familyId int) ([]types.Feed, error)
}

type service struct {
	r Repository
}

func NewFeedService(r Repository) Service {
	return &service{r}
}

// 특정 가족들의 피드를 조회하는 메소드
func (s *service) GetFamilyFeeds(familyId int) ([]types.Feed, error) {
	feeds, err := s.r.GetFamilyFeeds(familyId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return feeds, nil
}

// 전체 피드를 조회하는 메소드(공개로 설정된 것만 가져옴)
func (s *service) GetAllFeeds() ([]types.Feed, error) {
	feeds, err := s.r.GetAllFeeds()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return feeds, nil
}

// 수정된 피드 정보들을 업데이트하는 메소드
func (s *service) UpdateFeed(feed types.Feed) error {
	err := s.r.UpdateFeed(feed)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 특정 피드의 태그들을 삭제하는 메소드
func (s *service) DeleteFeedTags(feedId int) error {
	err := s.r.DeleteFeedTags(feedId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 특정 피드의 태그들을 불러오는 메소드
func (s *service) GetFeedTags(feedId int) ([]string, error) {
	tags, err := s.r.GetFeedTags(feedId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return tags, nil
}

// 피드 태그들을 등록하는 메소드
func (s *service) InsertFeedTag(feedId int, tag string) (int, error) {
	var id int
	id, err := s.r.InsertFeedTag(feedId, tag)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

// 피드를 등록하는 메소드
func (s *service) InsertFeed(feed types.Feed) (int, error) {
	id, err := s.r.InsertFeed(feed)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

// 피드를 삭제하는 메소드
func (s *service) DeleteFeed(feedId int) error {
	err := s.r.DeleteFeed(feedId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// id로 특정 피드 정보를 불러오는 메소드
func (s *service) GetFeed(feedId int) (types.Feed, error) {
	feed, err := s.r.GetFeed(feedId)
	if err != nil {
		log.Fatal(err)
		return types.Feed{}, err
	}
	err = s.r.UpFeedViews(feedId)
	if err != nil {
		log.Fatal(err)
		return types.Feed{}, err
	}
	return feed, nil
}

// 피드를 검색하는 메소드
func (s *service) SearchFeed(search string) ([]types.Feed, error) {
	feeds, err := s.r.SearchFeed(search)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return feeds, nil
}

// 사용자가 작성한 피드 정보를 불러오는 메소드
func (s *service) GetUserFeeds(userId int) ([]types.Feed, error) {
	feeds, err := s.r.GetUserFeeds(userId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return feeds, nil
}

// 피드 댓글을 등록하는 메소드
func (s *service) InsertFeedComment(comment types.FeedComment) (int, error) {
	id, err := s.r.InsertFeedComment(comment)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	err = s.r.UpFeedComments(comment.Feedid)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

// 피드 댓글을 삭제하는 메소드
func (s *service) DeleteFeedComment(comment types.FeedComment) error {
	err := s.r.DownFeedComments(comment.Feedid)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = s.r.DeleteFeedComment(comment.Feedcommentid)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 해당 피드의 댓글을 불러오는 메소드
func (s *service) GetOneFeedComments(feedId int) ([]types.FeedComment, error) {
	comments, err := s.r.GetOneFeedComments(feedId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return comments, nil
}

// 사용자가 쓴 피드댓글들을 불러오는 메소드
func (s *service) GetOneUserFeedComments(userId int) ([]types.FeedComment, error) {
	comments, err := s.r.GetOneUserFeedComments(userId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return comments, nil
}

// 사용자가 좋아요한 피드정보들을 불러오는 메소드
func (s *service) GetUserBookmarkFeed(userId int) ([]types.Feed, error) {
	feeds, err := s.r.GetUserBookmarkFeed(userId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return feeds, nil
}

// 특정 피드를 좋아요에 추가하는 메소드
func (s *service) InsertFeedBookmark(userId int, feedId int) (int, error) {
	err := s.r.UpFeedLikes(feedId)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	id, err := s.r.InsertFeedBookmark(userId, feedId)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil
}

// 특정 피드의 좋아요를 삭제하는 메소드
func (s *service) DeleteFeedBookmark(userId int, feedId int) error {
	err := s.r.DownFeedLikes(feedId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = s.r.DeleteFeedBookmark(userId, feedId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
