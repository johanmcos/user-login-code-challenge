import React, {FormEvent, useState} from 'react';
import logo from './logo.svg';
import './App.css';

type ActionType = "login" | "register"

const apiURL = "http://localhost:8080"

function App() {
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [actionType, setActionType] = useState<ActionType>("login")
    const [status, setStatus] = useState("waiting for action")

    const handleSubmit = async (evt: FormEvent) => {
        evt.preventDefault()
        setStatus("submitting request...")
        const url = `${apiURL}/${actionType}`
        console.log(`submitting to url ${url}...`)
        try {
            const response = await fetch(url, {
                method: "POST",
                body: JSON.stringify({username, password})
            })
            console.log("response promise resolved", response)
            const data = await response.text()
            console.log(data)
            setStatus(data)
            // setStatus(JSON.parse(data))
            // setStatus(response.ok ? "success": `${response.status}: ${response.statusText}`)
        } catch (err) {
            setStatus("Error encounterd " + err)
        }
    }

    const switchActionText = () => actionType === "login" ? "Don't have an account? Click here to register" : "Already have an account? Click here to log in"

    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo"/>
                <p>{status}</p>
                <form onSubmit={handleSubmit}>
                    <label>
                        Username:
                        <input type="text" value={username} onChange={evt => setUsername(evt.target.value)}/>
                    </label>
                    <br/>
                    <label>
                        Password:
                        <input type="password" value={password} onChange={evt => setPassword(evt.target.value)}/>
                    </label>
                    <br/>
                    <button type="submit" onClick={handleSubmit}>{actionType}</button>
                </form>
                <button
                    onClick={() => setActionType(actionType === "login" ? "register" : "login")}
                    type="button">{switchActionText()}
                </button>
            </header>
        </div>
    );
}

export default App;
