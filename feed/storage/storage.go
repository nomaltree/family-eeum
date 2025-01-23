package storage

import (
	"github.com/nomaltree/family-eeum/feed/types"
	"log"
)

//feed, feedcomment, feedbookmark에 관련된 쿼리함수 작성 페이지

// 전체 피드를 조회하는 메소드(공개로 설정된 것만 가져옴)
func (s *Storage) GetAllFeeds() ([]types.Feed, error) {
	var feed types.Feed
	var feeds []types.Feed
	iter, err := s.db.Query(`SELECT feed_id, mission_id, user_id, family_id, content, feed_image, comments, views, likes, regdate FROM feed_view where pub=true AND del=false order by regdate desc`)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&feed.FeedId, &feed.MissionId, &feed.Userid, &feed.Familyid, &feed.Content, &feed.FeedImage, &feed.Comments, &feed.Views, &feed.Likes, &feed.Regdate); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		comments, err := s.GetOneFeedComments(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedComments = comments
		tags, err := s.GetFeedTags(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedTags = tags
		user, err := s.GetOneUser(feed.Userid)
		if err != nil {
			log.Fatal(err)
		}
		feed.User = user
		users, err := s.GetBookmarkUsers(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		var likesUsers []types.User
		for _, user := range users {
			likesUser, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			likesUsers = append(likesUsers, likesUser)
		}
		feed.LikesUser = likesUsers
		feeds = append(feeds, feed)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return feeds, nil
}

// 수정된 피드 정보들을 업데이트하는 메소드
func (s *Storage) UpdateFeed(feed types.Feed) error {
	_, err := s.db.Query(`UPDATE feed SET feed_image=$1, content=$2 where feed_id=$3`, &feed.FeedImage, &feed.Content, &feed.FeedId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 특정 피드의 태그들을 삭제하는 메소드
func (s *Storage) DeleteFeedTags(feedId int) error {
	_, err := s.db.Query(`UPDATE feed_tag SET del=true where feed_id=$1`, feedId)
	if err != nil {
		log.Fatal("Error while trying to delete to DB: ", err)
	}
	return nil
}

// 특정 피드의 태그들을 불러오는 메소드
func (s *Storage) GetFeedTags(feedId int) ([]string, error) {
	var tag string
	var tags []string
	iter, err := s.db.Query(`SELECT content FROM feed_tag where feed_id=$1 AND del=false`, feedId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		err := iter.Scan(&tag)
		if err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		tags = append(tags, tag)
	}
	err = iter.Close()
	if err != nil {
		log.Fatal(err)
	}
	return tags, nil
}

// 피드 태그를 등록하는 메소드
func (s *Storage) InsertFeedTag(feedId int, tag string) (int, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO feed_tag (feed_id, content) VALUES ($1, $2) RETURNING tag_id`, feedId, tag).Scan(&id)
	if err != nil {
		log.Fatal("Error while trying to save to DB: ", err)
	}
	return id, nil
}

// 피드를 등록하는 메소드
func (s *Storage) InsertFeed(feed types.Feed) (int, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO feed (user_id, family_id, mission_id, feed_image, content) VALUES 
             ($1,$2,$3,$4,$5) RETURNING feed_id`, feed.Userid, feed.Familyid, feed.MissionId, feed.FeedImage, feed.Content).Scan(&id)
	if err != nil {
		log.Fatal("Error while trying to save to DB: ", err)
	}
	return id, nil
}

// 피드를 삭제하는 메소드
func (s *Storage) DeleteFeed(feedId int) error {
	_, err := s.db.Query(`UPDATE feed SET del=true where feed_id=$1`, feedId)
	if err != nil {
		log.Fatal("Error while trying to delete from DB 1: ", err)
	}
	_, err = s.db.Query(`UPDATE feed_comment SET del=true where feed_id=$1`, feedId)
	if err != nil {
		log.Fatal("Error while trying to delete from DB 2: ", err)
	}
	_, err = s.db.Query(`UPDATE feed_tag SET del=true where feed_id=$1`, feedId)
	if err != nil {
		log.Fatal("Error while trying to delete from DB 3: ", err)
	}
	_, err = s.db.Query(`UPDATE feed_bookmark SET del=true where feed_id=$1`, feedId)
	return nil
}

// id로 특정 피드 정보를 불러오는 메소드
func (s *Storage) GetFeed(feedId int) (types.Feed, error) {
	var feed types.Feed
	iter, err := s.db.Query(`SELECT feed_id, user_id, family_id, mission_id, feed_image, views, comments, content, regdate, likes FROM feed_view WHERE feed_id=$1 AND del=false`, feedId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&feed.FeedId, &feed.Userid, &feed.Familyid, &feed.MissionId, &feed.FeedImage, &feed.Views, &feed.Comments, &feed.Content, &feed.Regdate, &feed.Likes); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		comments, err := s.GetOneFeedComments(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedComments = comments
		tags, err := s.GetFeedTags(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedTags = tags
		user, err := s.GetOneUser(feed.Userid)
		if err != nil {
			log.Fatal(err)
		}
		feed.User = user
		users, err := s.GetBookmarkUsers(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		var likesUsers []types.User
		for _, user := range users {
			likesUser, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			likesUsers = append(likesUsers, likesUser)
		}
		feed.LikesUser = likesUsers
	}
	return feed, nil
}

// 피드를 검색하는 메소드
func (s *Storage) SearchFeed(search string) ([]types.Feed, error) {
	var feed types.Feed
	var feeds []types.Feed
	var searchs = "%" + search + "%"
	if search == "" {
		searchs = ""
	}
	iter, err := s.db.Query(`SELECT feed_id, user_id, family_id, mission_id, feed_image, views, comments, content, regdate, likes FROM feed_view where (content like $1 AND pub=true) AND del=false order by regdate desc`, searchs)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&feed.FeedId, &feed.Userid, &feed.Familyid, &feed.MissionId, &feed.FeedImage, &feed.Views, &feed.Comments, &feed.Content, &feed.Regdate, &feed.Likes); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		comments, err := s.GetOneFeedComments(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedComments = comments
		tags, err := s.GetFeedTags(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedTags = tags
		user, err := s.GetOneUser(feed.Userid)
		if err != nil {
			log.Fatal(err)
		}
		feed.User = user
		users, err := s.GetBookmarkUsers(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		var likesUsers []types.User
		for _, user := range users {
			likesUser, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			likesUsers = append(likesUsers, likesUser)
		}
		feed.LikesUser = likesUsers
		feeds = append(feeds, feed)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return feeds, nil
}

// 특정 가족들의 피드를 불러오는 메소드
func (s *Storage) GetFamilyFeeds(familyId int) ([]types.Feed, error) {
	var feed types.Feed
	var feeds []types.Feed
	iter, err := s.db.Query(`SELECT feed_id, user_id, family_id, mission_id, feed_image, views, comments, content, regdate, likes FROM feed_view where family_id=$1 AND del=false order by regdate desc`, familyId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&feed.FeedId, &feed.Userid, &feed.Familyid, &feed.MissionId, &feed.FeedImage, &feed.Views, &feed.Comments, &feed.Content, &feed.Regdate, &feed.Likes); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		comments, err := s.GetOneFeedComments(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedComments = comments
		tags, err := s.GetFeedTags(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedTags = tags
		user, err := s.GetOneUser(feed.Userid)
		if err != nil {
			log.Fatal(err)
		}
		feed.User = user
		users, err := s.GetBookmarkUsers(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		var likesUsers []types.User
		for _, user := range users {
			likesUser, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			likesUsers = append(likesUsers, likesUser)
		}
		feed.LikesUser = likesUsers
		feeds = append(feeds, feed)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return feeds, nil
}

// 사용자가 작성한 피드 정보를 불러오는 메소드
func (s *Storage) GetUserFeeds(userId int) ([]types.Feed, error) {
	var feed types.Feed
	var feeds []types.Feed
	iter, err := s.db.Query(`SELECT feed_id, user_id, family_id, mission_id, feed_image, views, comments, content, regdate, likes FROM feed_view where user_id=$1 AND del=false order by regdate desc`, userId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&feed.FeedId, &feed.Userid, &feed.Familyid, &feed.MissionId, &feed.FeedImage, &feed.Views, &feed.Comments, &feed.Content, &feed.Regdate, &feed.Likes); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		comments, err := s.GetOneFeedComments(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedComments = comments
		tags, err := s.GetFeedTags(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		feed.FeedTags = tags
		user, err := s.GetOneUser(feed.Userid)
		if err != nil {
			log.Fatal(err)
		}
		feed.User = user
		users, err := s.GetBookmarkUsers(feed.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		var likesUsers []types.User
		for _, user := range users {
			likesUser, err := s.GetOneUser(user)
			if err != nil {
				log.Fatal(err)
			}
			likesUsers = append(likesUsers, likesUser)
		}
		feed.LikesUser = likesUsers
		feeds = append(feeds, feed)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return feeds, nil
}

// 피드의 댓글 수 증가 메소드
func (s *Storage) UpFeedComments(feedId int) error {
	var comments int
	iter := s.db.QueryRow(`SELECT comments from feed where feed_id=$1`, feedId)
	iter.Scan(&comments)
	_, err := s.db.Query(`UPDATE feed SET comments=$1 where feed_id=$2`, comments+1, feedId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 피드의 댓글 수 감소 메소드
func (s *Storage) DownFeedComments(feedId int) error {
	var comments int
	iter := s.db.QueryRow(`SELECT comments FROM feed where feed_id=$1`, feedId)
	iter.Scan(&comments)
	_, err := s.db.Query(`UPDATE feed SET comments=$1 where feed_id=$2`, comments-1, feedId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 피드의 조회 수 증가 메소드
func (s *Storage) UpFeedViews(feedId int) error {
	var views int
	iter := s.db.QueryRow(`SELECT views from feed where feed_id=$1`, feedId)
	iter.Scan(&views)
	_, err := s.db.Query(`UPDATE feed SET views=$1 where feed_id=$2`, views+1, feedId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 피드의 좋아요 수 증가 메소드
func (s *Storage) UpFeedLikes(feedId int) error {
	var likes int
	iter := s.db.QueryRow(`SELECT likes from feed where feed_id=$1`, feedId)
	iter.Scan(&likes)
	_, err := s.db.Query(`UPDATE feed SET likes=$1 where feed_id=$2`, likes+1, feedId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// 피드의 좋아요 수 감소 메소드
func (s *Storage) DownFeedLikes(feedId int) error {
	var likes int
	iter := s.db.QueryRow(`SELECT likes from feed where feed_id=$1`, feedId)
	iter.Scan(&likes)
	_, err := s.db.Query(`UPDATE feed SET likes=$1 where feed_id=$2`, likes-1, feedId)
	if err != nil {
		log.Fatal("Error while trying to update to DB: ", err)
	}
	return nil
}

// -------------------------------------------
// 아래서부터는 feedcomment 테이블에 관한 쿼리 메소드
// -------------------------------------------

// 피드 댓글을 등록하는 메소드
func (s *Storage) InsertFeedComment(comment types.FeedComment) (int, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO feed_comment (user_id, feed_id, content, mission_id) VALUES ($1,$2,$3,$4) RETURNING feed_comment_id`,
		comment.Userid, comment.Feedid, comment.Content, comment.MissionId).Scan(&id)
	if err != nil {
		log.Fatal("Error while trying to save to DB: ", err)
	}
	return id, nil
}

// 특정 피드 댓글을 불러오는 메소드
func (s *Storage) GetOneFeedComment(feedCommentId int) (types.FeedComment, error) {
	var comment types.FeedComment
	iter, err := s.db.Query(`SELECT feed_comment_id, feed_id, user_id, content, mission_id, regdate FROM feed_comment WHERE feed_comment_id=$1 AND del=false`, feedCommentId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&comment.Feedcommentid, &comment.Feedid, &comment.Userid, &comment.Content, &comment.MissionId, &comment.Regdate); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
	}
	user, err := s.GetOneUser(comment.Userid)
	if err != nil {
		log.Fatal(err)
	}
	comment.User = user
	return comment, nil
}

// 피드 댓글을 삭제하는 메소드
func (s *Storage) DeleteFeedComment(feedCommentId int) error {
	_, err := s.db.Query(`UPDATE feed_comment SET del=true where feed_comment_id=$1`, feedCommentId)
	if err != nil {
		log.Fatal("Error while trying to delete from DB: ", err)
	}
	return nil
}

// 해당 피드의 댓글을 불러오는 메소드
func (s *Storage) GetOneFeedComments(feedId int) ([]types.FeedComment, error) {
	var comment types.FeedComment
	var comments []types.FeedComment
	iter, err := s.db.Query(`SELECT user_id, feed_id, content, mission_id, regdate FROM feed_comment where feed_id=$1 AND del=false order by regdate desc`, feedId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&comment.Userid, &comment.Feedid, &comment.Content, &comment.MissionId, &comment.Regdate); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		user, err := s.GetOneUser(comment.Userid)
		if err != nil {
			log.Fatal(err)
		}
		comment.User = user
		comments = append(comments, comment)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return comments, nil
}

// 사용자가 쓴 피드댓글들을 불러오는 메소드
func (s *Storage) GetOneUserFeedComments(userId int) ([]types.FeedComment, error) {
	var comment types.FeedComment
	var comments []types.FeedComment
	iter, err := s.db.Query(`SELECT user_id, feed_id, content, mission_id, regdate FROM feed_comment where user_id=$1 AND del=false order by regdate desc`, userId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&comment.Userid, &comment.Feedid, &comment.Content, &comment.MissionId, &comment.Regdate); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		user, err := s.GetOneUser(comment.Userid)
		if err != nil {
			log.Fatal(err)
		}
		comment.User = user
		comments = append(comments, comment)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return comments, nil
}

// --------------------------------------------
// 아래서부터는 feedbookmark 테이블에 관한 쿼리 메소드
// --------------------------------------------

// 사용자가 좋아요한 피드정보들을 불러오는 메소드
func (s *Storage) GetUserBookmarkFeed(userId int) ([]types.Feed, error) {
	var bookmark types.Feed
	var bookmarks []types.Feed
	iter, err := s.db.Query(`SELECT f.feed_id, f.user_id, f.family_id, f.mission_Id, f.feed_image, f.content, f.views, f.comments, f.likes, f.regdate FROM feed f INNER JOIN 
    						(SELECT * FROM feed_bookmark where user_id=$1) fb 
							on f.feed_id=fb.feed_id where f.del=false AND fb.del=false order by fb.regdate desc`, userId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&bookmark.FeedId, &bookmark.Userid, &bookmark.Familyid, &bookmark.MissionId, &bookmark.FeedImage, &bookmark.Content, &bookmark.Views, &bookmark.Comments, &bookmark.Likes, &bookmark.Regdate); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		comments, err := s.GetOneFeedComments(bookmark.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		bookmark.FeedComments = comments
		tags, err := s.GetFeedTags(bookmark.FeedId)
		if err != nil {
			log.Fatal(err)
		}
		bookmark.FeedTags = tags
		bookmarks = append(bookmarks, bookmark)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return bookmarks, nil
}

// 특정 피드를 좋아요에 추가하는 메소드
func (s *Storage) InsertFeedBookmark(userId int, feedId int) (int, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO feed_bookmark (feed_id, user_id) VALUES ($1,$2) RETURNING feed_bookmark_id`, feedId, userId).Scan(&id)
	if err != nil {
		log.Fatal("Error while trying to save to DB: ", err)
	}
	return id, nil
}

// 특정 피드의 좋아요를 삭제하는 메소드
func (s *Storage) DeleteFeedBookmark(userId int, feedId int) error {
	_, err := s.db.Query(`UPDATE feed_bookmark SET del=true where user_id=$1 AND feed_id=$2`, userId, feedId)
	if err != nil {
		log.Fatal("Error while trying to delete from DB: ", err)
	}
	return nil
}

func (s *Storage) GetBookmarkUsers(feedId int) ([]int, error) {
	var id int
	var ids []int
	iter, err := s.db.Query(`SELECT user_id FROM feed_bookmark WHERE feed_id=$1`, feedId)
	if err != nil {
		log.Fatal(err)
	}
	for iter.Next() {
		if err := iter.Scan(&id); err != nil {
			log.Fatal("Error while trying to get DB: ", err)
		}
		ids = append(ids, id)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return ids, nil
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
