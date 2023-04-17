# Getting Started

Install docker and run command
```docker-compose up``` to start docker

Run command ``go run main.go`` to start application

# API List

1. `POST /api/login`: login to application  
    How to test  
    endpoint `http://localhost:8080/api/login`  
    request body example  
       ```{
       "username": "tester01",
       "password": "1111"
       } ```
   
    response body example  
        ``` {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODE4MjUzMzMsImp0aSI6IjEiLCJpYXQiOjE2ODE3Mzg5MzN9.QFvBHCLFQ-NJBkChBB6PvjDXPCK-GXdbbqKzyDFQDho"
        } ```  

2. `POST /api/todos`: creates a new todo list  
   How to test  
   endpoint `http://localhost:8080/api/todos`  
   --header Authorization: Bearer <access_token>  
   request body example  
       ```{
       "message": "Buy groceries"
       }```  

3. `GET /api/todos`: retrieves todo list according to user  
   How to test  
   endpoint `http://localhost:8080/api/todos`  
   --header Authorization: Bearer <access_token>  


4. `PUT /api/todos/:id`: updates a todo list by ID  
   How to test  
   endpoint example `http://localhost:8080/api/todos/1`  
   --header Authorization: Bearer <access_token>  
   request body example  
       ```{
       "message": "Buy new groceries"
       }```

5. `DELETE /api/todos/:id`: delete a todo list by  ID  
   How to test  
   endpoint example `http://localhost:8080/api/todos/1`  
   --header Authorization: Bearer <access_token>  

# High-Level Design
The application is designed using the Hexagonal Architecture pattern, which provides a clean separation between the business logic and the infrastructure code.

Overall, this design provides a clean separation between the different components of the application, making it easy to test, maintain, and extend.  


# Stack
1. Gin as the web framework
2. Mariadb as the database for storing user and todo list data  