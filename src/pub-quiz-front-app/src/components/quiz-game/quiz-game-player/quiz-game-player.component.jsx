import React from 'react'

import './quiz-game-player.style.scss'

export default function QuizGamePlayer(props) {
  const {question, quiz, sendAnswer} = props
  const [isSubmited, setSubmited] = React.useState(false)

  const [answerData, setAnswerData] = React.useState("")

  function changeAnswer(e) {
    const { name, value, type, checked } = e.target
    let result = type === "checkbox" ? checked : value
    setAnswerData(value)
  }

  function selectAnswer(id) {
    if (question?.answer_type === "SELECT") {
      setAnswerData(id)
      setSubmited(true)
      sendAnswer(id)
    }
  }

  function submitAnswer() {
    setSubmited(true)
    sendAnswer(answerData)
  }

  React.useEffect(() => {
    if (quiz?.question_state === 0) setSubmited(false)
    if (quiz?.question_state === 4) setAnswerData(false)
  },[quiz])

  let selectAnswerClass = "answer"
  selectAnswerClass += quiz?.question_state !== 2 || isSubmited ? ' disabled' : ' '

  function isCorrectClass(id) {
    if (quiz?.question_state !== 4) return ''
    if (question?.answer_options?.[id]?.correct) return ' correct'
    else return ' wrong'
  }

  return (
    <div className='container quiz-game-player'>
      <div className='question-text-content'>
        <h2>{question?.question_text}</h2>
      </div>

      {(question?.answer_type === "SELECT" || question?.answer_type === "MULTIPLE") && (
        <div className={selectAnswerClass}>
          <div className="d-flex answer-options">
            <div className='answer-option-wrapper'>
              <div className={'answer-option' + isCorrectClass(0)}>
                <div onClick={() => {selectAnswer(0)}} className='answer-option-btn'></div>
                <div className="answer-content">
                  <span className='answer-label'>A: </span>
                  {question?.answer_options?.[0]?.value}
                </div>
              </div>
            </div>
            <div className='answer-option-wrapper'>
              <div className={'answer-option' + isCorrectClass(1)}>
                <div onClick={() => {selectAnswer("1")}} className='answer-option-btn'></div>
                <div className="answer-content">
                  <span className='answer-label'>B: </span>
                  {question?.answer_options?.[1]?.value}
                </div>
              </div>
            </div>
            <div className='answer-option-wrapper'>
              <div className={'answer-option' + isCorrectClass(2)}>
                <div onClick={() => {selectAnswer("2")}} className='answer-option-btn'></div>
                <div className="answer-content">
                  <span className='answer-label'>C: </span>
                  {question?.answer_options?.[2]?.value}
                </div>
              </div>
            </div>
            <div className='answer-option-wrapper'>
              <div className={'answer-option' + isCorrectClass(3)}>
                <div onClick={() => {selectAnswer("3")}} className='answer-option-btn'></div>
                <div className="answer-content">
                  <span className='answer-label'>D: </span>
                  {question?.answer_options?.[3]?.value}
                </div>
              </div>
            </div>
          </div>

        </div>
      )}

{(question?.answer_type === "NUMBER") && (
        <div className="d-flex">

          <div className="input-component">
            <input disabled={isSubmited} className='input'  onChange={changeAnswer} value={answerData} type="number" name="num-answer" id="num-answer" />
            <label className='placeholder' htmlFor="username">Answer</label>
          </div>
          <button onClick={submitAnswer} disabled={isSubmited} className="btn">Submit</button>

        </div>
      )}
      {(question?.answer_type === "TEXT") && (
        <div className="d-flex">
          <div className="input-component">
            <input disabled={isSubmited} className='input' onChange={changeAnswer} value={answerData} type="text" name="text-answer" id="text-answer" />
            <label className='placeholder' htmlFor="username">Answer</label>
          </div>
          <button onClick={submitAnswer} disabled={isSubmited} className="btn">Submit</button>
        </div>
      )}

    </div>
  )
}
