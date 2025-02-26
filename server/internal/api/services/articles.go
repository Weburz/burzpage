/*
Package services provides operations for managing articles within the system. It
includes functionality for retrieving, creating, updating, and deleting articles.

The primary interface, `ArticleService`, defines methods for interacting with article
data. The `ArticleServiceImpl` struct provides the concrete implementation of these
methods. These operations are used to manage articles, including article metadata like
titles, authors and publication status.

The package provides the following key functionalities:

  - GetAllArticles: Retrieves a list of all articles available in the system.
  - GetArticleByID: Fetches an article based on its unique identifier.
  - CreateArticle: Creates a new article by providing a title, author, and publication
    status.
  - UpdateArticle: Updates the details of an existing article, including title, author
    and publication status.
  - DeleteArticle: Removes an article from the system using its unique identifier.

This package is designed to handle typical CRUD (Create, Read, Update, Delete)
operations for articles, allowing the system to manage article data in a flexible
manner.

The package also includes a constructor function, `NewArticleService`, which initializes
and returns an instance of `ArticleServiceImpl` that implements the `ArticleService`
interface.
*/
package services

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/Weburz/burzcontent/server/internal/api/models"
)

/*
ArticleService defines the methods for interacting with articles in the system.

It includes CRUD operations for managing articles, such as retrieving a list of
articles, fetching a specific article by ID, creating a new article, updating
existing article details,
and deleting articles by their ID.
*/
type ArticleService interface {
	// GetAllArticles retrieves all the articles in the system.
	// It returns a slice of Article models and an error if any occurs.
	GetAllArticles() ([]models.Article, error)

	// GetArticleByID fetches a specific article by its unique ID.
	// It returns the Article model and an error if the article could not be found.
	GetArticleByID(id uuid.UUID) (models.Article, error)

	// CreateArticle creates a new article with the specified title, author, and
	// publication status.
	// It returns the newly created article model and an error if any occurs.
	CreateArticle(title, author string, isPublished bool) (models.Article, error)

	// UpdateArticle updates an existing article based on its ID.
	// The method accepts a unique ID, new title, new author, and publication status for
	// the update.
	// It returns the updated article and an error if any occurs.
	UpdateArticle(
		id uuid.UUID,
		title, author string,
		isPublished bool,
	) (models.Article, error)

	// DeleteArticle removes an article from the system using its unique ID.
	// It returns an error if the article could not be deleted (e.g., if it doesn't
	// exist).
	DeleteArticle(id uuid.UUID) error
}

/*
ArticleServiceImpl is the concrete implementation of the ArticleService interface.
It provides the actual logic for interacting with the article data.
*/
type ArticleServiceImpl struct{}

/*
NewArticleService creates and returns a new instance of ArticleServiceImpl,
which implements the ArticleService interface.
*/
func NewArticleService() *ArticleServiceImpl {
	return &ArticleServiceImpl{}
}

/*
GetAllArticles retrieves a list of all articles available in the system.

This method simulates fetching a collection of articles by generating a new
unique `articleID` and then returning a hardcoded list of articles. Each article
includes details such as ID, title, author, and publication status.

If there is an error generating the article ID (which should be rare), the function
returns an empty article with an error indicating the failure.

Returns:
  - A slice of `models.Article` representing the articles in the system.
  - An error, if there is an issue generating the article ID.
*/
func (as *ArticleServiceImpl) GetAllArticles() ([]models.Article, error) {
	articleID, err := uuid.NewV7()
	if err != nil {
		return []models.Article{
				{}},
			fmt.Errorf(
				"Unable to generate Article ID: %w\n",
				err,
			)
	}

	articles := []models.Article{
		{
			ID:          articleID,
			Title:       "Go Programming Basics",
			Author:      "John Doe",
			IsPublished: true,
		},
		{
			ID:          articleID,
			Title:       "Advanced Go Techniques",
			Author:      "Jane Smith",
			IsPublished: false,
		},
		{
			ID:          articleID,
			Title:       "Understanding Go Concurrency",
			Author:      "Alice Johnson",
			IsPublished: true,
		},
	}

	return articles, nil
}

/*
GetArticleByID retrieves a specific article by its unique ID.

This method simulates fetching an article based on the provided `id` by generating
a new article with hardcoded values. If there is an error generating the article ID,
it returns an empty article with no additional information.

Returns:
  - A `models.Article` representing the requested article.
  - An error, which is nil in this case as it always succeeds unless there is an issue
    with ID generation.
*/
func (as *ArticleServiceImpl) GetArticleByID(id uuid.UUID) (models.Article, error) {
	articleID, err := uuid.NewV7()
	if err != nil {
		return models.Article{}, nil
	}

	article := models.Article{
		ID:          articleID,
		Title:       "Go Programming Basics",
		Author:      "John Doe",
		IsPublished: true,
	}

	return article, nil
}

/*
CreateArticle creates a new article with the given title, author, and publication
status.

This method generates a unique article ID, then creates an article with the provided
title, author, and publication status. If there is an error generating the article ID,
it returns an empty article with no additional information.

Parameters:
  - title: The title of the article.
  - author: The author of the article.
  - isPublished: A boolean indicating whether the article is published or not.

Returns:
  - A `models.Article` representing the newly created article.
  - An error, which is nil in this case as the method always succeeds unless there is an
    issue with ID generation.
*/
func (as *ArticleServiceImpl) CreateArticle(
	title, author string,
	isPublished bool,
) (models.Article, error) {
	articleID, err := uuid.NewV7()
	if err != nil {
		return models.Article{}, nil
	}

	article := models.Article{
		ID:          articleID,
		Title:       title,
		Author:      author,
		IsPublished: isPublished,
	}

	return article, nil
}

/*
UpdateArticle updates the details of an existing article based on the provided ID.

This method generates a new unique article ID and updates the article with the given
title, author, and publication status. If there is an error generating the new article
ID, it returns an empty article with no additional information.

Parameters:
  - id: The unique identifier of the article to be updated.
  - title: The new title of the article.
  - author: The new author of the article.
  - isPublished: The new publication status of the article.

Returns:
  - A `models.Article` representing the updated article.
  - An error, which is nil in this case as the method always succeeds unless there is
    an issue with ID generation.
*/
func (as *ArticleServiceImpl) UpdateArticle(
	id uuid.UUID,
	title, author string,
	isPublished bool,
) (models.Article, error) {
	articleID, err := uuid.NewV7()
	if err != nil {
		return models.Article{}, nil
	}

	article := models.Article{
		ID:          articleID,
		Title:       title,
		Author:      author,
		IsPublished: isPublished,
	}

	return article, nil
}

/*
DeleteArticle removes an article from the system based on the provided article ID.

This method parses the provided article ID and attempts to delete the article. If the
ID is invalid, an error is returned indicating the issue. Once the article is
successfully deleted, a confirmation message is logged. The article's details are
printed to the console.

Parameters:
  - id: The unique identifier of the article to be deleted.

Returns:
  - An error if the article ID is invalid or if any other issues arise during the
    deletion process; nil if the deletion succeeds.
*/
func (as *ArticleServiceImpl) DeleteArticle(id uuid.UUID) error {
	articleID, err := uuid.Parse(id.String())
	if err != nil {
		return fmt.Errorf("Invalid Article ID: %w\n", err)
	}

	article := models.Article{ID: articleID}

	fmt.Printf("%v article is deleted!\n", article)
	return nil
}
