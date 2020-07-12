package repository

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"
	"github.com/jinzhu/gorm"
)

type postsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) onef_posts.PostsRepository {
	return &postsRepository{db}
}

func (repo *postsRepository) GetTopPosts(req *model.TopPostRequest) ([]model.TopPost, error) {
	var topPosts []model.TopPost
	if dbc := repo.db.Find(&topPosts).Limit(req.Count); dbc.Error != nil {
		return nil, dbc.Error
	}
	return topPosts, nil
}

func (repo *postsRepository) GetTrendingPosts(req *model.TrendingPostRequest) ([]model.TrendingPost, error) {
	var trendingPosts []model.TrendingPost
	if dbc := repo.db.Find(&trendingPosts).Limit(req.Count); dbc.Error != nil {
		return nil, dbc.Error
	}
	return trendingPosts, nil
}

func (repo *postsRepository) GetTimelinePosts(req *model.PostRequest) ([]model.Post, error) {
	var posts []model.Post
	if dbc := repo.db.Find(&posts).Limit(req.Count); dbc.Error != nil {
		return nil, dbc.Error
	}
	return posts, nil
}
