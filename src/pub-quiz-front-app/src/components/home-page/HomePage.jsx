import React, { Component } from 'react'
import { useNavigate } from 'react-router-dom';
import { useSelector, useDispatch } from 'react-redux'
import configData from '../../config'
import axios from 'axios';
import { setUser } from '../../store/userRedux'

import './home-page.styles.scss'

export default function HomePage() {
  let [quizCode, setQuizCode] = React.useState("")
  const navigate = useNavigate();
  const loggedUser = useSelector((state) => state.user.value)
  const dispatch = useDispatch()



  function handleChange(e) {
    const { value } = e.target
    setQuizCode(value)
  }

  async function createAnonymous() {
    let name = prompt("Pick a display name")
    let result = await axios({
      method: "POST",
      url: `${configData.AUTH_SERVICE_URL}/anonymous-users`,
      data: JSON.stringify({"name": name})
    })

    sessionStorage.setItem('token', result.data.jwt);
    dispatch(setUser(result.data))

  }

  async function join() {
    
    if (loggedUser == null) {
      await createAnonymous()
    }

    navigate("/game/" + quizCode)
  }

  return (
    <div>
      <h1>Pub QUIZ</h1>
      <div>
        <input placeholder='123 456' value={quizCode} onChange={handleChange} className='code-input' type="number" name="" max="999999" id="" />
        <button onClick={join} className='code-input-btn'>JOIN</button>
      </div>
    </div>
  )
}
