import React from 'react'

export default function AnswerTextEditor(props) {
  const {id, answer, setAnswer} = props


  function changeAnswerText(e) {
    const value = e.target.value
    setAnswer(value)
  }

  return (
    <div>
        <div className='d-flex'>
            <input onChange={changeAnswerText} value={answer?.value} className='choice-answer-input' type="text" name={'number-answer-' + id} id={'number-answer-' + id} />
        </div>     

    </div>
  )
}
