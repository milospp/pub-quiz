import axios from 'axios'
import React from 'react'
import { toast } from 'react-toastify'
import { QuestionEditor } from '../../components/quiz-creator/question-editor.component.jsx'
import configData from '../../config.js'
import { useNavigate } from 'react-router-dom';

import './quiz-creator.style.scss'


export function QuizCreator() {
  const navigate = useNavigate();

  const [quizInfo, setQuizInfo] = React.useState({
    quiz_name: "",
    start_schedule: "2022-08-23",
    start_schedule_time: "12:00",
  })

  const [questions, setQuestions] = React.useState([

    {
      question_text: 'Question #1',
      answer_type: "SELECT",
      answer_text: null,
      room_password: "",
      quiz_type: 0,
      answer_number: {value: 0},
      answer_options: [
        {id: 1, value: 'Number 6', correct: true},
        {id: 2, value: 'Number 61', correct: false},
        {id: 3, value: 'Number 62', correct: true},
        {id: 4, value: 'Number 123', correct: true}
      ],
  
    },
    {
      question_text: 'Question #2',
      answer_type: "NUMBER",
      answer_text: null,
      answer_number: 100,
  
    }
  ])

  function updateQuestion(question, id) {
    setQuestions(questions => {
      let newQuestions = [...questions];
      newQuestions[id] = question;
      return newQuestions;
    })
  }

  function renderQuestions() {
    return questions.map((question, index) => {
      return (
        <QuestionEditor key={index} question={question} setQuestion={updateQuestion} id={index} />
      )
    })
  }

  async function createQuiz() {
    let quizObj = {
      ...quizInfo,
      quiz_questions: questions
    }
    // TODO: FIX TIMEZONE
    quizObj.start_schedule = quizObj.start_schedule + "T" + quizObj.start_schedule_time + ":00Z"

    let res;
    try {
      res = await postQuiz(quizObj);
      navigate(`/game/${res.data.room_code}`)
      toast.success("Quiz Created")
    } catch (error) {
      toast.error("Cannot post quiz check it again")
    }

    console.log(res);
  }

  async function postQuiz(data) {
    return axios({
      method: "post",
      url: `${configData.QUIZ_SERVICE_URL}/quiz`,
      data: data,
    })
  }

  function addQuestion() {
    setQuestions(questions => {
      return [...questions, {
        question_text: 'Question #' + (questions.length + 1),
        answer_type: "SELECT",
        answer_text: null,
        answer_number: {value: 0},
        answer_options: [
          {id: 1, value: 'Answer 1', correct: true},
          {id: 2, value: 'Answer 2', correct: false},
          {id: 3, value: 'Answer 3', correct: true},
          {id: 4, value: 'Answer 4', correct: true}
        ],
      }]
    })
  }

  function handleChange(e) {
    const { name, value, type, checked } = e.target
    let result = type === "checkbox" ? checked : value
    setQuizInfo(prevData => {
      return {
        ...prevData,
        [name]: result
      }
    })
  }


  console.log("QUIZ CREATOR");

  return (
    <div className='quiz-creator container'>
      <h1>Quiz Creator</h1>


      <div className="basic-info">
        <div className='input-component'>
          <input id="firstname" className="input" name="quiz_name" type="text" onChange={handleChange} value={quizInfo.quiz_name} placeholder=" " />
          <div className="cut"></div>
          <label htmlFor="firstname" className="placeholder">Quiz name</label>
        </div>

        <div className="d-flex gap-1">

          <div className='input-component'>
            <input id="firstname" className="input" name="room_password" type="password" onChange={handleChange} value={quizInfo.room_password} placeholder=" " />
            <div className="cut"></div>
            <label htmlFor="firstname" className="placeholder">Room password</label>
          </div>

          <div className='input-component'>
            <select className='type-select' onChange={handleChange} name="quiz_type" value={quizInfo.quiz_type}>
              <option value="0">Private quiz</option>
              <option value="1">Public Quiz</option>
              <option value="2">Tournament</option>
            </select>

          </div>

        </div>

        <div className="d-flex gap-1">
          <div className='input-component'>
            <input id="firstname" className="input" name="start_schedule" onChange={handleChange} type="date" value={quizInfo.start_schedule} placeholder=" " />
            <div className="cut"></div>
            <label htmlFor="firstname" className="placeholder">Date</label>
          </div>

          <div className='input-component'>
            <input id="firstname" className="input" name="start_schedule_time" onChange={handleChange} value={quizInfo.start_schedule_time} type="time" placeholder=" " />
            <div className="cut"></div>
            <label htmlFor="firstname" className="placeholder">Time</label>
          </div>
        </div>
      </div>



      <div>
        <h4>Question</h4>
        { renderQuestions() }
        <button className="plus-question" onClick={addQuestion}>Plus</button>

        <button className='btn' onClick={createQuiz}>Create quiz</button>
      </div>


    </div>
  )
}
