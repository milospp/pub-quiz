import React from 'react'
import { Link } from 'react-router-dom'
import { useSelector } from 'react-redux'

import './quiz.style.scss'

export default function QuizCard(props) {
  let {quiz} = props
  const loggedUser = useSelector((state) => state.user.value)

  function state() {
    if (quiz?.quiz_type) return ""
    if (quiz.quiz_type === 0) return "Private"
    if (quiz.quiz_type === 1) return "Public"
    if (quiz.quiz_type === 21) return "Public Approved"
    if (quiz.quiz_type === 2) return "Private"
    
  }

  return (
    <div className='quiz-card'>
      <h3 className='quiz-name'>{quiz?.quiz_name}</h3>
      <ul>
        <li className='mb-1'>{state()}</li>
        <li>Start: {quiz?.start_schedule.slice(0,10)}</li>
      </ul>


      {loggedUser?.role === 1 && (<button className=''>Approve</button>)}

      <Link className='btn quiz-join-btn' to={'/game/' + quiz.room_code} >Join</Link>
    </div>
  )
}
