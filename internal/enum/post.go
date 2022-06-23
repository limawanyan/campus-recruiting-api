package enum

// PostSortType 帖子排序类型
type PostSortType int

// PostType 帖子类型
type PostType int

const (
	_          PostSortType = iota
	Hot                     // 最热
	NewReply                // 新回复
	NewRelease              // 新发布
)

const (
	_                   PostType = iota
	InterviewExperience          // 面经分享
	InterviewSubject             // 面题交流
	SalaryDisclosure             // 薪资爆料
)
