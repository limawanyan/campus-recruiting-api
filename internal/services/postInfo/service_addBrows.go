package postInfo

func (s *service) AddBrows(postId int32) (err error) {
	err = s.db.GetDbW().Exec("UPDATE post_info SET browse_num = browse_num+1 WHERE id = ?", postId).Error
	return
}
