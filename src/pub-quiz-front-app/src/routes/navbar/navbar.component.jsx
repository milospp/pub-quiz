import React, { Fragment } from 'react'
import { useSelector } from 'react-redux'
import {Link} from 'react-router-dom'

import './navbar.style.scss'

function Navbar() {

  const loggedUser = useSelector((state) => state.user.value)


  return (
    <nav className='navbar'>
      <div className='logo
      '>
        <h1 className='c-white'><Link to={'/'}>Logo</Link></h1>
      </div>
      <div className='links'>
        <ul>
          
          {loggedUser && loggedUser.username && (
            <Fragment>
              <li><Link to={'/create-quiz'}>Create quiz</Link></li>
              <li><Link to={'/quizzes'}>Quizzes</Link></li>
              <li>{loggedUser.firstname}</li>
              <li><Link to={'/login'}>Logout</Link></li>
            
            </Fragment>
          )}
          {!loggedUser?.username && (<li><Link to={'/login'}>Login</Link></li>)}
        </ul>
      </div>
    </nav>
  )
}

export default Navbar