import React from 'react'
import configData from '../../../config'
import axios from 'axios';
import { toast } from 'react-toastify';


export default function QuizStats(props) {
  let {quizId, connectedUsers} = props
  let [quizStats, setQuizStats] = React.useState(null)

  function getQuizStats(quizId) {
    return axios({
      method: "get",
      url: `${configData.GATEWAY_SERVICE_URL}/quiz-stats/${quizId}`,
      headers: {
          'Authorization': `Bearer ${window.sessionStorage.getItem("token")}` 
      }
    }).then(result => {
      setQuizStats(result.data)
    }).catch(error => {
      console.log(error)
      toast.warning("QUIZ NOT VALID");
    })
  }

  function getName(player) {
    try { 
      if (player.anonymous_user_id > 0) return player.anonymous_user.name
      return player.user.firstname + " " + player.user.lastname
    } catch {
      return "Player"
    }
  }

  function renderTopScoreRows() {
    // TODO: Refactor
    if (!quizStats) return
    return Object.values(quizStats).sort((a,b)=> b.points-a.points).map((player) => {
      let pl = connectedUsers.find(x=> x.id === player.player_id)

      return (<tr>
        <td>{getName(pl)}</td>
        <td>{player.points}</td>
      </tr>)
    })

  }
  function renderFirstCorrectRows() {
    // TODO: Refactor
    if (!quizStats) return
    return Object.values(quizStats).sort((a,b)=> b.first_correct_answers-a.first_correct_answers).map((player) => {
      let pl = connectedUsers.find(x=> x.id === player.player_id)

      return (<tr>
        <td>{getName(pl)}</td>
        <td>{player.first_correct_answers}</td>
      </tr>)
    })

  }
  function renderCorrectRows() {
    // TODO: Refactor
    if (!quizStats) return
    return Object.values(quizStats).sort((a,b)=> b.correct_answers-a.correct_answers).map((player) => {
      let pl = connectedUsers.find(x=> x.id === player.player_id)

      return (<tr>
        <td>{getName(pl)}</td>
        <td>{player.correct_answers}</td>
      </tr>)
    })

  }
  function renderFirstRows() {
    // TODO: Refactor
    if (!quizStats) return
    return Object.values(quizStats).sort((a,b)=> b.first_answers-a.first_answers).map((player) => {
      let pl = connectedUsers.find(x=> x.id === player.player_id)

      return (<tr>
        <td>{getName(pl)}</td>
        <td>{player.first_answers}</td>
      </tr>)
    })

  }

  React.useEffect(() => {
    if (quizId) getQuizStats(quizId)
  },[quizId])

  return (
    <div>

      <br />
      <hr />
      <br />

      <div className="d-flex gap-2 flex-wrap">
        <div>
          <h4>Score</h4>
          <table>
            <thead>
              <tr>
                <th>Player</th>
                <th>Score</th>
              </tr>
            </thead>
            <tbody>
              {renderTopScoreRows()}
            </tbody>
          </table>
          <br />
          <hr />
          <br />
        </div>

        <div>
          <h4>Correct answers</h4>
          <table>
            <thead>
              <tr>
                <th>Player</th>
                <th>Count</th>
              </tr>
            </thead>
            <tbody>
              {renderCorrectRows()}
            </tbody>
          </table>
          <br />
          <hr />
          <br />
        </div>

        <div>
          <h4>First Correct Answer</h4>
          <table>
            <thead>
              <tr>
                <th>Player</th>
                <th>Count</th>
              </tr>
            </thead>
            <tbody>
              {renderFirstCorrectRows()}
            </tbody>
          </table>
          <br />
          <hr />
          <br />
        </div>

        <div>
          <h4>First Answer</h4>
          <table>
            <thead>
              <tr>
                <th>Player</th>
                <th>Count</th>
              </tr>
            </thead>
            <tbody>
              {renderFirstRows()}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}
