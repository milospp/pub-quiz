
import React from 'react'
import { Outlet } from 'react-router-dom'
import Navbar from '../navbar/navbar.component'
export const HomeTemplate = () => {


  return (
    
    <div>
      <Navbar />
      <Outlet />
      <h5>Footer</h5>

    </div>
  )
}


