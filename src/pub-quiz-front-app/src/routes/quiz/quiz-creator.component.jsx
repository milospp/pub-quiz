import React from 'react'
import { QuestionEditor } from '../../components/quiz-creator/question-editor.component.jsx'

import './quiz-creator.style.scss'


export function QuizCreator() {

  const [questions, setQuestions] = React.useState([

    {
      questionText: 'Question #1',
      answerType: "SELECT",
      answerText: null,
      answerNumber: {value: 0},
      answersOptions: [
        {id: 1, value: 'Number 6', correct: true},
        {id: 2, value: 'Number 61', correct: false},
        {id: 3, value: 'Number 62', correct: true},
        {id: 4, value: 'Number 123', correct: true}
      ],
  
    },
    {
      questionText: 'Question #2',
      answerType: "NUMBER",
      answerText: null,
      answerNumber: {value: 100},
  
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

  function createQuiz() {
    console.log(questions);
  }

  function addQuestion() {
    setQuestions(questions => {
      return [...questions, {
        questionText: 'Question #' + (questions.length + 1),
        answerType: "SELECT",
        answerText: null,
        answerNumber: {value: 0},
        answersOptions: [
          {id: 1, value: 'Answer 1', correct: true},
          {id: 2, value: 'Answer 2', correct: false},
          {id: 3, value: 'Answer 3', correct: true},
          {id: 4, value: 'Answer 4', correct: true}
        ],
      }]
    })
  }

  console.log("QUIZ CREATOR");

  return (
    <div className='quiz-creator container'>
      <h1>Quiz Creator</h1>


      <div className="basic-info">
        <div className='input-component'>
          <input id="firstname" className="input" type="text" placeholder=" " />
          <div className="cut"></div>
          <label htmlFor="firstname" className="placeholder">Quiz name</label>
        </div>


        <div className="d-flex gap-1">
          <div className='input-component'>
            <input id="firstname" className="input" type="date" placeholder=" " />
            <div className="cut"></div>
            <label htmlFor="firstname" className="placeholder">Date</label>
          </div>

          <div className='input-component'>
            <input id="firstname" className="input" type="time" placeholder=" " />
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
