import React from 'react'
import {Link} from 'react-router-dom'

import './navbar.style.scss'

function Navbar() {
  return (
    <nav className='navbar'>
      <div className='logo
      '>
        <h1>Logo</h1>
      </div>
      <div className='links'>
        <ul>
          <li><Link to={'/create-quiz'}>Create quiz</Link></li>
          <li><Link to={'/questions'}>Questions</Link></li>
          <li><Link to={'/lobby'}>Lobby</Link></li>
          <li>Login</li>
        </ul>
      </div>
    </nav>
  )
}

export default Navbar