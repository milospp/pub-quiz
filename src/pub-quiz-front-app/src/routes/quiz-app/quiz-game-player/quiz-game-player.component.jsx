import React from 'react'

import './quiz-game-player.style.scss'

export default function QuizGamePlayer() {
  return (
    <div className='container quiz-game-player'>
      <div className='question-text-content'>
        <h2>Lorem, ipsum dolor sit amet consectetur adipisicing elit. Qui eveniet nesciunt voluptatem rerum, officiis illum culpa? Reprehenderit officiis corporis facere libero quasi earum, ad quos neque, deserunt ut nihil exercitationem?</h2>
      </div>

      <div className='answers'>
        <div className="d-flex answer-options">
          <div className='answer-option-wrapper'>
            <div className="answer-option">
              <div className='answer-option-btn'></div>
              <div className="answer-content">
                <span className='answer-label'>A: </span>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsa dignissimos vero, animi assumenda maiores mollitia, ab sint laborum illum perferendis amet, aliquam porro numquam placeat! Porro repellat itaque libero aliquam!
              </div>
            </div>
          </div>
          <div className='answer-option-wrapper'>
            <div className="answer-option">
              <div className='answer-option-btn'></div>
              <div className="answer-content">
                <span className='answer-label'>B: </span>
                Lorem ipsum dolor sit amet.
              </div>
            </div>
          </div>
          <div className='answer-option-wrapper'>
            <div className="answer-option">
              <div className='answer-option-btn'></div>
              <div className="answer-content">
                <span className='answer-label'>C: </span>
                Lorem ipsum dolor sit amet.
              </div>
            </div>
          </div>
          <div className='answer-option-wrapper'>
            <div className="answer-option">
              <div className='answer-option-btn'></div>
              <div className="answer-content">
                <span className='answer-label'>D: </span>
                Lorem ipsum dolor sit amet.
              </div>
            </div>
          </div>
        </div>

      </div>


    </div>
  )
}
