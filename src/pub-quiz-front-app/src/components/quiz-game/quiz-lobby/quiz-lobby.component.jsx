import React from 'react'

import './quiz-lobby.style.scss'

export default function QuizLobby(props) {
    const {startGame, connectedUsers} = props 

    function renderConnectedUsers() {
        return connectedUsers?.map((user, index) => {
            return (<li key={index}>
                <div className='player-info-box'>{user.firstname}</div>
            </li>)
        })
    }

  return (
    <div className='container quiz-lobby'>
        <div className='members'>
            <div className="d-flex players-info">
                <div className='mb-3'>
                    <h3>Connected users:</h3>
                    <ul>
                        {renderConnectedUsers()}
                    </ul>

                </div>
                {/* <div>
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

                </div> */}
            </div>
            

        </div>


        
    </div>
  )
}
