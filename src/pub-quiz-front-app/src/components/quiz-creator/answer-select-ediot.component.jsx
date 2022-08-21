import React from 'react'

export default function AnswerSelectEdiot(props) {
  const {multiple, id, answer, setAnswer} = props
  const [selectedData, setSelectedData] = React.useState([])

    

  function changedAnswer(e, id) {
    const checked = e.target.checked

      let newAnswerData = [...answer ]

      if (multiple) {
        newAnswerData[id].correct = checked;
      } else {
        Object.keys(newAnswerData).forEach((el, index) => {
          newAnswerData[index].correct = false
        });
        
        newAnswerData[id].correct = true;
      }

      setAnswer(newAnswerData)

  }

  function changeAnswerText(e, id) {
    const value = e.target.value

    let newAnswerOptionData = [...answer ]
    newAnswerOptionData[id].value = value
    setAnswer(newAnswerOptionData)

  }

  React.useEffect(() => {
  if (multiple) return;
    if (answer.filter(x => x.correct === true).length <= 1) return

    let newAnswerData = [...answer ]
    console.log("maltipl");

    let found = false
    newAnswerData.forEach(element => {
      if (found) {
        element.correct = false;
      } else if (element.correct === true) {
        found = true;
      }
      
    });
      
    setAnswer(newAnswerData)



  }, [multiple, answer, setAnswer])



  return (
    <div>
        <div className='d-flex'>
            <input onChange={(e) => changeAnswerText(e, 0)} value={answer[0].value} className='choice-answer-input' type="text" name="answerOption1" id="opt1" />
            <div>
                <input onChange={(e) => changedAnswer(e, 0)} className="ch-box-answer" checked={answer[0].correct} type="checkbox" name="answerCheck1" id={'answer-option' + id + '-1'} />
            </div>
        </div>
        <div className='d-flex'>
            <input onChange={(e) => changeAnswerText(e, 1)} value={answer[1].value} className='choice-answer-input' type="text" name="answerOption2" id="opt1" />
            <div>
                <input onChange={(e) => changedAnswer(e,1)} className="ch-box-answer" checked={answer[1].correct} type="checkbox" name="answerCheck2" id={'answer-option' + id + '-2'} />
            </div>
        </div>
        <div className='d-flex'>
            <input onChange={(e) => changeAnswerText(e, 2)} value={answer[2].value} className='choice-answer-input' type="text" name="answerOption3" id="opt1" />
            <div>
                <input onChange={(e) => changedAnswer(e, 2)} className="ch-box-answer" checked={answer[2].correct} type="checkbox" name="answerCheck3" id={'answer-option' + id + '-3'} />
            </div>
        </div>
        <div className='d-flex'>
            <input onChange={(e) => changeAnswerText(e, 3)} value={answer[3].value} className='choice-answer-input' type="text" name="answerOption4" id="opt1" />
            <div>
                <input onChange={(e) => changedAnswer(e, 3)} className="ch-box-answer" checked={answer[3].correct} type="checkbox" name="answerCheck4" id={'answer-option' + id + '-4'} />
            </div>
        </div>
        

    </div>
  )
}
