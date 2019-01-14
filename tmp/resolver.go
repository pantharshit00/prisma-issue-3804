package tmp

import (
	"context"
	"issue-3804/generated/prisma-client"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateDraft(ctx context.Context, title string, userId string) (*prisma.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) Publish(ctx context.Context, postId string) (*prisma.Post, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *prisma.Post) (*prisma.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) PublishedPosts(ctx context.Context) ([]prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, postId string) (*prisma.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) PostsByUser(ctx context.Context, userId string) ([]prisma.Post, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *prisma.User) ([]prisma.Post, error) {
	panic("not implemented")
}
