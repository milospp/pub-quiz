import React from 'react'
import PlayerInfoBox from '../player-info-box/player-info-box.component'
import { useSelector } from 'react-redux'

import './quiz-lobby.style.scss'

export default function QuizLobby(props) {
    const {startGame, connectedUsers, connection, quiz} = props 

    const loggedUser = useSelector((state) => state.user.value)
    const loggedPlayer = useSelector((state) => state.user.player)

    function renderConnectedUsers() {
        return connectedUsers?.map((player, index) => {
            return (<li key={index}>
                <PlayerInfoBox player={player} setPlayerRole={setPlayerRole} />
            </li>)
        })
    }

    function setPlayerRole(player, role) {
        let changeRoleData = {
            username: loggedUser?.username,
            room_id: quiz?.id,
            method: "CHANGLE_PLAYER_ROLE",
            data: {
              jwt: `Bearer ${window.sessionStorage.getItem("token")}`,
              role: parseInt(role),
              player_id: player.id,
            }
          }
          console.log(changeRoleData);
          connection.send(JSON.stringify(changeRoleData))
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
