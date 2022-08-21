import React from 'react'
import { Link } from 'react-router-dom'

export default function QuizNavbar() {
  return (
<nav className='navbar'>
      <div className='logo
      '>
        <h1>Logo</h1>
      </div>
      <div className='links'>
        <ul>
          <li><Link to={'/player'}>Player</Link></li>
          <li><Link to={'/lobby'}>Lobby</Link></li>
        </ul>
      </div>
    </nav>
  )
}