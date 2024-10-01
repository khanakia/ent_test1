## How to generate Ent Crud
Let try to generate the crud for the `saas/schema/post_category.go` schema

- Add this code in the schema file

  ```go
  func (PostCategory) Annotations() []schema.Annotation {
    return []schema.Annotation{
      entgql.RelayConnection(),
      entgql.QueryField("postCategories"),
      entgql.MultiOrder(),
      entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
    }
  }
  ```

- Run the command `task saas:entg`
- Edit this file `gqlsa/graph/gqlsaresolver/ent.resolvers.go` `func (r *queryResolver) PostCategories` update this func accordingly
- Now you will see the `input CreatePostCategoryInput` and update inputs in `gqlsa/graph/ent.graphql`
- Also queries `postCategories` must be generated to in ent.graphql
- Add the create and update mutation `gqlsa/graph/cms.graphql`
  ```graphql
  extend type Mutation {
    ...
    createPostCategory(input: CreatePostCategoryInput!): PostCategory!
    updatePostCategory(id: ID!, input: UpdatePostCategoryInput!): PostCategory!
  }
  ```
- Run the command `task gqlsa:gql`
- Go to the `gqlsa/graph/gqlsaresolver/cms.resolvers.go`
- Update the `CreatePostCategory` and `UpdatePostCategory` code accordingly
- Update the `gqlsa/graph/gqlsaresolver/ent.resolvers.go` `func (r *queryResolver) Node` to resolve the node type