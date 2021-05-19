package biz

import (
	"context"
	"time"
)

type Article struct {
	Id int64
	Title string
	CreatedBy int64
	UpdatedBy int64
	Content string
	Type int64
	CityId int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleRepo interface {
	CreateArticle(context.Context, *Article) error
	UpdateArticle(context.Context, *Article) error
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (uc *ArticleUsecase) Create(ctx context.Context, a *Article) error {
	return uc.repo.CreateArticle(ctx, a)
}

func (uc *ArticleUsecase) Update(ctx context.Context,a *Article) error {
	return uc.repo.UpdateArticle(ctx, a)
}
