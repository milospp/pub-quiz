
import React from 'react'

export const Login = () => {

  function submitLogin(e) {
    e.preventDefault()
    alert("login")
  }

  function submitRegister(e) {
    e.preventDefault()
    alert("register")

  }

  return (
    
    <div>
      <form onSubmit={submitLogin}>
        <div className="input-componet">
          <label htmlFor="username">Username</label>
          <input type="text" name="username" id="username" />
        </div>

        <div className="input-componet">
          <label htmlFor="username">Password</label>
          <input type="password" name="password" id="password" />
        </div>

        <button>Login</button>
      </form>


      <hr />

      <form onSubmit={submitRegister}>

        <div className="input-componet">
          <label htmlFor="firstname">First Name</label>
          <input type="text" name="firstname" id="firstname" />
        </div>
        <div className="input-componet">
          <label htmlFor="lastname">Last Name</label>
          <input type="text" name="lastname" id="lastname" />
        </div>
        <div className="input-componet">
          <label htmlFor="username">Username</label>
          <input type="text" name="username" id="username" />
        </div>
        <div className="input-componet">
          <label htmlFor="email">Email</label>
          <input type="email" name="email" id="email" />
        </div>
        <div className="input-componet">
          <label htmlFor="password">Password</label>
          <input type="password" name="password" id="password" />
        </div>
        <div className="input-componet">
          <label htmlFor="repeat-password">Repeat Password</label>
          <input type="password" name="repeat-password" id="repeat-password" />
        </div>

        <button>Register</button>


      </form>
    </div>
  )
}


