package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	PublishedAt string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleRepo interface {
	CreateArticle(context.Context, *Article) error
	UpdateArticle(context.Context, *Article) error
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Helper) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper("usecase/article", &logger)}
}

func (uc *ArticleUsecase) Create(ctx context.Context, a *Article) error {
	return uc.repo.CreateArticle(ctx, a)
}

func (uc *ArticleUsecase) Update(ctx context.Context,a *Article) error {
	return uc.repo.UpdateArticle(ctx, a)
}
