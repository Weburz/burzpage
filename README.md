# PixiePress

TODO: Write documentations

The database shape is as follows:

1. Articles
    - Title
    - Author (should relate to the "Users" entity)
    - Publication Date
    - Summary
    - Tags (should be synonymous with keywords)
    - Cover Image
    - Category (should relate to the "Categories" entity)
2. Categories
    - Name
    - Articles (should relate to a list of "Articles" entity)
    - Description
    - MetaImage (for SEO purposes)
3. Users
    - Name
    - Profile Picture
    - Biography
    - Articles (should to a list of the "Articles" entity)
    - Role (should relate to the "Roles" entity)
4. Roles
    - Name
    - Users (should relate to the "Users" entity)
