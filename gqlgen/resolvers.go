package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/renatospaka/api-book-store/pg"
)

// Resolver connects individual resolvers with the datalayer.
type Resolver struct {
	Repository pg.Repository
}

func (r *agentResolver) Authors(ctx context.Context, obj *pg.Agent) ([]pg.Author, error) {
	return r.Repository.ListAuthorsByAgentID(ctx, obj.ID)
}

func (r *authorResolver) Website(ctx context.Context, obj *pg.Author) (*string, error) {
	var w string
	if obj.Website.Valid {
		w = obj.Website.String
		return &w, nil
	}
	return nil, nil
}

func (r *authorResolver) Agent(ctx context.Context, obj *pg.Author) (*pg.Agent, error) {
	agent, err := r.Repository.GetAgent(ctx, obj.AgentID)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *authorResolver) Books(ctx context.Context, obj *pg.Author) ([]pg.Book, error) {
	return r.Repository.ListBooksByAuthorID(ctx, obj.ID)
}

func (r *bookResolver) Authors(ctx context.Context, obj *pg.Book) ([]pg.Author, error) {
	return r.Repository.ListAuthorsByBookID(ctx, obj.ID)
}

func (r *mutationResolver) CreateAgent(ctx context.Context, data AgentInput) (*pg.Agent, error) {
	agent, err := r.Repository.CreateAgent(ctx, pg.CreateAgentParams{
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *mutationResolver) UpdateAgent(ctx context.Context, id int64, data AgentInput) (*pg.Agent, error) {
	agent, err := r.Repository.UpdateAgent(ctx, pg.UpdateAgentParams{
		ID:    id,
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *mutationResolver) DeleteAgent(ctx context.Context, id int64) (*pg.Agent, error) {
	agent, err := r.Repository.DeleteAgent(ctx, id)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, data AuthorInput) (*pg.Author, error) {
	author, err := r.Repository.CreateAuthor(ctx, pg.CreateAuthorParams{
		Name:    data.Name,
		Website: pg.StringPtrToNullString(data.Website),
		AgentID: data.AgentID,
	})
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, id int64, data AuthorInput) (*pg.Author, error) {
	author, err := r.Repository.UpdateAuthor(ctx, pg.UpdateAuthorParams{
		ID:      id,
		Name:    data.Name,
		Website: pg.StringPtrToNullString(data.Website),
		AgentID: data.AgentID,
	})
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id int64) (*pg.Author, error) {
	author, err := r.Repository.DeleteAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, data BookInput) (*pg.Book, error) {
	return r.Repository.CreateBook(ctx, pg.CreateBookParams{
		Title:       data.Title,
		Description: data.Description,
		Cover:       data.Cover,
	}, data.AuthorIDs)
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int64, data BookInput) (*pg.Book, error) {
	return r.Repository.UpdateBook(ctx, pg.UpdateBookParams{
		ID:          id,
		Title:       data.Title,
		Description: data.Description,
		Cover:       data.Cover,
	}, data.AuthorIDs)
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int64) (*pg.Book, error) {
	// BookAuthors associations will cascade automatically.
	book, err := r.Repository.DeleteBook(ctx, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *queryResolver) Agent(ctx context.Context, id int64) (*pg.Agent, error) {
	agent, err := r.Repository.GetAgent(ctx, id)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (r *queryResolver) Agents(ctx context.Context) ([]pg.Agent, error) {
	return r.Repository.ListAgents(ctx)
}

func (r *queryResolver) Author(ctx context.Context, id int64) (*pg.Author, error) {
	author, err := r.Repository.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]pg.Author, error) {
	return r.Repository.ListAuthors(ctx)
}

func (r *queryResolver) Book(ctx context.Context, id int64) (*pg.Book, error) {
	book, err := r.Repository.GetBook(ctx, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]pg.Book, error) {
	return r.Repository.ListBooks(ctx)
}

// Agent returns AgentResolver implementation.
func (r *Resolver) Agent() AgentResolver { return &agentResolver{r} }

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Book returns BookResolver implementation.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type agentResolver struct{ *Resolver }
type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
