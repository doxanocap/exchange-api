import { useEffect, useState } from 'react'
import './App.css'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import api from './http'

const socket = new WebSocket('ws://localhost:8040/web-chat/pool')

function App() {
  const [count, setCount] = useState<number>(0)
  const [username, setUsername] = useState<string>('')
  const [email, setEmail] = useState<string>('')
  const [phoneNumber, setPhoneNumber] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [userList, setUserList] = useState()


  const signUp = () => {
    api.get("")
  }

  useEffect(() => {
    socket.onopen = () => {
      alert('[open] Connection established')
      alert('Sending to server')
      let msg = {
        sender_id: 1,
        recived_id: 5,
        message: 'hello my name is John',
      }

      socket.send(JSON.stringify(msg))
    }

    socket.onmessage = (event) => {
      alert(`[message] Data received from server: ${event.data}`)
    }

    socket.onclose = (event) => {
      if (event.wasClean) {
        alert(
          `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`,
        )
      } else {
        // e.g. server process killed or network down
        // event.code is usually 1006 in this case
        alert('[close] Connection died')
      }
    }

    socket.onerror = (error) => {
      console.log(error)
      alert(`[error] `)
    }
  }, [])

  return (
    <div>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
      <div className="signUp">
        <input
          type="text"
          placeholder="username"
          onClick={(e: React.FormEvent<HTMLInputElement>) =>
            setUsername(e.currentTarget.value)
          }
          value={username}
        />
        <input
          type="text"
          placeholder="email"
          onClick={(e: React.FormEvent<HTMLInputElement>) =>
            setEmail(e.currentTarget.value)
          }
          value={email}
        />
        <input
          type="text"
          placeholder="phone number"
          onClick={(e: React.FormEvent<HTMLInputElement>) =>
            setPhoneNumber(e.currentTarget.value)
          }
          value={phoneNumber}
        />
        <input
          type="text"
          placeholder="password"
          onClick={(e: React.FormEvent<HTMLInputElement>) =>
            setPassword(e.currentTarget.value)
          }
          value={password}
        />
        <button onClick={signUp}/>
      </div>
    </div>
  )
}

export default App
