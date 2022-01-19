package graph
import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/LunaChevalier/tutorial-go-react/graph/generated"
	"github.com/LunaChevalier/tutorial-go-react/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *gorm.DB
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }



func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
			ID:   "todo001",
			Text: "部屋の掃除",
			Done: false,
			User: &model.User{
					ID:   "user001",
					Name: "たろー",
			},
    },nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	r.DB.Find(&todos)
	return todos, nil
	
	// こんな形のmodelがレスポンスされる
	// return []*model.Todo{
	// 	&model.Todo{
	// 			ID:   "todo001",
	// 			Text: "部屋の掃除",
	// 			Done: false,
	// 			User: &model.User{
	// 					ID:   "user001",
	// 					Name: "たろ1",
	// 			},
	// 	},
	// 	&model.Todo{
	// 			ID:   "todo002",
	// 			Text: "買い物",
	// 			Done: true,
	// 			User: &model.User{
	// 					ID:   "user001",
	// 					Name: "たろ1",
	// 			},
	// 	},
	// },nil
}