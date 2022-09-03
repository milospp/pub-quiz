import axios from 'axios';
import React from 'react'
import { Link } from 'react-router-dom';
import { toast } from 'react-toastify';
import configData from '../../config';

export default function Quizzes() {
  let [quizzesData, setQuizzesData] = React.useState([]);

  function getQuizInfo() {
    return axios({
      method: "get",
      url: `${configData.QUIZ_SERVICE_URL}/quizzes`,

      headers: {
          'Authorization': `Bearer ${window.sessionStorage.getItem("token")}` 
      }
    })
    .then((result) => {
        // toast.success("Loaded Quiz");
        setQuizzesData(result.data)

    })
    .catch((error) => {            
        console.log(error);

        toast.warning("QUIZ NOT VALID");
        // history.push("/home")

    });
  }

  function renderQuizzes() {
    return quizzesData.map((quiz, index) => {
      return (
        <li key={index}><Link to={'/game/' + quiz.room_code} key={index} >{quiz.quiz_name}</Link></li>
      )
    })
  }

  React.useEffect(() => {
    getQuizInfo();
  }, [])

  return (
    <div className='container'>
        <h1>Quizzes</h1>
        <ul>
          {renderQuizzes()}
        </ul>
    </div>
  )
}
