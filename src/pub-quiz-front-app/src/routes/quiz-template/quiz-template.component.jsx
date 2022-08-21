
import React from 'react'
import { Outlet } from 'react-router-dom'
import QuizNavbar from '../quiz-navbar/quiz-navbar.component'

export const QuizTemplate = () => {


  return (
    
    <div>
      <QuizNavbar />
      <Outlet />
      <h5>Footer</h5>

    </div>
  )
}


