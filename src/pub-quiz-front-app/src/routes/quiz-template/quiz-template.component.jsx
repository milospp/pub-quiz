
import React from 'react'
import { Outlet } from 'react-router-dom'
import QuizNavbar from '../quiz-navbar/quiz-navbar.component'

export const QuizTemplate = () => {

  console.log("Quiz tempalte");
  return (
    
    <div>
      <QuizNavbar />
      <Outlet />
      <h5>Footer</h5>

    </div>
  )
}


