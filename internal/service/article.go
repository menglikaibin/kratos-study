
package service

import(
	"blog/internal/biz"
	"context"
	"time"

	pb "blog/api/blog"
)

type ArticleService struct {
	pb.UnimplementedArticleServer

	uc *biz.ArticleUsecase
}

func NewArticleService(uc *biz.ArticleUsecase) *ArticleService {
	return &ArticleService{
		uc: uc,
	}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	err := s.uc.Create(ctx, &biz.Article{
		Title: req.GetTitle(),
		Content: req.GetContent(),
		Type: req.GetType(),
		CreatedBy: req.GetCreatedBy(),
		UpdatedBy: req.GetUpdatedBy(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CityId: req.GetCityId(),
	})

	return &pb.CreateArticleReply{}, err
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	err := s.uc.Update(ctx, &biz.Article{
		Id: req.GetId(),
		Title: req.GetTitle(),
		Content: req.GetContent(),
		UpdatedBy: req.GetUpdatedBy(),
		UpdatedAt: time.Now(),
		CityId: req.GetCityId(),
		Type: req.GetType(),
	})
	return &pb.UpdateArticleReply{}, err
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	return &pb.DeleteArticleReply{}, nil
}
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	return &pb.GetArticleReply{}, nil
}
