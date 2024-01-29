# PixiePress

TODO: Write documentations

The database shape is as follows:

```mermaid
erDiagram
    ARTICLE {
        int articleId PK
        string title
        int authorId FK
        int publicationDate
        string summary
        string[] tags
        string coverImage
        string categoryName FK
    }
    AUTHOR {
        int authorId PK
        string username
        string profilePicture
        string biogpraphy
        int[] articleId FK
        string roleName FK
        string email
        string password
    }
    CATEGORY {
        string categoryName PK
        int[] articleId FK
        string description
        string metaImage
    }
    ROLE {
        string roleName PK
        int[] authorId FK
    }

    ARTICLE ||--|| CATEGORY : has
    CATEGORY ||--|{ ARTICLE : contains

    ARTICLE ||--|| AUTHOR : has
    AUTHOR ||--|{ ARTICLE : has

    AUTHOR ||--|| ROLE : has
    ROLE ||--|{ AUTHOR : contains
````

Planned features:

- Search capabilities on the frontend
- Comments section
