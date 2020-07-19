package repository

import (
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_posts"
	"github.com/jinzhu/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) onef_posts.PostRepository {
	return &postRepository{db}
}

func (repo *postRepository) Get(req *model.PostRequest) ([]model.Post, error) {
	var posts []model.Post
	if dbc := repo.db.Find(&posts); dbc.Error != nil {
		return nil, dbc.Error
	}
	return posts, nil
}

func (repo *postRepository) GetWithId(id int64) (*model.Post, error) {
	var post model.Post
	if dbc := repo.db.Where("id = ?", id).First(&post); dbc.Error != nil {
		return nil, dbc.Error
	}

	var creator model.User
	if dbc := repo.db.Where("id = ?", post.CreatorId).First(&creator); dbc.Error == nil {
		var profile model.UserProfile
		if dbc := repo.db.Where("user_id = ?", creator.Id).First(&profile); dbc.Error != nil {
			return nil, dbc.Error
		}

		creator.Profile = &profile
		post.Creator = &creator
	}

	var community model.Community
	if dbc := repo.db.Where("id = ?", post.CommunityId).First(&community); dbc.Error == nil {
		post.Community = &community
	}

	var language model.Language
	if dbc := repo.db.Where("id = ?", post.LanguageId).First(&language); dbc.Error == nil {
		post.Language = &language
	}

	var postComments []model.PostComment
	if dbc := repo.db.Where("post_id = ?", post.Id).Find(&postComments); dbc.Error == nil {
		post.PostComments = postComments
	}

	var postReaction *model.PostReaction
	if dbc := repo.db.Where("post_id = ?", post.Id).First(&postReaction); dbc.Error == nil {
		post.PostReaction = postReaction
	}
	return &post, nil
}

func (repo *postRepository) GetPostWithUuid(postUuid string) (*model.Post, error) {
	var post model.Post
	if dbc := repo.db.Where("uuid = ?", postUuid).First(&post); dbc.Error != nil {
		return nil, dbc.Error
	}

	var creator model.User
	if dbc := repo.db.Where("id = ?", post.CreatorId).First(&creator); dbc.Error == nil {
		var profile model.UserProfile
		if dbc := repo.db.Where("user_id = ?", creator.Id).First(&profile); dbc.Error != nil {
			return nil, dbc.Error
		}

		creator.Profile = &profile
		post.Creator = &creator
	}

	var community model.Community
	if dbc := repo.db.Where("id = ?", post.CommunityId).First(&community); dbc.Error == nil {
		post.Community = &community
	}

	var language model.Language
	if dbc := repo.db.Where("id = ?", post.LanguageId).First(&language); dbc.Error == nil {
		post.Language = &language
	}

	var postComments []model.PostComment
	if dbc := repo.db.Where("post_id = ?", post.Id).Find(&postComments); dbc.Error == nil {
		post.PostComments = postComments
	}

	var postReaction *model.PostReaction
	if dbc := repo.db.Where("post_id = ?", post.Id).First(&postReaction); dbc.Error == nil {
		post.PostReaction = postReaction
	}

	return &post, nil
}

func (repo *postRepository) Create(post *model.Post) error {
	if dbc := repo.db.Create(&post); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *postRepository) UpdateStatus(postUuuid string, status string) error {
	err := repo.db.Model(&model.Post{}).Where("uuid = ?", postUuuid).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *postRepository) Update(post *model.Post) error {
	err := repo.db.Where("id = ?", post.Id).Updates(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *postRepository) Delete(post *model.Post) error {
	err := repo.db.Where("id = ?", post.Id).Delete(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *postRepository) CreateCirclePost(circleId int64, postId int64) error {
	circlePost := model.CirclePost{CircleId: &circleId, PostId: &postId}
	if dbc := repo.db.Create(&circlePost); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}
