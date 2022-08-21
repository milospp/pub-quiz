import React from 'react'

import './quiz-lobby.style.scss'

export default function QuizLobby() {
  return (
    <div className='container quiz-lobby'>
        <div className='members'>
            <div className="d-flex players-info">
                <div>
                    <h3 className='mb-0'>Connected</h3>
                    <ul>
                        <li>
                            <div className='player-info-box'>Marko Marko</div>    
                            
                        </li>
                        <li>
                            <div className='player-info-box'>Marko Marko</div>    
                        </li>
                        <li>
                            <div className='player-info-box'>Marko Marko</div>   
                        </li>
                    </ul>

                </div>
                <div>
                    <h3 className='mb-0'>Invited</h3>
                    <ul>
                        <li>
                            <div className='player-info-box'>Marko Marko</div>    
                            
                        </li>
                        <li>
                            <div className='player-info-box'>Marko Marko</div>    
                        </li>
                        <li>
                            <div className='player-info-box'>Marko Marko</div>   
                        </li>
                    </ul>

                </div>
            </div>
            

        </div>


        <button className='btn'>Start Quiz</button>
        
    </div>
  )
}
