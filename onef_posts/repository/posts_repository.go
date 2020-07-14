package repository

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"
	"github.com/jinzhu/gorm"
)

type postsRepository struct {
	db       *gorm.DB
	postRepo onef_posts.PostRepository
}

func NewPostsRepository(db *gorm.DB, postRepo onef_posts.PostRepository) onef_posts.PostsRepository {
	return &postsRepository{db, postRepo}
}

func (repo *postsRepository) GetTopPosts(req *model.TopPostRequest) ([]model.TopPost, error) {
	var topPosts []model.TopPost
	if dbc := repo.db.Find(&topPosts).Limit(req.Count); dbc.Error != nil {
		return nil, dbc.Error
	}

	topPostList := make([]model.TopPost, 0)
	for _, topPost := range topPosts {
		if post, err := repo.postRepo.GetWithId(*topPost.PostId); err == nil {
			topPost.Post = post
		}

		topPostList = append(topPostList, topPost)
	}

	return topPostList, nil
}

func (repo *postsRepository) GetTrendingPosts(req *model.TrendingPostRequest) ([]model.TrendingPost, error) {
	var trendingPosts []model.TrendingPost
	if dbc := repo.db.Find(&trendingPosts).Limit(req.Count); dbc.Error != nil {
		return nil, dbc.Error
	}

	trendingPostList := make([]model.TrendingPost, 0)
	for _, trendingPost := range trendingPosts {
		if post, err := repo.postRepo.GetWithId(*trendingPost.PostId); err == nil {
			trendingPost.Post = post
		}

		trendingPostList = append(trendingPostList, trendingPost)
	}

	return trendingPostList, nil
}

func (repo *postsRepository) GetTimelinePosts(req *model.PostRequest) ([]model.Post, error) {
	var posts []model.Post
	if dbc := repo.db.Find(&posts).Limit(req.Count); dbc.Error != nil {
		return nil, dbc.Error
	}

	timelinePostList := make([]model.Post, 0)
	for _, timelinePost := range posts {
		if post, err := repo.postRepo.GetWithId(timelinePost.Id); err == nil {
			timelinePostList = append(timelinePostList, *post)
		} else {
			timelinePostList = append(timelinePostList, timelinePost)
		}
	}

	return timelinePostList, nil
}
