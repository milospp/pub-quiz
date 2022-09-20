
import React from 'react'
import configData from '../../config.js'
import axios from 'axios';
import { useSelector, useDispatch } from 'react-redux'
import { setUser } from '../../store/userRedux'
import { toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';

export const Login = () => {
  const navigate = useNavigate();
  const loggedUser = useSelector((state) => state.user.value)
  const dispatch = useDispatch()

  const [loginData, setLoginData] = React.useState({
    username: "milos",
    password: "Password2",
  });


  const [registerData, setRegisterData] = React.useState({
    firstname: "",
    lastname: "",
    email: "",
    username: "",
    password: "",
  });


  function submitLogin(e) {
    e.preventDefault()
    console.log(loginData);
    axios({
      method: "post",
      url: `${configData.GATEWAY_SERVICE_URL}/login`,
      data: loginData,
    }).then((result) => {
      console.log(result);
      sessionStorage.setItem('token', result.data);
      axios.defaults.headers.common['Authorization'] = `Bearer ${result.data}` 
      
      getUser();
      // toast.success("Successfuly");
      // getUser();
      navigate("/quizzes")
    })
  }

  function getUser() {
    axios({
        method: "get",
        url: `${configData.GATEWAY_SERVICE_URL}/profile`,
  
        headers: {
            'Authorization': `Bearer ${window.sessionStorage.getItem("token")}` 
        }
    })
    .then((result) => {
        toast.success("Successfuly");

        dispatch(setUser(result.data))
  
    })
    .catch((error) => {            
        console.log(error);
        dispatch(setUser(null))

        // toast.warning(error?.response?.data?.detail);
        // history.push("/home")
  
    });
  }


  function submitRegister(e) {
    e.preventDefault()
    console.log(registerData);
    axios({
      method: "post",
      url: `${configData.GATEWAY_SERVICE_URL}/register`,
      data: registerData,
    }).then((result) => {
      console.log(result);
      sessionStorage.setItem('token', result.data);
      axios.defaults.headers.common['Authorization'] = `Bearer ${result.data}` 
      
      submitLogin(e);
      // toast.success("Successfuly");
      // getUser();
      // history.push("/home")
    })
  }


  function changedValueLogin(event, setStateFunction) {
    const { name, value, type, checked } = event.target
    setStateFunction(prevFormData => {
      let result = type === "checkbox" ? checked : value

      return {
          ...prevFormData,
          [name]: result
        }
    });
  }
  React.useEffect(() => {
    dispatch(setUser(null))
  }, [])

  return (

    <div className='container-md'>
      <span>{loggedUser?.username}</span>

      <div className='mb-3'>
        <form onSubmit={submitLogin}>
          <div className="input-component">
            <input className='input' onChange={e => changedValueLogin(e, setLoginData)} value={loginData.username} type="text" name="username" id="username" />
            <label className='placeholder' htmlFor="username">Username</label>
          </div>

          <div className="input-component">
            <input className='input' onChange={e => changedValueLogin(e, setLoginData)} value={loginData.password} type="password" name="password" id="password" />
            <label className='placeholder' htmlFor="username">Password</label>
          </div>

          <button className='btn'>Login</button>
        </form>
      </div>


      <hr />

      <form onSubmit={submitRegister}>

        <div className="input-component">
          <input className='input' onChange={e => changedValueLogin(e, setRegisterData)} type="text" name="firstname" id="firstname" />
          <label className='placeholder' htmlFor="firstname">First Name</label>
        </div>
        <div className="input-component">
          <input className='input' onChange={e => changedValueLogin(e, setRegisterData)} type="text" name="lastname" id="lastname" />
          <label className='placeholder' htmlFor="lastname">Last Name</label>
        </div>
        <div className="input-component">
          <input className='input' onChange={e => changedValueLogin(e, setRegisterData)} type="text" name="username" id="username" />
          <label className='placeholder' htmlFor="username">Username</label>
        </div>
        <div className="input-component">
          <input className='input' onChange={e => changedValueLogin(e, setRegisterData)} type="email" name="email" id="email" />
          <label className='placeholder' htmlFor="email">Email</label>
        </div>
        <div className="input-component">
          <input className='input' onChange={e => changedValueLogin(e, setRegisterData)} type="password" name="password" id="password" />
          <label className='placeholder' htmlFor="password">Password</label>
        </div>
        <div className="input-component">
          <input className='input' onChange={e => changedValueLogin(e, setRegisterData)} type="password" name="repeat-password" id="repeat-password" />
          <label className='placeholder' htmlFor="repeat-password">Repeat Password</label>
        </div>

        <button className='btn'>Register</button>


      </form>
    </div>
  )
}


