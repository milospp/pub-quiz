import React, { Fragment } from 'react'
import { useParams } from 'react-router-dom'
import configData from '../../config'
import { toast } from 'react-toastify';
import axios from 'axios';
import QuizLobby from '../../components/quiz-game/quiz-lobby/quiz-lobby.component';
import QuizGamePlayer from '../../components/quiz-game/quiz-game-player/quiz-game-player.component';
import { setPlayer, setUser } from '../../store/userRedux'
import { useSelector, useDispatch } from 'react-redux'
import QuizStats from '../../components/quiz-game/quiz-stats/quiz-stats.component';
import QuizFinishedScreen from '../../components/quiz-game/quiz-finished-screen/quiz-finished-screen.component';
import { useNavigate } from 'react-router-dom';


export default function QuizGame() {
  const {code} = useParams();
  const loggedUser = useSelector((state) => state.user.value)
  const loggedPlayer = useSelector((state) => state.user.player)
  const dispatch = useDispatch()
  const navigate = useNavigate();


  let [quizState, setQuizState] = React.useState("LOBBY")
  let [quizData, setQuizData] = React.useState(null)
  let [quizConn, setQuizConn] = React.useState(null)

  let [connectedUsers, setConnectedUsers] = React.useState(null)
  let [questionData, setQuestionData] = React.useState(null)

  // console.log("Quiz tempalte");



  function getQuizInfo() {
    return axios({
      method: "get",
      url: `${configData.QUIZ_SERVICE_URL}/quiz/${code}`,

      headers: {
          'Authorization': `Bearer ${window.sessionStorage.getItem("token")}` 
      }
    });
  }

  function socketClosed(e) {
    console.log(e);
    dispatch(setPlayer(null))
    toast.warning("Connection closed!")
    setQuizConn(null)
  }
  
  function socketOpen(e) {
    console.log(e);
    sendAuth(e.target)
    toast.success("Connected into room")
  }
  function socketError(e) {
    console.log(e);
    toast.error("Connection Error!")
  }
  function socketMessage(msg) {
    console.log(msg);
    processMessage(msg.data)
  }

  function sendAuth(conn) {

    let authData = {
      username: loggedUser?.username,
      room_id: quizData?.id,
      method: "AUTH",
      data: {
        jwt: `Bearer ${window.sessionStorage.getItem("token")}`
      }
    }
    console.log(authData);
    conn.send(JSON.stringify(authData))
  }

  function connectSocket() {
    let ws = new WebSocket(configData.QUIZ_SOCKET_SERVICE_URL + `/${quizData?.id}`)
    ws.onopen = socketOpen
    ws.onclose = socketClosed
    ws.onerror = socketError
    ws.onmessage = socketMessage
    setQuizConn(ws)
  }

  function sendStartGame() {
    
    let startGameData = {
      username: loggedUser?.username,
      room_id: quizData?.id,
      method: "START_GAME",
      data: {
        jwt: `Bearer ${window.sessionStorage.getItem("token")}`
      }
    }
    console.log(startGameData);
    quizConn.send(JSON.stringify(startGameData))
  }

  function sendNextStage() {
    
    let startGameData = {
      username: loggedUser?.username,
      room_id: quizData?.id,
      method: "NEXT_STATE",
      data: {
        jwt: `Bearer ${window.sessionStorage.getItem("token")}`
      }
    }
    console.log(startGameData);
    quizConn.send(JSON.stringify(startGameData))
  }

  function restartGame() {
    let startGameData = {
      username: loggedUser?.username,
      room_id: quizData?.id,
      method: "RESTART",
      data: {
        jwt: `Bearer ${window.sessionStorage.getItem("token")}`
      }
    }
    console.log(startGameData);
    quizConn.send(JSON.stringify(startGameData))
  }

  function finishGame() {
    let startGameData = {
      username: loggedUser?.username,
      room_id: quizData?.id,
      method: "FINISH",
      data: {
        jwt: `Bearer ${window.sessionStorage.getItem("token")}`
      }
    }
    console.log(startGameData);
    quizConn.send(JSON.stringify(startGameData))
  }

  function showStats() {
    let startGameData = {
      username: loggedUser?.username,
      room_id: quizData?.id,
      method: "SHOW_STATS",
      data: {
        jwt: `Bearer ${window.sessionStorage.getItem("token")}`
      }
    }
    console.log(startGameData);
    quizConn.send(JSON.stringify(startGameData))
  }

  function hideStats() {
    let startGameData = {
      username: loggedUser?.username,
      room_id: quizData?.id,
      method: "HIDE_STATS",
      data: {
        jwt: `Bearer ${window.sessionStorage.getItem("token")}`
      }
    }
    console.log(startGameData);
    quizConn.send(JSON.stringify(startGameData))
  }


  function RenderControlls() {
    if (loggedPlayer?.role === 2) {
      if (quizState === "LOBBY") return (<div>
        <button className='btn' onClick={sendStartGame}>Start Quiz</button>
      </div>)
      if (quizState === "QUIZ") return (<div>
        <button className='btn' onClick={sendNextStage}>Next Stage</button>
        <button className='btn' onClick={restartGame}>Restart</button>
        <button className='btn' onClick={finishGame}>FinishGame</button>
        <button className='btn' onClick={showStats}>ShowStats</button>
      </div>)
      if (quizState === "STATS") return (<div>
        {/* <button className='btn' onClick={finishGame}>FinishGame</button> */}
        <button className='btn' onClick={hideStats}>HideStats</button>
      </div>)
      if (quizState === "FINISHED") return (<div>
        <button className='btn' onClick={restartGame}>Restart</button>
        <button className='btn' onClick={showStats}>ShowStats</button>
      </div>)
    }
  }

  function sendAnswer(question_id, answer) {
    let startGameData = {
      username: loggedUser?.username,
      room_id: quizData?.id,
      method: "ANSWER",
      data: {
        jwt: `Bearer ${window.sessionStorage.getItem("token")}`,
        question_id: question_id,
        answer: String(answer),
        timestamp: new Date().toISOString()
      }
    }
    console.log(startGameData);
    quizConn.send(JSON.stringify(startGameData))
  }

  function RenderGame() {
    if (!quizConn) return
    if (quizState === "LOBBY") return (<QuizLobby startGame={sendStartGame} connectedUsers={connectedUsers} connection={quizConn} quiz={quizData} active />)
    if (quizState === "QUIZ") return (<QuizGamePlayer question={questionData} quiz={quizData} sendAnswer={sendAnswer} active />)
    if (quizState === "STATS") return (<QuizStats quizId={quizData?.id} connectedUsers={connectedUsers} />)
    if (quizState === "FINISHED") return (<QuizFinishedScreen />)
  }

  React.useEffect(() => {
    // TODO: REFACTOR
    
    const loadData = async () => {

      try {
        let result = await getQuizInfo();

        if (loggedUser?.role !== 1 && result.data?.room_password?.Valid && result.data?.room_password.String !== "") {
          let pass = prompt("Room password")
          console.log(result.data?.room_password.String);
          if (pass !== result.data?.room_password.String) navigate("/")
        }

        setQuizData(result.data)


      } catch (error) {
        toast.warning("QUIZ NOT VALID");
        navigate("/")
      }
      // connectSocket();

    }
    loadData()
    
  }, [])

  React.useEffect(() => {
    return () => {
      dispatch(setPlayer(null))
      // quizConn?.close()
    };
  }, [quizConn])

  function processMessage(msg) {
    let message = JSON.parse(msg);
    console.log(message);

    switch (message.method) {
      case "UPDATE_USERS":
        console.log("added users");
        setConnectedUsers(message.data.players)
        updatePlayerState(message.data.players)
        break;
      case "START_GAME":
        setQuizState("QUIZ")
        break;
    
      case "QUIZ_STATE":
        console.log("new state");
        setQuizState(message.data.game_state.quiz_state)

        setQuizData((lastQuizData) => {
          let newQuizData = {...lastQuizData}
          newQuizData["quiz_state"] = message.data.game_state.quiz_state
          newQuizData["quiz_question"] = message.data.game_state.quiz_question
          newQuizData["question_state"] = message.data.game_state.question_state

          return newQuizData
        })

        setQuestionData((lastQuestionData) => {
          let newQuestionData = {
            ...lastQuestionData,
          }
          if (message.data.question_id) newQuestionData["id"] = message.data.question_id
          if (message.data.question_text) newQuestionData["question_text"] = message.data.question_text
          if (message.data.question_text) newQuestionData["question_text"] = message.data.question_text
          if (message.data.answer_type) newQuestionData["answer_type"] = message.data.answer_type
          if (message.data.question) newQuestionData={...message.data.question}
          if (message.data.game_state.question_state === 0) newQuestionData = null

          return newQuestionData
        })

        break;

      default:
        break;
    }
  }

  function updatePlayerState(players) {
    let anon = loggedUser?.anonymous_key === "" || loggedUser?.anonymous_key == null ? false : true
    players.forEach(pl => {
      if (anon && pl.anonymous_user_id === loggedUser.id) dispatch(setPlayer(pl))
      else if (!anon && pl.user_id ===loggedUser.id) dispatch(setPlayer(pl))
    });
  }

  return (
    <div className='container'>
      <h1>{quizData?.quiz_name}</h1>
      <div>
        {!quizConn && <button className='btn connect-join-btn' onClick={connectSocket}>Join</button>}
      </div>
      <small>{quizData?.room_code}</small>
      {RenderGame()}
      {RenderControlls()}
    </div>

  )
}
