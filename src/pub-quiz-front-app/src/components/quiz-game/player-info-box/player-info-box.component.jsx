import React from 'react'
import { useSelector } from 'react-redux'

import './player-info-box.style.scss'

export default function PlayerInfoBox(props) {
  let {player, role, setPlayerRole} = props
  const loggedPlayer = useSelector((state) => state.user.player)
  const loggedUser = useSelector((state) => state.user.value)

  function changedRole(event) {
    let value = event.target.value

    setPlayerRole(player, value)
  }

  function getStatus(player){
    switch (player.status) {
      case 0:
        return "offline"
      case 1:
        return "online"
      case 2:
        return "reserved"
      case 3:
        return "banned"
    
      default:
        return "offline"
    }
  }

  function getRole(player){
    switch (player.role) {
      case 0:
        return "Player"
      case 1:
        return "Spectator"
      case 2:
        return "Organisator"
      default:
        return "Player"
    }
  }


  function getName(player) {
    if (player.anonymous_user_id > 0) return player.anonymous_user.name
    return player.user.firstname + " " + player.user.lastname
  }



  function renderRole() {
    if (loggedPlayer?.role === 2 || loggedUser?.role === 1) {
    return (

      <div>
      <select onChange={changedRole} value={player.role} name="game_role" id="game_role">
        <option value="2">Organizator</option>
        <option value="0">Player</option>
        <option value="1">Spectator</option>
      </select>
      <button>x</button>
      </div>
    )} 

  return (<span>{getRole(player)}</span>)


  }

  return (
    <div className='player-info-box'>
      <div className="d-flex">
        <div>
          <div className="player-status">{getStatus(player)}</div>
          <div>
            {getName(player)}
          </div>
        </div>
        <div className='player-role'>
          {renderRole()}
        </div>
      </div>
    </div>
  )
}
