# User Login Code Challenge

Using a JavaScript framework of your choice (preferably [React](https://reactjs.org/)), create a simple login screen that allows users to enter their username and password and submit the login form to a backend process.

Create a backend (preferably using [GoLang](https://go.dev/), but not required) that processes the login information and checks if the username and password are valid. If the login information is valid, the backend should return a success message to the user. If the login information is invalid, the backend should return an error message to the user.


## Instructions
1. Click "Use this template" to create a copy of this repository in your personal github account.  
1. Update the README in your new repo with:
    * a `How-To` section containing any instructions needed to execute your program.
    * an `Assumptions` section containing documentation on any assumptions made while interpreting the requirements.
1. Send an email to Scoir (code_challenge@scoir.com) with a link to your newly created repo containing the completed exercise (preferably no later than one day before your next interview).

## Expectations
1. This exercise is meant to drive a conversation between you and Scoir's hiring team.  
1. Please invest only enough time needed to demonstrate your approach to problem solving and code design.  
1. Within reason, treat your solution as if it would become a production system.
1. If you have any questions, feel free to contact us at code_challenge@scoir.com

## How-to
- for backend:
  - cd to the `/backend` directory
  - run `go run backend.go`
- for frontend:
  - cd to the `frontend` directory
  - run `npm start` or similar
  - since we're using `react-scripts`, there are also options for building, testing, etc.

## Assumptions

- I decided to do this using the minimum number of external libraries.
  - Go backend uses no external libraries
  - Frontend uses react + typescript 
- The "database" is just a stand-in, there is no persistence so any users registered will disappear if the backend is restarted
- Anyone can register a username/password
- There are no restrictions on length or characters currently
- Limitations:
  - without TLS encryption, a man-in-the-middle can trivially steal user credentials
  - the fact that it will tell you if a username already exists potentially helps an attacker brute force a login
  - nothing stops users from using bad passwords or even no password

