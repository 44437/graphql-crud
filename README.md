```graphql
mutation createTodo {
  createTodo(input: { text: "todo-1", userId: "001" }) {
    todoId
    text
  }
}

mutation updateTodo {
  updateTodo(todoId:"1", userId:"002") {
    todoId
    text
  }
}

mutation deleteTodo {
  deleteTodo(todoId:"3", userId:"002") {
    todoId
    text
  }
}

query readTodos {
  todos {
    text
    done
    id
    user {
      name
      id
    }
  }
}
```
