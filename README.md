# hub

Hub is a programming language written in Go


# Hello World

```
print "Hello World"
```

# Example For loop

```
for i = 0, i < 1000, i++ {
  print i
}
```

# Example IF Condition

```
defun odd(n){
  if n mod 2 == 0 {
    return false
  }
  return true
}
```

# Example declaring var
```
var car = "Car"
```

# Example 1  support http handlers
```
//Configure builtin datasource
datasource MYSQL, MyConn {
  "url":"localhost"
  "port":3306
  "database":"todolist"
}
listen 8080 (
  post "/todo/new" >> MyConn.todoTable
  
  get "todo/:id" << MyConn.todoTable(id = :id)
  
  get "todo/:id/:date" << MyConn.todoTable(id = :id, created_at = :date)

  get "sum/:a/:b" defun(a,b) {
    return a + b
  }
  
  post "/schedule/" defun(sched){
    if sched.valid == true {
      MyConn.schedules << sched
      return {"message":"It was sched!"}
    }else {
      error("Invalid")
    }
    
  }
)
```

