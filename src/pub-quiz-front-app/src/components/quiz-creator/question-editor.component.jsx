import React from 'react'
import AnswerNumberEdior from './answer-number-edior.component';
import AnswerSelectEdiot from './answer-select-ediot.component'
import AnswerTextEditor from './answer-text-edior.component'

import './question-editor.style.scss'

export function QuestionEditor(props) {
  const {question, setQuestion, id} = props

  function changedType(e) {
    let newQuestionData = {...question}

    switch (e.target.value) {
      case "SELECT": 
        if (multiple)
          newQuestionData.answer_type = "SELECT"; 
        else
          newQuestionData.answer_type = "MULTIPLE"; 
        
        if (newQuestionData.answer_options == null)
          newQuestionData.answer_options = [
            {id: 1, value: 'Number 6', correct: true},
            {id: 2, value: 'Number 61', correct: false},
            {id: 3, value: 'Number 62', correct: false},
            {id: 4, value: 'Number 123', correct: false}
          ];

        break;

      case "NUMBER":
        newQuestionData.answer_type = "NUMBER"; 
        if (newQuestionData.answer_number == null)
          newQuestionData.answer_number = 0
        
        break;
      

      case "TEXT":
        newQuestionData.answer_type = "TEXT";
        if (newQuestionData.answer_text == null)
          newQuestionData.answer_text = ""
        
        break;

      default: newQuestionData.answer_type="SELECT"
    }
    setQuestion(newQuestionData, id)

  }

  function updateParent() {
    setQuestion(question, id);
  }

  function changedMultiple(e) {
    const {checked} = e.target
    let answer_type = checked ? 'MULTIPLE' : 'SELECT' 

    setQuestion({...question, answer_type: answer_type}, id)
  }

  function showMultipleCheckbox(show) {
    return
    // if (show) {
    //   return (
    //     <div className='multiple-cb'>
    //       <input onChange={changedMultiple} checked={question.answer_type === 'MULTIPLE'} type="checkbox" name={'multiple-select-' + id} id={'multiple-select-' + id} />
    //       <label htmlFor={'multiple-select-' + id}>Multiple answers</label>
    //     </div>
    //   )
    // } else {
    //   return
    // }
  }

  function updateSelectAnswer(answer_options) {
    let newQuestionData = {...question}
    newQuestionData.answer_options = answer_options
    setQuestion(newQuestionData, id)
  }

  function updateNumberAnswer(value) {
    let newQuestionData = {...question}
    newQuestionData.answer_number = value
    setQuestion(newQuestionData, id)
  }

  function updateTextAnswer(value) {
    let newQuestionData = {...question}
    newQuestionData.answer_text = value
    setQuestion(newQuestionData, id)
  }

  function handleChange(e) {
    const { name, value, type, checked } = e.target
    let result = type === "checkbox" ? checked : value
    setQuestion({
        ...question,
        [name]: result
    }, id)
  }
  
  let typeSelect = "SELECT";
  let multiple = false;

  if (question?.answer_type) {
    switch (question.answer_type) {
      case "SELECT":
        typeSelect = "SELECT"
        multiple = false
        break;
      case "MULTIPLE":
        typeSelect = "SELECT"
        multiple = true
        break;
      case "NUMBER":
        typeSelect = "NUMBER"
        break;
      case "TEXT":
        typeSelect = "TEXT"
        break;

      default:
        setQuestion({...question, answer_type: "SELECT"}, id)
        
    }


  }



  return (
    <div className="question-edit-box">
      <div className='input-component'>
        <input id="firstname" className="input" onChange={handleChange} value={question.question_text} type="text" name='question_text' placeholder=" " />
        <label htmlFor="firstname" className="placeholder">Questions</label>
      </div>
      <div className="d-flex mb-2">
        <select className='type-select' onChange={changedType} value={typeSelect}>
          <option value="SELECT">Ponudjeni odgovori</option>
          <option value="NUMBER">Broj</option>
          <option value="TEXT">Tekst</option>
        </select>

        { showMultipleCheckbox(typeSelect === 'SELECT') }
      </div>
      
      


      { typeSelect === 'SELECT' && <AnswerSelectEdiot id={id} answer={question.answer_options} setAnswer={updateSelectAnswer} multiple={multiple} /> }
      { typeSelect === 'NUMBER' && <AnswerNumberEdior id={id} answer={question.answer_number} setAnswer={updateNumberAnswer} /> }
      { typeSelect === 'TEXT' && <AnswerTextEditor id={id} answer={question.answer_text} setAnswer={updateTextAnswer} /> }

    </div>
  )
}
