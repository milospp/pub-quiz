import {Routes, Route} from 'react-router-dom';
import './App.css';
import './assets/style/global.style.scss'
import './assets/style/util.style.scss'
import './assets/style/comp.style.scss'

import { Home } from './routes/home/home.component';
import { HomeTemplate } from './routes/home-template/home-template.component';
import { QuizTemplate } from './routes/quiz-template/quiz-template.component'
import { Questions } from './routes/quiz/questions.component'
import { QuizCreator } from './routes/quiz/quiz-creator.component'
import { Login } from './routes/login/login.component'
import QuizLobby from './components/quiz-game/quiz-lobby/quiz-lobby.component';
import QuizGamePlayer from './components/quiz-game/quiz-game-player/quiz-game-player.component';
import QuizGame from './routes/quiz-game/quiz-game.component';
import { useDispatch } from 'react-redux';
import axios from 'axios';
import configData from './config';
import { setUser } from './store/userRedux'
import { toast } from 'react-toastify';
import Quizzes from './routes/quizzes/Quizzes';

function App() {
  const dispatch = useDispatch()

  dispatch(setUser(JSON.parse(sessionStorage.getItem('user'))))
  axios.defaults.headers.common['Authorization'] = `Bearer ${sessionStorage.getItem('token')}` 
  function getUser() {
    axios({
        method: "get",
        url: `${configData.AUTH_SERVICE_URL}/profile`,
  
        headers: {
            'Authorization': `Bearer ${window.sessionStorage.getItem("token")}` 
        }
    })
    .then((result) => {
        dispatch(setUser(result.data))
  
    })
    .catch((error) => {            
        console.log(error);
        dispatch(setUser(null))

        // toast.warning(error?.response?.data?.detail);
        // history.push("/home")
  
    });
  }

  getUser()

  return (
      <Routes>
        <Route path='/' element={<HomeTemplate />}>
          <Route path='/login' element={<Login />} /> 
          <Route path='/questions' element={<Questions />} /> 
          <Route path='/create-quiz' element={<QuizCreator />} /> 
          <Route path='/quizzes' element={<Quizzes />} /> 

        </Route> 
        <Route path='/' element={<QuizTemplate />}> 

          <Route path='/game/:code' element={<QuizGame />} /> 
          <Route path='/lobby' element={<QuizLobby />} /> 
          <Route path='/player' element={<QuizGamePlayer />} /> 

        </Route>

      </Routes>
  );
}

export default App;
