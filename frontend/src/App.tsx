import React, {FormEvent, useState} from 'react';
import logo from './logo.svg';
import './App.css';
import {TextEncoder} from "util";

type ActionType = "login"|"register"

const apiURL = "localhost:8080"

function App() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [actionType, setActionType] = useState<ActionType>("register")
  const [status, setStatus] = useState("waiting for action")

  const handleSubmit = async (event: FormEvent) => {
    const encoder = new TextEncoder()
    const data = encoder.encode(password)
    const hash = await crypto.subtle.digest('SHA-256', data)
    try {
      const response = await fetch(`${apiURL}/${actionType}`, {
        headers: {
          'Authorization': 'Basic ' + Buffer.from(username + ":" + hash).toString('base64'),
        }
      })
      let outcome = actionType + response.ok? "succeeded" : "failed"
      setStatus(outcome)
    } catch (err) {
      setStatus("Error encountered " + err)
    }
  }

  const loginForm = (type:ActionType):JSX.Element => {
    return (
        <form>
          <label>
            Username:
            <input type="text"/>
          </label>
          <label>
            Password:
            <input type="password"/>
          </label>
          <input type="submit" value="Submit"/>
        </form>
    )
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <body className="App-body">
        <label>
          action
          <select>
            
          </select>
          </input>
        </label>
      </body>
    </div>
  );
}

export default App;
