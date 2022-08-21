import React from 'react'
import AnswerNumberEdior from './answer-number-edior.component';
import AnswerSelectEdiot from './answer-select-ediot.component'
import AnswerTextEditor from './answer-text-edior.component'

import './question-editor.style.scss'

export function QuestionEditor(props) {
  const {question, setQuestion, id} = props

  function changedType(e) {
    let answerType
    let newQuestionData = {...question}

    switch (e.target.value) {
      case "SELECT": 
        if (multiple)
          newQuestionData.answerType = "SELECT"; 
        else
          newQuestionData.answerType = "MULTIPLE"; 
        
        if (newQuestionData.answersOptions == null)
          newQuestionData.answersOptions = [
            {id: 1, value: 'Number 6', correct: true},
            {id: 2, value: 'Number 61', correct: false},
            {id: 3, value: 'Number 62', correct: false},
            {id: 4, value: 'Number 123', correct: false}
          ];

        break;

      case "NUMBER":
        newQuestionData.answerType = "NUMBER"; 
        if (newQuestionData.answerNumber == null)
          newQuestionData.answerNumber = {value: 0, min:0, max:100000}
        
        break;
      

      case "TEXT":
        newQuestionData.answerType = "TEXT";
        if (newQuestionData.answerText == null)
          newQuestionData.answerText = {value: ''}
        
        
        break;

      default: newQuestionData.answerType="SELECT"
    }
    setQuestion(newQuestionData, id)

  }

  function updateParent() {
    setQuestion(question, id);
  }

  function changedMultiple(e) {
    const {checked} = e.target
    let answerType = checked ? 'MULTIPLE' : 'SELECT' 

    setQuestion({...question, answerType: answerType}, id)
  }

  function showMultipleCheckbox(show) {
    if (show) {
      return (
        <div className='multiple-cb'>
          <input onChange={changedMultiple} checked={question.answerType === 'MULTIPLE'} type="checkbox" name={'multiple-select-' + id} id={'multiple-select-' + id} />
          <label htmlFor={'multiple-select-' + id}>Multiple answers</label>
        </div>
      )
    } else {
      return
    }
  }

  function updateSelectAnswer(answersOptions) {
    let newQuestionData = {...question}
    newQuestionData.answersOptions = answersOptions
    setQuestion(newQuestionData, id)
  }

  function updateNumberAnswer(value) {
    let newQuestionData = {...question}
    newQuestionData.answerNumber.value = value
    setQuestion(newQuestionData, id)
  }

  function updateTextAnswer(value) {
    let newQuestionData = {...question}
    newQuestionData.answerText.value = value
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

  if (question?.answerType) {
    switch (question.answerType) {
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
        setQuestion({...question, answerType: "SELECT"}, id)
        
    }


  }



  return (
    <div className="question-edit-box">
      <div className='input-component'>
        <input id="firstname" className="input" onChange={handleChange} value={question.questionText} type="text" name='questionText' placeholder=" " />
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
      
      


      { typeSelect === 'SELECT' && <AnswerSelectEdiot id={id} answer={question.answersOptions} setAnswer={updateSelectAnswer} multiple={multiple} /> }
      { typeSelect === 'NUMBER' && <AnswerNumberEdior id={id} answer={question.answerNumber} setAnswer={updateNumberAnswer} /> }
      { typeSelect === 'TEXT' && <AnswerTextEditor id={id} answer={question.answerText} setAnswer={updateTextAnswer} /> }

    </div>
  )
}
