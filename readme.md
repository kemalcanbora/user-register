<style>
.center {
  margin-top: 30px;
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 50%;
}
</style>

<img src="https://www.chartboost.com/wp-content/uploads/2020/04/Rollic.png" height="150" width="300" class="center">

## Quickstart
 - First of all, please be sure Docker is installed and running.
 - If your Docker is running, you can use the following command to start the server:
    ```bash
    docker-compose build && up -d
    ```
## How to run all the tests?
 - You can run all the tests by running the following command:
    ```bash
     cd tests
     go test 
    ```
   Note: Please be sure to Docker is installed and running.

## How to use this API?
 - Probably you read the Quickstart above. Now we can start to use the API.
 - You can send a request to the server like this:
   - When you want to `ADD` a new user:
      ```bash
      curl -X PUT -d '{"first_name": "rollic", "email": "hr@rollic.com", "password": "securepasswd"}' -H 'Content-Type: application/json' http://localhost:8080/user
      ```
      Result:
      ```json
      {"id":"62d28cfad0f3da2632f7e2eb",
       "first_name":"rollic",
       "email":"hr@rollic.com"}
     ```
   - When you want to `Login`:
      ```bash
      curl -X POST -d '{"email":"hr@rollic.com", "password":"securepasswd"}' http://localhost:8080/user/login
      ```
        Result:
        ```json
        {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImhyQHJvbGxpYy5jb20iLCJleHAiOjE2NTc5Njc3MjF9.o1CG9W7p3XP_DNHdK6CpQNw1No7OK4imImw14hz6EoQ"}
     ```
   - When you want to `GET` all users:
       ```bash
       curl -X GET http://localhost:8080/users/all
       ```
       Result:
        ```json
        [{"id":"62c576ea438084dd21bd0e31",
          "first_name":"hakan",
          "last_name":"taşıyan",
          "email":"hakan@gmail.com"},
         {"id":"62c576ea438084dd21bd0e32",
          "first_name":"necati",
          "last_name":"ateş",
          "email":"necati@gmail.com"},
        {"id":"62d28cfad0f3da2632f7e2eb",
          "first_name":"rollic",
          "email":"hr@rollic.com"}]
        ```
     
   - When you want to `GET` a specific user:
       ```bash
       curl -X GET http://localhost:8080/user/62c576ea438084dd21bd0e31
       ```
       Result:
       ```json
            {"id":"62c576ea438084dd21bd0e31",
             "first_name":"hakan",
             "last_name":"taşıyan",
             "email":"hakan@gmail.com"}
        ```
   - When you want to `DELETE` a specific user:
       ```bash
       curl -X DELETE http://localhost:8080/user/62d28cfad0f3da2632f7e2eb
       ```
       Result:
       ```json
         {"message":"User deleted successfully!"}
       ```
   - When you want to `UPDATE` a user:
      - Note: If you want to update a user, you have to use login endpoint with the user username and password because we need a user `token`.
       ```bash
            curl -X POST -d '{"email":"necati@gmail.com", "password":"123456"}' http://localhost:8080/user/login
            example token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6Im5lY2F0aUBnbWFpbC5jb20iLCJleHAiOjE2NTc5Njg1MjZ9.DvpExXAxiLEPuXl1BgEZhl4QJHQWReEvTjuVrqT7Jj0"
            curl -XPATCH -H 'Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6Im5lY2F0aUBnbWFpbC5jb20iLCJleHAiOjE2NTc5Njg1MjZ9.DvpExXAxiLEPuXl1BgEZhl4QJHQWReEvTjuVrqT7Jj0' -d '{"first_name":"rose"}' 'http://localhost:8080/user'       
     ```
     Result: 
     ```json
         {"message":"User updated successfully!"}
     ```
     