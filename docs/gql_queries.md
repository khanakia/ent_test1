
```gql
mutation {
  authRegister(input:{
    email:"khanakia@gmail.com"
    password:"admin",
    firstName:"Aman",
    lastName:"khanakia"
  }) 
}

mutation {
  authRegisterConfirm(input:{
    email:"khanakia@gmail.com"
    id:"ulieg8yyd9k9zkmujk37al"
  }) {
    me {
      id
      email
    }
  }
}
```

```gql
mutation {
  authLogin(input:{
    userName:"khanakia@gmail.com"
    password:"admin"
  }) {
    token
    me {
      id
      phone
      email
      company
      firstName
      lastName
    }
  }
}
```

```gql
query {
  me {
    id
    email
    firstName
    lastName
  }
}
```

```gql
mutation{
  authForgotPassword(userName:"khanakia@gmail.com")
}
```

```gql
mutation{
  authResetPassword(token:"0k_T8hDw8CzQpOcqO7I4i", password:"admin") {
    token
    me {
      id
      email
      firstName
      lastName
    }
  }
}


```

```gql

```

```gql

```


## Super Admin
```gql
query {
	node(id:"PostType/01J37T7P4K20CEK9ZTM2N5VXE8")  {
    __typename
    ... on PostType {
      id
      metaTitle
      metaDescr
    }
  	
  }
}
```

```gql
{
  postTypes(first: 10, where:{
    id: "tana"
  } ) {
    edges {
      node {
        id
        name
      }
      cursor
    }
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
  }
}
```

```gql
mutation {
  updatePostType(id: "a_qr9uojchqv", input:{
    name: "test2",
  }) {
    id
    
  }
}
```

```gql
mutation CreateTodo {
  createPostType(
    input: {
      name:"post",
      slug:"post",
    }
  ) {
    id
    name
  }
}
```