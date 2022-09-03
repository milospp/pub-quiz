import React, { Fragment } from 'react'
import { useParams } from 'react-router-dom'
import configData from '../../config'
import { toast } from 'react-toastify';
import axios from 'axios';
import { useSelector } from 'react-redux';
import QuizLobby from '../../components/quiz-game/quiz-lobby/quiz-lobby.component';
import QuizGamePlayer from '../../components/quiz-game/quiz-game-player/quiz-game-player.component';


export default function QuizGame() {
  const {code} = useParams();
  const loggedUser = useSelector((state) => state.user.value)

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
    let ws = new WebSocket(configData.QUIZ_SOCKET_SERVICE_URL)
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

  function RenderControlls() {
    if (quizData?.organizer_id === loggedUser.id && quizConn) {
      if (quizState === "LOBBY") return (<div>
        <button className='btn' onClick={sendStartGame}>Start Quiz</button>
      </div>)
      if (quizState === "QUIZ") return (<div>
        <button className='btn' onClick={sendNextStage}>Next Stage</button>
      </div>)
    }
  }

  function sendAnswer() {

  }

  function RenderGame() {
    if (!quizConn) return
    if (quizState === "LOBBY") return (<QuizLobby startGame={sendStartGame} connectedUsers={connectedUsers} active />)
    if (quizState === "QUIZ") return (<QuizGamePlayer question={questionData} quiz={quizData} sendAnswer={sendAnswer} active />)
  }

  React.useEffect(() => {
    // TODO: REFACTOR
    const loadData = async () => {
      let result = await getQuizInfo();

      try {
        setQuizData(result.data)
      } catch (error) {
        toast.warning("QUIZ NOT VALID");
      }
      // connectSocket();

    }
    loadData()
    
  }, [])

  React.useEffect(() => {
    return () => {
      // quizConn?.close()
    };
  }, [quizConn])

  function processMessage(msg) {
    let message = JSON.parse(msg);
    console.log(message);

    switch (message.method) {
      case "UPDATE_USERS":
        console.log("added users");
        setConnectedUsers(message.data.users)
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

  return (
    <div className='container'>
      <h1>{quizData?.quiz_name}</h1>
      {!quizConn && <button onClick={connectSocket}>Connect socket</button>}
      <code>{quizData?.id}</code>
      {RenderGame()}
      {RenderControlls()}
    </div>

  )
}
