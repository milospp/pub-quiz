import axios from 'axios';
import React from 'react'
import { Link } from 'react-router-dom';
import { toast } from 'react-toastify';
import QuizCard from '../../components/quiz/quiz-card.component';
import configData from '../../config';
import { useSelector } from 'react-redux'

export default function Quizzes() {
  let [quizzesData, setQuizzesData] = React.useState([]);
  const loggedUser = useSelector((state) => state.user.value)

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
    if (loggedUser?.role == null) {
      return quizzesData.filter(x => x.quiz_type==1).map((quiz, index) => {
        return (
          <QuizCard quiz={quiz} key={index} />
        )
      });
    }

    if (loggedUser?.role === 0)
    return quizzesData.filter(x => x.quiz_type===0 || x.quiz_type===21 ||  x.quiz_type===2).map((quiz, index) => {
      return (
        <QuizCard quiz={quiz} key={index} />
      )
    })

    if (loggedUser?.role === 1)
    return quizzesData.map((quiz, index) => {
      return (
        <QuizCard quiz={quiz} key={index} />
      )
    })
  }

  React.useEffect(() => {
    getQuizInfo();
  }, [])

  return (
    <div className='container'>
        <h1 className='mb-1'>Quizzes</h1>
        <div className={'d-flex gap-1'}>
          {renderQuizzes()}
        </div>
    </div>
  )
}
