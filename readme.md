<h1 align="center"> Houseware-backend assignment task</h1>

<h2 align="left"> Setup the API & database</h2>
<p>1.Fork this repo</p>
<p>2.Clon this repo or download zip file</p>
<p>3. run "<code>go mod tidy</code>"  that install all dependencies of project in local system</p>
<p>4.run "<code>go run main.go</code>" This will run this api in on localhost:8080</p>
<h3 align="left" >database</h3>
<p>1.This API already connected to  mongodb atlas you can use it.This API is designed to connect with a MongoDB Atlas database for testing purposes. We can use this database to perform various operations and test the functionality of the API.</p>
<p>2.If you want to run the API on your own database, you can update the ApplyURI function in the database.go file with your MongoDB Atlas connection string. Once you have created a database with the name Auth, you can create a collection named credentials inside it to store the user credentials.</p>
<p>3.The `credentials` collection contains two administrative user accounts that have the necessary permissions to perform tasks such as adding or deleting other users.</p>


| admin User  | Password |
| ------------- | ------------- |
| root  | root01  |
| admin  | admin01  |

<h3 align="left" >add user</h3><br>
    <p>format  for user add</p>

     "username":"pqr"
	  "email":"pqr@gmail.com"
     "password":"pqr@01"
	  "role":"admin"/"non-admin"   // use admin only for giving admin level permission

<h3 align="left" >delete user</h3><br>
    <p>Just pass the user name</p>

     "username":"pqr"
	  



<h2 align="left"> Authorization+Authentication service in Golang</h2>
 
an authorization and authentication service in Golang ensures that only authorized users can access a system or application by verifying their identity and permissions. 


<h2 align="left"> Structure</h2>

##  /controller

####  admin.go

The "admin.go" file contains three functions - "adminadd", "admindelete", and "getalluser". These functions are designed to manage user accounts in a system.

The "adminadd" and "admindelete" functions first check whether the user has administrative level permissions before allowing them to add or delete a user account. This is an important security measure to ensure that only authorized users are able to modify user accounts.

The "getalluser" function returns a list of all user accounts present in the database. This function can be useful for administrative purposes such as auditing and managing user accounts.

####     auth.go
 The "auth.go" file is a critical component of a web application's authentication system. It contains several functions, including login, logout, refresh token, and home.

The login function takes a user's username and password and verifies them against the application's database. If the user is valid, the function generates a JSON Web Token (JWT) that has a validity of one hour. This token is used to authenticate the user in subsequent requests to protected endpoints.

The logout function invalidates the JWT token, ensuring that the user is logged out of the application.

The refresh token function generates a new JWT token with a validity of 24 hours. This function is useful for extending the user's authentication session beyond the initial one-hour validity period.

The home function is a simple endpoint used to test the authentication system. It verifies that the JWT token provided by the user is valid and returns a success message if the authentication is successful.

Overall, the auth.go file plays a critical role in ensuring that the web application's authentication system is secure and reliable, protecting sensitive user information and preventing unauthorized access.

## - /database
   ####     database.go

  The setup function in the database.go file connects to a MongoDB Atlas database and establishes a connection to the auth database, using the specified credential name collection. This function returns a database connection URL, which can be used by both users and administrators to perform further activities in the project.

This connection URL enables the project to interact with the MongoDB database and access its data. It serves as a gateway for users and administrators to perform operations such as inserting, updating, and querying data in the database.

## - /middleware
   ####     middleware.go
The middleware.go file includes two key functions: GenerateHasPassword and CompareHasPassword. These functions have been implemented to enhance the security of the project.
## - /models
   ####   models.go
The "models.go" file includes four important models that are utilized in various functions within the project.

The first model, "User", is used for adding new users to the system by the admin.

The second model, "Authentication", is used for the login process of the user where they need to provide their credentials.

The third model, "Name", is used by the admin to retrieve a list of all the registered users in the system.

The fourth model, "Claims", is used for the generation and validation of JWT (JSON Web Tokens) which ensures secure communication between the client and the server.
## - /routes
   ####   routes.go
Each route is associated with a specific HTTP method and controller function that will handle the incoming request. The routes are defined within the Router function, which is responsible for creating and returning an instance of the Gorilla Mux router.

Here's a brief explanation of each route:

<p>/Refresh: This route handles GET requests and is associated with the Refresh function in the controller package.</p>
<p>/Login: This route handles POST requests and is associated with the Login function in the controller package.</p>
<p>/Logout: This route handles GET requests and is associated with the Logout function in the controller package.</p>
<p>/Admin/add: This route handles POST requests and is associated with the AdminAdd function in the controller</p> <p>package. This route is intended to be used by administrators to add new users to the system.</p>
<p>/Admin/delete: This route handles DELETE requests and is associated with the AdminDelete function in the controller package. This route is intended to be used by administrators to delete existing users from the system.</p>
<p>/getUser: This route handles GET requests and is associated with the GetUser function in the controller package. This route is intended to be used by clients to retrieve information about a specific user.</p>

## main.go
the main.go file serves as the entry point for the application, where the router is created and the server is started to handle incoming HTTP requests.


